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

package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/ticks"
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

func (dr *DataRow) getColumnIndex(columnName string) (int, error) {
	column := dr.parent.ColumnByName(columnName)

	if column == nil {
		return -1, errors.New("column name \"" + columnName + "\" was not found in table \"" + dr.parent.Name() + "\"")
	}

	return column.Index(), nil
}

func (dr *DataRow) validateColumnType(columnIndex int, targetType int, read bool) (*DataColumn, error) {
	column := dr.parent.Column(columnIndex)

	if column == nil {
		return nil, errors.New("column index " + strconv.Itoa(columnIndex) + " is out of range for table \"" + dr.parent.Name() + "\"")
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

		return nil, fmt.Errorf("cannot %s \"%s\" value %s DataColumn \"%s\" for table \"%s\", column data type is \"%s\"", action, DataTypeEnum(targetType).String(), preposition, column.Name(), dr.parent.Name(), column.Type().String())
	}

	if !read && column.Computed() {
		return nil, errors.New("cannot assign value to DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\", column is computed with an expression")
	}

	return column, nil
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
// 			return nil, errors.New("expression defined for computed DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\" cannot produce a value")
// 		}

// 		dr.values[columnIndex] = parser
// 		return expressionTrees[0]
// 	}

// 	return dr.values[columnIndex].(*FilterExpressionParser).GetExpressionTrees()[0]
// }

func (dr *DataRow) getComputedValue(column *DataColumn) (interface{}, error) {
	// TODO: Evaluate expression using ANTLR grammar:
	// https://github.com/sttp/cppapi/blob/master/src/lib/filterexpressions/FilterExpressionSyntax.g4
	// expressionTree, err := dr.getExpressionTree(column)
	// sourceValue = expressionTree.Evaluate()

	// switch sourceValue.ValueType {
	// case ExpressionValueType.Boolean:
	// }

	return nil, nil
}

// Value reads the record value at the specified columnIndex.
func (dr *DataRow) Value(columnIndex int) (interface{}, error) {
	column, err := dr.validateColumnType(columnIndex, -1, true)

	if err != nil {
		return nil, err
	}

	if column.Computed() {
		return dr.getComputedValue(column)
	}

	return dr.values[columnIndex], nil
}

// ValueByName reads the record value for the specified columnName.
func (dr *DataRow) ValueByName(columnName string) (interface{}, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return nil, err
	}

	return dr.values[index], nil
}

// SetValue assigns the record value at the specified columnIndex.
func (dr *DataRow) SetValue(columnIndex int, value interface{}) error {
	_, err := dr.validateColumnType(columnIndex, -1, false)

	if err != nil {
		return err
	}

	dr.values[columnIndex] = value
	return nil
}

// SetValueByName assigns the record value for the specified columnName.
func (dr *DataRow) SetValueByName(columnName string, value interface{}) error {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return err
	}

	return dr.SetValue(index, value)
}

// ColumnValueAsString reads the record value for the specified data column
// converted to a string. For any errors, an empty string will be returned.
func (dr *DataRow) ColumnValueAsString(column *DataColumn) string {
	if column == nil {
		return ""
	}

	checkState := func(null bool, err error) (bool, string) {
		if err != nil {
			return true, ""
		}

		if null {
			return true, "<NULL>"
		}

		return false, ""
	}

	index := column.Index()

	switch column.Type() {
	case DataType.String:
		value, null, err := dr.StringValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return value
	case DataType.Boolean:
		value, null, err := dr.BooleanValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatBool(value)
	case DataType.DateTime:
		value, null, err := dr.DateTimeValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return value.Format(ticks.TimeFormat)
	case DataType.Single:
		value, null, err := dr.SingleValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatFloat(float64(value), 'f', 6, 32)
	case DataType.Double:
		value, null, err := dr.DoubleValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatFloat(value, 'f', 6, 64)
	case DataType.Decimal:
		value, null, err := dr.DecimalValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatFloat(value, 'f', 6, 64)
	case DataType.Guid:
		value, null, err := dr.GuidValue(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return value.String()
	case DataType.Int8:
		value, null, err := dr.Int8Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatInt(int64(value), 10)
	case DataType.Int16:
		value, null, err := dr.Int16Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatInt(int64(value), 10)
	case DataType.Int32:
		value, null, err := dr.Int32Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatInt(int64(value), 10)
	case DataType.Int64:
		value, null, err := dr.Int64Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatInt(value, 10)
	case DataType.UInt8:
		value, null, err := dr.UInt8Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatUint(uint64(value), 10)
	case DataType.UInt16:
		value, null, err := dr.UInt16Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatUint(uint64(value), 10)
	case DataType.UInt32:
		value, null, err := dr.UInt32Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatUint(uint64(value), 10)
	case DataType.UInt64:
		value, null, err := dr.UInt64Value(index)

		if invalid, result := checkState(null, err); invalid {
			return result
		}

		return strconv.FormatUint(value, 10)
	default:
		return ""
	}
}

// ValueAsString reads the record value at the specified columnIndex converted to a string.
// For columnIndex out of range or any other errors, an empty string will be returned.
func (dr *DataRow) ValueAsString(columnIndex int) string {
	return dr.ColumnValueAsString(dr.parent.Column(columnIndex))
}

// ValueAsStringByName reads the record value for the specified columnName converted to a string.
// For columnName not found or any other errors, an empty string will be returned.
func (dr *DataRow) ValueAsStringByName(columnName string) string {
	return dr.ColumnValueAsString(dr.parent.ColumnByName(columnName))
}

// StringValue gets the record value at the specified columnIndex cast as a string.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.String.
func (dr *DataRow) StringValue(columnIndex int) (string, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.String), true)

	if err != nil {
		return "", false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return "", false, err
		}

		if value == nil {
			return "", true, nil
		}

		return value.(string), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return "", true, nil
	}

	return value.(string), false, nil
}

