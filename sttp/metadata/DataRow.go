//******************************************************************************************************
//  DataRow.go - Gbtc
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

import (
	"fmt"
	"strconv"
	"time"
)

// DataRow represents a row, i.e., a record, in a DataTable defining a set of values for each
// defined DataColumn field in the DataTable columns collection.
type DataRow struct {
	parent *DataTable
	values []interface{}
}

func newDataRow(parent *DataTable) *DataRow {
	return &DataRow{
		parent: parent,
		values: make([]interface{}, parent.ColumnCount()),
	}
}

// Parent gets the parent DataTable of the DataRow.
func (dr *DataRow) Parent() *DataTable {
	return dr.parent
}

func (dr *DataRow) getColumnIndex(columnName string) int {
	column := dr.parent.ColumnByName(columnName)

	if column == nil {
		panic("Column name \"" + columnName + "\" was not found in table \"" + dr.parent.Name() + "\"")
	}

	return column.Index()
}

func (dr *DataRow) validateColumnType(columnIndex int, targetType int, read bool) *DataColumn {
	column := dr.parent.Column(columnIndex)

	if column == nil {
		panic("Column index " + strconv.Itoa(columnIndex) + " is out of range for table \"" + dr.parent.Name() + "\"")
	}

	if targetType > -1 && column.Type() != DataTypeEnum(targetType) {
		var action string
		var preposition string

		if read {
			action = "read"
			preposition = "from"
		} else {
			action = "assign"
			preposition = "to"
		}

		panic(fmt.Sprintf("Cannot %s \"%s\" value %s DataColumn \"%s\" for table \"%s\", column data type is \"%s\"", action, DataTypeEnum(targetType).Name(), preposition, column.Name(), dr.parent.Name(), column.Type().Name()))
	}

	if !read && column.Computed() {
		panic("Cannot assign value to DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\", column is computed with an expression")
	}

	return column
}

// func (dr *DataRow) getExpressionTree(column *DataColumn) (*ExpressionTree, error) {
// 	columnIndex := column.Index()

// 	if dr.values[columnIndex] == nil {
// 		dataTable := column.Parent()
// 		parser := NewFilterExpressionParser(column.Expression())

// 		parser.SetDataSet(dataTable.Parent())
// 		parser.SetPrimaryTableName(dataTable.Name())
// 		parser.SetTrackFilteredSignalIDs(false)
// 		parser.SetTrackFilteredRows(false)

// 		expressionTrees := parser.GetExpressionTrees()

// 		if len(expressionTrees) == 0 {
// 			return nil, errors.New("Expression defined for computed DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\" cannot produce a value")
// 		}

// 		dr.values[columnIndex] = parser
// 		return expressionTrees[0]
// 	}

// 	return dr.values[columnIndex].(*FilterExpressionParser).GetExpressionTrees()[0]
// }

func (dr *DataRow) getComputedValue(column *DataColumn) interface{} {
	// TODO: Evaluate expression using ANTLR grammar:
	// https://github.com/sttp/cppapi/blob/master/src/lib/filterexpressions/FilterExpressionSyntax.g4
	// expressionTree, err := dr.getExpressionTree(column)
	// sourceValue = expressionTree.Evaluate()

	// switch sourceValue.ValueType {
	// case ExpressionValueType.Boolean:
	// }

	return nil
}

// Value reads the record value at the specified columnIndex.
func (dr *DataRow) Value(columnIndex int) interface{} {
	column := dr.validateColumnType(columnIndex, -1, true)

	if column.Computed() {
		return dr.getComputedValue(column)
	}

	return dr.values[columnIndex]
}

// ValueByName reads the record value for the specified columnName.
func (dr *DataRow) ValueByName(columnName string) interface{} {
	return dr.values[dr.getColumnIndex(columnName)]
}

// SetValue assigns the record value at the specified columnIndex.
func (dr *DataRow) SetValue(columnIndex int, value interface{}) {
	dr.validateColumnType(columnIndex, -1, false)
	dr.values[columnIndex] = value
}

// SetValueByName assins the record value for the specified columnName.
func (dr *DataRow) SetValueByName(columnName string, value interface{}) {
	dr.SetValue(dr.getColumnIndex(columnName), value)
}

// ValueAsString gets the record value at the specified columnIndex as a string.
func (dr *DataRow) ValueAsString(columnIndex int) string {
	column := dr.validateColumnType(columnIndex, int(DataType.String), true)

	if column.Computed() {
		return dr.getComputedValue(column).(string)
	}

	return dr.values[columnIndex].(string)
}

// ValueAsStringByName gets the record value for the specified columnName as a string.
func (dr *DataRow) ValueAsStringByName(columnName string) string {
	return dr.ValueAsString(dr.getColumnIndex(columnName))
}

// ValueAsBool gets the record value at the specified columnIndex as a bool.
func (dr *DataRow) ValueAsBool(columnIndex int) bool {
	column := dr.validateColumnType(columnIndex, int(DataType.Boolean), true)

	if column.Computed() {
		return dr.getComputedValue(column).(bool)
	}

	return dr.values[columnIndex].(bool)
}

// ValueAsBoolByName gets the record value for the specified columnName as a bool.
func (dr *DataRow) ValueAsBoolByName(columnName string) bool {
	return dr.ValueAsBool(dr.getColumnIndex(columnName))
}

// ValueAsDateTime gets the record value at the specified columnIndex as a time.Time.
func (dr *DataRow) ValueAsDateTime(columnIndex int) time.Time {
	column := dr.validateColumnType(columnIndex, int(DataType.DateTime), true)

	if column.Computed() {
		return dr.getComputedValue(column).(time.Time)
	}

	return dr.values[columnIndex].(time.Time)
}

// ValueAsDateTimeByName gets the record value for the specified columnName as a time.Time.
func (dr *DataRow) ValueAsDateTimeByName(columnName string) time.Time {
	return dr.ValueAsDateTime(dr.getColumnIndex(columnName))
}

// TODO: Add remaining ValueAs<Type> methods
