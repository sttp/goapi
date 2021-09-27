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

package metadata

import "strings"

// DataTable represents a collection of DataColumn objects where each data column defines a name and
// a data type. Data columns can also be computed where its value would be derived from other columns
// and functions (https://sttp.github.io/documentation/filter-expressions/) defined in an expression.
// Note that this implementation uses a case-insensitive map for DataColumn name lookups. Internally
// this is accomplished using ToUpper to keep things simple and efficient, however, this implies that
// case-insensitivity will be effectively restricted to ASCII-based column names.
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

// GetRowValue reads the row record value at the specified columnIndex as a string,
// if columnIndex is out of range, an empty string will be returned.
func (dt *DataTable) GetRowValue(rowIndex int, columnIndex int) string {
	row := dt.Row(rowIndex)

	if row == nil {
		return ""
	}

	return row.GetValue(columnIndex)
}

// GetRowValueByName reads the row record value at the specified columnName as a string,
// if columnName is not found, an empty string will be returned.
func (dt *DataTable) GetRowValueByName(rowIndex int, columnName string) string {
	row := dt.Row(rowIndex)

	if row == nil {
		return ""
	}

	return row.GetValueByName(columnName)
}
