package gen

import (
	"Riddle/src/ir"
	"fmt"
	"github.com/llir/llvm/ir/types"
)
import "github.com/llir/llvm/ir/constant"

type StmtGen struct {
	ir.BaseStmtVisitor
}

func (v *StmtGen) Program(stmt *ir.ProgramStmt) interface{} {
	for e := stmt.Stmts.Front(); e != nil; e = e.Next() {
		if s, ok := e.Value.(ir.BaseStmt); ok {
			s.Accept(v)
		} else {
			fmt.Println("stmt is not a BaseStmt")
		}
	}
	return nil
}

func (v *StmtGen) Integer(stmt *ir.IntegerStmt) interface{} {
	return constant.NewInt(types.NewInt(32), int64(stmt.GetValue()))
}

func (v *StmtGen) Float(stmt *ir.FloatStmt) interface{} {
	return constant.NewFloat(types.Float, stmt.GetValue())
}

func (v *StmtGen) String(stmt *ir.StringStmt) interface{} {
	return constant.CharArray{Typ: types.NewArray(stmt.Length(), types.I8), X: []byte(stmt.GetValue())}
}

func (v *StmtGen) Boolean(stmt *ir.BooleanStmt) interface{} {
	return constant.NewBool(stmt.GetValue())
}