// StringValueByName gets the record value for the specified columnName cast as a string.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.String.
func (dr *DataRow) StringValueByName(columnName string) (string, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return "", false, err
	}

	return dr.StringValue(index)
}

// BooleanValue gets the record value at the specified columnIndex cast as a bool.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Boolean.
func (dr *DataRow) BooleanValue(columnIndex int) (bool, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Boolean), true)

	if err != nil {
		return false, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return false, false, err
		}

		if value == nil {
			return false, true, nil
		}

		return value.(bool), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return false, true, nil
	}

	return value.(bool), false, nil
}

// BooleanValueByName gets the record value for the specified columnName cast as a bool.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Boolean.
func (dr *DataRow) BooleanValueByName(columnName string) (bool, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return false, false, err
	}

	return dr.BooleanValue(index)
}

// DateTimeValue gets the record value at the specified columnIndex cast as a time.Time.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.DateTime.
func (dr *DataRow) DateTimeValue(columnIndex int) (time.Time, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.DateTime), true)

	if err != nil {
		return time.Time{}, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return time.Time{}, false, err
		}

		if value == nil {
			return time.Time{}, true, nil
		}

		return value.(time.Time), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return time.Time{}, true, nil
	}

	return value.(time.Time), false, nil
}

// DateTimeValueByName gets the record value for the specified columnName cast as a time.Time.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.DateTime.
func (dr *DataRow) DateTimeValueByName(columnName string) (time.Time, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return time.Time{}, false, err
	}

	return dr.DateTimeValue(index)
}

// SingleValue gets the record value at the specified columnIndex cast as a float32.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Single.
func (dr *DataRow) SingleValue(columnIndex int) (float32, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Single), true)

	if err != nil {
		return 0.0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0.0, false, err
		}

		if value == nil {
			return 0.0, true, nil
		}

		return value.(float32), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0.0, true, nil
	}

	return value.(float32), false, nil
}

// SingleValueByName gets the record value for the specified columnName cast as a float32.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Single.
func (dr *DataRow) SingleValueByName(columnName string) (float32, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0.0, false, err
	}

	return dr.SingleValue(index)
}

// DoubleValue gets the record value at the specified columnIndex cast as a float64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Double.
func (dr *DataRow) DoubleValue(columnIndex int) (float64, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Double), true)

	if err != nil {
		return 0.0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0.0, false, err
		}

		if value == nil {
			return 0.0, true, nil
		}

		return value.(float64), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0.0, true, nil
	}

	return value.(float64), false, nil
}

// DoubleValueByName gets the record value for the specified columnName cast as a float64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Double.
func (dr *DataRow) DoubleValueByName(columnName string) (float64, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0.0, false, err
	}

	return dr.DoubleValue(index)
}

// DecimalValue gets the record value at the specified columnIndex cast as a float64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Decimal.
func (dr *DataRow) DecimalValue(columnIndex int) (float64, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Decimal), true)

	if err != nil {
		return 0.0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0.0, false, err
		}

		if value == nil {
			return 0.0, true, nil
		}

		return value.(float64), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0.0, true, nil
	}

	return value.(float64), false, nil
}

// DecimalValueByName gets the record value for the specified columnName cast as a float64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Decimal.
func (dr *DataRow) DecimalValueByName(columnName string) (float64, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0.0, false, err
	}

	return dr.DecimalValue(index)
}

