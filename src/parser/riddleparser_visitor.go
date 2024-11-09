// Code generated from E:/Riddle-Go/RiddleParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RiddleParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RiddleParser.
type RiddleParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RiddleParser#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by RiddleParser#null_cnt.
	VisitNull_cnt(ctx *Null_cntContext) interface{}

	// Visit a parse tree produced by RiddleParser#statement_ed.
	VisitStatement_ed(ctx *Statement_edContext) interface{}

	// Visit a parse tree produced by RiddleParser#statement.
	VisitStatement(ctx *StatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#bodyExpr.
	VisitBodyExpr(ctx *BodyExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#packStatement.
	VisitPackStatement(ctx *PackStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#importStatement.
	VisitImportStatement(ctx *ImportStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#varDefineStatement.
	VisitVarDefineStatement(ctx *VarDefineStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#argsExpr.
	VisitArgsExpr(ctx *ArgsExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#defineArgs.
	VisitDefineArgs(ctx *DefineArgsContext) interface{}

	// Visit a parse tree produced by RiddleParser#funcDefine.
	VisitFuncDefine(ctx *FuncDefineContext) interface{}

	// Visit a parse tree produced by RiddleParser#forStatement.
	VisitForStatement(ctx *ForStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#whileStatement.
	VisitWhileStatement(ctx *WhileStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#continueStatement.
	VisitContinueStatement(ctx *ContinueStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#breakStatement.
	VisitBreakStatement(ctx *BreakStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#returnStatement.
	VisitReturnStatement(ctx *ReturnStatementContext) interface{}

	// Visit a parse tree produced by RiddleParser#classDefine.
	VisitClassDefine(ctx *ClassDefineContext) interface{}

	// Visit a parse tree produced by RiddleParser#tryExpr.
	VisitTryExpr(ctx *TryExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#catchExpr.
	VisitCatchExpr(ctx *CatchExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#funcExpr.
	VisitFuncExpr(ctx *FuncExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#objectExpr.
	VisitObjectExpr(ctx *ObjectExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#blendExpr.
	VisitBlendExpr(ctx *BlendExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#exprPtrParser.
	VisitExprPtrParser(ctx *ExprPtrParserContext) interface{}

	// Visit a parse tree produced by RiddleParser#selfSubRightExpr.
	VisitSelfSubRightExpr(ctx *SelfSubRightExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#andAssignExpr.
	VisitAndAssignExpr(ctx *AndAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#nullExpr.
	VisitNullExpr(ctx *NullExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#modExpr.
	VisitModExpr(ctx *ModExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#castExpr.
	VisitCastExpr(ctx *CastExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#lShrExpr.
	VisitLShrExpr(ctx *LShrExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#negativeExpr.
	VisitNegativeExpr(ctx *NegativeExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#numberExpr.
	VisitNumberExpr(ctx *NumberExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#mulAssignExpr.
	VisitMulAssignExpr(ctx *MulAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#xorAssignExpr.
	VisitXorAssignExpr(ctx *XorAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#subExpr.
	VisitSubExpr(ctx *SubExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#greaterEqualExpr.
	VisitGreaterEqualExpr(ctx *GreaterEqualExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#addAssignExpr.
	VisitAddAssignExpr(ctx *AddAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#orAssignExpr.
	VisitOrAssignExpr(ctx *OrAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#bitXorExpr.
	VisitBitXorExpr(ctx *BitXorExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#ptrExpr.
	VisitPtrExpr(ctx *PtrExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#stringExpr.
	VisitStringExpr(ctx *StringExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#greaterExpr.
	VisitGreaterExpr(ctx *GreaterExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#divAssignExpr.
	VisitDivAssignExpr(ctx *DivAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#notEqualExpr.
	VisitNotEqualExpr(ctx *NotEqualExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#selfSubLeftExpr.
	VisitSelfSubLeftExpr(ctx *SelfSubLeftExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#divExpr.
	VisitDivExpr(ctx *DivExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#bitAndExpr.
	VisitBitAndExpr(ctx *BitAndExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#assignExpr.
	VisitAssignExpr(ctx *AssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#selfAddRightExpr.
	VisitSelfAddRightExpr(ctx *SelfAddRightExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#subAssignExpr.
	VisitSubAssignExpr(ctx *SubAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#bracketExpr.
	VisitBracketExpr(ctx *BracketExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#booleanExpr.
	VisitBooleanExpr(ctx *BooleanExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#aShrExpr.
	VisitAShrExpr(ctx *AShrExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#lessExpr.
	VisitLessExpr(ctx *LessExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#orExpr.
	VisitOrExpr(ctx *OrExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#selfAddLeftExpr.
	VisitSelfAddLeftExpr(ctx *SelfAddLeftExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#lShrAssignExpr.
	VisitLShrAssignExpr(ctx *LShrAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#lessEqualExpr.
	VisitLessEqualExpr(ctx *LessEqualExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#aShrAssignExpr.
	VisitAShrAssignExpr(ctx *AShrAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#bitOrExpr.
	VisitBitOrExpr(ctx *BitOrExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#notExpr.
	VisitNotExpr(ctx *NotExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#shlAssignExpr.
	VisitShlAssignExpr(ctx *ShlAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#modAssignExpr.
	VisitModAssignExpr(ctx *ModAssignExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#positiveExpr.
	VisitPositiveExpr(ctx *PositiveExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#shlExpr.
	VisitShlExpr(ctx *ShlExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#equalExpr.
	VisitEqualExpr(ctx *EqualExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#andExpr.
	VisitAndExpr(ctx *AndExprContext) interface{}

	// Visit a parse tree produced by RiddleParser#id.
	VisitId(ctx *IdContext) interface{}

	// Visit a parse tree produced by RiddleParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by RiddleParser#boolean.
	VisitBoolean(ctx *BooleanContext) interface{}

	// Visit a parse tree produced by RiddleParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by RiddleParser#float.
	VisitFloat(ctx *FloatContext) interface{}

	// Visit a parse tree produced by RiddleParser#integer.
	VisitInteger(ctx *IntegerContext) interface{}

	// Visit a parse tree produced by RiddleParser#templateArg.
	VisitTemplateArg(ctx *TemplateArgContext) interface{}

	// Visit a parse tree produced by RiddleParser#templateArgs.
	VisitTemplateArgs(ctx *TemplateArgsContext) interface{}

	// Visit a parse tree produced by RiddleParser#typeName.
	VisitTypeName(ctx *TypeNameContext) interface{}
}
