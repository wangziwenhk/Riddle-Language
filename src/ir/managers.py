import abc

from llvmlite import ir
from src.ir.ir_types import Class


class BaseManager(abc.ABC):
    @abc.abstractmethod
    def set(self, key, value):
        pass

    @abc.abstractmethod
    def push(self):
        pass

    @abc.abstractmethod
    def pop(self):
        pass

    @abc.abstractmethod
    def deep(self):
        pass


class VarManager(BaseManager):
    def __init__(self):
        self.vars: dict[str, list[ir.Value]] = {}
        self.defined: list[list] = []

    def get(self, item: str) -> ir.Value:
        # 该变量不存在
        if item not in self.vars:
            raise KeyError('The variable does not exist')
        return self.vars[item][-1]

    def set(self, key: str, value: ir.Value) -> None:
        self.vars[key].append(value)

    # 进入下一个作用域
    def push(self) -> None:
        self.defined.append([])

    # 退出当前作用域
    def pop(self) -> None:
        # 销毁当前作用域中声明和定义的变量
        for i in self.defined[-1]:
            self.vars[i].pop()
            # 如果一个名称没有对应的变量，则删除这个 map
            if len(self.vars[i]) == 0:
                self.vars.pop(i)

        self.defined.pop()

    def deep(self) -> int:
        return len(self.defined)


class ClassManager(BaseManager):
    def __init__(self):
        self.classes: dict[str, list[Class]] = {}
        self.defined: list[list[str]] = []



    # 仅用于获取类
    def getClass(self, item: str) -> Class:
        if item not in self.classes:
            raise KeyError('The class does not exist')
        return self.classes[item][-1]

    # 获取type(baseType)
    def getType(self, item: str) -> ir.Type:
        base_type = {
            'int': ir.IntType(32),
            'float': ir.FloatType(),
            'void': ir.VoidType(),
        }
        if item in base_type:
            return base_type[item]

        return self.getClass(item).struct

    def set(self, key: str, value: Class) -> None:
        self.classes[key].append(value)

    def push(self) -> None:
        self.defined.append([])

    def pop(self) -> None:
        for i in self.defined[-1]:
            self.classes.pop(i)
            if len(self.classes) == 0:
                self.defined.pop()

        self.defined.pop()

    def deep(self) -> int:
        return len(self.defined)
