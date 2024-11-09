// Code generated from E:/Riddle-Go/RiddleParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RiddleParser

import "github.com/antlr4-go/antlr/v4"

type BaseRiddleParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRiddleParserVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNull_cnt(ctx *Null_cntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitStatement_ed(ctx *Statement_edContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitStatement(ctx *StatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBodyExpr(ctx *BodyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitPackStatement(ctx *PackStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitImportStatement(ctx *ImportStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitVarDefineStatement(ctx *VarDefineStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitArgsExpr(ctx *ArgsExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitDefineArgs(ctx *DefineArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitFuncDefine(ctx *FuncDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitForStatement(ctx *ForStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitWhileStatement(ctx *WhileStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitContinueStatement(ctx *ContinueStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBreakStatement(ctx *BreakStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitReturnStatement(ctx *ReturnStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitClassDefine(ctx *ClassDefineContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitTryExpr(ctx *TryExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitCatchExpr(ctx *CatchExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitFuncExpr(ctx *FuncExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitObjectExpr(ctx *ObjectExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBlendExpr(ctx *BlendExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitExprPtrParser(ctx *ExprPtrParserContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitSelfSubRightExpr(ctx *SelfSubRightExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAndAssignExpr(ctx *AndAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNullExpr(ctx *NullExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitModExpr(ctx *ModExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitCastExpr(ctx *CastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitLShrExpr(ctx *LShrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNegativeExpr(ctx *NegativeExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNumberExpr(ctx *NumberExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitMulAssignExpr(ctx *MulAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitXorAssignExpr(ctx *XorAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitSubExpr(ctx *SubExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitGreaterEqualExpr(ctx *GreaterEqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAddAssignExpr(ctx *AddAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitOrAssignExpr(ctx *OrAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBitXorExpr(ctx *BitXorExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitPtrExpr(ctx *PtrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitStringExpr(ctx *StringExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitGreaterExpr(ctx *GreaterExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitDivAssignExpr(ctx *DivAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNotEqualExpr(ctx *NotEqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitSelfSubLeftExpr(ctx *SelfSubLeftExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitDivExpr(ctx *DivExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBitAndExpr(ctx *BitAndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAssignExpr(ctx *AssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitSelfAddRightExpr(ctx *SelfAddRightExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitSubAssignExpr(ctx *SubAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBracketExpr(ctx *BracketExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBooleanExpr(ctx *BooleanExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAShrExpr(ctx *AShrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitLessExpr(ctx *LessExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitOrExpr(ctx *OrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitSelfAddLeftExpr(ctx *SelfAddLeftExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitLShrAssignExpr(ctx *LShrAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitLessEqualExpr(ctx *LessEqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAShrAssignExpr(ctx *AShrAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBitOrExpr(ctx *BitOrExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNotExpr(ctx *NotExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitShlAssignExpr(ctx *ShlAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitModAssignExpr(ctx *ModAssignExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitPositiveExpr(ctx *PositiveExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitShlExpr(ctx *ShlExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitEqualExpr(ctx *EqualExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitAndExpr(ctx *AndExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitId(ctx *IdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitBoolean(ctx *BooleanContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitFloat(ctx *FloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitInteger(ctx *IntegerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitTemplateArg(ctx *TemplateArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitTemplateArgs(ctx *TemplateArgsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRiddleParserVisitor) VisitTypeName(ctx *TypeNameContext) interface{} {
	return v.VisitChildren(ctx)
}
