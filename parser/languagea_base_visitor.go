// Code generated from LanguageA.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LanguageA

import "github.com/antlr/antlr4/runtime/Go/antlr"

type BaseLanguageAVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseLanguageAVisitor) VisitTranslation(ctx *TranslationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageAVisitor) VisitTopLevelStatement(ctx *TopLevelStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageAVisitor) VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageAVisitor) VisitExpressionBinary(ctx *ExpressionBinaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageAVisitor) VisitExpressionConstant(ctx *ExpressionConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseLanguageAVisitor) VisitBinaryOperator(ctx *BinaryOperatorContext) interface{} {
	return v.VisitChildren(ctx)
}
