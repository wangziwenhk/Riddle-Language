from antlr4.tree.Tree import TerminalNodeImpl
from llvmlite import ir

from src.ir.builders import Builder
from src.ir.errors import ParserError
from src.ir.ir_types import Class
from src.parser.RiddleParser import RiddleParser as Parser, RiddleParser
from src.parser.RiddleParserVisitor import RiddleParserVisitor


class GenVisitor(RiddleParserVisitor):
    def __init__(self, name: str = ''):
        self.module: ir.Module = ir.Module(name)
        self.builder: Builder = Builder(self.module)
        self.parent: list[ir.Function | Class] = []

    def visitProgram(self, ctx: Parser.ProgramContext):
        self.builder.push()
        self.visitChildren(ctx)
        self.builder.pop()

    def visitFuncDefine(self, ctx: RiddleParser.FuncDefineContext):
        if ctx.funcName is None:
            raise RuntimeError("FuncName is None")
        func_name: str = ctx.funcName.text
        func_args: dict[str, ir.Type] = self.visit(ctx.args)
        if ctx.returnType is None:
            return_type = ir.VoidType()
        else:
            return_type = self.visitTypeName(ctx.returnType)

        function = self.builder.create_function(func_name, return_type, func_args)
        self.builder.add_func_call(func_name, function)
        self.builder.push()

        for i in function.args:
            self.builder.add_func_args(i.name, i)

        self.parent.append(function)
        self.visit(ctx.body)
        self.builder.pop()
        self.parent.pop(-1)

    def visitDefineArgs(self, ctx: RiddleParser.DefineArgsContext) -> dict[str, ir.Type]:
        args: dict[str, ir.types] = {}
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
                args[temp_name] = self.builder.get_type(temp_type)
                temp_name = ""
                temp_type = ""

        return args

    def visitVarDefineStatement(self, ctx: RiddleParser.VarDefineStatementContext):
        if ctx.name is None:
            raise RuntimeError("name is None")
        name: str = ctx.name.text
        value = self.visit(ctx.value) if ctx.value is not None else None
        typ: ir.Type
        if ctx.type_ is not None:
            typ = self.visit(ctx.type_)
        elif value is not None:
            typ = value.type
        else:
            raise RuntimeError("None Type")
        self.builder.create_variable(typ, name, value)

    def visitIfStatement(self, ctx: RiddleParser.IfStatementContext):
        if not isinstance(self.parent[-1], ir.Function):
            raise ParserError("if statement only run in function")

        function: ir.Function = self.parent[-1]
        self.builder.push()
        cond: ir.Value = self.visit(ctx.cond)

        old_block = self.builder.get_now_block()
        then_block = function.append_basic_block('if.then')
        else_block = function.append_basic_block('if.else')

        self.builder.cond_branch(cond, then_block, else_block)

        self.builder.set_insert_block(then_block)
        self.visit(ctx.body)
        # 防止在 if 语句中 return
        try:
            self.builder.branch(old_block)
        except AssertionError as e:
            pass
        self.builder.set_insert_block(old_block)

        self.builder.set_insert_block(else_block)
        if ctx.hasElse:
            self.visit(ctx.elseBody)
        # 防止在 else 语句中 return
        try:
            self.builder.branch(old_block)
        except AssertionError as e:
            pass

        self.builder.set_insert_block(old_block)
        self.builder.pop()

    def visitForStatement(self, ctx: RiddleParser.ForStatementContext):
        if not isinstance(self.parent[-1], ir.Function):
            raise ParserError("if statement only run in function")

        function: ir.Function = self.parent[-1]

        old_block = self.builder.get_now_block()
        cond_block = function.append_basic_block("for.cond")
        loop_block = function.append_basic_block("for.loop")

        self.builder.push()
        # 初始化循环变量
        if ctx.init is not None:
            self.visit(ctx.init)

        self.builder.branch(cond_block)
        self.builder.set_insert_block(cond_block)
        cond: ir.Value = self.visit(ctx.termCond)
        self.builder.cond_branch(cond, loop_block,old_block)

        self.builder.set_insert_block(loop_block)
        self.visit(ctx.body)
        self.builder.branch(cond_block)

        self.builder.set_insert_block(old_block)

        self.builder.pop()

    def visitReturnStatement(self, ctx: RiddleParser.ReturnStatementContext):
        self.builder.create_return(self.visit(ctx.result))

    def visitTypeName(self, ctx: RiddleParser.TypeNameContext) -> ir.Type:
        name = ctx.getText()
        return self.builder.get_type(name)

    def visitInteger(self, ctx: RiddleParser.IntegerContext) -> ir.Value:
        if ctx.value is None:
            raise RuntimeError("Value is None")
        return self.builder.get_int(ctx.value)

    def visitFloat(self, ctx: RiddleParser.FloatContext) -> ir.Value:
        if ctx.value is None:
            raise RuntimeError("Value is None")
        return self.builder.get_float(ctx.value)

    def visitBoolean(self, ctx: RiddleParser.BooleanContext) -> ir.Value:
        if ctx.value is None:
            raise RuntimeError("Value is None")
        return self.builder.get_bool(ctx.value)

    def visitString(self, ctx: RiddleParser.StringContext) -> str:
        return eval(ctx.getText())

    def visitStatement_ed(self, ctx: RiddleParser.Statement_edContext):
        return self.visit(ctx.children[0])

    def visitFuncExpr(self, ctx: RiddleParser.FuncExprContext) -> ir.Value:
        if ctx.funcName is None:
            raise RuntimeError("funcName is None")

        name = ctx.funcName.text
        function: ir.Function = self.builder.get_func(name)
        args = self.visit(ctx.args)

        return self.builder.call(function, args)

    def visitArgsExpr(self, ctx: RiddleParser.ArgsExprContext) -> list[ir.Value]:
        args: list[ir.Value] = []
        for i in ctx.children:
            if isinstance(i, TerminalNodeImpl):
                continue

            args.append(self.visit(i))

        return args

    def visitObjectExpr(self, ctx: RiddleParser.ObjectExprContext):
        name = ctx.getText()
        return self.builder.get_var(name)

    def visitPtrExpr(self, ctx: RiddleParser.PtrExprContext):
        ptr = self.visit(ctx.children[0])
        if not isinstance(ptr, ir.Value):
            raise RuntimeError("Ptr value is not a IR")

        if isinstance(ptr, ir.Instruction) and ptr.opname == 'alloca':
            return self.builder.create_load(ptr)
        else:
            return ptr
