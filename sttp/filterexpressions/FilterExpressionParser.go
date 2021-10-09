//******************************************************************************************************
//  FilterExpressionParser.go - Gbtc
//
//  Copyright © 2021, Grid Protection Alliance.  All Rights Reserved.
//
//  Licensed to the Grid Protection Alliance (GPA) under one or more contributor license agreements. See
//  the NOTICE file distributed with this work for additional information regarding copyright ownership.
//  The GPA licenses this file to you under the MIT License (MIT), the "License"; you may not use this
//  file except in compliance with the License. You may obtain a copy of the License at:
//
//      http://opensource.org/licenses/MIT
//
//  Unless agreed to in writing, the subject software distributed under the License is distributed on an
//  "AS-IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. Refer to the
//  License for the specific language governing permissions and limitations.
//
//  Code Modification History:
//  ----------------------------------------------------------------------------------------------------
//  10/07/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package filterexpressions

import (
	"errors"
	"math"
	"strconv"
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/araddon/dateparse"
	"github.com/shopspring/decimal"
	"github.com/sttp/goapi/sttp/data"
	"github.com/sttp/goapi/sttp/filterexpressions/parser"
	"github.com/sttp/goapi/sttp/guid"
)

// FilterExpressionParser represents a parser for STTP filter expressions.
type FilterExpressionParser struct {
	*parser.BaseFilterExpressionSyntaxListener

	inputStream   *antlr.InputStream
	lexer         *parser.FilterExpressionSyntaxLexer
	tokens        *antlr.CommonTokenStream
	parser        *parser.FilterExpressionSyntaxParser
	errorListener *CallbackErrorListener

	tableIDFields map[string]*TableIDFields

	filteredRows   []*data.DataRow
	filteredRowSet data.DataRowHashSet

	filteredSignalIDs   []guid.Guid
	filteredSignalIDSet guid.HashSet

	filterExpressionStatementCount int

	activeExpressionTree *ExpressionTree
	expressionTrees      []*ExpressionTree
	expressions          map[antlr.ParserRuleContext]Expression

	// DataSet defines the source metadata used for parsing the filter expression.
	DataSet *data.DataSet

	// PrimaryTableName defines the name of the table to use in the DataSet when filter
	// expressions do not specify a table name, e.g., direct signal identification. See:
	// https://sttp.github.io/documentation/filter-expressions/#direct-signal-identification
	PrimaryTableName string

	// TrackFilteredRows enables tracking of matching rows during filter expression evaluation.
	// Value defaults to true. Set value to false and set TrackFilteredSignalIDs to true if
	// only signal IDs are needed post filter expression evaluation.
	TrackFilteredRows bool

	// TrackFilteredSignalIDs enables tracking of matching signal IDs during filter expression
	// evaluation. Value defaults to false.
	TrackFilteredSignalIDs bool
}

// NewFilterExpressionParser creates a new FilterExpressionParser.
func NewFilterExpressionParser(filterExpression string, suppressConsoleErrorOutput bool) *FilterExpressionParser {
	fep := new(FilterExpressionParser)

	fep.inputStream = antlr.NewInputStream(filterExpression)
	fep.lexer = parser.NewFilterExpressionSyntaxLexer(fep.inputStream)
	fep.tokens = antlr.NewCommonTokenStream(fep.lexer, 0)
	fep.parser = parser.NewFilterExpressionSyntaxParser(fep.tokens)
	fep.errorListener = NewCallbackErrorListener()
	fep.tableIDFields = make(map[string]*TableIDFields)
	fep.TrackFilteredRows = true

	if suppressConsoleErrorOutput {
		fep.parser.RemoveErrorListeners()
	}

	fep.parser.AddErrorListener(fep.errorListener)

	return fep
}

// FilteredRows gets the rows matching the parsed filter expression.
func (fep *FilterExpressionParser) FilteredRows() []*data.DataRow {
	return fep.filteredRows
}

// FilteredRows gets the unique row set matching the parsed filter expression.
func (fep *FilterExpressionParser) FilteredRowSet() data.DataRowHashSet {
	fep.initializeSetOperations()
	return fep.filteredRowSet
}

// FilteredSignalIDs gets the Guid-based signalIDs matching the parsed filter expression.
func (fep *FilterExpressionParser) FilteredSignalIDs() []guid.Guid {
	return fep.filteredSignalIDs
}

// FilteredSignalIDs gets the unique Guid-based signalID set matching the parsed filter expression.
func (fep *FilterExpressionParser) FilteredSignalIDSet() guid.HashSet {
	fep.initializeSetOperations()
	return fep.filteredSignalIDSet
}

// TableIDFields gets the table ID fields associated with the specified tableName; or nil if not found.
func (fep *FilterExpressionParser) TableIDFields(tableName string) *TableIDFields {
	tableIDFields, ok := fep.tableIDFields[tableName]

	if ok {
		return tableIDFields
	}

	return nil
}

// RegisterTableIDFields associates the tableIDFields value with the specified tableName.
func (fep *FilterExpressionParser) RegisterTableIDFields(tableName string, tableIDFields *TableIDFields) {
	fep.tableIDFields[tableName] = tableIDFields
}

// RegisterParsingExceptionCallback registers a callback for receiving parsing exception messsages.
func (fep *FilterExpressionParser) RegisterParsingExceptionCallback(callback func(message string)) {
	fep.errorListener.ParsingExceptionCallback = callback
}

func (fep *FilterExpressionParser) visitParseTreeNodes() error {
	var err error

	defer func() {
		if r := recover(); r != nil {
			switch rt := r.(type) {
			case string:
				err = errors.New(rt)
			case error:
				err = rt
			default:
				err = errors.New("unknown panic")
			}
		}
	}()

	// Create a parse tree and start visiting listener methods
	walker := antlr.NewParseTreeWalker()
	parseTree := fep.parser.Parse()
	walker.Walk(fep, parseTree)

	return err
}

