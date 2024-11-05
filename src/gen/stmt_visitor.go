package gen

import (
	"Riddle/src/ir"
	"Riddle/src/parser"
	"fmt"
	"github.com/antlr4-go/antlr/v4"
)

type StmtVisitor struct {
	*parser.BaseRiddleParserVisitor
}

func NewStmtVisitor() *StmtVisitor {
	return &StmtVisitor{
		BaseRiddleParserVisitor: &parser.BaseRiddleParserVisitor{},
	}
}

func (v *StmtVisitor) Visit(tree antlr.ParseTree) interface{} {
	fmt.Printf("Visit %T\n", tree)
	return tree.Accept(v)
}

func (v *StmtVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	program := ir.NewProgramStmt()
	for _, i := range ctx.GetChildren() {
		if child, ok := i.(antlr.ParseTree); ok {
			stmt := child.Accept(v)
			program.AddStmt(stmt)
		} else {
			fmt.Println("Non-ParseTree child encountered")
		}

	}
	return program
}

func (v *StmtVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	fmt.Println("VisitStatement")
	return nil
}
