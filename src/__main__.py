import sys
from src.build import unit
from src.build import graph

arguments = sys.argv[1:]

if __name__ == '__main__':
    code = None
    # 如果在测试模式则读取测试文件
    if arguments[0] == '-test':
        with open('test/main.red', 'r', encoding='utf-8') as file:
            code = file.read()

        g = graph.BuildGraph()
        u = unit.Unit(code)

        g.add_node(u)
        g.parser_graph()

        with open('test/main.ll', 'w', encoding='utf-8') as file:
            file.write(str(g.get_node("main").get_module()))
        print(g.get_node("main").get_module())
