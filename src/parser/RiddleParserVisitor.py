# Generated from C:/Users/wangz/PycharmProjects/Riddle-Language/RiddleParser.g4 by ANTLR 4.13.2
from antlr4 import *
if "." in __name__:
    from .RiddleParser import RiddleParser
else:
    from RiddleParser import RiddleParser

# This class defines a complete generic visitor for a parse tree produced by RiddleParser.

class RiddleParserVisitor(ParseTreeVisitor):

    # Visit a parse tree produced by RiddleParser#program.
    def visitProgram(self, ctx:RiddleParser.ProgramContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#null_cnt.
    def visitNull_cnt(self, ctx:RiddleParser.Null_cntContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#statement_ed.
    def visitStatement_ed(self, ctx:RiddleParser.Statement_edContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#statement.
    def visitStatement(self, ctx:RiddleParser.StatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#bodyExpr.
    def visitBodyExpr(self, ctx:RiddleParser.BodyExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#packStatement.
    def visitPackStatement(self, ctx:RiddleParser.PackStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#importStatement.
    def visitImportStatement(self, ctx:RiddleParser.ImportStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#varDefineStatement.
    def visitVarDefineStatement(self, ctx:RiddleParser.VarDefineStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#argsExpr.
    def visitArgsExpr(self, ctx:RiddleParser.ArgsExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#defineArgs.
    def visitDefineArgs(self, ctx:RiddleParser.DefineArgsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#funcDefine.
    def visitFuncDefine(self, ctx:RiddleParser.FuncDefineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#forStatement.
    def visitForStatement(self, ctx:RiddleParser.ForStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#whileStatement.
    def visitWhileStatement(self, ctx:RiddleParser.WhileStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#continueStatement.
    def visitContinueStatement(self, ctx:RiddleParser.ContinueStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#breakStatement.
    def visitBreakStatement(self, ctx:RiddleParser.BreakStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#ifStatement.
    def visitIfStatement(self, ctx:RiddleParser.IfStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#returnStatement.
    def visitReturnStatement(self, ctx:RiddleParser.ReturnStatementContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#classDefine.
    def visitClassDefine(self, ctx:RiddleParser.ClassDefineContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#tryExpr.
    def visitTryExpr(self, ctx:RiddleParser.TryExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#catchExpr.
    def visitCatchExpr(self, ctx:RiddleParser.CatchExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#funcExpr.
    def visitFuncExpr(self, ctx:RiddleParser.FuncExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#objectExpr.
    def visitObjectExpr(self, ctx:RiddleParser.ObjectExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#blendExpr.
    def visitBlendExpr(self, ctx:RiddleParser.BlendExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#exprPtrParser.
    def visitExprPtrParser(self, ctx:RiddleParser.ExprPtrParserContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#selfSubRightExpr.
    def visitSelfSubRightExpr(self, ctx:RiddleParser.SelfSubRightExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#andAssignExpr.
    def visitAndAssignExpr(self, ctx:RiddleParser.AndAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#nullExpr.
    def visitNullExpr(self, ctx:RiddleParser.NullExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#modExpr.
    def visitModExpr(self, ctx:RiddleParser.ModExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#castExpr.
    def visitCastExpr(self, ctx:RiddleParser.CastExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#lShrExpr.
    def visitLShrExpr(self, ctx:RiddleParser.LShrExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#negativeExpr.
    def visitNegativeExpr(self, ctx:RiddleParser.NegativeExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#numberExpr.
    def visitNumberExpr(self, ctx:RiddleParser.NumberExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#mulAssignExpr.
    def visitMulAssignExpr(self, ctx:RiddleParser.MulAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#xorAssignExpr.
    def visitXorAssignExpr(self, ctx:RiddleParser.XorAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#subExpr.
    def visitSubExpr(self, ctx:RiddleParser.SubExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#greaterEqualExpr.
    def visitGreaterEqualExpr(self, ctx:RiddleParser.GreaterEqualExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#addAssignExpr.
    def visitAddAssignExpr(self, ctx:RiddleParser.AddAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#orAssignExpr.
    def visitOrAssignExpr(self, ctx:RiddleParser.OrAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#bitXorExpr.
    def visitBitXorExpr(self, ctx:RiddleParser.BitXorExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#ptrExpr.
    def visitPtrExpr(self, ctx:RiddleParser.PtrExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#stringExpr.
    def visitStringExpr(self, ctx:RiddleParser.StringExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#greaterExpr.
    def visitGreaterExpr(self, ctx:RiddleParser.GreaterExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#divAssignExpr.
    def visitDivAssignExpr(self, ctx:RiddleParser.DivAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#mulExpr.
    def visitMulExpr(self, ctx:RiddleParser.MulExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#notEqualExpr.
    def visitNotEqualExpr(self, ctx:RiddleParser.NotEqualExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#selfSubLeftExpr.
    def visitSelfSubLeftExpr(self, ctx:RiddleParser.SelfSubLeftExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#divExpr.
    def visitDivExpr(self, ctx:RiddleParser.DivExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#bitAndExpr.
    def visitBitAndExpr(self, ctx:RiddleParser.BitAndExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#assignExpr.
    def visitAssignExpr(self, ctx:RiddleParser.AssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#selfAddRightExpr.
    def visitSelfAddRightExpr(self, ctx:RiddleParser.SelfAddRightExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#subAssignExpr.
    def visitSubAssignExpr(self, ctx:RiddleParser.SubAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#bracketExpr.
    def visitBracketExpr(self, ctx:RiddleParser.BracketExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#booleanExpr.
    def visitBooleanExpr(self, ctx:RiddleParser.BooleanExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#aShrExpr.
    def visitAShrExpr(self, ctx:RiddleParser.AShrExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#lessExpr.
    def visitLessExpr(self, ctx:RiddleParser.LessExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#orExpr.
    def visitOrExpr(self, ctx:RiddleParser.OrExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#selfAddLeftExpr.
    def visitSelfAddLeftExpr(self, ctx:RiddleParser.SelfAddLeftExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#lShrAssignExpr.
    def visitLShrAssignExpr(self, ctx:RiddleParser.LShrAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#lessEqualExpr.
    def visitLessEqualExpr(self, ctx:RiddleParser.LessEqualExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#aShrAssignExpr.
    def visitAShrAssignExpr(self, ctx:RiddleParser.AShrAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#bitOrExpr.
    def visitBitOrExpr(self, ctx:RiddleParser.BitOrExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#notExpr.
    def visitNotExpr(self, ctx:RiddleParser.NotExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#addExpr.
    def visitAddExpr(self, ctx:RiddleParser.AddExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#shlAssignExpr.
    def visitShlAssignExpr(self, ctx:RiddleParser.ShlAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#modAssignExpr.
    def visitModAssignExpr(self, ctx:RiddleParser.ModAssignExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#positiveExpr.
    def visitPositiveExpr(self, ctx:RiddleParser.PositiveExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#shlExpr.
    def visitShlExpr(self, ctx:RiddleParser.ShlExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#equalExpr.
    def visitEqualExpr(self, ctx:RiddleParser.EqualExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#andExpr.
    def visitAndExpr(self, ctx:RiddleParser.AndExprContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#id.
    def visitId(self, ctx:RiddleParser.IdContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#number.
    def visitNumber(self, ctx:RiddleParser.NumberContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#boolean.
    def visitBoolean(self, ctx:RiddleParser.BooleanContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#string.
    def visitString(self, ctx:RiddleParser.StringContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#float.
    def visitFloat(self, ctx:RiddleParser.FloatContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#integer.
    def visitInteger(self, ctx:RiddleParser.IntegerContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#templateArg.
    def visitTemplateArg(self, ctx:RiddleParser.TemplateArgContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#templateArgs.
    def visitTemplateArgs(self, ctx:RiddleParser.TemplateArgsContext):
        return self.visitChildren(ctx)


    # Visit a parse tree produced by RiddleParser#typeName.
    def visitTypeName(self, ctx:RiddleParser.TypeNameContext):
        return self.visitChildren(ctx)



del RiddleParser