func (fep *FilterExpressionParser) initializeSetOperations() {
	// As an optimization, set operations are not engaged until second filter expression statement
	// is encountered, only then will duplicate results be a concern. Note that only using an
	// HastSet is not an option because results can be sorted with the "ORDER BY" clause.
	if fep.TrackFilteredRows && fep.filteredRowSet == nil {
		count := len(fep.filteredRows)
		fep.filteredRowSet = make(data.DataRowHashSet, count)

		for i := 0; i < count; i++ {
			fep.filteredRowSet.Add(fep.filteredRows[i])
		}
	}

	if fep.TrackFilteredSignalIDs && fep.filteredSignalIDSet == nil {
		count := len(fep.filteredSignalIDs)
		fep.filteredSignalIDSet = make(guid.HashSet, count)

		for i := 0; i < count; i++ {
			fep.filteredSignalIDSet.Add(fep.filteredSignalIDs[i])
		}
	}
}

func (fep *FilterExpressionParser) addMatchedRow(row *data.DataRow, signalIDColumnIndex int) {
	if fep.filterExpressionStatementCount > 1 {
		// Set operations
		if fep.TrackFilteredRows && fep.filteredRowSet.Add(row) {
			fep.filteredRows = append(fep.filteredRows, row)
		}

		if fep.TrackFilteredSignalIDs {
			signalIDField, null, err := row.GuidValue(signalIDColumnIndex)

			if !null && err != nil && !signalIDField.IsZero() && fep.filteredSignalIDSet.Add(signalIDField) {
				fep.filteredSignalIDs = append(fep.filteredSignalIDs, signalIDField)
			}
		}
	} else {
		// Vector only operations
		if fep.TrackFilteredRows {
			fep.filteredRows = append(fep.filteredRows, row)
		}

		if fep.TrackFilteredSignalIDs {
			signalIDField, null, err := row.GuidValue(signalIDColumnIndex)

			if !null && err != nil && !signalIDField.IsZero() {
				fep.filteredSignalIDs = append(fep.filteredSignalIDs, signalIDField)
			}
		}
	}
}

func (fep *FilterExpressionParser) mapMatchedFieldRow(primaryTable *data.DataTable, columnName string, matchValue string, signalIDColumnIndex int) {
	column := primaryTable.ColumnByName(columnName)

	if column == nil {
		return
	}

	matchValue = strings.ToUpper(matchValue)
	columnIndex := column.Index()

	for i := 0; i < primaryTable.RowCount(); i++ {
		row := primaryTable.Row(i)

		if row == nil {
			continue
		}

		value, null, err := row.StringValue(columnIndex)

		if !null && err == nil && matchValue == strings.ToUpper(value) {
			fep.addMatchedRow(row, signalIDColumnIndex)
			return
		}
	}
}

func (fep *FilterExpressionParser) tryGetExpr(ctx antlr.ParserRuleContext, value *Expression) bool {
	var ok bool
	*value, ok = fep.expressions[ctx]
	return ok
}

func (fep *FilterExpressionParser) addExpr(ctx antlr.ParserRuleContext, value Expression) {
	fep.expressions[ctx] = value
}

// Evaluate parses each statement in the filter expression and tracks the results.
// Calling evaluate will yield filtered rows and/or signal IDs that match the
// specified STTP filter expression.
func (fep *FilterExpressionParser) Evaluate() error {
	if fep.DataSet == nil {
		return errors.New("cannot evaluate filter expression, no DataSet has been defined")
	}

	if !fep.TrackFilteredRows && !fep.TrackFilteredSignalIDs {
		return errors.New("no use in evaluating filter expression, neither filtered rows nor signal IDs have been set for tracking")
	}

	fep.filterExpressionStatementCount = 0
	fep.filteredRows = make([]*data.DataRow, 0)
	fep.filteredRowSet = nil
	fep.filteredSignalIDs = make([]guid.Guid, 0)
	fep.filteredSignalIDSet = nil
	fep.expressionTrees = make([]*ExpressionTree, 0)
	fep.expressions = make(map[antlr.ParserRuleContext]Expression)

	if err := fep.visitParseTreeNodes(); err != nil {
		return err
	}

	// Each statement in the filter expression will have its own expression tree, evaluate each
	for _, expressionTree := range fep.expressionTrees {
		matchedRows, err := expressionTree.Select()

		if err != nil {
			return err
		}

		signalIDColumnIndex := -1

		if fep.TrackFilteredSignalIDs {
			table := expressionTree.Table()
			primaryTableIDFields := fep.tableIDFields[table.Name()]

			if primaryTableIDFields == nil {
				return errors.New("failed to find ID fields record for table \"" + table.Name() + "\"")
			}

			signalIDColumn := table.ColumnByName(primaryTableIDFields.SignalIDFieldName)

			if signalIDColumn == nil {
				return errors.New("failed to find signal ID field \"" + primaryTableIDFields.SignalIDFieldName + "\" for table \"" + table.Name() + "\"")
			}

			signalIDColumnIndex = signalIDColumn.Index()
		}

		for _, matchedRow := range matchedRows {
			fep.addMatchedRow(matchedRow, signalIDColumnIndex)
		}
	}

	return nil
}

// GetExpressionTrees gets the expression trees, parsing the filter expression if needed.
func (fep *FilterExpressionParser) GetExpressionTrees() ([]*ExpressionTree, error) {
	if len(fep.expressionTrees) == 0 {
		if err := fep.visitParseTreeNodes(); err != nil {
			return nil, err
		}
	}

	return fep.expressionTrees, nil
}

/*
   filterExpressionStatement
    : identifierStatement
    | filterStatement
    | expression
    ;
*/

