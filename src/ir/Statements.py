import enum


class StmtTypeID(enum.Enum):
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
    ObjStmtID = 13
    BlockStmtID = 14
    LabelStmtID = 15
    BinaryOpStmtID = 16
    UnaryOpStmtID = 17
    # 没有任何效果的语句
    NoneStmtID = -1


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
