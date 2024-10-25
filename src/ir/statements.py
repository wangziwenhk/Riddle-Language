import enum


class StmtTypeID(enum.Enum):
    # 没有任何效果的语句
    NoneStmtID = -99
    # 主语句
    ProgramStmt = -1
    # 常量语句
    IntegerStmtID = 0
    FloatStmtID = 1
    StringStmtID = 2
    BooleanStmtID = 3
    NullStmtID = 4
    # 有实际行为的语句
    VarDefineStmtID = 5
    ForStmtID = 6
    WhileStmtID = 7
    FuncDefineStmtID = 8
    FuncCallStmtID = 9
    ReturnStmtID = 10
    IfStmtID = 11
    TryStmtID = 12
    ObjectStmtID = 13
    BlockStmtID = 14
    LabelStmtID = 15
    BinaryOpStmtID = 16
    UnaryOpStmtID = 17
    BodyStmtID = 18


# 基本类型
class BaseStmt:
    def __init__(self, type_id: StmtTypeID) -> None:
        self.type_id = type_id

    def isConstant(self) -> bool:
        return False


# 常量父类
class ConstantStmt(BaseStmt):
    def __init__(self, type_id: StmtTypeID) -> None:
        super().__init__(type_id)

    def isConstant(self) -> bool:
        return True


class ProgramStmt(BaseStmt):
    def __init__(self, package: str = '') -> None:
        super().__init__(StmtTypeID.ProgramStmt)
        self.body: list[BaseStmt] = []
        self.package: str = package


# 存储 int 常量的语句
class IntegerStmt(ConstantStmt):
    def __init__(self, value: int) -> None:
        super().__init__(StmtTypeID.IntegerStmtID)
        self._value = value

    def get_value(self) -> int:
        return self._value


# 存储 float 常量的语句
class FloatStmt(ConstantStmt):
    def __init__(self, value: float) -> None:
        super().__init__(StmtTypeID.FloatStmtID)
        self._value: float = value

    def get_value(self) -> float:
        return self._value


# 存储 string 常量的语句
class StringStmt(ConstantStmt):
    def __init__(self, value: str) -> None:
        super().__init__(StmtTypeID.StringStmtID)
        self._value = value

    def get_value(self) -> str:
        return self._value


# 存储 bool 常量的语句
class BooleanStmt(ConstantStmt):
    def __init__(self, value: bool) -> None:
        super().__init__(StmtTypeID.BooleanStmtID)
        self._value = value

    def get_value(self) -> bool:
        return self._value


# 存储 None
class NullStmt(ConstantStmt):
    def __init__(self) -> None:
        super().__init__(StmtTypeID.NullStmtID)

    @staticmethod
    def get_value() -> None:
        return None


# 无效果的语句
class NoneStmt(BaseStmt):
    def __init__(self) -> None:
        super().__init__(StmtTypeID.NoneStmtID)


class VarDefineStmt(BaseStmt):
    def __init__(self, name: str, typ: str, value: BaseStmt) -> None:
        super().__init__(StmtTypeID.VarDefineStmtID)
        self._name = name
        self._type = typ
        self._value = value

    def get_name(self) -> str:
        return self._name

    def get_type(self) -> str:
        return self._type

    def get_value(self) -> BaseStmt:
        return self._value


class FuncDefineStmt(BaseStmt):
    class Arg:
        def __init__(self, name: str, typ: str, default: BaseStmt = NoneStmt()) -> None:
            self.name: str = name
            self.type: str = typ
            self.default: BaseStmt = default

    def __init__(self, name: str, return_type: str, args: list[Arg], body: BaseStmt) -> None:
        super().__init__(StmtTypeID.FuncDefineStmtID)
        self.name = name
        self.return_type = return_type
        self.args = args
        self.body = body


class BodyStmt(BaseStmt):
    def __init__(self, body: list[BaseStmt]) -> None:
        super().__init__(StmtTypeID.BodyStmtID)
        self.body = body


class ReturnStmt(BaseStmt):
    def __init__(self, result: BaseStmt) -> None:
        super().__init__(StmtTypeID.ReturnStmtID)
        self.result = result


class ObjectStmt(BaseStmt):
    def __init__(self, name: str):
        super().__init__(StmtTypeID.ObjectStmtID)
        self.name = name
