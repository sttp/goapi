// Code generated from FilterExpressionSyntax.g4 by ANTLR 4.9.2. DO NOT EDIT.

package parser // FilterExpressionSyntax

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseFilterExpressionSyntaxListener is a complete listener for a parse tree produced by FilterExpressionSyntaxParser.
type BaseFilterExpressionSyntaxListener struct{}

var _ FilterExpressionSyntaxListener = &BaseFilterExpressionSyntaxListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseFilterExpressionSyntaxListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseFilterExpressionSyntaxListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterParse is called when production parse is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterParse(ctx *ParseContext) {}

// ExitParse is called when production parse is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitParse(ctx *ParseContext) {}

// EnterErr is called when production err is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterErr(ctx *ErrContext) {}

// ExitErr is called when production err is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitErr(ctx *ErrContext) {}

// EnterFilterExpressionStatementList is called when production filterExpressionStatementList is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterFilterExpressionStatementList(ctx *FilterExpressionStatementListContext) {
}

// ExitFilterExpressionStatementList is called when production filterExpressionStatementList is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitFilterExpressionStatementList(ctx *FilterExpressionStatementListContext) {
}

// EnterFilterExpressionStatement is called when production filterExpressionStatement is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterFilterExpressionStatement(ctx *FilterExpressionStatementContext) {
}

// ExitFilterExpressionStatement is called when production filterExpressionStatement is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitFilterExpressionStatement(ctx *FilterExpressionStatementContext) {
}

// EnterIdentifierStatement is called when production identifierStatement is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterIdentifierStatement(ctx *IdentifierStatementContext) {
}

// ExitIdentifierStatement is called when production identifierStatement is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitIdentifierStatement(ctx *IdentifierStatementContext) {
}

// EnterFilterStatement is called when production filterStatement is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterFilterStatement(ctx *FilterStatementContext) {}

// ExitFilterStatement is called when production filterStatement is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitFilterStatement(ctx *FilterStatementContext) {}

// EnterTopLimit is called when production topLimit is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterTopLimit(ctx *TopLimitContext) {}

// ExitTopLimit is called when production topLimit is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitTopLimit(ctx *TopLimitContext) {}

// EnterOrderingTerm is called when production orderingTerm is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterOrderingTerm(ctx *OrderingTermContext) {}

// ExitOrderingTerm is called when production orderingTerm is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitOrderingTerm(ctx *OrderingTermContext) {}

// EnterExpressionList is called when production expressionList is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterExpressionList(ctx *ExpressionListContext) {}

// ExitExpressionList is called when production expressionList is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitExpressionList(ctx *ExpressionListContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitExpression(ctx *ExpressionContext) {}

// EnterPredicateExpression is called when production predicateExpression is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterPredicateExpression(ctx *PredicateExpressionContext) {
}

// ExitPredicateExpression is called when production predicateExpression is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitPredicateExpression(ctx *PredicateExpressionContext) {
}

// EnterValueExpression is called when production valueExpression is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterValueExpression(ctx *ValueExpressionContext) {}

// ExitValueExpression is called when production valueExpression is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitValueExpression(ctx *ValueExpressionContext) {}

// EnterNotOperator is called when production notOperator is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterNotOperator(ctx *NotOperatorContext) {}

// ExitNotOperator is called when production notOperator is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitNotOperator(ctx *NotOperatorContext) {}

// EnterUnaryOperator is called when production unaryOperator is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterUnaryOperator(ctx *UnaryOperatorContext) {}

// ExitUnaryOperator is called when production unaryOperator is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitUnaryOperator(ctx *UnaryOperatorContext) {}

// EnterExactMatchModifier is called when production exactMatchModifier is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterExactMatchModifier(ctx *ExactMatchModifierContext) {
}

// ExitExactMatchModifier is called when production exactMatchModifier is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitExactMatchModifier(ctx *ExactMatchModifierContext) {}

// EnterComparisonOperator is called when production comparisonOperator is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterComparisonOperator(ctx *ComparisonOperatorContext) {
}

// ExitComparisonOperator is called when production comparisonOperator is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitComparisonOperator(ctx *ComparisonOperatorContext) {}

// EnterLogicalOperator is called when production logicalOperator is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterLogicalOperator(ctx *LogicalOperatorContext) {}

// ExitLogicalOperator is called when production logicalOperator is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitLogicalOperator(ctx *LogicalOperatorContext) {}

// EnterBitwiseOperator is called when production bitwiseOperator is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterBitwiseOperator(ctx *BitwiseOperatorContext) {}

// ExitBitwiseOperator is called when production bitwiseOperator is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitBitwiseOperator(ctx *BitwiseOperatorContext) {}

// EnterMathOperator is called when production mathOperator is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterMathOperator(ctx *MathOperatorContext) {}

// ExitMathOperator is called when production mathOperator is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitMathOperator(ctx *MathOperatorContext) {}

// EnterFunctionName is called when production functionName is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterFunctionName(ctx *FunctionNameContext) {}

// ExitFunctionName is called when production functionName is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitFunctionName(ctx *FunctionNameContext) {}

// EnterFunctionExpression is called when production functionExpression is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterFunctionExpression(ctx *FunctionExpressionContext) {
}

// ExitFunctionExpression is called when production functionExpression is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitFunctionExpression(ctx *FunctionExpressionContext) {}

// EnterLiteralValue is called when production literalValue is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterLiteralValue(ctx *LiteralValueContext) {}

// ExitLiteralValue is called when production literalValue is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitLiteralValue(ctx *LiteralValueContext) {}

// EnterTableName is called when production tableName is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterTableName(ctx *TableNameContext) {}

// ExitTableName is called when production tableName is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitTableName(ctx *TableNameContext) {}

// EnterColumnName is called when production columnName is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterColumnName(ctx *ColumnNameContext) {}

// ExitColumnName is called when production columnName is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitColumnName(ctx *ColumnNameContext) {}

// EnterOrderByColumnName is called when production orderByColumnName is entered.
func (s *BaseFilterExpressionSyntaxListener) EnterOrderByColumnName(ctx *OrderByColumnNameContext) {}

// ExitOrderByColumnName is called when production orderByColumnName is exited.
func (s *BaseFilterExpressionSyntaxListener) ExitOrderByColumnName(ctx *OrderByColumnNameContext) {}