// EnterFilterExpressionStatement is called when production filterExpressionStatement is entered.
func (fep *FilterExpressionParser) EnterFilterExpressionStatement(context *parser.FilterExpressionStatementContext) {
	// One filter expression can contain multiple filter statements separated by semi-colon,
	// so we track each as an independent expression tree
	fep.expressions = make(map[antlr.ParserRuleContext]Expression)
	fep.activeExpressionTree = nil
	fep.filterExpressionStatementCount++

	// Encountering second filter expression statement necessitates the use of set operations
	// to prevent possible result duplications
	if fep.filterExpressionStatementCount == 2 {
		fep.initializeSetOperations()
	}
}

/*
   filterStatement
    : K_FILTER ( K_TOP topLimit )? tableName K_WHERE expression ( K_ORDER K_BY orderingTerm ( ',' orderingTerm )* )?
    ;

   topLimit
    : ( '-' | '+' )? INTEGER_LITERAL
    ;

   orderingTerm
    : exactMatchModifier? columnName ( K_ASC | K_DESC )?
    ;
*/

// EnterFilterStatement is called when production filterStatement is entered.
func (fep *FilterExpressionParser) EnterFilterStatement(context *parser.FilterStatementContext) {
	tableName := context.TableName().GetText()
	table := fep.DataSet.Table(tableName)

	if table == nil {
		panic("failed to find table \"" + tableName + "\"")
	}

	fep.activeExpressionTree = NewExpressionTree(table)
	fep.expressionTrees = append(fep.expressionTrees, fep.activeExpressionTree)

	if context.K_TOP() != nil {
		topLimit, err := strconv.Atoi(context.TopLimit().GetText())

		if err == nil {
			fep.activeExpressionTree.TopLimit = topLimit
		} else {
			fep.activeExpressionTree.TopLimit = -1
		}
	}

	if context.K_ORDER() != nil && context.K_BY() != nil {
		orderingTerms := context.AllOrderingTerm()

		for i := 0; i < len(orderingTerms); i++ {
			orderingTerm := orderingTerms[i].(*parser.OrderingTermContext)
			orderByColumnName := orderingTerm.OrderByColumnName().GetText()
			orderByColumn := table.ColumnByName(orderByColumnName)

			if orderByColumn == nil {
				panic("failed to find order by field \"" + orderByColumnName + "\" for table \"" + table.Name() + "\"")
			}

			fep.activeExpressionTree.OrderByTerms = append(fep.activeExpressionTree.OrderByTerms, &OrderByTerm{
				Column:     orderByColumn,
				Ascending:  orderingTerm.K_DESC() == nil,
				ExactMatch: orderingTerm.ExactMatchModifier() == nil,
			})
		}
	}
}

/*
   identifierStatement
    : GUID_LITERAL
    | MEASUREMENT_KEY_LITERAL
    | POINT_TAG_LITERAL
    ;
*/

// ExitIdentifierStatement is called when production identifierStatement is exited.
//gocyclo: ignore
func (fep *FilterExpressionParser) ExitIdentifierStatement(context *parser.IdentifierStatementContext) {
	signalID := guid.Empty

	if context.GUID_LITERAL() != nil {
		signalID = parseGuidLiteral(context.GUID_LITERAL().GetText())

		if !fep.TrackFilteredRows && !fep.TrackFilteredSignalIDs {
			// Handle edge case of encountering standalone Guid when not tracking rows or table identifiers.
			// In this scenario the filter expression parser would only be used to generate expression trees
			// for general expression parsing, e.g., for a DataColumn expression, so here the Guid should be
			// treated as a literal expression value instead of an identifier to track:
			fep.EnterExpression(nil)
			fep.activeExpressionTree.Root = newValueExpression(ExpressionValueType.Guid, signalID)
			return
		}

		if fep.TrackFilteredSignalIDs && !signalID.IsZero() {
			if fep.filterExpressionStatementCount > 1 {
				if fep.filteredSignalIDSet.Add(signalID) {
					fep.filteredSignalIDs = append(fep.filteredSignalIDs, signalID)
				}
			} else {
				fep.filteredSignalIDs = append(fep.filteredSignalIDs, signalID)
			}
		}

		if !fep.TrackFilteredRows {
			return
		}
	}

	primaryTable := fep.DataSet.Table(fep.PrimaryTableName)

	if primaryTable == nil {
		return
	}

	primaryTableIDFields, ok := fep.tableIDFields[fep.PrimaryTableName]

	if !ok || primaryTableIDFields == nil {
		return
	}

	signalIDColumn := primaryTable.ColumnByName(primaryTableIDFields.SignalIDFieldName)

	if signalIDColumn == nil {
		return
	}

	signalIDColumnIndex := signalIDColumn.Index()

	if fep.TrackFilteredRows && !signalID.IsZero() {
		// Map matching row for manually specified Guid
		for i := 0; i < primaryTable.RowCount(); i++ {
			row := primaryTable.Row(i)

			if row == nil {
				continue
			}

			value, null, err := row.GuidValue(signalIDColumnIndex)

			if !null && err == nil && value == signalID {
				if fep.filterExpressionStatementCount > 1 {
					if fep.filteredRowSet.Add(row) {
						fep.filteredRows = append(fep.filteredRows, row)
					}
				} else {
					fep.filteredRows = append(fep.filteredRows, row)
				}

				return
			}
		}

		return
	}

	if context.MEASUREMENT_KEY_LITERAL() != nil {
		fep.mapMatchedFieldRow(primaryTable, primaryTableIDFields.MeasurementKeyFieldName, context.MEASUREMENT_KEY_LITERAL().GetText(), signalIDColumnIndex)
		return
	}

	if context.POINT_TAG_LITERAL() != nil {
		fep.mapMatchedFieldRow(primaryTable, primaryTableIDFields.PointTagFieldName, parsePointTagLiteral(context.POINT_TAG_LITERAL().GetText()), signalIDColumnIndex)
	}
}

