# -*- coding: utf-8 -*-
from llvmlite import ir

from src.ir.managers import VarManager, ClassManager, FunctionManager
from src.ir.ir_types import Class, BuildArg


class Builder:

    def __init__(self, module: ir.Module):
        self._module = module
        self._context = module.context
        self._llvm_builder: ir.IRBuilder | None = None
        self._var_manager = VarManager()
        self._class_manager = ClassManager()
        self._func_manager = FunctionManager()

    def get_module(self) -> ir.Module:
        return self._module

    def get_context(self) -> ir.Context:
        return self._context

    def create_function(self, name: str, return_type: ir.types = ir.VoidType(), args: list[BuildArg] | None = None):
        if args is None:
            args = []

        # 获取类型
        args_type: list[ir.Type] = []
        args_name: list[str] = []
        for i in args:
            args_name.append(i.name)
            args_type.append(i.type)

        # 创建函数
        func_type = ir.FunctionType(return_type, args_type)
        function = ir.Function(self._module, func_type, name)

        # 命名
        for i in function.args:
            if not isinstance(i, ir.Argument):
                raise RuntimeError("UNKNOWN ERROR")
            i.name = args_name[0]
            args_name.pop(0)

        # 进入函数代码块
        block = function.append_basic_block('entry')
        self._llvm_builder = ir.IRBuilder(block)

        return function

    def create_add(self, x: ir.Value, y: ir.Value) -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')
        return self._llvm_builder.add(x, y)

    def create_sub(self, x: ir.Value, y: ir.Value) -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')
        return self._llvm_builder.sub(x, y)

    # 创建一个乘法
    def create_mul(self, x: ir.Value, y: ir.Value) -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')
        return self._llvm_builder.mul(x, y)

    # 创建一个赋值指令
    def create_store(self, value: ir.Value, ptr: ir.Value) -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')
        return self._llvm_builder.store(value, ptr)

    # 创建一个 load 指令
    def create_load(self, ptr: ir.Value, name: str = '') -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')
        return self._llvm_builder.load(ptr, name)

    # 创建一个局部变量
    def create_allocate(self, typ: ir.Type, name: str = '') -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')
        return self._llvm_builder.alloca(typ, name=name)

    # 创建一个通用变量
    def create_variable(self, typ: ir.Type, name: str = '', value: ir.Value = None, is_const: bool = False) -> ir.Value:

        if self._var_manager.deep() == 1:
            result = self.create_global_var(typ, name, value, is_const)
            self._var_manager.set(name, result)
            return result
        else:
            if self._llvm_builder is None:
                raise RuntimeError('Cannot create statement outside the allowed scope')
            result = self._llvm_builder.alloca(typ, name=name)
            self.create_store(value, result)
            self._var_manager.set(name, result)
            return result

    def add_func_args(self, name: str, var: ir.Value) -> None:
        self._var_manager.set(name, var)

    # 创建一个全局变量
    def create_global_var(self, typ: ir.Type, name: str, value: ir.Value = None, is_const: bool = False) -> ir.Value:
        global_var = ir.GlobalVariable(self._module, typ, name)
        # 初始化
        global_var.initializer = value
        # 是否可变
        global_var.global_constant = is_const
        return global_var

    # 创建一个返回语句
    def create_return(self, value: ir.Value = None) -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')

        if value is None:
            return self._llvm_builder.ret_void()

        return self._llvm_builder.ret(value)

    # 进入下一个作用域
    def push(self):
        self._var_manager.push()
        self._class_manager.push()
        self._func_manager.push()

    # 推出当前作用域
    def pop(self):
        self._var_manager.pop()
        self._class_manager.pop()
        self._func_manager.pop()

    def get_type(self, name: str) -> ir.Type:
        return self._class_manager.get_type(name)

    def get_func(self, name: str) -> ir.Function:
        return self._func_manager.get(name)

    def get_class(self, name: str) -> Class:
        return self._class_manager.get_class(name)

    @staticmethod
    def get_int(value: int) -> ir.Value:
        return ir.Constant(ir.IntType(32), value)

    @staticmethod
    def get_float(value: float) -> ir.Value:
        return ir.Constant(ir.FloatType(), value)

    @staticmethod
    def get_bool(value: bool) -> ir.Value:
        return ir.Constant(ir.IntType(1), int(value))

    def cond_branch(self, cond: ir.Value, then_block: ir.Block, else_block: ir.Block) -> None:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')

        self._llvm_builder.cbranch(cond, then_block, else_block)

    def get_now_block(self) -> ir.Block:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')

        return self._llvm_builder.block

    def add_func_call(self, name: str, func: ir.Function, args: list[BuildArg] = []) -> None:

        self._func_manager.set(name, func, args)

    def set_insert_block(self, block: ir.Block) -> None:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')

        self._llvm_builder.position_at_end(block)

    def branch(self, block: ir.Block) -> None:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')

        self._llvm_builder.branch(block)

    def call(self, name: str, args: list[ir.Value], result_name: str = '', cconv=None, tail: bool = False,
             fastmath: tuple = (), attrs: tuple = (), arg_attrs=None) -> ir.Value:
        if self._llvm_builder is None:
            raise RuntimeError('Cannot create statement outside the allowed scope')

        function = self._func_manager.get(name)

        return self._llvm_builder.call(function, args, result_name, cconv, tail, fastmath, attrs, arg_attrs)

    def get_var(self, name: str) -> ir.Value:
        return self._var_manager.get(name)
