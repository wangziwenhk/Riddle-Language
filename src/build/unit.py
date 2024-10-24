from antlr4 import CommonTokenStream, InputStream

from src.parser.RiddleLexer import RiddleLexer
from src.parser.RiddleParser import RiddleParser
from src.visitors.generater import GenStmt, StmtVisitor


def get_parser_tree(code: str):
    input_stream = InputStream(code)
    lexer = RiddleLexer(input_stream)
    token_stream = CommonTokenStream(lexer)
    parser = RiddleParser(token_stream)

    return parser.program()


class Unit:
    def __init__(self, code: str) -> None:
        # parser package
        self.name = 'main'
        self.code: str = code
        self.gen = StmtVisitor(self.name)
        self.module = None

    def parser(self):
        root = self.gen.visit(get_parser_tree(self.code))
        p = GenStmt(self.name)
        p.accept(root)
        self.module = p.module

    def get_module(self):
        return self.module
