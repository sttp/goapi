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
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/sttp/goapi/sttp/guid"
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
		return -1, errors.New("Column name \"" + columnName + "\" was not found in table \"" + dr.parent.Name() + "\"")
	}

	return column.Index(), nil
}

func (dr *DataRow) validateColumnType(columnIndex int, targetType int, read bool) (*DataColumn, error) {
	column := dr.parent.Column(columnIndex)

	if column == nil {
		return nil, errors.New("Column index " + strconv.Itoa(columnIndex) + " is out of range for table \"" + dr.parent.Name() + "\"")
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
// 			return nil, errors.New("Expression defined for computed DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\" cannot produce a value")
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

// GetValue reads the record value at the specified columnIndex as a string,
// if columnIndex is out of range, an empty string will be returned.
func (dr *DataRow) GetValue(columnIndex int) string {
	column := dr.parent.Column(columnIndex)

	if column == nil {
		return ""
	}

	index := column.Index()

	switch column.Type() {
	case DataType.String:
		value, err := dr.ValueAsString(index)

		if err != nil {
			return ""
		}

		return value
	case DataType.Boolean:
		value, err := dr.ValueAsBool(index)

		if err != nil {
			return ""
		}

		return strconv.FormatBool(value)
	case DataType.DateTime:
		value, err := dr.ValueAsDateTime(index)

		if err != nil {
			return ""
		}
		return value.String()
	case DataType.Single:
		value, err := dr.ValueAsSingle(index)

		if err != nil {
			return ""
		}
		return strconv.FormatFloat(float64(value), 'f', 6, 32)
	case DataType.Decimal:
		fallthrough
	case DataType.Double:
		value, err := dr.ValueAsDouble(index)

		if err != nil {
			return ""
		}

		return strconv.FormatFloat(value, 'f', 6, 64)
	case DataType.Guid:
		value, err := dr.ValueAsGuid(index)

		if err != nil {
			return ""
		}

		return value.String()
	case DataType.Int8:
		value, err := dr.ValueAsInt8(index)

		if err != nil {
			return ""
		}

		return strconv.FormatInt(int64(value), 10)
	case DataType.Int16:
		value, err := dr.ValueAsInt16(index)

		if err != nil {
			return ""
		}

		return strconv.FormatInt(int64(value), 10)
	case DataType.Int32:
		value, err := dr.ValueAsInt32(index)

		if err != nil {
			return ""
		}

		return strconv.FormatInt(int64(value), 10)
	case DataType.Int64:
		value, err := dr.ValueAsInt64(index)

		if err != nil {
			return ""
		}

		return strconv.FormatInt(value, 10)
	case DataType.UInt8:
		value, err := dr.ValueAsUInt8(index)

		if err != nil {
			return ""
		}

		return strconv.FormatUint(uint64(value), 10)
	case DataType.UInt16:
		value, err := dr.ValueAsUInt16(index)

		if err != nil {
			return ""
		}

		return strconv.FormatUint(uint64(value), 10)
	case DataType.UInt32:
		value, err := dr.ValueAsUInt32(index)

		if err != nil {
			return ""
		}

		return strconv.FormatUint(uint64(value), 10)
	case DataType.UInt64:
		value, err := dr.ValueAsUInt64(index)

		if err != nil {
			return ""
		}

		return strconv.FormatUint(value, 10)
	default:
		return ""
	}
}

// GetValue reads the record value at the specified columnName as a string,
// if columnName is not found, an empty string will be returned.
func (dr *DataRow) GetValueByName(columnName string) string {
	column := dr.parent.ColumnByName(columnName)

	if column == nil {
		return ""
	}

	return dr.GetValue(column.Index())
}

// ValueAsString gets the record value at the specified columnIndex cast as a string.
func (dr *DataRow) ValueAsString(columnIndex int) (string, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.String), true)

	if err != nil {
		return "", err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return "", err
		}

		return value.(string), nil
	}

	return dr.values[columnIndex].(string), nil
}

// ValueAsStringByName gets the record value for the specified columnName cast as a string.
func (dr *DataRow) ValueAsStringByName(columnName string) (string, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return "", err
	}

	return dr.ValueAsString(index)
}

// ValueAsBool gets the record value at the specified columnIndex cast as a bool.
func (dr *DataRow) ValueAsBool(columnIndex int) (bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Boolean), true)

	if err != nil {
		return false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return false, err
		}

		return value.(bool), nil
	}

	return dr.values[columnIndex].(bool), nil
}

// ValueAsBoolByName gets the record value for the specified columnName cast as a bool.
func (dr *DataRow) ValueAsBoolByName(columnName string) (bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return false, err
	}

	return dr.ValueAsBool(index)
}

// ValueAsDateTime gets the record value at the specified columnIndex cast as a time.Time.
func (dr *DataRow) ValueAsDateTime(columnIndex int) (time.Time, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.DateTime), true)

	if err != nil {
		return time.Time{}, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return time.Time{}, err
		}

		return value.(time.Time), nil
	}

	return dr.values[columnIndex].(time.Time), nil
}

// ValueAsDateTimeByName gets the record value for the specified columnName cast as a time.Time.
func (dr *DataRow) ValueAsDateTimeByName(columnName string) (time.Time, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return time.Time{}, err
	}

	return dr.ValueAsDateTime(index)
}

// ValueAsSingle gets the record value at the specified columnIndex cast as a float32.
func (dr *DataRow) ValueAsSingle(columnIndex int) (float32, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Single), true)

	if err != nil {
		return 0.0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0.0, err
		}

		return value.(float32), nil
	}

	return dr.values[columnIndex].(float32), nil
}