/*
   expression
    : notOperator expression
    | expression logicalOperator expression
    | predicateExpression
    ;
*/

// EnterExpression is called when production expression is entered.
func (fep *FilterExpressionParser) EnterExpression(*parser.ExpressionContext) {
	// Handle case of encountering a standalone expression, i.e., an expression not
	// within a filter statement context
	if fep.activeExpressionTree == nil {
		table := fep.DataSet.Table(fep.PrimaryTableName)

		if table == nil {
			panic("failed to find table \"" + fep.PrimaryTableName + "\"")
		}

		fep.activeExpressionTree = NewExpressionTree(table)
		fep.expressionTrees = append(fep.expressionTrees, fep.activeExpressionTree)
	}
}

/*
   expression
    : notOperator expression
    | expression logicalOperator expression
    | predicateExpression
    ;
*/

// ExitExpression is called when production expression is exited.
func (fep *FilterExpressionParser) ExitExpression(context *parser.ExpressionContext) {
	var value Expression

	// Check for predicate expressions (see explicit visit function)
	predicateExpression := context.PredicateExpression()

	if predicateExpression != nil {
		if fep.tryGetExpr(predicateExpression, &value) {
			fep.addExpr(context, value)
			return
		}

		panic("failed to find predicate expression \"" + predicateExpression.GetText() + "\"")
	}

	// Check for not operator expressions
	notOperator := context.NotOperator()

	if notOperator != nil {
		expressions := context.AllExpression()

		if len(expressions) != 1 {
			panic("not operator expression is malformed: \"" + context.GetText() + "\"")
		}

		if !fep.tryGetExpr(expressions[0], &value) {
			panic("failed to find not operator expression \"" + context.GetText() + "\"")
		}

		fep.addExpr(context, NewUnaryExpression(ExpressionUnaryType.Not, value))
		return
	}

	// Check for logical operator expressions
	logicalOperator := context.LogicalOperator()

	if logicalOperator != nil {
		logicalOperatorContext := logicalOperator.(*parser.LogicalOperatorContext)
		var leftValue, rightValue Expression
		var operatorType ExpressionOperatorTypeEnum
		expressions := context.AllExpression()

		if len(expressions) != 2 {
			panic("operator expression, in logical operator expression context, is malformed: \"" + context.GetText() + "\"")
		}

		if !fep.tryGetExpr(expressions[0], &leftValue) {
			panic("failed to find left operator expression \"" + expressions[0].GetText() + "\"")
		}

		if !fep.tryGetExpr(expressions[1], &rightValue) {
			panic("failed to find right operator expression \"" + expressions[1].GetText() + "\"")
		}

		operatorSymbol := logicalOperatorContext.GetText()

		// Check for boolean operations
		if logicalOperatorContext.K_AND() != nil || operatorSymbol == "&&" {
			operatorType = ExpressionOperatorType.And
		} else if logicalOperatorContext.K_OR() != nil || operatorSymbol == "||" {
			operatorType = ExpressionOperatorType.Or
		} else {
			panic("unexpected logical operator \"" + operatorSymbol + "\"")
		}

		fep.addExpr(context, NewOperatorExpression(operatorType, leftValue, rightValue))
		return
	}

	panic("unexpected expression \"" + context.GetText() + "\"")
}

/*
   predicateExpression
    : predicateExpression notOperator? K_IN exactMatchModifier? '(' expressionList ')'
    | predicateExpression K_IS notOperator? K_NULL
    | predicateExpression comparisonOperator predicateExpression
    | predicateExpression notOperator? K_LIKE exactMatchModifier? predicateExpression
    | valueExpression
    ;
*/

