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

	"github.com/araddon/dateparse"
	"github.com/shopspring/decimal"
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
		return -1, errors.New("column name \"" + columnName + "\" was not found in table \"" + dr.parent.Name() + "\"")
	}

	return column.Index(), nil
}

func (dr *DataRow) validateColumnType(columnIndex, targetType int, read bool) (*DataColumn, error) {
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

func (dr *DataRow) expressionTree(column *DataColumn) (*ExpressionTree, error) {
	columnIndex := column.Index()

	if dr.values[columnIndex] == nil {
		dataTable := column.Parent()
		parser := NewFilterExpressionParser(column.Expression(), true)

		parser.DataSet = dataTable.Parent()
		parser.PrimaryTableName = dataTable.Name()
		parser.TrackFilteredSignalIDs = false
		parser.TrackFilteredRows = false

		expressionTrees, err := parser.ExpressionTrees()

		if err != nil {
			return nil, errors.New("failed to parse expression defined for computed DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\": " + err.Error())
		}

		if len(expressionTrees) == 0 {
			return nil, errors.New("expression defined for computed DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\" cannot produce a value")
		}

		dr.values[columnIndex] = expressionTrees[0]
		return expressionTrees[0], nil
	}

	return dr.values[columnIndex].(*ExpressionTree), nil
}

func (dr *DataRow) getComputedValue(column *DataColumn) (interface{}, error) {
	expressionTree, err := dr.expressionTree(column)

	if err != nil {
		return nil, err
	}

	sourceValue, err := expressionTree.Evaluate(dr)

	if err != nil {
		return nil, errors.New("failed to evaluate expression defined for computed DataColumn \"" + column.Name() + "\" for table \"" + dr.parent.Name() + "\": " + err.Error())
	}

	targetType := column.Type()

	switch sourceValue.ValueType() {
	case ExpressionValueType.Boolean:
		return convertFromBoolean(sourceValue.booleanValue(), targetType)
	case ExpressionValueType.Int32:
		return convertFromInt32(sourceValue.int32Value(), targetType)
	case ExpressionValueType.Int64:
		return convertFromInt64(sourceValue.int64Value(), targetType)
	case ExpressionValueType.Decimal:
		return convertFromDecimal(sourceValue.decimalValue(), targetType)
	case ExpressionValueType.Double:
		return convertFromDouble(sourceValue.doubleValue(), targetType)
	case ExpressionValueType.String:
		return convertFromString(sourceValue.stringValue(), targetType)
	case ExpressionValueType.Guid:
		return convertFromGuid(sourceValue.guidValue(), targetType)
	case ExpressionValueType.DateTime:
		return convertFromDateTime(sourceValue.dateTimeValue(), targetType)
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

//gocyclo:ignore
func convertFromBoolean(value bool, targetType DataTypeEnum) (interface{}, error) {
	var valueAsInt int

	if value {
		valueAsInt = 1
	}

	switch targetType {
	case DataType.String:
		return strconv.FormatBool(value), nil
	case DataType.Boolean:
		return value, nil
	case DataType.Single:
		return float32(valueAsInt), nil
	case DataType.Double:
		return float64(valueAsInt), nil
	case DataType.Decimal:
		return decimal.NewFromInt(int64(valueAsInt)), nil
	case DataType.Int8:
		return int8(valueAsInt), nil
	case DataType.Int16:
		return int16(valueAsInt), nil
	case DataType.Int32:
		return int32(valueAsInt), nil
	case DataType.Int64:
		return int64(valueAsInt), nil
	case DataType.UInt8:
		return uint8(valueAsInt), nil
	case DataType.UInt16:
		return uint16(valueAsInt), nil
	case DataType.UInt32:
		return uint32(valueAsInt), nil
	case DataType.UInt64:
		return uint64(valueAsInt), nil
	case DataType.DateTime:
		fallthrough
	case DataType.Guid:
		return nil, errors.New("cannot convert \"Boolean\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromInt32(value int32, targetType DataTypeEnum) (interface{}, error) {
	switch targetType {
	case DataType.String:
		return strconv.FormatInt(int64(value), 10), nil
	case DataType.Boolean:
		return value != 0, nil
	case DataType.Single:
		return float32(value), nil
	case DataType.Double:
		return float64(value), nil
	case DataType.Decimal:
		return decimal.NewFromInt32(value), nil
	case DataType.Int8:
		return int8(value), nil
	case DataType.Int16:
		return int16(value), nil
	case DataType.Int32:
		return value, nil
	case DataType.Int64:
		return int64(value), nil
	case DataType.UInt8:
		return uint8(value), nil
	case DataType.UInt16:
		return uint16(value), nil
	case DataType.UInt32:
		return uint32(value), nil
	case DataType.UInt64:
		return uint64(value), nil
	case DataType.DateTime:
		fallthrough
	case DataType.Guid:
		return nil, errors.New("cannot convert \"Int32\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromInt64(value int64, targetType DataTypeEnum) (interface{}, error) {
	switch targetType {
	case DataType.String:
		return strconv.FormatInt(value, 10), nil
	case DataType.Boolean:
		return value != 0, nil
	case DataType.Single:
		return float32(value), nil
	case DataType.Double:
		return float64(value), nil
	case DataType.Decimal:
		return decimal.NewFromInt(value), nil
	case DataType.Int8:
		return int8(value), nil
	case DataType.Int16:
		return int16(value), nil
	case DataType.Int32:
		return int32(value), nil
	case DataType.Int64:
		return value, nil
	case DataType.UInt8:
		return uint8(value), nil
	case DataType.UInt16:
		return uint16(value), nil
	case DataType.UInt32:
		return uint32(value), nil
	case DataType.UInt64:
		return uint64(value), nil
	case DataType.DateTime:
		fallthrough
	case DataType.Guid:
		return nil, errors.New("cannot convert \"Int64\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromDecimal(value decimal.Decimal, targetType DataTypeEnum) (interface{}, error) {
	switch targetType {
	case DataType.String:
		return value.String(), nil
	case DataType.Boolean:
		return !value.Equal(decimal.Zero), nil
	case DataType.Single:
		f64, _ := value.Float64()
		return float32(f64), nil
	case DataType.Double:
		f64, _ := value.Float64()
		return f64, nil
	case DataType.Decimal:
		return value, nil
	case DataType.Int8:
		return int8(value.IntPart()), nil
	case DataType.Int16:
		return int16(value.IntPart()), nil
	case DataType.Int32:
		return int32(value.IntPart()), nil
	case DataType.Int64:
		return value.IntPart(), nil
	case DataType.UInt8:
		return uint8(value.IntPart()), nil
	case DataType.UInt16:
		return uint16(value.IntPart()), nil
	case DataType.UInt32:
		return uint32(value.IntPart()), nil
	case DataType.UInt64:
		return uint64(value.IntPart()), nil
	case DataType.DateTime:
		fallthrough
	case DataType.Guid:
		return nil, errors.New("cannot convert \"Decimal\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromDouble(value float64, targetType DataTypeEnum) (interface{}, error) {
	switch targetType {
	case DataType.String:
		return strconv.FormatFloat(value, 'f', 6, 64), nil
	case DataType.Boolean:
		return value != 0.0, nil
	case DataType.Single:
		return float32(value), nil
	case DataType.Double:
		return value, nil
	case DataType.Decimal:
		return decimal.NewFromFloat(value), nil
	case DataType.Int8:
		return int8(value), nil
	case DataType.Int16:
		return int16(value), nil
	case DataType.Int32:
		return int32(value), nil
	case DataType.Int64:
		return int64(value), nil
	case DataType.UInt8:
		return uint8(value), nil
	case DataType.UInt16:
		return uint16(value), nil
	case DataType.UInt32:
		return uint32(value), nil
	case DataType.UInt64:
		return uint64(value), nil
	case DataType.DateTime:
		fallthrough
	case DataType.Guid:
		return nil, errors.New("cannot convert \"Double\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromString(value string, targetType DataTypeEnum) (interface{}, error) {
	switch targetType {
	case DataType.String:
		return value, nil
	case DataType.Boolean:
		return strconv.ParseBool(value)
	case DataType.DateTime:
		return dateparse.ParseAny(value)
	case DataType.Single:
		f64, err := strconv.ParseFloat(value, 64)
		return float32(f64), err
	case DataType.Double:
		return strconv.ParseFloat(value, 64)
	case DataType.Decimal:
		return decimal.NewFromString(value)
	case DataType.Guid:
		return guid.Parse(value)
	case DataType.Int8:
		i, err := strconv.ParseInt(value, 10, 8)
		return int8(i), err
	case DataType.Int16:
		i, err := strconv.ParseInt(value, 10, 16)
		return int16(i), err
	case DataType.Int32:
		i, err := strconv.ParseInt(value, 10, 32)
		return int32(i), err
	case DataType.Int64:
		i, err := strconv.ParseInt(value, 10, 64)
		return i, err
	case DataType.UInt8:
		ui, err := strconv.ParseUint(value, 10, 8)
		return uint8(ui), err
	case DataType.UInt16:
		ui, err := strconv.ParseUint(value, 10, 16)
		return uint16(ui), err
	case DataType.UInt32:
		ui, err := strconv.ParseUint(value, 10, 32)
		return uint32(ui), err
	case DataType.UInt64:
		ui, err := strconv.ParseUint(value, 10, 64)
		return ui, err
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromGuid(value guid.Guid, targetType DataTypeEnum) (interface{}, error) {
	switch targetType {
	case DataType.String:
		return value.String(), nil
	case DataType.Guid:
		return value, nil
	case DataType.Boolean:
		fallthrough
	case DataType.DateTime:
		fallthrough
	case DataType.Single:
		fallthrough
	case DataType.Double:
		fallthrough
	case DataType.Decimal:
		fallthrough
	case DataType.Int8:
		fallthrough
	case DataType.Int16:
		fallthrough
	case DataType.Int32:
		fallthrough
	case DataType.Int64:
		fallthrough
	case DataType.UInt8:
		fallthrough
	case DataType.UInt16:
		fallthrough
	case DataType.UInt32:
		fallthrough
	case DataType.UInt64:
		return nil, errors.New("cannot convert \"Guid\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
}

//gocyclo:ignore
func convertFromDateTime(value time.Time, targetType DataTypeEnum) (interface{}, error) {
	seconds := value.Unix()

	switch targetType {
	case DataType.String:
		return value.Format(DateTimeFormat), nil
	case DataType.Boolean:
		return seconds == 0, nil
	case DataType.DateTime:
		return value, nil
	case DataType.Single:
		return float32(seconds), nil
	case DataType.Double:
		return float64(seconds), nil
	case DataType.Decimal:
		return decimal.NewFromInt(seconds), nil
	case DataType.Int8:
		return int8(seconds), nil
	case DataType.Int16:
		return int16(seconds), nil
	case DataType.Int32:
		return int32(seconds), nil
	case DataType.Int64:
		return seconds, nil
	case DataType.UInt8:
		return uint8(seconds), nil
	case DataType.UInt16:
		return uint16(seconds), nil
	case DataType.UInt32:
		return uint32(seconds), nil
	case DataType.UInt64:
		return uint64(seconds), nil
	case DataType.Guid:
		return nil, errors.New("cannot convert \"DateTime\" expression value to \"" + targetType.String() + "\" column")
	default:
		return nil, errors.New("unexpected column data type encountered")
	}
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
//gocyclo:ignore
func (dr *DataRow) ColumnValueAsString(column *DataColumn) string {
	if column == nil {
		return ""
	}

	index := column.Index()

	switch column.Type() {
	case DataType.String:
		return dr.stringValueFromString(index)
	case DataType.Boolean:
		return dr.stringValueFromBoolean(index)
	case DataType.DateTime:
		return dr.stringValueFromDateTime(index)
	case DataType.Single:
		return dr.stringValueFromSingle(index)
	case DataType.Double:
		return dr.stringValueFromDouble(index)
	case DataType.Decimal:
		return dr.stringValueFromDecimal(index)
	case DataType.Guid:
		return dr.stringValueFromGuid(index)
	case DataType.Int8:
		return dr.stringValueFromInt8(index)
	case DataType.Int16:
		return dr.stringValueFromInt16(index)
	case DataType.Int32:
		return dr.stringValueFromInt32(index)
	case DataType.Int64:
		return dr.stringValueFromInt64(index)
	case DataType.UInt8:
		return dr.stringValueFromUInt8(index)
	case DataType.UInt16:
		return dr.stringValueFromUInt16(index)
	case DataType.UInt32:
		return dr.stringValueFromUInt32(index)
	case DataType.UInt64:
		return dr.stringValueFromUInt64(index)
	default:
		return ""
	}
}

func checkState(null bool, err error) (bool, string) {
	if err != nil {
		return true, ""
	}

	if null {
		return true, "<NULL>"
	}

	return false, ""
}

func (dr *DataRow) stringValueFromString(index int) string {
	value, null, err := dr.StringValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return value
}

func (dr *DataRow) stringValueFromBoolean(index int) string {
	value, null, err := dr.BooleanValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatBool(value)
}

func (dr *DataRow) stringValueFromDateTime(index int) string {
	value, null, err := dr.DateTimeValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return value.Format(DateTimeFormat)
}

func (dr *DataRow) stringValueFromSingle(index int) string {
	value, null, err := dr.SingleValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatFloat(float64(value), 'f', 6, 32)
}

func (dr *DataRow) stringValueFromDouble(index int) string {
	value, null, err := dr.DoubleValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatFloat(value, 'f', 6, 64)
}

func (dr *DataRow) stringValueFromDecimal(index int) string {
	value, null, err := dr.DecimalValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return value.String()
}

func (dr *DataRow) stringValueFromGuid(index int) string {
	value, null, err := dr.GuidValue(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return value.String()
}

func (dr *DataRow) stringValueFromInt8(index int) string {
	value, null, err := dr.Int8Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatInt(int64(value), 10)
}

func (dr *DataRow) stringValueFromInt16(index int) string {
	value, null, err := dr.Int16Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatInt(int64(value), 10)
}

func (dr *DataRow) stringValueFromInt32(index int) string {
	value, null, err := dr.Int32Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatInt(int64(value), 10)
}

func (dr *DataRow) stringValueFromInt64(index int) string {
	value, null, err := dr.Int64Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatInt(value, 10)
}

func (dr *DataRow) stringValueFromUInt8(index int) string {
	value, null, err := dr.UInt8Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatUint(uint64(value), 10)
}

func (dr *DataRow) stringValueFromUInt16(index int) string {
	value, null, err := dr.UInt16Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatUint(uint64(value), 10)
}

func (dr *DataRow) stringValueFromUInt32(index int) string {
	value, null, err := dr.UInt32Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatUint(uint64(value), 10)
}

func (dr *DataRow) stringValueFromUInt64(index int) string {
	value, null, err := dr.UInt64Value(index)

	if invalid, result := checkState(null, err); invalid {
		return result
	}

	return strconv.FormatUint(value, 10)
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

// DecimalValue gets the record value at the specified columnIndex cast as a decimal.Decimal.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Decimal.
func (dr *DataRow) DecimalValue(columnIndex int) (decimal.Decimal, bool, error) {
	column, err := dr.validateColumnType(columnIndex, int(DataType.Decimal), true)

	if err != nil {
		return decimal.Zero, false, err
	}

	if column.Computed() {
		value, err := dr.getComputedValue(column)

		if err != nil {
			return decimal.Zero, false, err
		}

		if value == nil {
			return decimal.Zero, true, nil
		}

		return value.(decimal.Decimal), false, nil
	}

	value := dr.values[columnIndex]

	if value == nil {
		return decimal.Zero, true, nil
	}

	return value.(decimal.Decimal), false, nil
}

// DecimalValueByName gets the record value for the specified columnName cast as a decimal.Decimal.
// Second parameter in tuple return value indicates if original value was nil.
// An error will be returned if column type is not DataType.Decimal.
func (dr *DataRow) DecimalValueByName(columnName string) (decimal.Decimal, bool, error) {
	index, err := dr.getColumnIndex(columnName)

	if err != nil {
		return decimal.Zero, false, err
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

// CompareDataRowColumns returns an integer comparing two DataRow column values for the
// specified column index. The result will be 0 if a==b, -1 if a < b, and +1 if a > b.
//gocyclo:ignore
func CompareDataRowColumns(leftRow, rightRow *DataRow, columnIndex int, exactMatch bool) (int, error) {
	leftColumn := leftRow.parent.Column(columnIndex)
	rightColumn := rightRow.parent.Column(columnIndex)

	if leftColumn == nil || rightColumn == nil {
		return 0, errors.New("cannot compare, column index out of range")
	}

	leftType := leftColumn.Type()
	rightType := rightColumn.Type()

	if leftType != rightType {
		return 0, errors.New("cannot compare, types do not match")
	}

	switch leftType {
	case DataType.String:
		leftValue, leftNull, leftErr := leftRow.StringValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.StringValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if exactMatch {
				return strings.Compare(leftValue, rightValue), nil
			}

			return strings.Compare(strings.ToUpper(leftValue), strings.ToUpper(rightValue)), nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Boolean:
		leftValue, leftNull, leftErr := leftRow.BooleanValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.BooleanValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue && !rightValue {
				return -1, nil
			}

			if !leftValue && rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.DateTime:
		leftValue, leftNull, leftErr := leftRow.DateTimeValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.DateTimeValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue.Before(rightValue) {
				return -1, nil
			}

			if leftValue.After(rightValue) {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Single:
		leftValue, leftNull, leftErr := leftRow.SingleValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.SingleValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Double:
		leftValue, leftNull, leftErr := leftRow.DoubleValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.DoubleValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Decimal:
		leftValue, leftNull, leftErr := leftRow.DecimalValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.DecimalValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue.Cmp(rightValue) < 0 {
				return -1, nil
			}

			if leftValue.Cmp(rightValue) > 0 {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Guid:
		leftValue, leftNull, leftErr := leftRow.GuidValue(columnIndex)
		rightValue, rightNull, rightErr := rightRow.GuidValue(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue.Compare(rightValue) < 0 {
				return -1, nil
			}

			if leftValue.Compare(rightValue) > 0 {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Int8:
		leftValue, leftNull, leftErr := leftRow.Int8Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.Int8Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Int16:
		leftValue, leftNull, leftErr := leftRow.Int16Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.Int16Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Int32:
		leftValue, leftNull, leftErr := leftRow.Int32Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.Int32Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.Int64:
		leftValue, leftNull, leftErr := leftRow.Int64Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.Int64Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.UInt8:
		leftValue, leftNull, leftErr := leftRow.UInt8Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.UInt8Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.UInt16:
		leftValue, leftNull, leftErr := leftRow.UInt16Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.UInt16Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.UInt32:
		leftValue, leftNull, leftErr := leftRow.UInt32Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.UInt32Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	case DataType.UInt64:
		leftValue, leftNull, leftErr := leftRow.UInt64Value(columnIndex)
		rightValue, rightNull, rightErr := rightRow.UInt64Value(columnIndex)
		leftHasValue := !leftNull && leftErr == nil
		rightHasValue := !rightNull && rightErr == nil

		if leftHasValue && rightHasValue {
			if leftValue < rightValue {
				return -1, nil
			}

			if leftValue > rightValue {
				return 1, nil
			}

			return 0, nil
		}

		return nullCompare(leftHasValue, rightHasValue), nil
	default:
		return 0, errors.New("unexpected column data type encountered")
	}
}

func nullCompare(leftHasValue, rightHasValue bool) int {
	if !leftHasValue && !rightHasValue {
		return 0
	}

	if leftHasValue {
		return 1
	}

	return -1
}
