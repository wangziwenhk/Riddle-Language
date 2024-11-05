package ir

import "container/list"

type ProgramStmt struct {
	Stmts *list.List
}

func (p *ProgramStmt) AddStmt(stmt interface{}) {
	p.Stmts.PushBack(stmt)
}

func NewProgramStmt() *ProgramStmt {
	return &ProgramStmt{Stmts: list.New()}
}
