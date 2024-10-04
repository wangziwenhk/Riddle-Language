import enum

from llvmlite import ir


class TypeID(enum.Enum):
    VariableType = 1
    ClassType = 2


class BaseType:
    def __init__(self, typeid: TypeID):
        self.type = typeid

    def getType(self) -> TypeID:
        return self.type


class Class(BaseType):
    def __init__(self, context: ir.Context, name: str, types: dict[str, ir.Type] | None = None, packed: bool = False):
        super().__init__(TypeID.ClassType)
        self.name: str = name
        self.struct: ir.IdentifiedStructType = ir.IdentifiedStructType(context, name, packed)
        self.members: dict[int, str] = {}
        if types is not None:
            for i in types:
                print(i)

    def getName(self) -> str:
        return self.name

    def add_member(self, name: str, typ: ir.Type):
        self.struct.set_body(self.struct.elements.append(typ))
        self.members[len(self.members)] = name