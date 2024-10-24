from llvmlite import ir

import src.ir.statements as stmts
from src.ir.builders import Builder, BuildArg
from src.parser.RiddleParser import RiddleParser
from src.parser.RiddleParserVisitor import RiddleParserVisitor


class StmtVisitor(RiddleParserVisitor):
    def __init__(self, name: str = '') -> None:
        self.name = name

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


class GenStmt:
    def __init__(self, name: str = '') -> None:
        self.module = ir.Module(name)
        self.builder = Builder(self.module)
        self.parent = []

    def accept(self, stmt: stmts.BaseStmt):
        if isinstance(stmt, stmts.ProgramStmt):
            return self.program(stmt)

        if isinstance(stmt, stmts.IntegerStmt):
            return self.integer(stmt)

    def program(self, stmt: stmts.ProgramStmt):
        for i in stmt.body:
            self.accept(i)

    def integer(self, stmt: stmts.IntegerStmt):
        return self.builder.get_int(stmt.get_value())

    def double(self, stmt: stmts.FloatStmt):
        return self.builder.get_float(stmt.get_value())

    def float(self, stmt: stmts.FloatStmt):
        return self.builder.get_float(stmt.get_value())

    def funcDefine(self, stmt: stmts.FuncDefineStmt):
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
