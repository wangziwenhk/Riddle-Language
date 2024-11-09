// Code generated from E:/Riddle-Go/RiddleParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RiddleParser

import "github.com/antlr4-go/antlr/v4"

// RiddleParserListener is a complete listener for a parse tree produced by RiddleParser.
type RiddleParserListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterNull_cnt is called when entering the null_cnt production.
	EnterNull_cnt(c *Null_cntContext)

	// EnterStatement_ed is called when entering the statement_ed production.
	EnterStatement_ed(c *Statement_edContext)

	// EnterStatement is called when entering the statement production.
	EnterStatement(c *StatementContext)

	// EnterBodyExpr is called when entering the bodyExpr production.
	EnterBodyExpr(c *BodyExprContext)

	// EnterPackStatement is called when entering the packStatement production.
	EnterPackStatement(c *PackStatementContext)

	// EnterImportStatement is called when entering the importStatement production.
	EnterImportStatement(c *ImportStatementContext)

	// EnterVarDefineStatement is called when entering the varDefineStatement production.
	EnterVarDefineStatement(c *VarDefineStatementContext)

	// EnterArgsExpr is called when entering the argsExpr production.
	EnterArgsExpr(c *ArgsExprContext)

	// EnterDefineArgs is called when entering the defineArgs production.
	EnterDefineArgs(c *DefineArgsContext)

	// EnterFuncDefine is called when entering the funcDefine production.
	EnterFuncDefine(c *FuncDefineContext)

	// EnterForStatement is called when entering the forStatement production.
	EnterForStatement(c *ForStatementContext)

	// EnterWhileStatement is called when entering the whileStatement production.
	EnterWhileStatement(c *WhileStatementContext)

	// EnterContinueStatement is called when entering the continueStatement production.
	EnterContinueStatement(c *ContinueStatementContext)

	// EnterBreakStatement is called when entering the breakStatement production.
	EnterBreakStatement(c *BreakStatementContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterReturnStatement is called when entering the returnStatement production.
	EnterReturnStatement(c *ReturnStatementContext)

	// EnterClassDefine is called when entering the classDefine production.
	EnterClassDefine(c *ClassDefineContext)

	// EnterTryExpr is called when entering the tryExpr production.
	EnterTryExpr(c *TryExprContext)

	// EnterCatchExpr is called when entering the catchExpr production.
	EnterCatchExpr(c *CatchExprContext)

	// EnterFuncExpr is called when entering the funcExpr production.
	EnterFuncExpr(c *FuncExprContext)

	// EnterObjectExpr is called when entering the objectExpr production.
	EnterObjectExpr(c *ObjectExprContext)

	// EnterBlendExpr is called when entering the blendExpr production.
	EnterBlendExpr(c *BlendExprContext)

	// EnterExprPtrParser is called when entering the exprPtrParser production.
	EnterExprPtrParser(c *ExprPtrParserContext)

	// EnterSelfSubRightExpr is called when entering the selfSubRightExpr production.
	EnterSelfSubRightExpr(c *SelfSubRightExprContext)

	// EnterAndAssignExpr is called when entering the andAssignExpr production.
	EnterAndAssignExpr(c *AndAssignExprContext)

	// EnterNullExpr is called when entering the nullExpr production.
	EnterNullExpr(c *NullExprContext)

	// EnterModExpr is called when entering the modExpr production.
	EnterModExpr(c *ModExprContext)

	// EnterCastExpr is called when entering the castExpr production.
	EnterCastExpr(c *CastExprContext)

	// EnterLShrExpr is called when entering the lShrExpr production.
	EnterLShrExpr(c *LShrExprContext)

	// EnterNegativeExpr is called when entering the negativeExpr production.
	EnterNegativeExpr(c *NegativeExprContext)

	// EnterNumberExpr is called when entering the numberExpr production.
	EnterNumberExpr(c *NumberExprContext)

	// EnterMulAssignExpr is called when entering the mulAssignExpr production.
	EnterMulAssignExpr(c *MulAssignExprContext)

	// EnterXorAssignExpr is called when entering the xorAssignExpr production.
	EnterXorAssignExpr(c *XorAssignExprContext)

	// EnterSubExpr is called when entering the subExpr production.
	EnterSubExpr(c *SubExprContext)

	// EnterGreaterEqualExpr is called when entering the greaterEqualExpr production.
	EnterGreaterEqualExpr(c *GreaterEqualExprContext)

	// EnterAddAssignExpr is called when entering the addAssignExpr production.
	EnterAddAssignExpr(c *AddAssignExprContext)

	// EnterOrAssignExpr is called when entering the orAssignExpr production.
	EnterOrAssignExpr(c *OrAssignExprContext)

	// EnterBitXorExpr is called when entering the bitXorExpr production.
	EnterBitXorExpr(c *BitXorExprContext)

	// EnterPtrExpr is called when entering the ptrExpr production.
	EnterPtrExpr(c *PtrExprContext)

	// EnterStringExpr is called when entering the stringExpr production.
	EnterStringExpr(c *StringExprContext)

	// EnterGreaterExpr is called when entering the greaterExpr production.
	EnterGreaterExpr(c *GreaterExprContext)

	// EnterDivAssignExpr is called when entering the divAssignExpr production.
	EnterDivAssignExpr(c *DivAssignExprContext)

	// EnterMulExpr is called when entering the mulExpr production.
	EnterMulExpr(c *MulExprContext)

	// EnterNotEqualExpr is called when entering the notEqualExpr production.
	EnterNotEqualExpr(c *NotEqualExprContext)

	// EnterSelfSubLeftExpr is called when entering the selfSubLeftExpr production.
	EnterSelfSubLeftExpr(c *SelfSubLeftExprContext)

	// EnterDivExpr is called when entering the divExpr production.
	EnterDivExpr(c *DivExprContext)

	// EnterBitAndExpr is called when entering the bitAndExpr production.
	EnterBitAndExpr(c *BitAndExprContext)

	// EnterAssignExpr is called when entering the assignExpr production.
	EnterAssignExpr(c *AssignExprContext)

	// EnterSelfAddRightExpr is called when entering the selfAddRightExpr production.
	EnterSelfAddRightExpr(c *SelfAddRightExprContext)

	// EnterSubAssignExpr is called when entering the subAssignExpr production.
	EnterSubAssignExpr(c *SubAssignExprContext)

	// EnterBracketExpr is called when entering the bracketExpr production.
	EnterBracketExpr(c *BracketExprContext)

	// EnterBooleanExpr is called when entering the booleanExpr production.
	EnterBooleanExpr(c *BooleanExprContext)

	// EnterAShrExpr is called when entering the aShrExpr production.
	EnterAShrExpr(c *AShrExprContext)

	// EnterLessExpr is called when entering the lessExpr production.
	EnterLessExpr(c *LessExprContext)

	// EnterOrExpr is called when entering the orExpr production.
	EnterOrExpr(c *OrExprContext)

	// EnterSelfAddLeftExpr is called when entering the selfAddLeftExpr production.
	EnterSelfAddLeftExpr(c *SelfAddLeftExprContext)

	// EnterLShrAssignExpr is called when entering the lShrAssignExpr production.
	EnterLShrAssignExpr(c *LShrAssignExprContext)

	// EnterLessEqualExpr is called when entering the lessEqualExpr production.
	EnterLessEqualExpr(c *LessEqualExprContext)

	// EnterAShrAssignExpr is called when entering the aShrAssignExpr production.
	EnterAShrAssignExpr(c *AShrAssignExprContext)

	// EnterBitOrExpr is called when entering the bitOrExpr production.
	EnterBitOrExpr(c *BitOrExprContext)

	// EnterNotExpr is called when entering the notExpr production.
	EnterNotExpr(c *NotExprContext)

	// EnterAddExpr is called when entering the addExpr production.
	EnterAddExpr(c *AddExprContext)

	// EnterShlAssignExpr is called when entering the shlAssignExpr production.
	EnterShlAssignExpr(c *ShlAssignExprContext)

	// EnterModAssignExpr is called when entering the modAssignExpr production.
	EnterModAssignExpr(c *ModAssignExprContext)

	// EnterPositiveExpr is called when entering the positiveExpr production.
	EnterPositiveExpr(c *PositiveExprContext)

	// EnterShlExpr is called when entering the shlExpr production.
	EnterShlExpr(c *ShlExprContext)

	// EnterEqualExpr is called when entering the equalExpr production.
	EnterEqualExpr(c *EqualExprContext)

	// EnterAndExpr is called when entering the andExpr production.
	EnterAndExpr(c *AndExprContext)

	// EnterId is called when entering the id production.
	EnterId(c *IdContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterBoolean is called when entering the boolean production.
	EnterBoolean(c *BooleanContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterFloat is called when entering the float production.
	EnterFloat(c *FloatContext)

	// EnterInteger is called when entering the integer production.
	EnterInteger(c *IntegerContext)

	// EnterTemplateArg is called when entering the templateArg production.
	EnterTemplateArg(c *TemplateArgContext)

	// EnterTemplateArgs is called when entering the templateArgs production.
	EnterTemplateArgs(c *TemplateArgsContext)

	// EnterTypeName is called when entering the typeName production.
	EnterTypeName(c *TypeNameContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitNull_cnt is called when exiting the null_cnt production.
	ExitNull_cnt(c *Null_cntContext)

	// ExitStatement_ed is called when exiting the statement_ed production.
	ExitStatement_ed(c *Statement_edContext)

	// ExitStatement is called when exiting the statement production.
	ExitStatement(c *StatementContext)

	// ExitBodyExpr is called when exiting the bodyExpr production.
	ExitBodyExpr(c *BodyExprContext)

	// ExitPackStatement is called when exiting the packStatement production.
	ExitPackStatement(c *PackStatementContext)

	// ExitImportStatement is called when exiting the importStatement production.
	ExitImportStatement(c *ImportStatementContext)

	// ExitVarDefineStatement is called when exiting the varDefineStatement production.
	ExitVarDefineStatement(c *VarDefineStatementContext)

	// ExitArgsExpr is called when exiting the argsExpr production.
	ExitArgsExpr(c *ArgsExprContext)

	// ExitDefineArgs is called when exiting the defineArgs production.
	ExitDefineArgs(c *DefineArgsContext)

	// ExitFuncDefine is called when exiting the funcDefine production.
	ExitFuncDefine(c *FuncDefineContext)

	// ExitForStatement is called when exiting the forStatement production.
	ExitForStatement(c *ForStatementContext)

	// ExitWhileStatement is called when exiting the whileStatement production.
	ExitWhileStatement(c *WhileStatementContext)

	// ExitContinueStatement is called when exiting the continueStatement production.
	ExitContinueStatement(c *ContinueStatementContext)

	// ExitBreakStatement is called when exiting the breakStatement production.
	ExitBreakStatement(c *BreakStatementContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitReturnStatement is called when exiting the returnStatement production.
	ExitReturnStatement(c *ReturnStatementContext)

	// ExitClassDefine is called when exiting the classDefine production.
	ExitClassDefine(c *ClassDefineContext)

	// ExitTryExpr is called when exiting the tryExpr production.
	ExitTryExpr(c *TryExprContext)

	// ExitCatchExpr is called when exiting the catchExpr production.
	ExitCatchExpr(c *CatchExprContext)

	// ExitFuncExpr is called when exiting the funcExpr production.
	ExitFuncExpr(c *FuncExprContext)

	// ExitObjectExpr is called when exiting the objectExpr production.
	ExitObjectExpr(c *ObjectExprContext)

	// ExitBlendExpr is called when exiting the blendExpr production.
	ExitBlendExpr(c *BlendExprContext)

	// ExitExprPtrParser is called when exiting the exprPtrParser production.
	ExitExprPtrParser(c *ExprPtrParserContext)

	// ExitSelfSubRightExpr is called when exiting the selfSubRightExpr production.
	ExitSelfSubRightExpr(c *SelfSubRightExprContext)

	// ExitAndAssignExpr is called when exiting the andAssignExpr production.
	ExitAndAssignExpr(c *AndAssignExprContext)

	// ExitNullExpr is called when exiting the nullExpr production.
	ExitNullExpr(c *NullExprContext)

	// ExitModExpr is called when exiting the modExpr production.
	ExitModExpr(c *ModExprContext)

	// ExitCastExpr is called when exiting the castExpr production.
	ExitCastExpr(c *CastExprContext)

	// ExitLShrExpr is called when exiting the lShrExpr production.
	ExitLShrExpr(c *LShrExprContext)

	// ExitNegativeExpr is called when exiting the negativeExpr production.
	ExitNegativeExpr(c *NegativeExprContext)

	// ExitNumberExpr is called when exiting the numberExpr production.
	ExitNumberExpr(c *NumberExprContext)

	// ExitMulAssignExpr is called when exiting the mulAssignExpr production.
	ExitMulAssignExpr(c *MulAssignExprContext)

	// ExitXorAssignExpr is called when exiting the xorAssignExpr production.
	ExitXorAssignExpr(c *XorAssignExprContext)

	// ExitSubExpr is called when exiting the subExpr production.
	ExitSubExpr(c *SubExprContext)

	// ExitGreaterEqualExpr is called when exiting the greaterEqualExpr production.
	ExitGreaterEqualExpr(c *GreaterEqualExprContext)

	// ExitAddAssignExpr is called when exiting the addAssignExpr production.
	ExitAddAssignExpr(c *AddAssignExprContext)

	// ExitOrAssignExpr is called when exiting the orAssignExpr production.
	ExitOrAssignExpr(c *OrAssignExprContext)

	// ExitBitXorExpr is called when exiting the bitXorExpr production.
	ExitBitXorExpr(c *BitXorExprContext)

	// ExitPtrExpr is called when exiting the ptrExpr production.
	ExitPtrExpr(c *PtrExprContext)

	// ExitStringExpr is called when exiting the stringExpr production.
	ExitStringExpr(c *StringExprContext)

	// ExitGreaterExpr is called when exiting the greaterExpr production.
	ExitGreaterExpr(c *GreaterExprContext)

	// ExitDivAssignExpr is called when exiting the divAssignExpr production.
	ExitDivAssignExpr(c *DivAssignExprContext)

	// ExitMulExpr is called when exiting the mulExpr production.
	ExitMulExpr(c *MulExprContext)

	// ExitNotEqualExpr is called when exiting the notEqualExpr production.
	ExitNotEqualExpr(c *NotEqualExprContext)

	// ExitSelfSubLeftExpr is called when exiting the selfSubLeftExpr production.
	ExitSelfSubLeftExpr(c *SelfSubLeftExprContext)

	// ExitDivExpr is called when exiting the divExpr production.
	ExitDivExpr(c *DivExprContext)

	// ExitBitAndExpr is called when exiting the bitAndExpr production.
	ExitBitAndExpr(c *BitAndExprContext)

	// ExitAssignExpr is called when exiting the assignExpr production.
	ExitAssignExpr(c *AssignExprContext)

	// ExitSelfAddRightExpr is called when exiting the selfAddRightExpr production.
	ExitSelfAddRightExpr(c *SelfAddRightExprContext)

	// ExitSubAssignExpr is called when exiting the subAssignExpr production.
	ExitSubAssignExpr(c *SubAssignExprContext)

	// ExitBracketExpr is called when exiting the bracketExpr production.
	ExitBracketExpr(c *BracketExprContext)

	// ExitBooleanExpr is called when exiting the booleanExpr production.
	ExitBooleanExpr(c *BooleanExprContext)

	// ExitAShrExpr is called when exiting the aShrExpr production.
	ExitAShrExpr(c *AShrExprContext)

	// ExitLessExpr is called when exiting the lessExpr production.
	ExitLessExpr(c *LessExprContext)

	// ExitOrExpr is called when exiting the orExpr production.
	ExitOrExpr(c *OrExprContext)

	// ExitSelfAddLeftExpr is called when exiting the selfAddLeftExpr production.
	ExitSelfAddLeftExpr(c *SelfAddLeftExprContext)

	// ExitLShrAssignExpr is called when exiting the lShrAssignExpr production.
	ExitLShrAssignExpr(c *LShrAssignExprContext)

	// ExitLessEqualExpr is called when exiting the lessEqualExpr production.
	ExitLessEqualExpr(c *LessEqualExprContext)

	// ExitAShrAssignExpr is called when exiting the aShrAssignExpr production.
	ExitAShrAssignExpr(c *AShrAssignExprContext)

	// ExitBitOrExpr is called when exiting the bitOrExpr production.
	ExitBitOrExpr(c *BitOrExprContext)

	// ExitNotExpr is called when exiting the notExpr production.
	ExitNotExpr(c *NotExprContext)

	// ExitAddExpr is called when exiting the addExpr production.
	ExitAddExpr(c *AddExprContext)

	// ExitShlAssignExpr is called when exiting the shlAssignExpr production.
	ExitShlAssignExpr(c *ShlAssignExprContext)

	// ExitModAssignExpr is called when exiting the modAssignExpr production.
	ExitModAssignExpr(c *ModAssignExprContext)

	// ExitPositiveExpr is called when exiting the positiveExpr production.
	ExitPositiveExpr(c *PositiveExprContext)

	// ExitShlExpr is called when exiting the shlExpr production.
	ExitShlExpr(c *ShlExprContext)

	// ExitEqualExpr is called when exiting the equalExpr production.
	ExitEqualExpr(c *EqualExprContext)

	// ExitAndExpr is called when exiting the andExpr production.
	ExitAndExpr(c *AndExprContext)

	// ExitId is called when exiting the id production.
	ExitId(c *IdContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitBoolean is called when exiting the boolean production.
	ExitBoolean(c *BooleanContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitFloat is called when exiting the float production.
	ExitFloat(c *FloatContext)

	// ExitInteger is called when exiting the integer production.
	ExitInteger(c *IntegerContext)

	// ExitTemplateArg is called when exiting the templateArg production.
	ExitTemplateArg(c *TemplateArgContext)

	// ExitTemplateArgs is called when exiting the templateArgs production.
	ExitTemplateArgs(c *TemplateArgsContext)

	// ExitTypeName is called when exiting the typeName production.
	ExitTypeName(c *TypeNameContext)
}
