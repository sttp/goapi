//******************************************************************************************************
//  DataTable.go - Gbtc
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
//  09/23/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package data

import (
	"errors"
	"fmt"
	"strings"

	"github.com/sttp/goapi/sttp/format"
)

// DataTable represents a collection of DataColumn objects where each data column defines a name and
// a data type. Data columns can also be computed where its value would be derived from other columns
// and functions (https://sttp.github.io/documentation/filter-expressions/) defined in an expression.
// Note that this implementation uses a case-insensitive map for DataColumn name lookups. Internally,
// case-insensitive lookups are accomplished using `strings.ToUpper`.
type DataTable struct {
	parent        *DataSet
	name          string
	columnIndexes map[string]int
	columns       []*DataColumn
	rows          []*DataRow
}

func newDataTable(parent *DataSet, name string) *DataTable {
	return &DataTable{
		parent:        parent,
		name:          name,
		columnIndexes: make(map[string]int),
	}
}

// Parent gets the parent DataSet of the DataTable.
func (dt *DataTable) Parent() *DataSet {
	return dt.parent
}

// Name gets the name of the DataTable.
func (dt *DataTable) Name() string {
	return dt.name
}

// InitColumns initializes the internal column collection to the specified length.
// Any existing columns will be deleted.
func (dt *DataTable) InitColumns(length int) {
	dt.columns = make([]*DataColumn, 0, length)
	dt.columnIndexes = make(map[string]int, length)
}

// AddColumn adds the specified column to the DataTable.
func (dt *DataTable) AddColumn(column *DataColumn) {
	column.index = len(dt.columns)
	dt.columnIndexes[strings.ToUpper(column.Name())] = column.index
	dt.columns = append(dt.columns, column)
}

// Column gets the DataColumn at the specified columnIndex if the index is in range;
// otherwise, nil is returned.
func (dt *DataTable) Column(columnIndex int) *DataColumn {
	if columnIndex < 0 || columnIndex >= len(dt.columns) {
		return nil
	}

	return dt.columns[columnIndex]
}

// ColumnByName gets the DataColumn for the specified columnName if the name exists;
// otherwise, nil is returned. Lookup is case-insensitive.
func (dt *DataTable) ColumnByName(columnName string) *DataColumn {
	if columnIndex, ok := dt.columnIndexes[strings.ToUpper(columnName)]; ok {
		return dt.Column(columnIndex)
	}

	return nil
}

// ColumnIndex gets the index for the specified columnName if the name exists;
// otherwise, -1 is returned. Lookup is case-insensitive.
func (dt *DataTable) ColumnIndex(columnName string) int {
	column := dt.ColumnByName(columnName)

	if column == nil {
		return -1
	}

	return column.Index()
}

// CreateColumn creates a new DataColumn associated with the DataTable.
// Use AddColumn to add the new column to the DataTable.
func (dt *DataTable) CreateColumn(name string, dataType DataTypeEnum, expression string) *DataColumn {
	return newDataColumn(dt, name, dataType, expression)
}

// CloneColumn creates a copy of the specified source DataColumn associated with the DataTable.
func (dt *DataTable) CloneColumn(source *DataColumn) *DataColumn {
	return dt.CreateColumn(source.Name(), source.Type(), source.Expression())
}

// ColumnCount gets the total number columns defined in the DataTable.
func (dt *DataTable) ColumnCount() int {
	return len(dt.columns)
}

// InitRows initializes the internal row collection to the specified length.
// Any existing rows will be deleted.
func (dt *DataTable) InitRows(length int) {
	dt.rows = make([]*DataRow, 0, length)
}

// AddRow adds the specified row to the DataTable.
func (dt *DataTable) AddRow(row *DataRow) {
	dt.rows = append(dt.rows, row)
}

// Row gets the DataRow at the specified rowIndex if the index is in range;
// otherwise, nil is returned.
func (dt *DataTable) Row(rowIndex int) *DataRow {
	if rowIndex < 0 || rowIndex >= len(dt.rows) {
		return nil
	}

	return dt.rows[rowIndex]
}

// CreateRow creates a new DataRow associated with the DataTable.
// Use AddRow to add the new row to the DataTable.
func (dt *DataTable) CreateRow() *DataRow {
	return newDataRow(dt)
}

// CloneRow creates a copy of the specified source DataRow associated with the DataTable.
func (dt *DataTable) CloneRow(source *DataRow) *DataRow {
	row := dt.CreateRow()

	for i := 0; i < len(dt.columns); i++ {
		value, _ := source.Value(i)
		row.SetValue(i, value)
	}

	return row
}

// RowCount gets the total number of rows defined in the DataTable.
func (dt *DataTable) RowCount() int {
	return len(dt.rows)
}

// RowValueAsString reads the row record value at the specified columnIndex converted to a string.
// For columnIndex out of range or any other errors, an empty string will be returned.
func (dt *DataTable) RowValueAsString(rowIndex, columnIndex int) string {
	row := dt.Row(rowIndex)

	if row == nil {
		return ""
	}

	return row.ValueAsString(columnIndex)
}

// RowValueAsStringByName reads the row record value for the specified columnName converted to a string.
// For columnName not found or any other errors, an empty string will be returned.
func (dt *DataTable) RowValueAsStringByName(rowIndex int, columnName string) string {
	row := dt.Row(rowIndex)

	if row == nil {
		return ""
	}

	return row.ValueAsStringByName(columnName)
}

// String get a representation of the DataTable as a string.
func (dt *DataTable) String() string {
	var image strings.Builder

	image.WriteString(dt.name)
	image.WriteString(" [")

	for i := 0; i < len(dt.columns); i++ {
		if i > 0 {
			image.WriteString(", ")
		}

		image.WriteString(dt.columns[i].String())
	}

	image.WriteString("] x ")
	image.WriteString(format.Int(len(dt.rows)))
	image.WriteString(" rows")

	return image.String()
}

// Select returns the rows matching the filterExpression criteria in the specified sort order. The filterExpression parameter
// should be in the syntax of a SQL WHERE expression but should not include the WHERE keyword. The sortOrder parameter defines
// field names, separated by commas, that exist in the DataTable used to order the results. Each field specified in the
// sortOrder can have an ASC or DESC suffix; defaults to ASC when no suffix is provided. When sortOrder is an empty string,
// records will be returned in natural order. Set limit parameter to -1 for all matching rows. When filterExpression is an
// empty string, all records will be returned; any specified sortOrder and limit will still be respected.
func (dt *DataTable) Select(filterExpression string, sortOrder string, limit int) ([]*DataRow, error) {
	if len(filterExpression) == 0 {
		filterExpression = "True" // Return all records
	}

	if limit > 0 {
		filterExpression = fmt.Sprintf("FILTER TOP %d %s WHERE %s", limit, dt.name, filterExpression)
	} else {
		filterExpression = fmt.Sprintf("FILTER %s WHERE %s", dt.name, filterExpression)
	}

	if len(sortOrder) > 0 {
		filterExpression = fmt.Sprintf("%s ORDER BY %s", filterExpression, sortOrder)
	}

	expressionTree, err := GenerateExpressionTree(dt, filterExpression, true)

	if err != nil {
		return nil, errors.New("failed to parse filter expression, " + err.Error())
	}

	return expressionTree.Select(dt)
}