// ExitPredicateExpression is called when production predicateExpression is exited.
//gocyclo: ignore
func (fep *FilterExpressionParser) ExitPredicateExpression(context *parser.PredicateExpressionContext) {
	var value Expression

	// Check for value expressions (see explicit visit function)
	valueExpression := context.ValueExpression()

	if valueExpression != nil {
		if fep.tryGetExpr(valueExpression, &value) {
			fep.addExpr(context, value)
			return
		}

		panic("failed to find value expression \"" + valueExpression.GetText() + "\"")
	}

	notOperator := context.NotOperator()
	exactMatchModifier := context.ExactMatchModifier()

	// Check for IN expressions
	if context.K_IN() != nil {
		predicates := context.AllPredicateExpression()

		// IN expression expects one predicate
		if len(predicates) != 1 {
			panic("\"IN\" expression is malformed: \"" + context.GetText() + "\"")
		}

		if !fep.tryGetExpr(predicates[0], &value) {
			panic("failed to find \"IN\" predicate expression \"" + predicates[0].GetText() + "\"")
		}

		expressionListContext := context.ExpressionList().(*parser.ExpressionListContext)
		expressions := expressionListContext.AllExpression()
		argumentCount := len(expressions)

		if argumentCount < 1 {
			panic("not enough expressions found for \"IN\" operation")
		}

		arguments := make([]Expression, 0, argumentCount)

		for i := 0; i < argumentCount; i++ {
			var argument Expression

			if fep.tryGetExpr(expressions[i], &argument) {
				arguments = append(arguments, argument)
			} else {
				panic("failed to find argument expression " + strconv.Itoa(i) + " \"" + expressions[i].GetText() + "\" for \"IN\" operation")
			}
		}

		fep.addExpr(context, NewInListExpression(value, arguments, notOperator != nil, exactMatchModifier != nil))
		return
	}

	// Check for IS NULL expressions
	if context.K_IS() != nil && context.K_NULL() != nil {
		var operatorType ExpressionOperatorTypeEnum

		if notOperator == nil {
			operatorType = ExpressionOperatorType.IsNull
		} else {
			operatorType = ExpressionOperatorType.IsNotNull
		}

		predicates := context.AllPredicateExpression()

		// IS NULL expression expects one predicate
		if len(predicates) != 1 {
			panic("\"IS NULL\" expression is malformed: \"" + context.GetText() + "\"")
		}

		if fep.tryGetExpr(predicates[0], &value) {
			fep.addExpr(context, NewOperatorExpression(operatorType, value, nil))
			return
		}

		panic("failed to find \"IS NULL\" predicate expression \"" + predicates[0].GetText() + "\"")
	}

	// Remaining operators require two predicate expressions
	predicates := context.AllPredicateExpression()

	if len(predicates) != 2 {
		panic("operator expression, in predicate expression context, is malformed: \"" + context.GetText() + "\"")
	}

	var leftValue, rightValue Expression
	var operatorType ExpressionOperatorTypeEnum

	if !fep.tryGetExpr(predicates[0], &leftValue) {
		panic("failed to find left operator predicate expression \"" + predicates[0].GetText() + "\"")
	}

	if !fep.tryGetExpr(predicates[1], &rightValue) {
		panic("failed to find right operator predicate expression \"" + predicates[1].GetText() + "\"")
	}

	// Check for comparison operator expressions
	comparisonOperator := context.ComparisonOperator()

	if comparisonOperator != nil {
		operatorSymbol := comparisonOperator.GetText()

		// Check for comparison operations
		switch operatorSymbol {
		case "<":
			operatorType = ExpressionOperatorType.LessThan
		case "<=":
			operatorType = ExpressionOperatorType.LessThanOrEqual
		case ">":
			operatorType = ExpressionOperatorType.GreaterThan
		case ">=":
			operatorType = ExpressionOperatorType.GreaterThanOrEqual
		case "=", "==":
			operatorType = ExpressionOperatorType.Equal
		case "===":
			operatorType = ExpressionOperatorType.EqualExactMatch
		case "<>", "!=":
			operatorType = ExpressionOperatorType.NotEqual
		case "!==":
			operatorType = ExpressionOperatorType.NotEqualExactMatch
		default:
			panic("Unexpected comparison operator \"" + operatorSymbol + "\"")
		}

		fep.addExpr(context, NewOperatorExpression(operatorType, leftValue, rightValue))
		return
	}

	// Check for LIKE expressions
	if context.K_LIKE() != nil {
		var operatorType ExpressionOperatorTypeEnum

		if exactMatchModifier == nil {
			if notOperator == nil {
				operatorType = ExpressionOperatorType.Like
			} else {
				operatorType = ExpressionOperatorType.NotLike
			}
		} else {
			if notOperator == nil {
				operatorType = ExpressionOperatorType.LikeExactMatch
			} else {
				operatorType = ExpressionOperatorType.NotLikeExactMatch
			}
		}

		fep.addExpr(context, NewOperatorExpression(operatorType, leftValue, rightValue))
		return
	}

	panic("unexpected predicate expression \"" + context.GetText() + "\"")
}

/*
   valueExpression
    : literalValue
    | columnName
    | functionExpression
    | unaryOperator valueExpression
    | '(' expression ')'
    | valueExpression mathOperator valueExpression
    | valueExpression bitwiseOperator valueExpression
	;
*/

