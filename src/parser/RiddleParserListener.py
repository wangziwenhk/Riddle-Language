# Generated from C:/Users/wangz/PycharmProjects/Riddle-Language/RiddleParser.g4 by ANTLR 4.13.1
from antlr4 import *
if "." in __name__:
    from .RiddleParser import RiddleParser
else:
    from RiddleParser import RiddleParser

# This class defines a complete listener for a parse tree produced by RiddleParser.
class RiddleParserListener(ParseTreeListener):

    # Enter a parse tree produced by RiddleParser#program.
    def enterProgram(self, ctx:RiddleParser.ProgramContext):
        pass

    # Exit a parse tree produced by RiddleParser#program.
    def exitProgram(self, ctx:RiddleParser.ProgramContext):
        pass


    # Enter a parse tree produced by RiddleParser#statement_ed.
    def enterStatement_ed(self, ctx:RiddleParser.Statement_edContext):
        pass

    # Exit a parse tree produced by RiddleParser#statement_ed.
    def exitStatement_ed(self, ctx:RiddleParser.Statement_edContext):
        pass


    # Enter a parse tree produced by RiddleParser#statement.
    def enterStatement(self, ctx:RiddleParser.StatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#statement.
    def exitStatement(self, ctx:RiddleParser.StatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#packStatement.
    def enterPackStatement(self, ctx:RiddleParser.PackStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#packStatement.
    def exitPackStatement(self, ctx:RiddleParser.PackStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#importStatement.
    def enterImportStatement(self, ctx:RiddleParser.ImportStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#importStatement.
    def exitImportStatement(self, ctx:RiddleParser.ImportStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#varDefineStatement.
    def enterVarDefineStatement(self, ctx:RiddleParser.VarDefineStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#varDefineStatement.
    def exitVarDefineStatement(self, ctx:RiddleParser.VarDefineStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#argsExpr.
    def enterArgsExpr(self, ctx:RiddleParser.ArgsExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#argsExpr.
    def exitArgsExpr(self, ctx:RiddleParser.ArgsExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#defineArgs.
    def enterDefineArgs(self, ctx:RiddleParser.DefineArgsContext):
        pass

    # Exit a parse tree produced by RiddleParser#defineArgs.
    def exitDefineArgs(self, ctx:RiddleParser.DefineArgsContext):
        pass


    # Enter a parse tree produced by RiddleParser#funcDefine.
    def enterFuncDefine(self, ctx:RiddleParser.FuncDefineContext):
        pass

    # Exit a parse tree produced by RiddleParser#funcDefine.
    def exitFuncDefine(self, ctx:RiddleParser.FuncDefineContext):
        pass


    # Enter a parse tree produced by RiddleParser#funcBody.
    def enterFuncBody(self, ctx:RiddleParser.FuncBodyContext):
        pass

    # Exit a parse tree produced by RiddleParser#funcBody.
    def exitFuncBody(self, ctx:RiddleParser.FuncBodyContext):
        pass


    # Enter a parse tree produced by RiddleParser#forStatement.
    def enterForStatement(self, ctx:RiddleParser.ForStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#forStatement.
    def exitForStatement(self, ctx:RiddleParser.ForStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#whileStatement.
    def enterWhileStatement(self, ctx:RiddleParser.WhileStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#whileStatement.
    def exitWhileStatement(self, ctx:RiddleParser.WhileStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#ifStatement.
    def enterIfStatement(self, ctx:RiddleParser.IfStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#ifStatement.
    def exitIfStatement(self, ctx:RiddleParser.IfStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#returnStatement.
    def enterReturnStatement(self, ctx:RiddleParser.ReturnStatementContext):
        pass

    # Exit a parse tree produced by RiddleParser#returnStatement.
    def exitReturnStatement(self, ctx:RiddleParser.ReturnStatementContext):
        pass


    # Enter a parse tree produced by RiddleParser#classDefine.
    def enterClassDefine(self, ctx:RiddleParser.ClassDefineContext):
        pass

    # Exit a parse tree produced by RiddleParser#classDefine.
    def exitClassDefine(self, ctx:RiddleParser.ClassDefineContext):
        pass


    # Enter a parse tree produced by RiddleParser#classBody.
    def enterClassBody(self, ctx:RiddleParser.ClassBodyContext):
        pass

    # Exit a parse tree produced by RiddleParser#classBody.
    def exitClassBody(self, ctx:RiddleParser.ClassBodyContext):
        pass


    # Enter a parse tree produced by RiddleParser#funcExpr.
    def enterFuncExpr(self, ctx:RiddleParser.FuncExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#funcExpr.
    def exitFuncExpr(self, ctx:RiddleParser.FuncExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#objectExpr.
    def enterObjectExpr(self, ctx:RiddleParser.ObjectExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#objectExpr.
    def exitObjectExpr(self, ctx:RiddleParser.ObjectExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#blendExpr.
    def enterBlendExpr(self, ctx:RiddleParser.BlendExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#blendExpr.
    def exitBlendExpr(self, ctx:RiddleParser.BlendExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#exprPtrParser.
    def enterExprPtrParser(self, ctx:RiddleParser.ExprPtrParserContext):
        pass

    # Exit a parse tree produced by RiddleParser#exprPtrParser.
    def exitExprPtrParser(self, ctx:RiddleParser.ExprPtrParserContext):
        pass


    # Enter a parse tree produced by RiddleParser#selfSubRightExpr.
    def enterSelfSubRightExpr(self, ctx:RiddleParser.SelfSubRightExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#selfSubRightExpr.
    def exitSelfSubRightExpr(self, ctx:RiddleParser.SelfSubRightExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#andAssignExpr.
    def enterAndAssignExpr(self, ctx:RiddleParser.AndAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#andAssignExpr.
    def exitAndAssignExpr(self, ctx:RiddleParser.AndAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#nullExpr.
    def enterNullExpr(self, ctx:RiddleParser.NullExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#nullExpr.
    def exitNullExpr(self, ctx:RiddleParser.NullExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#modExpr.
    def enterModExpr(self, ctx:RiddleParser.ModExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#modExpr.
    def exitModExpr(self, ctx:RiddleParser.ModExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#castExpr.
    def enterCastExpr(self, ctx:RiddleParser.CastExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#castExpr.
    def exitCastExpr(self, ctx:RiddleParser.CastExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#lShrExpr.
    def enterLShrExpr(self, ctx:RiddleParser.LShrExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#lShrExpr.
    def exitLShrExpr(self, ctx:RiddleParser.LShrExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#negativeExpr.
    def enterNegativeExpr(self, ctx:RiddleParser.NegativeExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#negativeExpr.
    def exitNegativeExpr(self, ctx:RiddleParser.NegativeExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#numberExpr.
    def enterNumberExpr(self, ctx:RiddleParser.NumberExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#numberExpr.
    def exitNumberExpr(self, ctx:RiddleParser.NumberExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#mulAssignExpr.
    def enterMulAssignExpr(self, ctx:RiddleParser.MulAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#mulAssignExpr.
    def exitMulAssignExpr(self, ctx:RiddleParser.MulAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#xorAssignExpr.
    def enterXorAssignExpr(self, ctx:RiddleParser.XorAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#xorAssignExpr.
    def exitXorAssignExpr(self, ctx:RiddleParser.XorAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#subExpr.
    def enterSubExpr(self, ctx:RiddleParser.SubExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#subExpr.
    def exitSubExpr(self, ctx:RiddleParser.SubExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#greaterEqualExpr.
    def enterGreaterEqualExpr(self, ctx:RiddleParser.GreaterEqualExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#greaterEqualExpr.
    def exitGreaterEqualExpr(self, ctx:RiddleParser.GreaterEqualExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#addAssignExpr.
    def enterAddAssignExpr(self, ctx:RiddleParser.AddAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#addAssignExpr.
    def exitAddAssignExpr(self, ctx:RiddleParser.AddAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#orAssignExpr.
    def enterOrAssignExpr(self, ctx:RiddleParser.OrAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#orAssignExpr.
    def exitOrAssignExpr(self, ctx:RiddleParser.OrAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#bitXorExpr.
    def enterBitXorExpr(self, ctx:RiddleParser.BitXorExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#bitXorExpr.
    def exitBitXorExpr(self, ctx:RiddleParser.BitXorExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#ptrExpr.
    def enterPtrExpr(self, ctx:RiddleParser.PtrExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#ptrExpr.
    def exitPtrExpr(self, ctx:RiddleParser.PtrExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#stringExpr.
    def enterStringExpr(self, ctx:RiddleParser.StringExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#stringExpr.
    def exitStringExpr(self, ctx:RiddleParser.StringExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#greaterExpr.
    def enterGreaterExpr(self, ctx:RiddleParser.GreaterExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#greaterExpr.
    def exitGreaterExpr(self, ctx:RiddleParser.GreaterExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#divAssignExpr.
    def enterDivAssignExpr(self, ctx:RiddleParser.DivAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#divAssignExpr.
    def exitDivAssignExpr(self, ctx:RiddleParser.DivAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#mulExpr.
    def enterMulExpr(self, ctx:RiddleParser.MulExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#mulExpr.
    def exitMulExpr(self, ctx:RiddleParser.MulExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#notEqualExpr.
    def enterNotEqualExpr(self, ctx:RiddleParser.NotEqualExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#notEqualExpr.
    def exitNotEqualExpr(self, ctx:RiddleParser.NotEqualExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#selfSubLeftExpr.
    def enterSelfSubLeftExpr(self, ctx:RiddleParser.SelfSubLeftExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#selfSubLeftExpr.
    def exitSelfSubLeftExpr(self, ctx:RiddleParser.SelfSubLeftExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#divExpr.
    def enterDivExpr(self, ctx:RiddleParser.DivExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#divExpr.
    def exitDivExpr(self, ctx:RiddleParser.DivExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#bitAndExpr.
    def enterBitAndExpr(self, ctx:RiddleParser.BitAndExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#bitAndExpr.
    def exitBitAndExpr(self, ctx:RiddleParser.BitAndExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#assignExpr.
    def enterAssignExpr(self, ctx:RiddleParser.AssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#assignExpr.
    def exitAssignExpr(self, ctx:RiddleParser.AssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#selfAddRightExpr.
    def enterSelfAddRightExpr(self, ctx:RiddleParser.SelfAddRightExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#selfAddRightExpr.
    def exitSelfAddRightExpr(self, ctx:RiddleParser.SelfAddRightExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#subAssignExpr.
    def enterSubAssignExpr(self, ctx:RiddleParser.SubAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#subAssignExpr.
    def exitSubAssignExpr(self, ctx:RiddleParser.SubAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#bracketExpr.
    def enterBracketExpr(self, ctx:RiddleParser.BracketExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#bracketExpr.
    def exitBracketExpr(self, ctx:RiddleParser.BracketExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#booleanExpr.
    def enterBooleanExpr(self, ctx:RiddleParser.BooleanExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#booleanExpr.
    def exitBooleanExpr(self, ctx:RiddleParser.BooleanExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#aShrExpr.
    def enterAShrExpr(self, ctx:RiddleParser.AShrExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#aShrExpr.
    def exitAShrExpr(self, ctx:RiddleParser.AShrExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#lessExpr.
    def enterLessExpr(self, ctx:RiddleParser.LessExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#lessExpr.
    def exitLessExpr(self, ctx:RiddleParser.LessExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#orExpr.
    def enterOrExpr(self, ctx:RiddleParser.OrExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#orExpr.
    def exitOrExpr(self, ctx:RiddleParser.OrExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#selfAddLeftExpr.
    def enterSelfAddLeftExpr(self, ctx:RiddleParser.SelfAddLeftExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#selfAddLeftExpr.
    def exitSelfAddLeftExpr(self, ctx:RiddleParser.SelfAddLeftExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#lShrAssignExpr.
    def enterLShrAssignExpr(self, ctx:RiddleParser.LShrAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#lShrAssignExpr.
    def exitLShrAssignExpr(self, ctx:RiddleParser.LShrAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#lessEqualExpr.
    def enterLessEqualExpr(self, ctx:RiddleParser.LessEqualExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#lessEqualExpr.
    def exitLessEqualExpr(self, ctx:RiddleParser.LessEqualExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#aShrAssignExpr.
    def enterAShrAssignExpr(self, ctx:RiddleParser.AShrAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#aShrAssignExpr.
    def exitAShrAssignExpr(self, ctx:RiddleParser.AShrAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#bitOrExpr.
    def enterBitOrExpr(self, ctx:RiddleParser.BitOrExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#bitOrExpr.
    def exitBitOrExpr(self, ctx:RiddleParser.BitOrExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#notExpr.
    def enterNotExpr(self, ctx:RiddleParser.NotExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#notExpr.
    def exitNotExpr(self, ctx:RiddleParser.NotExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#addExpr.
    def enterAddExpr(self, ctx:RiddleParser.AddExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#addExpr.
    def exitAddExpr(self, ctx:RiddleParser.AddExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#shlAssignExpr.
    def enterShlAssignExpr(self, ctx:RiddleParser.ShlAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#shlAssignExpr.
    def exitShlAssignExpr(self, ctx:RiddleParser.ShlAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#modAssignExpr.
    def enterModAssignExpr(self, ctx:RiddleParser.ModAssignExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#modAssignExpr.
    def exitModAssignExpr(self, ctx:RiddleParser.ModAssignExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#positiveExpr.
    def enterPositiveExpr(self, ctx:RiddleParser.PositiveExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#positiveExpr.
    def exitPositiveExpr(self, ctx:RiddleParser.PositiveExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#shlExpr.
    def enterShlExpr(self, ctx:RiddleParser.ShlExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#shlExpr.
    def exitShlExpr(self, ctx:RiddleParser.ShlExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#equalExpr.
    def enterEqualExpr(self, ctx:RiddleParser.EqualExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#equalExpr.
    def exitEqualExpr(self, ctx:RiddleParser.EqualExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#andExpr.
    def enterAndExpr(self, ctx:RiddleParser.AndExprContext):
        pass

    # Exit a parse tree produced by RiddleParser#andExpr.
    def exitAndExpr(self, ctx:RiddleParser.AndExprContext):
        pass


    # Enter a parse tree produced by RiddleParser#id.
    def enterId(self, ctx:RiddleParser.IdContext):
        pass

    # Exit a parse tree produced by RiddleParser#id.
    def exitId(self, ctx:RiddleParser.IdContext):
        pass


    # Enter a parse tree produced by RiddleParser#number.
    def enterNumber(self, ctx:RiddleParser.NumberContext):
        pass

    # Exit a parse tree produced by RiddleParser#number.
    def exitNumber(self, ctx:RiddleParser.NumberContext):
        pass


    # Enter a parse tree produced by RiddleParser#boolean.
    def enterBoolean(self, ctx:RiddleParser.BooleanContext):
        pass

    # Exit a parse tree produced by RiddleParser#boolean.
    def exitBoolean(self, ctx:RiddleParser.BooleanContext):
        pass


    # Enter a parse tree produced by RiddleParser#string.
    def enterString(self, ctx:RiddleParser.StringContext):
        pass

    # Exit a parse tree produced by RiddleParser#string.
    def exitString(self, ctx:RiddleParser.StringContext):
        pass


    # Enter a parse tree produced by RiddleParser#float.
    def enterFloat(self, ctx:RiddleParser.FloatContext):
        pass

    # Exit a parse tree produced by RiddleParser#float.
    def exitFloat(self, ctx:RiddleParser.FloatContext):
        pass


    # Enter a parse tree produced by RiddleParser#integer.
    def enterInteger(self, ctx:RiddleParser.IntegerContext):
        pass

    # Exit a parse tree produced by RiddleParser#integer.
    def exitInteger(self, ctx:RiddleParser.IntegerContext):
        pass


    # Enter a parse tree produced by RiddleParser#templateArg.
    def enterTemplateArg(self, ctx:RiddleParser.TemplateArgContext):
        pass

    # Exit a parse tree produced by RiddleParser#templateArg.
    def exitTemplateArg(self, ctx:RiddleParser.TemplateArgContext):
        pass


    # Enter a parse tree produced by RiddleParser#templateArgs.
    def enterTemplateArgs(self, ctx:RiddleParser.TemplateArgsContext):
        pass

    # Exit a parse tree produced by RiddleParser#templateArgs.
    def exitTemplateArgs(self, ctx:RiddleParser.TemplateArgsContext):
        pass


    # Enter a parse tree produced by RiddleParser#typeName.
    def enterTypeName(self, ctx:RiddleParser.TypeNameContext):
        pass

    # Exit a parse tree produced by RiddleParser#typeName.
    def exitTypeName(self, ctx:RiddleParser.TypeNameContext):
        pass



del RiddleParser