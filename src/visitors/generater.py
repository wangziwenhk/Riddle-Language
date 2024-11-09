from antlr4.tree.Tree import TerminalNodeImpl
from llvmlite import ir

import src.ir.statements as stmts
from src.ir.builders import Builder, BuildArg
from src.ir.statements import FuncDefineStmt
from src.parser.RiddleParser import RiddleParser
from src.parser.RiddleParserVisitor import RiddleParserVisitor


class StmtVisitor(RiddleParserVisitor):
    def __init__(self, name: str = '') -> None:
        self.name = name

    def visitProgram(self, ctx: RiddleParser.ProgramContext):
        program = stmts.ProgramStmt(self.name)
        for i in ctx.children:
            result = self.visit(i)
            program.body.append(result)
        return program

    def visitStatement_ed(self, ctx: RiddleParser.Statement_edContext):
        return self.visit(ctx.children[0])

    def visitInteger(self, ctx: RiddleParser.IntegerContext) -> stmts.IntegerStmt:
        if ctx.value is None:
            raise RuntimeError("Value is None")
        return stmts.IntegerStmt(ctx.value)

    def visitFloat(self, ctx: RiddleParser.FloatContext) -> stmts.FloatStmt:
        if ctx.value is None:
            raise RuntimeError("Value is None")
        return stmts.FloatStmt(ctx.value)

    def visitBoolean(self, ctx: RiddleParser.BooleanContext) -> stmts.BooleanStmt:
        if ctx.value is None:
            raise RuntimeError("Value is None")
        return stmts.BooleanStmt(ctx.value)

    def visitString(self, ctx: RiddleParser.StringContext) -> stmts.StringStmt:
        return stmts.StringStmt(eval(ctx.getText()))

    def visitDefineArgs(self, ctx: RiddleParser.DefineArgsContext) -> list[stmts.FuncDefineStmt.Arg]:
        args: list[stmts.FuncDefineStmt.Arg] = []
        temp_name: str = ""
        temp_type: str = ""
        if ctx.children is None:
            return args
        for i in ctx.children:
            # 参数名称
            if isinstance(i, TerminalNodeImpl):
                if i.getSymbol().type == RiddleParser.Identifier:
                    temp_name = i.getText()

            # 参数类型
            if isinstance(i, RiddleParser.TypeNameContext):
                temp_type = i.getText()

            # 压入参数
            if temp_name != "" and temp_type != "":
                args.append(stmts.FuncDefineStmt.Arg(temp_name, temp_type))
                temp_name = ""
                temp_type = ""

        return args

    def visitBodyExpr(self, ctx: RiddleParser.BodyExprContext):
        body: list[stmts.BaseStmt] = []
        for i in ctx.children:
            body.append(self.visit(i))

        return stmts.BodyStmt(body)

    def visitFuncDefine(self, ctx: RiddleParser.FuncDefineContext) -> stmts.FuncDefineStmt:
        if ctx.funcName is None:
            raise RuntimeError("FuncName is None")
        name = ctx.funcName.text

        return_type = ctx.returnType.getText()
        args = self.visit(ctx.args)
        body = self.visit(ctx.body)

        return stmts.FuncDefineStmt(name, return_type, args, body)

    def visitReturnStatement(self, ctx: RiddleParser.ReturnStatementContext):
        result = self.visit(ctx.result)
        return stmts.ReturnStmt(result)

    def visitObjectExpr(self, ctx: RiddleParser.ObjectExprContext):
        name = ctx.getText()
        return stmts.ObjectStmt(name)

    def visitVarDefineStatement(self, ctx: RiddleParser.VarDefineStatementContext):
        if ctx.name is None:
            raise RuntimeError("Name is None")
        name = ctx.name.text
        value = self.visit(ctx.value)
        typ = None
        if ctx.type_:
            typ = ctx.type_.text

        return stmts.VarDefineStmt(name, typ, value)


class GenStmt:
    def __init__(self, name: str = '') -> None:
        self.module = ir.Module(name)
        self.builder = Builder(self.module)
        self.parent: list[ir.Function] = []

        self.dispatch_table = {
            stmts.ProgramStmt: self.program,
            stmts.IntegerStmt: self.integer,
            stmts.FloatStmt: self.float,
            stmts.FuncDefineStmt: self.funcDefine,
            stmts.BodyStmt: self.body,
            stmts.ReturnStmt: self.return_,
            stmts.ObjectStmt: self.object_,
            stmts.VarDefineStmt: self.varDefine,
            stmts.FuncCallStmt: self.funcCall,
        }

    def accept(self, stmt: stmts.BaseStmt):
        # noinspection PyTypeChecker
        handler = self.dispatch_table.get(type(stmt))

        if handler is not None:
            return handler(stmt)

    def program(self, stmt: stmts.ProgramStmt) -> None:
        self.builder.push()
        for i in stmt.body:
            self.accept(i)
        self.builder.pop()

    def integer(self, stmt: stmts.IntegerStmt) -> ir.Value:
        return self.builder.get_int(stmt.get_value())

    def double(self, stmt: stmts.FloatStmt) -> ir.Value:
        return self.builder.get_float(stmt.get_value())

    def float(self, stmt: stmts.FloatStmt) -> ir.Value:
        return self.builder.get_float(stmt.get_value())

    def funcDefine(self, stmt: stmts.FuncDefineStmt) -> None:
        # 获取 function 相关选项
        arg_types: list[BuildArg] = []
        for i in stmt.args:
            name = i.name
            typ = self.builder.get_type(i.type)
            default = i.default
            arg_types.append(BuildArg(name, typ, default))

        func_name = stmt.name
        return_type = self.builder.get_type(stmt.return_type)
        function = self.builder.create_function(func_name, return_type, arg_types)

        # 添加 function call
        self.builder.add_func_call(func_name, function)
        self.builder.push()

        for i in function.args:
            self.builder.add_func_args(i.name, i)

        self.parent.append(function)
        self.accept(stmt.body)
        self.builder.pop()
        self.parent.pop(-1)
        return None

    def body(self, stmt: stmts.BodyStmt) -> None:
        for i in stmt.body:
            self.accept(i)
        return None

    def return_(self, stmt: stmts.ReturnStmt) -> None:
        result = self.accept(stmt.result)
        self.builder.create_return(result)
        return None

    def object_(self, stmt: stmts.ObjectStmt) -> ir.Value:
        name = stmt.name
        return self.builder.get_var(name)

    def varDefine(self, stmt: stmts.VarDefineStmt) -> None:
        name = stmt.name
        value: ir.Value = self.accept(stmt.value)
        typ: ir.Type
        if stmt.type is None:
            # noinspection PyUnresolvedReferences
            typ = value.type
        else:
            typ = self.builder.get_type(stmt.type)

        self.builder.create_variable(typ, name, value)
        return None

    def funcCall(self, stmt: stmts.FuncCallStmt):
        self.builder.call(stmt.name, stmt.args)
