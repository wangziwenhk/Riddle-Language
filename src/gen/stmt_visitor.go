package gen

import (
	"Riddle/src/config"
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
	if config.IsDebug {
		fmt.Printf("Visit %T\n", tree)
	}
	return tree.Accept(v)
}

func (v *StmtVisitor) VisitProgram(ctx *parser.ProgramContext) interface{} {
	program := ir.NewProgramStmt()
	for _, i := range ctx.GetChildren() {
		if child, ok := i.(antlr.ParseTree); ok {
			stmt := child.Accept(v)
			program.AddStmt(stmt.(*ir.BaseStmt))
		} else {
			fmt.Println("Non-ParseTree child encountered")
		}
	}
	return program
}

func (v *StmtVisitor) VisitStatement(ctx *parser.StatementContext) interface{} {
	s := ctx.GetChild(0).(antlr.ParseTree)
	return s.Accept(v)
}

func (v *StmtVisitor) VisitInteger(ctx *parser.IntegerContext) *ir.IntegerStmt {
	return ir.NewIntegerStmt(ctx.GetValue())
}

func (v *StmtVisitor) VisitFloat(ctx *parser.FloatContext) *ir.FloatStmt {
	return ir.NewFloatStmt(ctx.GetValue())
}

func (v *StmtVisitor) VisitBoolean(ctx *parser.BooleanContext) *ir.BooleanStmt {
	return ir.NewBooleanStmt(ctx.GetValue())
}

func (v *StmtVisitor) VisitString(ctx *parser.StringContext) *ir.StringStmt {
	str := ctx.GetText()
	return ir.NewStringStmt(str[1 : len(str)-1])
}
