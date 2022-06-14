// Code generated from LanguageA.g4 by ANTLR 4.10.1. DO NOT EDIT.

package parser // LanguageA

import "github.com/antlr/antlr4/runtime/Go/antlr"

// A complete Visitor for a parse tree produced by LanguageAParser.
type LanguageAVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by LanguageAParser#translation.
	VisitTranslation(ctx *TranslationContext) interface{}

	// Visit a parse tree produced by LanguageAParser#topLevelStatement.
	VisitTopLevelStatement(ctx *TopLevelStatementContext) interface{}

	// Visit a parse tree produced by LanguageAParser#variableDeclaration.
	VisitVariableDeclaration(ctx *VariableDeclarationContext) interface{}

	// Visit a parse tree produced by LanguageAParser#ExpressionBinary.
	VisitExpressionBinary(ctx *ExpressionBinaryContext) interface{}

	// Visit a parse tree produced by LanguageAParser#ExpressionConstant.
	VisitExpressionConstant(ctx *ExpressionConstantContext) interface{}

	// Visit a parse tree produced by LanguageAParser#binaryOperator.
	VisitBinaryOperator(ctx *BinaryOperatorContext) interface{}
}