// ExitValueExpression is called when production valueExpression is exited.
//gocyclo: ignore
func (fep *FilterExpressionParser) ExitValueExpression(context *parser.ValueExpressionContext) {
	var value Expression

	// Check for literal values (see explicit visit function)
	literalValue := context.LiteralValue()

	if literalValue != nil {
		if fep.tryGetExpr(literalValue, &value) {
			fep.addExpr(context, value)
			return
		}

		panic("failed to find literal value \"" + literalValue.GetText() + "\"")
	}

	// Check for column names (see explicit visit function)
	columnName := context.ColumnName()

	if columnName != nil {
		if fep.tryGetExpr(columnName, &value) {
			fep.addExpr(context, value)
			return
		}

		panic("failed to find column name \"" + columnName.GetText() + "\"")
	}

	// Check for function expressions (see explicit visit function)
	functionExpression := context.FunctionExpression()

	if functionExpression != nil {
		if fep.tryGetExpr(functionExpression, &value) {
			fep.addExpr(context, value)
			return
		}

		panic("failed to find function expression \"" + functionExpression.GetText() + "\"")
	}

	// Check for unary operators
	unaryOperator := context.UnaryOperator()

	if unaryOperator != nil {
		valueExpressions := context.AllValueExpression()

		if len(valueExpressions) != 1 {
			panic("unary operator value expression is malformed: \"" + context.GetText() + "\"")
		}

		if fep.tryGetExpr(valueExpressions[0], &value) {
			var unaryType ExpressionUnaryTypeEnum
			unaryOperatorContext := unaryOperator.(*parser.UnaryOperatorContext)

			// Check for unary operators
			if unaryOperatorContext.K_NOT() == nil {
				operatorSymbol := unaryOperatorContext.GetText()

				switch operatorSymbol {
				case "+":
					unaryType = ExpressionUnaryType.Plus
				case "-":
					unaryType = ExpressionUnaryType.Minus
				case "~", "!":
					unaryType = ExpressionUnaryType.Not
				default:
					panic("unexpected unary operator type \"" + operatorSymbol + "\"")
				}
			} else {
				unaryType = ExpressionUnaryType.Not
			}

			fep.addExpr(context, NewUnaryExpression(unaryType, value))
			return
		}

		panic("failed to find unary operator value expression \"" + context.GetText() + "\"")
	}

	// Check for sub-expressions, i.e., "(" expression ")"
	expression := context.Expression()

	if expression != nil {
		if fep.tryGetExpr(expression, &value) {
			fep.addExpr(context, value)
			return
		}

		panic("failed to find sub-expression \"" + expression.GetText() + "\"")
	}

	// Remaining operators require two value expressions
	valueExpressions := context.AllValueExpression()

	if len(valueExpressions) != 2 {
		panic("operator expression, in value expression context, is malformed: \"" + context.GetText() + "\"")
	}

	var leftValue, rightValue Expression
	var operatorType ExpressionOperatorTypeEnum

	if !fep.tryGetExpr(valueExpressions[0], &leftValue) {
		panic("Failed to find left operator value expression \"" + valueExpressions[0].GetText() + "\"")
	}

	if !fep.tryGetExpr(valueExpressions[1], &rightValue) {
		panic("Failed to find right operator value expression \"" + valueExpressions[1].GetText() + "\"")
	}

	// Check for math operator expressions
	mathOperator := context.MathOperator()

	if mathOperator != nil {
		operatorSymbol := mathOperator.GetText()

		// Check for arithmetic operators
		switch operatorSymbol {
		case "*":
			operatorType = ExpressionOperatorType.Multiply
		case "/":
			operatorType = ExpressionOperatorType.Divide
		case "%":
			operatorType = ExpressionOperatorType.Modulus
		case "+":
			operatorType = ExpressionOperatorType.Add
		case "-":
			operatorType = ExpressionOperatorType.Subtract
		default:
			panic("unexpected math operator \"" + operatorSymbol + "\"")
		}

		fep.addExpr(context, NewOperatorExpression(operatorType, leftValue, rightValue))
		return
	}

	// Check for bitwise operator expressions
	bitwiseOperator := context.BitwiseOperator()

	if bitwiseOperator != nil {
		bitwiseOperatorContext := bitwiseOperator.(*parser.BitwiseOperatorContext)

		// Check for bitwise operators
		if bitwiseOperatorContext.K_XOR() == nil {
			operatorSymbol := bitwiseOperatorContext.GetText()

			switch operatorSymbol {
			case "<<":
				operatorType = ExpressionOperatorType.BitShiftLeft
			case ">>":
				operatorType = ExpressionOperatorType.BitShiftRight
			case "&":
				operatorType = ExpressionOperatorType.BitwiseAnd
			case "|":
				operatorType = ExpressionOperatorType.BitwiseOr
			case "^":
				operatorType = ExpressionOperatorType.BitwiseXor
			default:
				panic("unexpected bitwise operator \"" + operatorSymbol + "\"")
			}
		} else {
			operatorType = ExpressionOperatorType.BitwiseXor
		}

		fep.addExpr(context, NewOperatorExpression(operatorType, leftValue, rightValue))
		return
	}

	panic("unexpected value expression \"" + context.GetText() + "\"")
}

/*
   literalValue
    : INTEGER_LITERAL
    | NUMERIC_LITERAL
    | STRING_LITERAL
    | DATETIME_LITERAL
    | GUID_LITERAL
    | BOOLEAN_LITERAL
    | K_NULL
    ;
*/

// ExitLiteralValue is called when production literalValue is exited.
//gocyclo: ignore
func (fep *FilterExpressionParser) ExitLiteralValue(context *parser.LiteralValueContext) {
	var result *ValueExpression

	if integerLiteral := context.INTEGER_LITERAL(); integerLiteral != nil {
		literal := integerLiteral.GetText()

		if i64, err := strconv.ParseInt(literal, 10, 64); err == nil {
			if i64 < math.MinInt32 || i64 > math.MaxInt32 {
				result = newValueExpression(ExpressionValueType.Int64, i64)
			} else {
				result = newValueExpression(ExpressionValueType.Int32, int32(i64))
			}
		} else if ui64, err := strconv.ParseUint(literal, 10, 64); err == nil {
			if ui64 > math.MaxInt64 {
				result = parseNumericLiteral(literal)
			} else if ui64 > math.MaxInt32 {
				result = newValueExpression(ExpressionValueType.Int64, int64(ui64))
			} else {
				result = newValueExpression(ExpressionValueType.Int32, int32(ui64))
			}
		} else if d, err := decimal.NewFromString(literal); err == nil {
			result = newValueExpression(ExpressionValueType.Decimal, d)
		} else {
			result = newValueExpression(ExpressionValueType.String, literal)
		}
	} else if numericLiteral := context.NUMERIC_LITERAL(); numericLiteral != nil {
		literal := numericLiteral.GetText()

		// Real literals using scientific notation are parsed as double
		if strings.ContainsAny(literal, "Ee") {
			if f64, err := strconv.ParseFloat(literal, 64); err == nil {
				result = newValueExpression(ExpressionValueType.Double, f64)
			} else {
				result = parseNumericLiteral(literal)
			}
		} else {
			result = parseNumericLiteral(literal)
		}
	} else if stringLiteral := context.STRING_LITERAL(); stringLiteral != nil {
		result = newValueExpression(ExpressionValueType.String, parseStringLiteral(stringLiteral.GetText()))
	} else if dataTimeLiteral := context.DATETIME_LITERAL(); dataTimeLiteral != nil {
		result = newValueExpression(ExpressionValueType.DateTime, parseDateTimeLiteral(dataTimeLiteral.GetText()))
	} else if guidLiteral := context.GUID_LITERAL(); guidLiteral != nil {
		result = newValueExpression(ExpressionValueType.DateTime, parseGuidLiteral(guidLiteral.GetText()))
	} else if booleanLiteral := context.BOOLEAN_LITERAL(); booleanLiteral != nil {
		if strings.ToUpper(booleanLiteral.GetText()) == "TRUE" {
			result = True
		} else {
			result = False
		}
	} else if context.K_NULL() != nil {
		result = NullValue(ExpressionValueType.Undefined)
	}

	if result != nil {
		fep.addExpr(context, result)
	}
}

