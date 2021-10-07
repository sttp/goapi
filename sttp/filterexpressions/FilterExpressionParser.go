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
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
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

// Evaluate parses each statement in the filter expression and tracks the results.
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

	fep.visitParseTreeNodes()

	// Each statement in the filter expression will have its own expression tree, evaluate each
	for _, expressionTree := range fep.expressionTrees {
		matchedRows := fep.Select(expressionTree)
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

// func (fep *FilterExpressionParser) tryGetExpr(context antlr.ParserRuleContext) (Expression, bool) {
// 	if expression, ok := fep.expressions[context]; ok {
// 		return expression, true
// 	}

// 	return nil, false
// }

// GetExpressionTrees gets the expression trees, parsing the filter expression if needed.
func (fep *FilterExpressionParser) GetExpressionTrees() []*ExpressionTree {
	if len(fep.expressionTrees) == 0 {
		fep.visitParseTreeNodes()
	}

	return fep.expressionTrees
}

/*
   filterExpressionStatement
    : identifierStatement
    | filterStatement
    | expression
    ;
*/

// EnterFilterExpressionStatement is called when production filterExpressionStatement is entered.
func (fep *FilterExpressionParser) EnterFilterExpressionStatement(ctx *parser.FilterExpressionStatementContext) {
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
func (fep *FilterExpressionParser) EnterFilterStatement(ctx *parser.FilterStatementContext) {
	tableName := ctx.TableName().GetText()
	table := fep.DataSet.Table(tableName)

	if table == nil {
		panic("failed to find table \"" + tableName + "\"")
	}

	fep.activeExpressionTree = NewExpressionTree(table)
	fep.expressionTrees = append(fep.expressionTrees, fep.activeExpressionTree)

	if ctx.K_TOP() != nil {
		topLimit, err := strconv.Atoi(ctx.TopLimit().GetText())

		if err == nil {
			fep.activeExpressionTree.TopLimit = topLimit
		} else {
			fep.activeExpressionTree.TopLimit = -1
		}
	}

	if ctx.K_ORDER() != nil && ctx.K_BY() != nil {
		orderingTerms := ctx.AllOrderingTerm()

		for i := 0; i < len(orderingTerms); i++ {
			orderingTerm := orderingTerms[i].(*parser.OrderingTermContext)
			orderByColumnName := orderingTerm.OrderByColumnName().GetText()
			orderByColumn := table.ColumnByName(orderByColumnName)

			if orderByColumn == nil {
				panic("Failed to find order by field \"" + orderByColumnName + "\" for table \"" + table.Name() + "\"")
			}

			fep.activeExpressionTree.OrderByTerms = append(fep.activeExpressionTree.OrderByTerms, &OrderByTerm{
				Column:     orderByColumn,
				Ascending:  orderingTerm.K_DESC() == nil,
				ExactMatch: orderingTerm.ExactMatchModifier() == nil,
			})
		}
	}
}

// Select evaluates the specified expression tree returning matching rows.
func (fep *FilterExpressionParser) Select(expressionTree *ExpressionTree) []*data.DataRow {
	matchedRows := make([]*data.DataRow, 0)

	// TODO...

	return matchedRows
}
