// Code generated from E:/Riddle-Go/RiddleParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // RiddleParser

import "github.com/antlr4-go/antlr/v4"

// BaseRiddleParserListener is a complete listener for a parse tree produced by RiddleParser.
type BaseRiddleParserListener struct{}

var _ RiddleParserListener = &BaseRiddleParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRiddleParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRiddleParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRiddleParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRiddleParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterProgram is called when production program is entered.
func (s *BaseRiddleParserListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseRiddleParserListener) ExitProgram(ctx *ProgramContext) {}

// EnterNull_cnt is called when production null_cnt is entered.
func (s *BaseRiddleParserListener) EnterNull_cnt(ctx *Null_cntContext) {}

// ExitNull_cnt is called when production null_cnt is exited.
func (s *BaseRiddleParserListener) ExitNull_cnt(ctx *Null_cntContext) {}

// EnterStatement_ed is called when production statement_ed is entered.
func (s *BaseRiddleParserListener) EnterStatement_ed(ctx *Statement_edContext) {}

// ExitStatement_ed is called when production statement_ed is exited.
func (s *BaseRiddleParserListener) ExitStatement_ed(ctx *Statement_edContext) {}

// EnterStatement is called when production statement is entered.
func (s *BaseRiddleParserListener) EnterStatement(ctx *StatementContext) {}

// ExitStatement is called when production statement is exited.
func (s *BaseRiddleParserListener) ExitStatement(ctx *StatementContext) {}

// EnterBodyExpr is called when production bodyExpr is entered.
func (s *BaseRiddleParserListener) EnterBodyExpr(ctx *BodyExprContext) {}

// ExitBodyExpr is called when production bodyExpr is exited.
func (s *BaseRiddleParserListener) ExitBodyExpr(ctx *BodyExprContext) {}

// EnterPackStatement is called when production packStatement is entered.
func (s *BaseRiddleParserListener) EnterPackStatement(ctx *PackStatementContext) {}

// ExitPackStatement is called when production packStatement is exited.
func (s *BaseRiddleParserListener) ExitPackStatement(ctx *PackStatementContext) {}

// EnterImportStatement is called when production importStatement is entered.
func (s *BaseRiddleParserListener) EnterImportStatement(ctx *ImportStatementContext) {}

// ExitImportStatement is called when production importStatement is exited.
func (s *BaseRiddleParserListener) ExitImportStatement(ctx *ImportStatementContext) {}

// EnterVarDefineStatement is called when production varDefineStatement is entered.
func (s *BaseRiddleParserListener) EnterVarDefineStatement(ctx *VarDefineStatementContext) {}

// ExitVarDefineStatement is called when production varDefineStatement is exited.
func (s *BaseRiddleParserListener) ExitVarDefineStatement(ctx *VarDefineStatementContext) {}

// EnterArgsExpr is called when production argsExpr is entered.
func (s *BaseRiddleParserListener) EnterArgsExpr(ctx *ArgsExprContext) {}

// ExitArgsExpr is called when production argsExpr is exited.
func (s *BaseRiddleParserListener) ExitArgsExpr(ctx *ArgsExprContext) {}

// EnterDefineArgs is called when production defineArgs is entered.
func (s *BaseRiddleParserListener) EnterDefineArgs(ctx *DefineArgsContext) {}

// ExitDefineArgs is called when production defineArgs is exited.
func (s *BaseRiddleParserListener) ExitDefineArgs(ctx *DefineArgsContext) {}

// EnterFuncDefine is called when production funcDefine is entered.
func (s *BaseRiddleParserListener) EnterFuncDefine(ctx *FuncDefineContext) {}

// ExitFuncDefine is called when production funcDefine is exited.
func (s *BaseRiddleParserListener) ExitFuncDefine(ctx *FuncDefineContext) {}

// EnterForStatement is called when production forStatement is entered.
func (s *BaseRiddleParserListener) EnterForStatement(ctx *ForStatementContext) {}

// ExitForStatement is called when production forStatement is exited.
func (s *BaseRiddleParserListener) ExitForStatement(ctx *ForStatementContext) {}

// EnterWhileStatement is called when production whileStatement is entered.
func (s *BaseRiddleParserListener) EnterWhileStatement(ctx *WhileStatementContext) {}