func parseNumericLiteral(literal string) *ValueExpression {
	if d, err := decimal.NewFromString(literal); err == nil {
		return newValueExpression(ExpressionValueType.Decimal, d)
	}

	if f64, err := strconv.ParseFloat(literal, 64); err == nil {
		return newValueExpression(ExpressionValueType.Double, f64)
	}

	return newValueExpression(ExpressionValueType.String, literal)
}

/*
   columnName
    : IDENTIFIER
    ;
*/

// ExitColumnName is called when production columnName is exited.
func (fep *FilterExpressionParser) ExitColumnName(context *parser.ColumnNameContext) {
	columnName := context.IDENTIFIER().GetText()
	dataColumn := fep.activeExpressionTree.Table().ColumnByName(columnName)

	if dataColumn == nil {
		panic("failed to find column \"" + columnName + "\" in table \"" + fep.activeExpressionTree.Table().Name() + "\"")
	}

	fep.addExpr(context, NewColumnExpression(dataColumn))
}

/*
   functionExpression
    : functionName '(' expressionList? ')'
    ;
*/

// ExitFunctionExpression is called when production functionExpression is exited.
//gocyclo: ignore
func (fep *FilterExpressionParser) ExitFunctionExpression(context *parser.FunctionExpressionContext) {
	functionNameContext := context.FunctionName().(*parser.FunctionNameContext)
	var functionType ExpressionFunctionTypeEnum

	switch {
	case functionNameContext.K_ABS() != nil:
		functionType = ExpressionFunctionType.Abs
	case functionNameContext.K_CEILING() != nil:
		functionType = ExpressionFunctionType.Ceiling
	case functionNameContext.K_COALESCE() != nil:
		functionType = ExpressionFunctionType.Coalesce
	case functionNameContext.K_CONVERT() != nil:
		functionType = ExpressionFunctionType.Convert
	case functionNameContext.K_CONTAINS() != nil:
		functionType = ExpressionFunctionType.Contains
	case functionNameContext.K_DATEADD() != nil:
		functionType = ExpressionFunctionType.DateAdd
	case functionNameContext.K_DATEDIFF() != nil:
		functionType = ExpressionFunctionType.DateDiff
	case functionNameContext.K_DATEPART() != nil:
		functionType = ExpressionFunctionType.DatePart
	case functionNameContext.K_ENDSWITH() != nil:
		functionType = ExpressionFunctionType.EndsWith
	case functionNameContext.K_FLOOR() != nil:
		functionType = ExpressionFunctionType.Floor
	case functionNameContext.K_IIF() != nil:
		functionType = ExpressionFunctionType.IIf
	case functionNameContext.K_INDEXOF() != nil:
		functionType = ExpressionFunctionType.IndexOf
	case functionNameContext.K_ISDATE() != nil:
		functionType = ExpressionFunctionType.IsDate
	case functionNameContext.K_ISINTEGER() != nil:
		functionType = ExpressionFunctionType.IsInteger
	case functionNameContext.K_ISGUID() != nil:
		functionType = ExpressionFunctionType.IsGuid
	case functionNameContext.K_ISNULL() != nil:
		functionType = ExpressionFunctionType.IsNull
	case functionNameContext.K_ISNUMERIC() != nil:
		functionType = ExpressionFunctionType.IsNumeric
	case functionNameContext.K_LASTINDEXOF() != nil:
		functionType = ExpressionFunctionType.LastIndexOf
	case functionNameContext.K_LEN() != nil:
		functionType = ExpressionFunctionType.Len
	case functionNameContext.K_LOWER() != nil:
		functionType = ExpressionFunctionType.Lower
	case functionNameContext.K_MAXOF() != nil:
		functionType = ExpressionFunctionType.MaxOf
	case functionNameContext.K_MINOF() != nil:
		functionType = ExpressionFunctionType.MinOf
	case functionNameContext.K_NOW() != nil:
		functionType = ExpressionFunctionType.Now
	case functionNameContext.K_NTHINDEXOF() != nil:
		functionType = ExpressionFunctionType.NthIndexOf
	case functionNameContext.K_POWER() != nil:
		functionType = ExpressionFunctionType.Power
	case functionNameContext.K_REGEXMATCH() != nil:
		functionType = ExpressionFunctionType.RegExMatch
	case functionNameContext.K_REGEXVAL() != nil:
		functionType = ExpressionFunctionType.RegExVal
	case functionNameContext.K_REPLACE() != nil:
		functionType = ExpressionFunctionType.Replace
	case functionNameContext.K_REVERSE() != nil:
		functionType = ExpressionFunctionType.Reverse
	case functionNameContext.K_ROUND() != nil:
		functionType = ExpressionFunctionType.Round
	case functionNameContext.K_SPLIT() != nil:
		functionType = ExpressionFunctionType.Split
	case functionNameContext.K_SQRT() != nil:
		functionType = ExpressionFunctionType.Sqrt
	case functionNameContext.K_STARTSWITH() != nil:
		functionType = ExpressionFunctionType.StartsWith
	case functionNameContext.K_STRCOUNT() != nil:
		functionType = ExpressionFunctionType.StrCount
	case functionNameContext.K_STRCMP() != nil:
		functionType = ExpressionFunctionType.StrCmp
	case functionNameContext.K_SUBSTR() != nil:
		functionType = ExpressionFunctionType.SubStr
	case functionNameContext.K_TRIM() != nil:
		functionType = ExpressionFunctionType.Trim
	case functionNameContext.K_TRIMLEFT() != nil:
		functionType = ExpressionFunctionType.TrimLeft
	case functionNameContext.K_TRIMRIGHT() != nil:
		functionType = ExpressionFunctionType.TrimRight
	case functionNameContext.K_UPPER() != nil:
		functionType = ExpressionFunctionType.Upper
	case functionNameContext.K_UTCNOW() != nil:
		functionType = ExpressionFunctionType.UtcNow
	default:
		panic("unexpected function type \"" + functionNameContext.GetText() + "\"")
	}

	expressionList := context.ExpressionList()
	var arguments []Expression

	if expressionList != nil {
		expressionListContext := expressionList.(*parser.ExpressionListContext)
		expressions := expressionListContext.AllExpression()
		argumentCount := len(expressions)
		arguments = make([]Expression, 0, argumentCount)

		for i := 0; i < argumentCount; i++ {
			var argument Expression

			if fep.tryGetExpr(expressions[i], &argument) {
				arguments = append(arguments, argument)
			} else {
				panic("failed to find argument expression " + strconv.Itoa(i) + " \"" + expressions[i].GetText() + "\" for function \"" + functionNameContext.GetText() + "\"")
			}
		}
	} else {
		arguments = make([]Expression, 0)
	}

	fep.addExpr(context, NewFunctionExpression(functionType, arguments))
}

