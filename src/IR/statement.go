package ir

import "container/list"

type BaseStmtVisitor interface {
	Program(stmt *ProgramStmt) interface{}
	Integer(stmt *IntegerStmt) interface{}
	Float(stmt *FloatStmt) interface{}
	String(stmt *StringStmt) interface{}
	Boolean(stmt *BooleanStmt) interface{}
}

type BaseStmt interface {
	Accept(BaseStmtVisitor) interface{}
}

type ProgramStmt struct {
	Stmts *list.List
}

func NewProgramStmt() *ProgramStmt {
	return &ProgramStmt{Stmts: list.New()}
}

func (p *ProgramStmt) Accept(gen BaseStmtVisitor) interface{} {
	return gen.Program(p)
}

func (p *ProgramStmt) AddStmt(stmt *BaseStmt) {
	p.Stmts.PushBack(stmt)
}

type ConstantStmt interface {
	GetValue() interface{}
}

type IntegerStmt struct {
	value int
}

func NewIntegerStmt(value int) *IntegerStmt {
	return &IntegerStmt{value: value}
}

func (s *IntegerStmt) Accept(gen BaseStmtVisitor) interface{} {
	return gen.Integer(s)
}

func (s *IntegerStmt) GetValue() int {
	return s.value
}

type FloatStmt struct {
	value float64
}

func NewFloatStmt(value float64) *FloatStmt {
	return &FloatStmt{value: value}
}

func (s *FloatStmt) Accept(gen BaseStmtVisitor) interface{} {
	return gen.Float(s)
}

func (s *FloatStmt) GetValue() float64 {
	return s.value
}

type BooleanStmt struct {
	value bool
}

func NewBooleanStmt(value bool) *BooleanStmt {
	return &BooleanStmt{value: value}
}

func (s *BooleanStmt) Accept(gen BaseStmtVisitor) interface{} {
	return gen.Boolean(s)
}

func (s *BooleanStmt) GetValue() bool {
	return s.value
}

type StringStmt struct {
	value string
}

func NewStringStmt(value string) *StringStmt {
	return &StringStmt{value: value}
}

func (s *StringStmt) Accept(gen BaseStmtVisitor) interface{} {
	return gen.String(s)
}

func (s *StringStmt) GetValue() string {
	return s.value
}

func (s *StringStmt) Length() uint64 {
	return uint64(len(s.value))
}
