// Code generated from FilterExpressionSyntax.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FilterExpressionSyntax

import "github.com/antlr/antlr4/runtime/Go/antlr"

// FilterExpressionSyntaxListener is a complete listener for a parse tree produced by FilterExpressionSyntaxParser.
type FilterExpressionSyntaxListener interface {
	antlr.ParseTreeListener

	// EnterParse is called when entering the parse production.
	EnterParse(c *ParseContext)

	// EnterErr is called when entering the err production.
	EnterErr(c *ErrContext)

	// EnterFilterExpressionStatementList is called when entering the filterExpressionStatementList production.
	EnterFilterExpressionStatementList(c *FilterExpressionStatementListContext)

	// EnterFilterExpressionStatement is called when entering the filterExpressionStatement production.
	EnterFilterExpressionStatement(c *FilterExpressionStatementContext)

	// EnterIdentifierStatement is called when entering the identifierStatement production.
	EnterIdentifierStatement(c *IdentifierStatementContext)

	// EnterFilterStatement is called when entering the filterStatement production.
	EnterFilterStatement(c *FilterStatementContext)

	// EnterTopLimit is called when entering the topLimit production.
	EnterTopLimit(c *TopLimitContext)

	// EnterOrderingTerm is called when entering the orderingTerm production.
	EnterOrderingTerm(c *OrderingTermContext)

	// EnterExpressionList is called when entering the expressionList production.
	EnterExpressionList(c *ExpressionListContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterPredicateExpression is called when entering the predicateExpression production.
	EnterPredicateExpression(c *PredicateExpressionContext)

	// EnterValueExpression is called when entering the valueExpression production.
	EnterValueExpression(c *ValueExpressionContext)

	// EnterNotOperator is called when entering the notOperator production.
	EnterNotOperator(c *NotOperatorContext)

	// EnterUnaryOperator is called when entering the unaryOperator production.
	EnterUnaryOperator(c *UnaryOperatorContext)

	// EnterExactMatchModifier is called when entering the exactMatchModifier production.
	EnterExactMatchModifier(c *ExactMatchModifierContext)

	// EnterComparisonOperator is called when entering the comparisonOperator production.
	EnterComparisonOperator(c *ComparisonOperatorContext)

	// EnterLogicalOperator is called when entering the logicalOperator production.
	EnterLogicalOperator(c *LogicalOperatorContext)

	// EnterBitwiseOperator is called when entering the bitwiseOperator production.
	EnterBitwiseOperator(c *BitwiseOperatorContext)

	// EnterMathOperator is called when entering the mathOperator production.
	EnterMathOperator(c *MathOperatorContext)

	// EnterFunctionName is called when entering the functionName production.
	EnterFunctionName(c *FunctionNameContext)

	// EnterFunctionExpression is called when entering the functionExpression production.
	EnterFunctionExpression(c *FunctionExpressionContext)

	// EnterLiteralValue is called when entering the literalValue production.
	EnterLiteralValue(c *LiteralValueContext)

	// EnterTableName is called when entering the tableName production.
	EnterTableName(c *TableNameContext)

	// EnterColumnName is called when entering the columnName production.
	EnterColumnName(c *ColumnNameContext)

	// EnterOrderByColumnName is called when entering the orderByColumnName production.
	EnterOrderByColumnName(c *OrderByColumnNameContext)

	// ExitParse is called when exiting the parse production.
	ExitParse(c *ParseContext)

	// ExitErr is called when exiting the err production.
	ExitErr(c *ErrContext)

	// ExitFilterExpressionStatementList is called when exiting the filterExpressionStatementList production.
	ExitFilterExpressionStatementList(c *FilterExpressionStatementListContext)

	// ExitFilterExpressionStatement is called when exiting the filterExpressionStatement production.
	ExitFilterExpressionStatement(c *FilterExpressionStatementContext)

	// ExitIdentifierStatement is called when exiting the identifierStatement production.
	ExitIdentifierStatement(c *IdentifierStatementContext)

	// ExitFilterStatement is called when exiting the filterStatement production.
	ExitFilterStatement(c *FilterStatementContext)

	// ExitTopLimit is called when exiting the topLimit production.
	ExitTopLimit(c *TopLimitContext)

	// ExitOrderingTerm is called when exiting the orderingTerm production.
	ExitOrderingTerm(c *OrderingTermContext)

	// ExitExpressionList is called when exiting the expressionList production.
	ExitExpressionList(c *ExpressionListContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitPredicateExpression is called when exiting the predicateExpression production.
	ExitPredicateExpression(c *PredicateExpressionContext)

	// ExitValueExpression is called when exiting the valueExpression production.
	ExitValueExpression(c *ValueExpressionContext)

	// ExitNotOperator is called when exiting the notOperator production.
	ExitNotOperator(c *NotOperatorContext)

	// ExitUnaryOperator is called when exiting the unaryOperator production.
	ExitUnaryOperator(c *UnaryOperatorContext)

	// ExitExactMatchModifier is called when exiting the exactMatchModifier production.
	ExitExactMatchModifier(c *ExactMatchModifierContext)

	// ExitComparisonOperator is called when exiting the comparisonOperator production.
	ExitComparisonOperator(c *ComparisonOperatorContext)

	// ExitLogicalOperator is called when exiting the logicalOperator production.
	ExitLogicalOperator(c *LogicalOperatorContext)

	// ExitBitwiseOperator is called when exiting the bitwiseOperator production.
	ExitBitwiseOperator(c *BitwiseOperatorContext)

	// ExitMathOperator is called when exiting the mathOperator production.
	ExitMathOperator(c *MathOperatorContext)

	// ExitFunctionName is called when exiting the functionName production.
	ExitFunctionName(c *FunctionNameContext)

	// ExitFunctionExpression is called when exiting the functionExpression production.
	ExitFunctionExpression(c *FunctionExpressionContext)

	// ExitLiteralValue is called when exiting the literalValue production.
	ExitLiteralValue(c *LiteralValueContext)

	// ExitTableName is called when exiting the tableName production.
	ExitTableName(c *TableNameContext)

	// ExitColumnName is called when exiting the columnName production.
	ExitColumnName(c *ColumnNameContext)

	// ExitOrderByColumnName is called when exiting the orderByColumnName production.
	ExitOrderByColumnName(c *OrderByColumnNameContext)
}