// GuidValue gets the record value at the specified columnIndex cast as a guid.Guid.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Guid.
func (dr *DataRow) GuidValue(columnIndex int) (guid.Guid, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Guid), true)

	if err != nil {
		return guid.Guid{}, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return guid.Guid{}, false, err
		}

		if value == nil {
			return guid.Guid{}, true, nil
		}

		return value.(guid.Guid), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return guid.Guid{}, true, nil
	}

	return value.(guid.Guid), false, nil
}

// GuidValueByName gets the record value for the specified columnName cast as a guid.Guid.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Guid.
func (dr *DataRow) GuidValueByName(columnName string) (guid.Guid, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return guid.Guid{}, false, err
	}

	return dr.GuidValue(index)
}

// Int8Value gets the record value at the specified columnIndex cast as an int8.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int8.
func (dr *DataRow) Int8Value(columnIndex int) (int8, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int8), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(int8), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(int8), false, nil
}

// Int8ValueByName gets the record value for the specified columnName cast as an int8.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int8.
func (dr *DataRow) Int8ValueByName(columnName string) (int8, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.Int8Value(index)
}

// Int16Value gets the record value at the specified columnIndex cast as an int16.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int16.
func (dr *DataRow) Int16Value(columnIndex int) (int16, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int16), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(int16), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(int16), false, nil
}

// Int16ValueByName gets the record value for the specified columnName cast as an int16.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int16.
func (dr *DataRow) Int16ValueByName(columnName string) (int16, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.Int16Value(index)
}

// Int32Value gets the record value at the specified columnIndex cast as an int32.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int32.
func (dr *DataRow) Int32Value(columnIndex int) (int32, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int32), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(int32), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(int32), false, nil
}

// Int32ValueByName gets the record value for the specified columnName cast as an int32.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int32.
func (dr *DataRow) Int32ValueByName(columnName string) (int32, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.Int32Value(index)
}

// Int64Value gets the record value at the specified columnIndex cast as an int64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int64.
func (dr *DataRow) Int64Value(columnIndex int) (int64, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int64), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(int64), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(int64), false, nil
}

// Int64ValueByName gets the record value for the specified columnName cast as an int64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Int64.
func (dr *DataRow) Int64ValueByName(columnName string) (int64, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.Int64Value(index)
}

// UInt8Value gets the record value at the specified columnIndex cast as an uint8.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt8.
func (dr *DataRow) UInt8Value(columnIndex int) (uint8, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt8), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(uint8), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(uint8), false, nil
}

// UInt8ValueByName gets the record value for the specified columnName cast as an uint8.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt8.
func (dr *DataRow) UInt8ValueByName(columnName string) (uint8, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.UInt8Value(index)
}

// UInt16Value gets the record value at the specified columnIndex cast as an uint16.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt16.
func (dr *DataRow) UInt16Value(columnIndex int) (uint16, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt16), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(uint16), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(uint16), false, nil
}

// UInt16ValueByName gets the record value for the specified columnName cast as an uint16.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt16.
func (dr *DataRow) UInt16ValueByName(columnName string) (uint16, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.UInt16Value(index)
}

// UInt32Value gets the record value at the specified columnIndex cast as an uint32.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt32.
func (dr *DataRow) UInt32Value(columnIndex int) (uint32, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt32), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(uint32), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(uint32), false, nil
}

// UInt32ValueByName gets the record value for the specified columnName cast as an uint32.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt32.
func (dr *DataRow) UInt32ValueByName(columnName string) (uint32, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.UInt32Value(index)
}

// UInt64Value gets the record value at the specified columnIndex cast as an uint64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt64.
func (dr *DataRow) UInt64Value(columnIndex int) (uint64, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt64), true)

	if err != nil {
		return 0, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, false, err
		}

		if value == nil {
			return 0, true, nil
		}

		return value.(uint64), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return 0, true, nil
	}

	return value.(uint64), false, nil
}

// UInt64ValueByName gets the record value for the specified columnName cast as an uint64.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.UInt64.
func (dr *DataRow) UInt64ValueByName(columnName string) (uint64, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, false, err
	}

	return dr.UInt64Value(index)
}

// String get a representation of the DataRow as a string.
func (dr *DataRow) String() string {
	var image strings.Builder

	image.WriteRune('[')

	for i := 0; i < dr.parent.ColumnCount(); i++ {
		if i > 0 {
			image.WriteString(", ")
		}

		stringColumn := dr.parent.Column(i).Type() == DataType.String

		if stringColumn {
			image.WriteRune('"')
		}

		image.WriteString(dr.ValueAsString(i))

		if stringColumn {
			image.WriteRune('"')
		}
	}

	image.WriteRune(']')

	return image.String()
}
