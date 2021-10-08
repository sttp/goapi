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
	"strings"
	"time"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	"github.com/araddon/dateparse"
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

/*
   identifierStatement
    : GUID_LITERAL
    | MEASUREMENT_KEY_LITERAL
    | POINT_TAG_LITERAL
    ;
*/

// ExitIdentifierStatement is called when production identifierStatement is exited.
func (fep *FilterExpressionParser) ExitIdentifierStatement(ctx *parser.IdentifierStatementContext) {
	signalID := guid.Empty

	if ctx.GUID_LITERAL() != nil {
		signalID = parseGuidLiteral(ctx.GUID_LITERAL().GetText())

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

	if ctx.MEASUREMENT_KEY_LITERAL() != nil {
		fep.mapMatchedFieldRow(primaryTable, primaryTableIDFields.MeasurementKeyFieldName, ctx.MEASUREMENT_KEY_LITERAL().GetText(), signalIDColumnIndex)
		return
	}

	if ctx.POINT_TAG_LITERAL() != nil {
		fep.mapMatchedFieldRow(primaryTable, primaryTableIDFields.PointTagFieldName, parsePointTagLiteral(ctx.POINT_TAG_LITERAL().GetText()), signalIDColumnIndex)
	}
}

// Select evaluates the specified expression tree returning matching rows.
func (fep *FilterExpressionParser) Select(expressionTree *ExpressionTree) []*data.DataRow {
	matchedRows := make([]*data.DataRow, 0)

	// TODO...

	return matchedRows
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