// ValueAsSingleByName gets the record value for the specified columnName cast as a float32.
func (dr *DataRow) ValueAsSingleByName(columnName string) (float32, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0.0, err
	}

	return dr.ValueAsSingle(index)
}

// ValueAsDouble gets the record value at the specified columnIndex cast as a float64.
func (dr *DataRow) ValueAsDouble(columnIndex int) (float64, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Double), true)

	if err != nil {
		return 0.0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0.0, err
		}

		return value.(float64), nil
	}

	return dr.values[columnIndex].(float64), nil
}

// ValueAsDoubleByName gets the record value for the specified columnName cast as a float64.
func (dr *DataRow) ValueAsDoubleByName(columnName string) (float64, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0.0, err
	}

	return dr.ValueAsDouble(index)
}

// ValueAsDecimal gets the record value at the specified columnIndex cast as a float64.
func (dr *DataRow) ValueAsDecimal(columnIndex int) (float64, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Decimal), true)

	if err != nil {
		return 0.0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0.0, err
		}

		return value.(float64), nil
	}

	return dr.values[columnIndex].(float64), nil
}

// ValueAsDecimalByName gets the record value for the specified columnName cast as a float64.
func (dr *DataRow) ValueAsDecimalByName(columnName string) (float64, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0.0, err
	}

	return dr.ValueAsDecimal(index)
}

// ValueAsGuid gets the record value at the specified columnIndex cast as a guid.Guid.
func (dr *DataRow) ValueAsGuid(columnIndex int) (guid.Guid, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Guid), true)

	if err != nil {
		return guid.Guid{}, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return guid.Guid{}, err
		}

		return value.(guid.Guid), nil
	}

	return dr.values[columnIndex].(guid.Guid), nil
}

// ValueAsGuidByName gets the record value for the specified columnName cast as a guid.Guid.
func (dr *DataRow) ValueAsGuidByName(columnName string) (guid.Guid, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return guid.Guid{}, err
	}

	return dr.ValueAsGuid(index)
}

// ValueAsInt8 gets the record value at the specified columnIndex cast as a int8.
func (dr *DataRow) ValueAsInt8(columnIndex int) (int8, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int8), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(int8), nil
	}

	return dr.values[columnIndex].(int8), nil
}

// ValueAsInt8ByName gets the record value for the specified columnName cast as a int8.
func (dr *DataRow) ValueAsInt8ByName(columnName string) (int8, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsInt8(index)
}

// ValueAsInt16 gets the record value at the specified columnIndex cast as a int16.
func (dr *DataRow) ValueAsInt16(columnIndex int) (int16, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int16), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(int16), nil
	}

	return dr.values[columnIndex].(int16), nil
}

// ValueAsInt16ByName gets the record value for the specified columnName cast as a int16.
func (dr *DataRow) ValueAsInt16ByName(columnName string) (int16, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsInt16(index)
}

// ValueAsInt32 gets the record value at the specified columnIndex cast as a int32.
func (dr *DataRow) ValueAsInt32(columnIndex int) (int32, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int32), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(int32), nil
	}

	return dr.values[columnIndex].(int32), nil
}

// ValueAsInt32ByName gets the record value for the specified columnName cast as a int32.
func (dr *DataRow) ValueAsInt32ByName(columnName string) (int32, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsInt32(index)
}

// ValueAsInt64 gets the record value at the specified columnIndex cast as a int64.
func (dr *DataRow) ValueAsInt64(columnIndex int) (int64, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Int64), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(int64), nil
	}

	return dr.values[columnIndex].(int64), nil
}

// ValueAsInt64ByName gets the record value for the specified columnName cast as a int64.
func (dr *DataRow) ValueAsInt64ByName(columnName string) (int64, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsInt64(index)
}

// ValueAsUInt8 gets the record value at the specified columnIndex cast as a uint8.
func (dr *DataRow) ValueAsUInt8(columnIndex int) (uint8, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt8), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(uint8), nil
	}

	return dr.values[columnIndex].(uint8), nil
}

// ValueAsUInt8ByName gets the record value for the specified columnName cast as a uint8.
func (dr *DataRow) ValueAsUInt8ByName(columnName string) (uint8, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsUInt8(index)
}

// ValueAsUInt16 gets the record value at the specified columnIndex cast as a uint16.
func (dr *DataRow) ValueAsUInt16(columnIndex int) (uint16, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt16), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(uint16), nil
	}

	return dr.values[columnIndex].(uint16), nil
}

// ValueAsUInt16ByName gets the record value for the specified columnName cast as a uint16.
func (dr *DataRow) ValueAsUInt16ByName(columnName string) (uint16, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsUInt16(index)
}

// ValueAsUInt32 gets the record value at the specified columnIndex cast as a uint32.
func (dr *DataRow) ValueAsUInt32(columnIndex int) (uint32, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt32), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(uint32), nil
	}

	return dr.values[columnIndex].(uint32), nil
}

// ValueAsUInt32ByName gets the record value for the specified columnName cast as a uint32.
func (dr *DataRow) ValueAsUInt32ByName(columnName string) (uint32, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsUInt32(index)
}

// ValueAsUInt64 gets the record value at the specified columnIndex cast as a uint64.
func (dr *DataRow) ValueAsUInt64(columnIndex int) (uint64, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.UInt64), true)

	if err != nil {
		return 0, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return 0, err
		}

		return value.(uint64), nil
	}

	return dr.values[columnIndex].(uint64), nil
}

// ValueAsUInt64ByName gets the record value for the specified columnName cast as a uint64.
func (dr *DataRow) ValueAsUInt64ByName(columnName string) (uint64, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return 0, err
	}

	return dr.ValueAsUInt64(index)
}
