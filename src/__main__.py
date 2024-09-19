import sys

from antlr4 import CommonTokenStream, InputStream

from src.parser.RiddleLexer import RiddleLexer
from src.parser.RiddleParser import RiddleParser
from src.visitors.generater import GenVisitor

arguments = sys.argv[1:]


def get_parser_tree(code: str):
    input_stream = InputStream(code)
    lexer = RiddleLexer(input_stream)
    token_stream = CommonTokenStream(lexer)
    parser = RiddleParser(token_stream)

    return parser.program()


if __name__ == '__main__':
    code = None
    # 如果在测试模式则读取测试文件
    if arguments[0] == '-test':
        with open('../test/main.red', 'r', encoding='utf-8') as file:
            code = file.read()

        gen = GenVisitor('main')
        gen.visit(get_parser_tree(code))
        print(gen.module)