// ExitWhileStatement is called when production whileStatement is exited.
func (s *BaseRiddleParserListener) ExitWhileStatement(ctx *WhileStatementContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BaseRiddleParserListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BaseRiddleParserListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterReturnStatement is called when production returnStatement is entered.
func (s *BaseRiddleParserListener) EnterReturnStatement(ctx *ReturnStatementContext) {}

// ExitReturnStatement is called when production returnStatement is exited.
func (s *BaseRiddleParserListener) ExitReturnStatement(ctx *ReturnStatementContext) {}

// EnterClassDefine is called when production classDefine is entered.
func (s *BaseRiddleParserListener) EnterClassDefine(ctx *ClassDefineContext) {}

// ExitClassDefine is called when production classDefine is exited.
func (s *BaseRiddleParserListener) ExitClassDefine(ctx *ClassDefineContext) {}

// EnterTryExpr is called when production tryExpr is entered.
func (s *BaseRiddleParserListener) EnterTryExpr(ctx *TryExprContext) {}

// ExitTryExpr is called when production tryExpr is exited.
func (s *BaseRiddleParserListener) ExitTryExpr(ctx *TryExprContext) {}

// EnterCatchExpr is called when production catchExpr is entered.
func (s *BaseRiddleParserListener) EnterCatchExpr(ctx *CatchExprContext) {}

// ExitCatchExpr is called when production catchExpr is exited.
func (s *BaseRiddleParserListener) ExitCatchExpr(ctx *CatchExprContext) {}

// EnterFuncExpr is called when production funcExpr is entered.
func (s *BaseRiddleParserListener) EnterFuncExpr(ctx *FuncExprContext) {}

// ExitFuncExpr is called when production funcExpr is exited.
func (s *BaseRiddleParserListener) ExitFuncExpr(ctx *FuncExprContext) {}

// EnterObjectExpr is called when production objectExpr is entered.
func (s *BaseRiddleParserListener) EnterObjectExpr(ctx *ObjectExprContext) {}

// ExitObjectExpr is called when production objectExpr is exited.
func (s *BaseRiddleParserListener) ExitObjectExpr(ctx *ObjectExprContext) {}

// EnterBlendExpr is called when production blendExpr is entered.
func (s *BaseRiddleParserListener) EnterBlendExpr(ctx *BlendExprContext) {}

// ExitBlendExpr is called when production blendExpr is exited.
func (s *BaseRiddleParserListener) ExitBlendExpr(ctx *BlendExprContext) {}

// EnterExprPtrParser is called when production exprPtrParser is entered.
func (s *BaseRiddleParserListener) EnterExprPtrParser(ctx *ExprPtrParserContext) {}

// ExitExprPtrParser is called when production exprPtrParser is exited.
func (s *BaseRiddleParserListener) ExitExprPtrParser(ctx *ExprPtrParserContext) {}

// EnterSelfSubRightExpr is called when production selfSubRightExpr is entered.
func (s *BaseRiddleParserListener) EnterSelfSubRightExpr(ctx *SelfSubRightExprContext) {}

// ExitSelfSubRightExpr is called when production selfSubRightExpr is exited.
func (s *BaseRiddleParserListener) ExitSelfSubRightExpr(ctx *SelfSubRightExprContext) {}

// EnterAndAssignExpr is called when production andAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterAndAssignExpr(ctx *AndAssignExprContext) {}

// ExitAndAssignExpr is called when production andAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitAndAssignExpr(ctx *AndAssignExprContext) {}

// EnterNullExpr is called when production nullExpr is entered.
func (s *BaseRiddleParserListener) EnterNullExpr(ctx *NullExprContext) {}

// ExitNullExpr is called when production nullExpr is exited.
func (s *BaseRiddleParserListener) ExitNullExpr(ctx *NullExprContext) {}

// EnterModExpr is called when production modExpr is entered.
func (s *BaseRiddleParserListener) EnterModExpr(ctx *ModExprContext) {}

// ExitModExpr is called when production modExpr is exited.
func (s *BaseRiddleParserListener) ExitModExpr(ctx *ModExprContext) {}

// EnterCastExpr is called when production castExpr is entered.
func (s *BaseRiddleParserListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production castExpr is exited.
func (s *BaseRiddleParserListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterLShrExpr is called when production lShrExpr is entered.
func (s *BaseRiddleParserListener) EnterLShrExpr(ctx *LShrExprContext) {}

// ExitLShrExpr is called when production lShrExpr is exited.
func (s *BaseRiddleParserListener) ExitLShrExpr(ctx *LShrExprContext) {}

// EnterNegativeExpr is called when production negativeExpr is entered.
func (s *BaseRiddleParserListener) EnterNegativeExpr(ctx *NegativeExprContext) {}

// ExitNegativeExpr is called when production negativeExpr is exited.
func (s *BaseRiddleParserListener) ExitNegativeExpr(ctx *NegativeExprContext) {}

// EnterNumberExpr is called when production numberExpr is entered.
func (s *BaseRiddleParserListener) EnterNumberExpr(ctx *NumberExprContext) {}

// ExitNumberExpr is called when production numberExpr is exited.
func (s *BaseRiddleParserListener) ExitNumberExpr(ctx *NumberExprContext) {}

// EnterMulAssignExpr is called when production mulAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterMulAssignExpr(ctx *MulAssignExprContext) {}

// ExitMulAssignExpr is called when production mulAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitMulAssignExpr(ctx *MulAssignExprContext) {}

// EnterXorAssignExpr is called when production xorAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterXorAssignExpr(ctx *XorAssignExprContext) {}

// ExitXorAssignExpr is called when production xorAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitXorAssignExpr(ctx *XorAssignExprContext) {}

// EnterSubExpr is called when production subExpr is entered.
func (s *BaseRiddleParserListener) EnterSubExpr(ctx *SubExprContext) {}

// ExitSubExpr is called when production subExpr is exited.
func (s *BaseRiddleParserListener) ExitSubExpr(ctx *SubExprContext) {}

// EnterGreaterEqualExpr is called when production greaterEqualExpr is entered.
func (s *BaseRiddleParserListener) EnterGreaterEqualExpr(ctx *GreaterEqualExprContext) {}

// ExitGreaterEqualExpr is called when production greaterEqualExpr is exited.
func (s *BaseRiddleParserListener) ExitGreaterEqualExpr(ctx *GreaterEqualExprContext) {}

// EnterAddAssignExpr is called when production addAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterAddAssignExpr(ctx *AddAssignExprContext) {}

// ExitAddAssignExpr is called when production addAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitAddAssignExpr(ctx *AddAssignExprContext) {}

// EnterOrAssignExpr is called when production orAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterOrAssignExpr(ctx *OrAssignExprContext) {}

// ExitOrAssignExpr is called when production orAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitOrAssignExpr(ctx *OrAssignExprContext) {}

// EnterBitXorExpr is called when production bitXorExpr is entered.
func (s *BaseRiddleParserListener) EnterBitXorExpr(ctx *BitXorExprContext) {}

// ExitBitXorExpr is called when production bitXorExpr is exited.
func (s *BaseRiddleParserListener) ExitBitXorExpr(ctx *BitXorExprContext) {}

// EnterPtrExpr is called when production ptrExpr is entered.
func (s *BaseRiddleParserListener) EnterPtrExpr(ctx *PtrExprContext) {}

// ExitPtrExpr is called when production ptrExpr is exited.
func (s *BaseRiddleParserListener) ExitPtrExpr(ctx *PtrExprContext) {}

// EnterStringExpr is called when production stringExpr is entered.
func (s *BaseRiddleParserListener) EnterStringExpr(ctx *StringExprContext) {}

// ExitStringExpr is called when production stringExpr is exited.
func (s *BaseRiddleParserListener) ExitStringExpr(ctx *StringExprContext) {}

// EnterGreaterExpr is called when production greaterExpr is entered.
func (s *BaseRiddleParserListener) EnterGreaterExpr(ctx *GreaterExprContext) {}

// ExitGreaterExpr is called when production greaterExpr is exited.
func (s *BaseRiddleParserListener) ExitGreaterExpr(ctx *GreaterExprContext) {}

// EnterDivAssignExpr is called when production divAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterDivAssignExpr(ctx *DivAssignExprContext) {}

// ExitDivAssignExpr is called when production divAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitDivAssignExpr(ctx *DivAssignExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseRiddleParserListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseRiddleParserListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterNotEqualExpr is called when production notEqualExpr is entered.
func (s *BaseRiddleParserListener) EnterNotEqualExpr(ctx *NotEqualExprContext) {}

// ExitNotEqualExpr is called when production notEqualExpr is exited.
func (s *BaseRiddleParserListener) ExitNotEqualExpr(ctx *NotEqualExprContext) {}

// EnterSelfSubLeftExpr is called when production selfSubLeftExpr is entered.
func (s *BaseRiddleParserListener) EnterSelfSubLeftExpr(ctx *SelfSubLeftExprContext) {}

// ExitSelfSubLeftExpr is called when production selfSubLeftExpr is exited.
func (s *BaseRiddleParserListener) ExitSelfSubLeftExpr(ctx *SelfSubLeftExprContext) {}

// EnterDivExpr is called when production divExpr is entered.
func (s *BaseRiddleParserListener) EnterDivExpr(ctx *DivExprContext) {}

// ExitDivExpr is called when production divExpr is exited.
func (s *BaseRiddleParserListener) ExitDivExpr(ctx *DivExprContext) {}

// EnterBitAndExpr is called when production bitAndExpr is entered.
func (s *BaseRiddleParserListener) EnterBitAndExpr(ctx *BitAndExprContext) {}

// ExitBitAndExpr is called when production bitAndExpr is exited.
func (s *BaseRiddleParserListener) ExitBitAndExpr(ctx *BitAndExprContext) {}

// EnterAssignExpr is called when production assignExpr is entered.
func (s *BaseRiddleParserListener) EnterAssignExpr(ctx *AssignExprContext) {}

// ExitAssignExpr is called when production assignExpr is exited.
func (s *BaseRiddleParserListener) ExitAssignExpr(ctx *AssignExprContext) {}

// EnterSelfAddRightExpr is called when production selfAddRightExpr is entered.
func (s *BaseRiddleParserListener) EnterSelfAddRightExpr(ctx *SelfAddRightExprContext) {}

// ExitSelfAddRightExpr is called when production selfAddRightExpr is exited.
func (s *BaseRiddleParserListener) ExitSelfAddRightExpr(ctx *SelfAddRightExprContext) {}

// EnterSubAssignExpr is called when production subAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterSubAssignExpr(ctx *SubAssignExprContext) {}

// ExitSubAssignExpr is called when production subAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitSubAssignExpr(ctx *SubAssignExprContext) {}

// EnterBracketExpr is called when production bracketExpr is entered.
func (s *BaseRiddleParserListener) EnterBracketExpr(ctx *BracketExprContext) {}

// ExitBracketExpr is called when production bracketExpr is exited.
func (s *BaseRiddleParserListener) ExitBracketExpr(ctx *BracketExprContext) {}

// EnterBooleanExpr is called when production booleanExpr is entered.
func (s *BaseRiddleParserListener) EnterBooleanExpr(ctx *BooleanExprContext) {}

// ExitBooleanExpr is called when production booleanExpr is exited.
func (s *BaseRiddleParserListener) ExitBooleanExpr(ctx *BooleanExprContext) {}

// EnterAShrExpr is called when production aShrExpr is entered.
func (s *BaseRiddleParserListener) EnterAShrExpr(ctx *AShrExprContext) {}

// ExitAShrExpr is called when production aShrExpr is exited.
func (s *BaseRiddleParserListener) ExitAShrExpr(ctx *AShrExprContext) {}

// EnterLessExpr is called when production lessExpr is entered.
func (s *BaseRiddleParserListener) EnterLessExpr(ctx *LessExprContext) {}

// ExitLessExpr is called when production lessExpr is exited.
func (s *BaseRiddleParserListener) ExitLessExpr(ctx *LessExprContext) {}

// EnterOrExpr is called when production orExpr is entered.
func (s *BaseRiddleParserListener) EnterOrExpr(ctx *OrExprContext) {}

// ExitOrExpr is called when production orExpr is exited.
func (s *BaseRiddleParserListener) ExitOrExpr(ctx *OrExprContext) {}

// EnterSelfAddLeftExpr is called when production selfAddLeftExpr is entered.
func (s *BaseRiddleParserListener) EnterSelfAddLeftExpr(ctx *SelfAddLeftExprContext) {}

// ExitSelfAddLeftExpr is called when production selfAddLeftExpr is exited.
func (s *BaseRiddleParserListener) ExitSelfAddLeftExpr(ctx *SelfAddLeftExprContext) {}

// EnterLShrAssignExpr is called when production lShrAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterLShrAssignExpr(ctx *LShrAssignExprContext) {}

// ExitLShrAssignExpr is called when production lShrAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitLShrAssignExpr(ctx *LShrAssignExprContext) {}

// EnterLessEqualExpr is called when production lessEqualExpr is entered.
func (s *BaseRiddleParserListener) EnterLessEqualExpr(ctx *LessEqualExprContext) {}

// ExitLessEqualExpr is called when production lessEqualExpr is exited.
func (s *BaseRiddleParserListener) ExitLessEqualExpr(ctx *LessEqualExprContext) {}

// EnterAShrAssignExpr is called when production aShrAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterAShrAssignExpr(ctx *AShrAssignExprContext) {}

// ExitAShrAssignExpr is called when production aShrAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitAShrAssignExpr(ctx *AShrAssignExprContext) {}

// EnterBitOrExpr is called when production bitOrExpr is entered.
func (s *BaseRiddleParserListener) EnterBitOrExpr(ctx *BitOrExprContext) {}

// ExitBitOrExpr is called when production bitOrExpr is exited.
func (s *BaseRiddleParserListener) ExitBitOrExpr(ctx *BitOrExprContext) {}

// EnterNotExpr is called when production notExpr is entered.
func (s *BaseRiddleParserListener) EnterNotExpr(ctx *NotExprContext) {}

// ExitNotExpr is called when production notExpr is exited.
func (s *BaseRiddleParserListener) ExitNotExpr(ctx *NotExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseRiddleParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseRiddleParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterShlAssignExpr is called when production shlAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterShlAssignExpr(ctx *ShlAssignExprContext) {}

// ExitShlAssignExpr is called when production shlAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitShlAssignExpr(ctx *ShlAssignExprContext) {}

// EnterModAssignExpr is called when production modAssignExpr is entered.
func (s *BaseRiddleParserListener) EnterModAssignExpr(ctx *ModAssignExprContext) {}

// ExitModAssignExpr is called when production modAssignExpr is exited.
func (s *BaseRiddleParserListener) ExitModAssignExpr(ctx *ModAssignExprContext) {}

// EnterPositiveExpr is called when production positiveExpr is entered.
func (s *BaseRiddleParserListener) EnterPositiveExpr(ctx *PositiveExprContext) {}

// ExitPositiveExpr is called when production positiveExpr is exited.
func (s *BaseRiddleParserListener) ExitPositiveExpr(ctx *PositiveExprContext) {}

// EnterShlExpr is called when production shlExpr is entered.
func (s *BaseRiddleParserListener) EnterShlExpr(ctx *ShlExprContext) {}

// ExitShlExpr is called when production shlExpr is exited.
func (s *BaseRiddleParserListener) ExitShlExpr(ctx *ShlExprContext) {}

// EnterEqualExpr is called when production equalExpr is entered.
func (s *BaseRiddleParserListener) EnterEqualExpr(ctx *EqualExprContext) {}

// ExitEqualExpr is called when production equalExpr is exited.
func (s *BaseRiddleParserListener) ExitEqualExpr(ctx *EqualExprContext) {}

// EnterAndExpr is called when production andExpr is entered.
func (s *BaseRiddleParserListener) EnterAndExpr(ctx *AndExprContext) {}

// ExitAndExpr is called when production andExpr is exited.
func (s *BaseRiddleParserListener) ExitAndExpr(ctx *AndExprContext) {}

// EnterId is called when production id is entered.
func (s *BaseRiddleParserListener) EnterId(ctx *IdContext) {}

// ExitId is called when production id is exited.
func (s *BaseRiddleParserListener) ExitId(ctx *IdContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseRiddleParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseRiddleParserListener) ExitNumber(ctx *NumberContext) {}

// EnterBoolean is called when production boolean is entered.
func (s *BaseRiddleParserListener) EnterBoolean(ctx *BooleanContext) {}

// ExitBoolean is called when production boolean is exited.
func (s *BaseRiddleParserListener) ExitBoolean(ctx *BooleanContext) {}

// EnterString is called when production string is entered.
func (s *BaseRiddleParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseRiddleParserListener) ExitString(ctx *StringContext) {}

// EnterFloat is called when production float is entered.
func (s *BaseRiddleParserListener) EnterFloat(ctx *FloatContext) {}

// ExitFloat is called when production float is exited.
func (s *BaseRiddleParserListener) ExitFloat(ctx *FloatContext) {}

// EnterInteger is called when production integer is entered.
func (s *BaseRiddleParserListener) EnterInteger(ctx *IntegerContext) {}

// ExitInteger is called when production integer is exited.
func (s *BaseRiddleParserListener) ExitInteger(ctx *IntegerContext) {}

// EnterTemplateArg is called when production templateArg is entered.
func (s *BaseRiddleParserListener) EnterTemplateArg(ctx *TemplateArgContext) {}

// ExitTemplateArg is called when production templateArg is exited.
func (s *BaseRiddleParserListener) ExitTemplateArg(ctx *TemplateArgContext) {}

// EnterTemplateArgs is called when production templateArgs is entered.
func (s *BaseRiddleParserListener) EnterTemplateArgs(ctx *TemplateArgsContext) {}

// ExitTemplateArgs is called when production templateArgs is exited.
func (s *BaseRiddleParserListener) ExitTemplateArgs(ctx *TemplateArgsContext) {}

// EnterTypeName is called when production typeName is entered.
func (s *BaseRiddleParserListener) EnterTypeName(ctx *TypeNameContext) {}

// ExitTypeName is called when production typeName is exited.
func (s *BaseRiddleParserListener) ExitTypeName(ctx *TypeNameContext) {}