// GenerateExpressionTreesFromDataSet produces a set of expression trees for the provided filterExpression and dataSet.
// One expression tree will be produced per filter expression statement encountered in the specified filterExpression.
func GenerateExpressionTreesFromDataSet(dataSet *data.DataSet, primaryTable string, filterExpression string, suppressConsoleErrorOutput bool) ([]*ExpressionTree, error) {
	parser := NewFilterExpressionParser(filterExpression, suppressConsoleErrorOutput)

	parser.DataSet = dataSet
	parser.PrimaryTableName = primaryTable
	parser.TrackFilteredRows = false

	return parser.GetExpressionTrees()
}

// GenerateExpressionTrees produces a set of expression trees for the provided filterExpression and dataTable.
// One expression tree will be produced per filter expression statement encountered in the specified filterExpression.
func GenerateExpressionTrees(dataTable *data.DataTable, filterExpression string, suppressConsoleErrorOutput bool) ([]*ExpressionTree, error) {
	return GenerateExpressionTreesFromDataSet(dataTable.Parent(), dataTable.Name(), filterExpression, suppressConsoleErrorOutput)
}

// GenerateExpressionTree gets the first produced expression trees for the provided filterExpression and dataTable.
func GenerateExpressionTree(dataTable *data.DataTable, filterExpression string, suppressConsoleErrorOutput bool) (*ExpressionTree, error) {
	expressionTrees, err := GenerateExpressionTrees(dataTable, filterExpression, suppressConsoleErrorOutput)

	if err != nil {
		return nil, err
	}

	if len(expressionTrees) > 0 {
		return expressionTrees[0], nil
	}

	return nil, errors.New("no expression trees generated with filter expression \"" + filterExpression + "\" for table \"" + dataTable.Name() + "\"")
}

// Evaluates the provided filterExpression on the specified dataRow.
func Evaluate(dataRow *data.DataRow, filterExpression string, suppressConsoleErrorOutput bool) (*ValueExpression, error) {
	expressionTree, err := GenerateExpressionTree(dataRow.Parent(), filterExpression, suppressConsoleErrorOutput)

	if err != nil {
		return nil, err
	}

	return expressionTree.Evaluate(dataRow)
}

// SelectFromDataSet returns the rows matching the provided filterExpression and dataSet.
func SelectFromDataSet(dataSet *data.DataSet, filterExpression string, primaryTable string, tableIDFields *TableIDFields, suppressConsoleErrorOutput bool) ([]*data.DataRow, error) {
	parser := NewFilterExpressionParser(filterExpression, suppressConsoleErrorOutput)

	parser.DataSet = dataSet

	if len(primaryTable) > 0 {
		parser.PrimaryTableName = primaryTable

		if tableIDFields == nil {
			parser.RegisterTableIDFields(primaryTable, DefaultTableIDFields)
		} else {
			parser.RegisterTableIDFields(primaryTable, tableIDFields)
		}
	}

	if err := parser.Evaluate(); err != nil {
		return nil, err
	}

	return parser.FilteredRows(), nil
}

// Select returns the rows matching the provided filterExpression and dataTable.
func Select(dataTable *data.DataTable, filterExpression string, tableIDFields *TableIDFields, suppressConsoleErrorOutput bool) ([]*data.DataRow, error) {
	return SelectFromDataSet(dataTable.Parent(), filterExpression, dataTable.Name(), tableIDFields, suppressConsoleErrorOutput)
}

func parseStringLiteral(stringLiteral string) string {
	// Remove any surrounding quotes from string, ANTLR grammar already
	// ensures strings starting with quote also ends with one
	if stringLiteral[0] == '\'' {
		return stringLiteral[1 : len(stringLiteral)-1]
	}

	return stringLiteral
}

func parseGuidLiteral(guidLiteral string) guid.Guid {
	// Remove any quotes from GUID (boost currently only handles optional braces),
	// ANTLR grammar already ensures GUID starting with quote also ends with one
	if guidLiteral[0] == '\'' {
		guidLiteral = guidLiteral[1 : len(guidLiteral)-1]
	}

	g, _ := guid.Parse(guidLiteral)
	return g
}

func parseDateTimeLiteral(dateTimeLiteral string) time.Time {
	// Remove any surrounding '#' symbols from date/time, ANTLR grammar already
	// ensures date/time starting with '#' symbol will also end with one
	if dateTimeLiteral[0] == '#' {
		dateTimeLiteral = dateTimeLiteral[1 : len(dateTimeLiteral)-1]
	}

	dt, _ := dateparse.ParseAny(dateTimeLiteral)
	return dt
}

func parsePointTagLiteral(pointTagLiteral string) string {
	// Remove any double-quotes from point tag literal, ANTLR grammar already
	// ensures tag starting with quote also ends with one
	if pointTagLiteral[0] == '"' {
		return pointTagLiteral[1 : len(pointTagLiteral)-1]
	}

	return pointTagLiteral
}
