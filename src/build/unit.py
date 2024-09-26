from llvmlite import ir
from antlr4 import CommonTokenStream, InputStream

from src.parser.RiddleLexer import RiddleLexer
from src.parser.RiddleParser import RiddleParser

from src.visitors.generater import GenVisitor


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
        self.gen = GenVisitor(self.name)

    def parser(self):
        self.gen.visit(get_parser_tree(self.code))

    def get_module(self):
        return self.gen.module