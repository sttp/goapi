//******************************************************************************************************
//  ExpressionTree.go - Gbtc
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
//  09/30/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package filterexpressions

import (
	"errors"
	"math"
	"strconv"
	"strings"

	"github.com/sttp/goapi/sttp/data"
)

// ExpressionTree represents a tree of expressions for evaluation.
type ExpressionTree struct {
	currentRow *data.DataRow
	table      *data.DataTable

	// TopLimit represents the parsed value associated with the "TOP" keyword.
	TopLimit int32

	// OrderByTerms represents the order by elements parsed from the "ORDER BY" keyword.
	OrderByTerms []OrderByTerm

	// Root is the starting Expression from the parsed expression evaluation, or nil if there is not one.
	// This is the root expression of the ExpressionTree.
	Root Expression
}

// NewExpressionTree creates a new expression tree.
func NewExpressionTree(table *data.DataTable) *ExpressionTree {
	return &ExpressionTree{
		table:    table,
		TopLimit: -1,
	}
}

// Table gets the reference to the data table associated with the ExpressionTree.
func (et *ExpressionTree) Table() *data.DataTable {
	return et.table
}

// Evaluate executes the filter expression parser for the specified row for the ExpressionTree.
func (et *ExpressionTree) Evaluate(row *data.DataRow) (*ValueExpression, error) {
	et.currentRow = row
	return et.evaluate(et.Root)
}

func (et *ExpressionTree) evaluate(expression Expression) (*ValueExpression, error) {
	return et.evaluateAs(expression, ExpressionValueType.Boolean)
}

func (et *ExpressionTree) evaluateAs(expression Expression, targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	if expression == nil {
		return NullValue(targetValueType), nil
	}

	switch expression.Type() {
	case ExpressionType.Value:
		valueExpression := expression.(*ValueExpression)

		// Change Undefined NULL values to Nullable of target type
		if valueExpression.ValueType() == ExpressionValueType.Undefined {
			return NullValue(targetValueType), nil
		}

		return valueExpression, nil
	case ExpressionType.Unary:
		return et.evaluateUnary(expression)
	case ExpressionType.Column:
		return et.evaluateColumn(expression)
	case ExpressionType.InList:
		return et.evaluateInList(expression)
	case ExpressionType.Function:
		return et.evaluateFunction(expression)
	case ExpressionType.Operator:
		return et.evaluateOperator(expression)
	default:
		return nil, errors.New("unexpected expression type encountered")
	}
}

func (et *ExpressionTree) evaluateUnary(expression Expression) (*ValueExpression, error) {
	unaryExpression := expression.(*UnaryExpression)
	var err error

	var unaryValue *ValueExpression

	if unaryValue, err = et.evaluate(unaryExpression.Value()); err != nil {
		return nil, err
	}

	unaryValueType := unaryValue.ValueType()

	// If unary value is Null, result is Null
	if unaryValue.IsNull() {
		return NullValue(unaryValueType), nil
	}

	switch unaryValueType {
	case ExpressionValueType.Boolean:
		return unaryExpression.unaryBoolean(unaryValue.booleanValue())
	case ExpressionValueType.Int32:
		return unaryExpression.unaryInt32(unaryValue.int32Value())
	case ExpressionValueType.Int64:
		return unaryExpression.unaryInt64(unaryValue.int64Value())
	case ExpressionValueType.Decimal:
		return unaryExpression.unaryDecimal(unaryValue.decimalValue())
	case ExpressionValueType.Double:
		return unaryExpression.unaryDouble(unaryValue.doubleValue())
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply unary \"" + unaryExpression.UnaryType().String() + "\" operator to \"" + unaryValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) evaluateColumn(expression Expression) (*ValueExpression, error) {
	columnExpression := expression.(*ColumnExpression)
	var column *data.DataColumn
	var err error

	if column = columnExpression.DataColumn(); column == nil {
		return nil, errors.New("encountered column expression with undefined data column reference")
	}

	columnIndex := column.Index()
	var valueType ExpressionValueTypeEnum
	var value interface{}
	var isNull bool

	// Map column DataType to ExpressionType, storing equivalent literal value (can be nil)
	switch column.Type() {
	case data.DataType.String:
		valueType = ExpressionValueType.String
		value, isNull, err = et.currentRow.StringValue(columnIndex)
	case data.DataType.Boolean:
		valueType = ExpressionValueType.Boolean
		value, isNull, err = et.currentRow.BooleanValue(columnIndex)
	case data.DataType.DateTime:
		valueType = ExpressionValueType.DateTime
		value, isNull, err = et.currentRow.DateTimeValue(columnIndex)
	case data.DataType.Single:
		var f32 float32
		valueType = ExpressionValueType.Double
		f32, isNull, err = et.currentRow.SingleValue(columnIndex)
		value = float64(f32)
	case data.DataType.Double:
		valueType = ExpressionValueType.Double
		value, isNull, err = et.currentRow.DoubleValue(columnIndex)
	case data.DataType.Decimal:
		valueType = ExpressionValueType.Decimal
		value, isNull, err = et.currentRow.DecimalValue(columnIndex)
	case data.DataType.Guid:
		valueType = ExpressionValueType.Guid
		value, isNull, err = et.currentRow.GuidValue(columnIndex)
	case data.DataType.Int8:
		var i8 int8
		valueType = ExpressionValueType.Int32
		i8, isNull, err = et.currentRow.Int8Value(columnIndex)
		value = int32(i8)
	case data.DataType.Int16:
		var i16 int16
		valueType = ExpressionValueType.Int32
		i16, isNull, err = et.currentRow.Int16Value(columnIndex)
		value = int32(i16)
	case data.DataType.Int32:
		valueType = ExpressionValueType.Int32
		value, isNull, err = et.currentRow.Int32Value(columnIndex)
	case data.DataType.Int64:
		valueType = ExpressionValueType.Int64
		value, isNull, err = et.currentRow.Int64Value(columnIndex)
	case data.DataType.UInt8:
		var ui8 uint8
		valueType = ExpressionValueType.Int32
		ui8, isNull, err = et.currentRow.UInt8Value(columnIndex)
		value = int32(ui8)
	case data.DataType.UInt16:
		var ui16 uint16
		valueType = ExpressionValueType.Int32
		ui16, isNull, err = et.currentRow.UInt16Value(columnIndex)
		value = int32(ui16)
	case data.DataType.UInt32:
		var ui32 uint32
		valueType = ExpressionValueType.Int64
		ui32, isNull, err = et.currentRow.UInt32Value(columnIndex)
		value = int64(ui32)
	case data.DataType.UInt64:
		var ui64 uint64
		ui64, isNull, err = et.currentRow.UInt64Value(columnIndex)

		if ui64 > math.MaxInt64 {
			valueType = ExpressionValueType.Double
			value = float64(ui64)
		} else {
			valueType = ExpressionValueType.Int64
			value = int64(ui64)
		}
	default:
		return nil, errors.New("unexpected column data type encountered")
	}

	if err != nil {
		return nil, err
	}

	if isNull {
		return NullValue(valueType), nil
	}

	return newValueExpression(valueType, value), nil
}

func (et *ExpressionTree) evaluateInList(expression Expression) (*ValueExpression, error) {
	inListExpression := expression.(*InListExpression)
	var inListValue *ValueExpression
	var err error

	if inListValue, err = et.evaluate(inListExpression.Value()); err != nil {
		return nil, err
	}

	// If in list test value is Null, result is Null
	if inListValue.IsNull() {
		return NullValue(inListValue.ValueType()), nil
	}

	hasNotKeyWord := inListExpression.HasNotKeyword()
	exactMatch := inListExpression.ExtactMatch()
	arguments := inListExpression.Arguments()

	var argumentValue, result *ValueExpression
	var valueType ExpressionValueTypeEnum

	for i := 0; i < len(arguments); i++ {
		if argumentValue, err = et.evaluate(arguments[i]); err != nil {
			return nil, err
		}

		if valueType, err = ExpressionOperatorType.Equal.deriveComparisonOperationValueType(inListValue.ValueType(), argumentValue.ValueType()); err != nil {
			return nil, err
		}

		if result, err = et.equalOp(inListValue, argumentValue, valueType, exactMatch); err != nil {
			return nil, err
		}

		if result.booleanValue() {
			if hasNotKeyWord {
				return False, nil
			}

			return True, nil
		}
	}

	if hasNotKeyWord {
		return True, nil
	}

	return False, nil
}

func (et *ExpressionTree) evaluateFunction(expression Expression) (*ValueExpression, error) {
	functionExpression := expression.(*FunctionExpression)
	arguments := functionExpression.Arguments()

	switch functionExpression.FunctionType() {
	case ExpressionFunctionType.Abs:
		return et.evaluateAbs(arguments)
	case ExpressionFunctionType.Ceiling:
		return et.evaluateCeiling(arguments)
	case ExpressionFunctionType.Coalesce:
		return et.evaluateCoalesce(arguments)
	case ExpressionFunctionType.Convert:
		return et.evaluateConvert(arguments)
	case ExpressionFunctionType.Contains:
		return et.evaluateContains(arguments)
	case ExpressionFunctionType.DateAdd:
		return et.evaluateDateAdd(arguments)
	case ExpressionFunctionType.DateDiff:
		return et.evaluateDateDiff(arguments)
	case ExpressionFunctionType.DatePart:
		return et.evaluateDatePart(arguments)
	case ExpressionFunctionType.EndsWith:
		return et.evaluateEndsWith(arguments)
	case ExpressionFunctionType.Floor:
		return et.evaluateFloor(arguments)
	case ExpressionFunctionType.IIf:
		return et.evaluateIIf(arguments)
	case ExpressionFunctionType.IndexOf:
		return et.evaluateIndexOf(arguments)
	case ExpressionFunctionType.IsDate:
		return et.evaluateIsDate(arguments)
	case ExpressionFunctionType.IsInteger:
		return et.evaluateIsInteger(arguments)
	case ExpressionFunctionType.IsGuid:
		return et.evaluateIsGuid(arguments)
	case ExpressionFunctionType.IsNull:
		return et.evaluateIsNull(arguments)
	case ExpressionFunctionType.IsNumeric:
		return et.evaluateIsNumeric(arguments)
	case ExpressionFunctionType.LastIndexOf:
		return et.evaluateLastIndexOf(arguments)
	case ExpressionFunctionType.Len:
		return et.evaluateLen(arguments)
	case ExpressionFunctionType.Lower:
		return et.evaluateLower(arguments)
	case ExpressionFunctionType.MaxOf:
		return et.evaluateMaxOf(arguments)
	case ExpressionFunctionType.MinOf:
		return et.evaluateMinOf(arguments)
	case ExpressionFunctionType.NthIndexOf:
		return et.evaluateNthIndexOf(arguments)
	case ExpressionFunctionType.Now:
		return et.evaluateNow(arguments)
	case ExpressionFunctionType.Power:
		return et.evaluatePower(arguments)
	case ExpressionFunctionType.RegExMatch:
		return et.evaluateRegExMatch(arguments)
	case ExpressionFunctionType.RegExVal:
		return et.evaluateRegExVal(arguments)
	case ExpressionFunctionType.Replace:
		return et.evaluateReplace(arguments)
	case ExpressionFunctionType.Reverse:
		return et.evaluateReverse(arguments)
	case ExpressionFunctionType.Round:
		return et.evaluateRound(arguments)
	case ExpressionFunctionType.Split:
		return et.evaluateSplit(arguments)
	case ExpressionFunctionType.Sqrt:
		return et.evaluateSqrt(arguments)
	case ExpressionFunctionType.StartsWith:
		return et.evaluateStartsWith(arguments)
	case ExpressionFunctionType.StrCount:
		return et.evaluateStrCount(arguments)
	case ExpressionFunctionType.StrCmp:
		return et.evaluateStrCmp(arguments)
	case ExpressionFunctionType.SubStr:
		return et.evaluateSubStr(arguments)
	case ExpressionFunctionType.Trim:
		return et.evaluateTrim(arguments)
	case ExpressionFunctionType.TrimLeft:
		return et.evaluateTrimLeft(arguments)
	case ExpressionFunctionType.TrimRight:
		return et.evaluateTrimRight(arguments)
	case ExpressionFunctionType.Upper:
		return et.evaluateUpper(arguments)
	case ExpressionFunctionType.UtcNow:
		return et.evaluateUtcNow(arguments)
	default:
		return nil, errors.New("unexpected function type encountered")
	}
}

func (et *ExpressionTree) evaluateAbs(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Abs\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.Double); err != nil {
		return nil, err
	}

	return et.abs(sourceValue)
}

func (et *ExpressionTree) evaluateCeiling(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Ceiling\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.Double); err != nil {
		return nil, err
	}

	return et.ceiling(sourceValue)
}

func (et *ExpressionTree) evaluateCoalesce(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 {
		return nil, errors.New("\"Coalesce\" function expects at least 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	// Not pre-evaluating Coalesce arguments - arguments will be evaluated only up to first non-null value
	return et.coalesce(arguments)
}

func (et *ExpressionTree) evaluateConvert(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"Convert\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, targetType *ValueExpression
	var err error

	if sourceValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, err
	}

	if targetType, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.convert(sourceValue, targetType)
}

func (et *ExpressionTree) evaluateContains(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"Contains\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.contains(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.contains(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateDateAdd(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 3 {
		return nil, errors.New("\"DateAdd\" function expects 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, addValue, intervalType *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.DateTime); err != nil {
		return nil, err
	}

	if addValue, err = et.evaluateAs(arguments[1], ExpressionValueType.Int32); err != nil {
		return nil, err
	}

	if intervalType, err = et.evaluateAs(arguments[2], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.dateAdd(sourceValue, addValue, intervalType)
}

func (et *ExpressionTree) evaluateDateDiff(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 3 {
		return nil, errors.New("\"DateDiff\" function expects 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var leftValue, rightValue, intervalType *ValueExpression
	var err error

	if leftValue, err = et.evaluateAs(arguments[0], ExpressionValueType.DateTime); err != nil {
		return nil, err
	}

	if rightValue, err = et.evaluateAs(arguments[1], ExpressionValueType.DateTime); err != nil {
		return nil, err
	}

	if intervalType, err = et.evaluateAs(arguments[2], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.dateDiff(leftValue, rightValue, intervalType)
}

func (et *ExpressionTree) evaluateDatePart(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"DatePart\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, intervalType *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.DateTime); err != nil {
		return nil, err
	}

	if intervalType, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.datePart(sourceValue, intervalType)
}

func (et *ExpressionTree) evaluateEndsWith(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"EndsWith\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.endsWith(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.endsWith(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateFloor(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Floor\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.Double); err != nil {
		return nil, err
	}

	return et.floor(sourceValue)
}

func (et *ExpressionTree) evaluateIIf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 3 {
		return nil, errors.New("\"IIf\" function expects 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var testValue *ValueExpression
	var err error

	if testValue, err = et.evaluate(arguments[1]); err != nil {
		return nil, err
	}

	// Not pre-evaluating IIf result value arguments - only evaluating desired path
	return et.iif(testValue, arguments[1], arguments[2])
}

func (et *ExpressionTree) evaluateIndexOf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"IndexOf\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.indexOf(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.indexOf(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateIsDate(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsDate\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, err
	}

	return et.isDate(sourceValue)
}

func (et *ExpressionTree) evaluateIsInteger(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsInteger\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, err
	}

	return et.isInteger(sourceValue)
}

func (et *ExpressionTree) evaluateIsGuid(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsGuid\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, err
	}

	return et.isGuid(sourceValue)
}

func (et *ExpressionTree) evaluateIsNull(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"IsNull\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var testValue, defaultValue *ValueExpression
	var err error

	if testValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if defaultValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.isNull(testValue, defaultValue)
}

func (et *ExpressionTree) evaluateIsNumeric(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsNumeric\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, err
	}

	return et.isNumeric(sourceValue)
}

func (et *ExpressionTree) evaluateLastIndexOf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"LastIndexOf\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.lastIndexOf(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.lastIndexOf(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateLen(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Len\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.len(sourceValue)
}

func (et *ExpressionTree) evaluateLower(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Lower\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.lower(sourceValue)
}

func (et *ExpressionTree) evaluateMaxOf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 {
		return nil, errors.New("\"MaxOf\" function expects at least 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	return et.maxOf(arguments)
}

func (et *ExpressionTree) evaluateMinOf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 {
		return nil, errors.New("\"MinOf\" function expects at least 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	return et.minOf(arguments)
}

func (et *ExpressionTree) evaluateNthIndexOf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 3 || len(arguments) > 4 {
		return nil, errors.New("\"NthIndexOf\" function expects 3 or 4 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue, indexValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if indexValue, err = et.evaluateAs(arguments[2], ExpressionValueType.Int32); err != nil {
		return nil, err
	}

	if len(arguments) == 3 {
		return et.nthIndexOf(sourceValue, testValue, indexValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[3], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.nthIndexOf(sourceValue, testValue, indexValue, ignoreCase)
}

func (et *ExpressionTree) evaluateNow(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) > 0 {
		return nil, errors.New("\"Now\" function expects 0 arguments, received " + strconv.Itoa(len(arguments)))
	}

	return et.now()
}

func (et *ExpressionTree) evaluatePower(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"Power\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, exponentValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.Double); err != nil {
		return nil, err
	}

	if exponentValue, err = et.evaluateAs(arguments[1], ExpressionValueType.Int32); err != nil {
		return nil, err
	}

	return et.power(sourceValue, exponentValue)
}

func (et *ExpressionTree) evaluateRegExMatch(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"RegExMatch\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var regexValue, testValue *ValueExpression
	var err error

	if regexValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.regExMatch(regexValue, testValue)
}

func (et *ExpressionTree) evaluateRegExVal(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"RegExVal\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var regexValue, testValue *ValueExpression
	var err error

	if regexValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.regExVal(regexValue, testValue)
}

func (et *ExpressionTree) evaluateReplace(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 3 || len(arguments) > 4 {
		return nil, errors.New("\"Replace\" function expects 3 or 4 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue, replaceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if replaceValue, err = et.evaluateAs(arguments[2], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.replace(sourceValue, testValue, replaceValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[3], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.replace(sourceValue, testValue, replaceValue, ignoreCase)
}

func (et *ExpressionTree) evaluateReverse(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Reverse\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.reverse(sourceValue)
}

func (et *ExpressionTree) evaluateRound(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Round\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.Double); err != nil {
		return nil, err
	}

	return et.round(sourceValue)
}

func (et *ExpressionTree) evaluateSplit(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 3 || len(arguments) > 4 {
		return nil, errors.New("\"Split\" function expects 3 or 4 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, delimeterValue, indexValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if delimeterValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if indexValue, err = et.evaluateAs(arguments[2], ExpressionValueType.Int32); err != nil {
		return nil, err
	}

	if len(arguments) == 3 {
		return et.split(sourceValue, delimeterValue, indexValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[3], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.split(sourceValue, delimeterValue, indexValue, ignoreCase)
}

func (et *ExpressionTree) evaluateSqrt(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Sqrt\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.Double); err != nil {
		return nil, err
	}

	return et.sqrt(sourceValue)
}

func (et *ExpressionTree) evaluateStartsWith(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"StartsWith\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.startsWith(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.startsWith(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateStrCount(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"StrCount\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.strCount(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.strCount(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateStrCmp(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"StrCmp\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var leftValue, rightValue *ValueExpression
	var err error

	if leftValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if rightValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.strCmp(leftValue, rightValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, err
	}

	return et.strCmp(leftValue, rightValue, ignoreCase)
}

func (et *ExpressionTree) evaluateSubStr(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"SubStr\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, indexValue, lengthValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	if indexValue, err = et.evaluateAs(arguments[1], ExpressionValueType.Int32); err != nil {
		return nil, err
	}

	if len(arguments) == 2 {
		return et.subStr(sourceValue, indexValue, NullValue(ExpressionValueType.Int32))
	}

	if lengthValue, err = et.evaluateAs(arguments[2], ExpressionValueType.Int32); err != nil {
		return nil, err
	}

	return et.subStr(sourceValue, indexValue, lengthValue)
}

func (et *ExpressionTree) evaluateTrim(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Trim\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.trim(sourceValue)
}

func (et *ExpressionTree) evaluateTrimLeft(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"TrimLeft\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.trimLeft(sourceValue)
}

func (et *ExpressionTree) evaluateTrimRight(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"TrimRight\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.trimRight(sourceValue)
}

func (et *ExpressionTree) evaluateUpper(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"Upper\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, err
	}

	return et.upper(sourceValue)
}

func (et *ExpressionTree) evaluateUtcNow(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) > 0 {
		return nil, errors.New("\"UtcNow\" function expects 0 arguments, received " + strconv.Itoa(len(arguments)))
	}

	return et.utcNow()
}

func (et *ExpressionTree) evaluateOperator(expression Expression) (*ValueExpression, error) {
	operatorExpression := expression.(*OperatorExpression)
	var err error

	var leftValue *ValueExpression

	if leftValue, err = et.evaluate(operatorExpression.LeftValue()); err != nil {
		return nil, err
	}

	var rightValue *ValueExpression

	if rightValue, err = et.evaluate(operatorExpression.RightValue()); err != nil {
		return nil, err
	}

	var valueType ExpressionValueTypeEnum

	if valueType, err = ExpressionOperatorType.Equal.deriveOperationValueType(leftValue.ValueType(), rightValue.ValueType()); err != nil {
		return nil, err
	}

	switch operatorExpression.OperatorType() {
	case ExpressionOperatorType.Multiply:
		return et.multiplyOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.Divide:
		return et.divideOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.Modulus:
		return et.modulusOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.Add:
		return et.addOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.Subtract:
		return et.subtractOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.BitShiftLeft:
		return et.bitShiftLeftOp(leftValue, rightValue)
	case ExpressionOperatorType.BitShiftRight:
		return et.bitShiftRightOp(leftValue, rightValue)
	case ExpressionOperatorType.BitwiseAnd:
		return et.bitwiseAndOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.BitwiseOr:
		return et.bitwiseOrOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.BitwiseXor:
		return et.bitwiseXorOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.LessThan:
		return et.lessThanOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.LessThanOrEqual:
		return et.lessThanOrEqualOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.GreaterThan:
		return et.greaterThanOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.GreaterThanOrEqual:
		return et.greaterThanOrEqualOp(leftValue, rightValue, valueType)
	case ExpressionOperatorType.Equal:
		return et.equalOp(leftValue, rightValue, valueType, false)
	case ExpressionOperatorType.EqualExactMatch:
		return et.equalOp(leftValue, rightValue, valueType, true)
	case ExpressionOperatorType.NotEqual:
		return et.notEqualOp(leftValue, rightValue, valueType, false)
	case ExpressionOperatorType.NotEqualExactMatch:
		return et.notEqualOp(leftValue, rightValue, valueType, true)
	case ExpressionOperatorType.IsNull:
		return et.isNullOp(leftValue)
	case ExpressionOperatorType.IsNotNull:
		return et.isNotNullOp(leftValue)
	case ExpressionOperatorType.Like:
		return et.likeOp(leftValue, rightValue, false)
	case ExpressionOperatorType.LikeExactMatch:
		return et.likeOp(leftValue, rightValue, true)
	case ExpressionOperatorType.NotLike:
		return et.notLikeOp(leftValue, rightValue, false)
	case ExpressionOperatorType.NotLikeExactMatch:
		return et.notLikeOp(leftValue, rightValue, true)
	case ExpressionOperatorType.And:
		return et.andOp(leftValue, rightValue)
	case ExpressionOperatorType.Or:
		return et.orOp(leftValue, rightValue)
	default:
		return nil, errors.New("unexpected operator type encountered")
	}
}

// Filter Expression Function Implementations

func (et *ExpressionTree) abs(sourceValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Abs\" function argument must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, sourceValue.booleanValue()), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, abs32(sourceValue.int32Value())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, abs64(sourceValue.int64Value())), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Abs(sourceValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Abs(sourceValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func abs32(value int32) int32 {
	if value < 0 {
		return -value
	}
	return value
}

func abs64(value int64) int64 {
	if value < 0 {
		return -value
	}
	return value
}

func (et *ExpressionTree) ceiling(sourceValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Ceiling\" function argument must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	if sourceValue.ValueType().IsIntegerType() {
		return sourceValue, nil
	}

	var err error

	switch sourceValue.ValueType() {
	case ExpressionValueType.Decimal:
		var f64 float64
		if f64, err = sourceValue.DoubleValue(); err != nil {
			return nil, err
		}
		return NewValueExpression(ExpressionValueType.Decimal, math.Ceil(f64)), nil
	case ExpressionValueType.Double:
		var f64 float64
		if f64, err = sourceValue.DoubleValue(); err != nil {
			return nil, err
		}
		return NewValueExpression(ExpressionValueType.Double, math.Ceil(f64)), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) coalesce(arguments []Expression) (*ValueExpression, error) {
	testValue, err := et.evaluate(arguments[0])

	if err != nil {
		return nil, err
	}

	if !testValue.IsNull() {
		return testValue, nil
	}

	for i := 1; i < len(arguments); i++ {
		listValue, err := et.evaluate(arguments[i])

		if err != nil {
			return nil, err
		}

		if !listValue.IsNull() {
			return listValue, nil
		}
	}

	return testValue, nil
}

func (et *ExpressionTree) convert(sourceValue *ValueExpression, targetType *ValueExpression) (*ValueExpression, error) {
	if targetType.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Convert\" function target type, second argument, must be a string")
	}

	if targetType.IsNull() {
		return nil, errors.New("\"Convert\" function target type, second argument, is null")
	}

	targetTypeName, err := targetType.StringValue()

	if err != nil {
		return nil, err
	}

	targetTypeName = strings.ToUpper(targetTypeName)

	// Remove any "System." prefix:       01234567
	if strings.HasPrefix(targetTypeName, "SYSTEM.") && len(targetTypeName) > 7 {
		targetTypeName = targetTypeName[7:]
	}

	targetValueType := ExpressionValueType.Undefined
	var foundValueType bool

	for i := 0; i < ExpressionValueTypeLen(); i++ {
		valueType := ExpressionValueTypeEnum(i)

		if targetTypeName == strings.ToUpper(valueType.String()) {
			targetValueType = valueType
			foundValueType = true
			break
		}
	}

	if !foundValueType {
		// Handle a few common aliases
		if targetTypeName == "SINGLE" || strings.HasPrefix(targetTypeName, "FLOAT") {
			targetValueType = ExpressionValueType.Double
			foundValueType = true
		} else if targetTypeName == "BOOL" {
			targetValueType = ExpressionValueType.Boolean
			foundValueType = true
		} else if strings.HasPrefix(targetTypeName, "INT") || strings.HasPrefix(targetTypeName, "UINT") {
			targetValueType = ExpressionValueType.Int64
			foundValueType = true
		} else if targetTypeName == "DATE" || targetTypeName == "TIME" {
			targetValueType = ExpressionValueType.DateTime
			foundValueType = true
		} else if targetTypeName == "UUID" {
			targetValueType = ExpressionValueType.Guid
			foundValueType = true
		}
	}

	if !foundValueType || targetValueType == ExpressionValueType.Undefined {
		target, _ := targetType.StringValue()
		return nil, errors.New("specified \"Convert\" function target type \"" + target + "\", second argument, is not supported")
	}

	return sourceValue.Convert(targetValueType)
}

func (et *ExpressionTree) contains(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) dateAdd(sourceValue *ValueExpression, addValue *ValueExpression, intervalType *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) dateDiff(leftValue *ValueExpression, rightValue *ValueExpression, intervalType *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) datePart(sourceValue *ValueExpression, intervalType *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) endsWith(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) floor(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) iif(testValue *ValueExpression, leftResultValue Expression, rightResultValue Expression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) indexOf(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isDate(testValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isInteger(testValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isGuid(testValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isNull(testValue *ValueExpression, defaultValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isNumeric(testValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) lastIndexOf(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) len(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) lower(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) maxOf(arguments []Expression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) minOf(arguments []Expression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) now() (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) nthIndexOf(sourceValue *ValueExpression, testValue *ValueExpression, indexValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) power(sourceValue *ValueExpression, exponentValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) regExMatch(regexValue *ValueExpression, testValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) regExVal(regexValue *ValueExpression, testValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) replace(sourceValue *ValueExpression, testValue *ValueExpression, replaceValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) reverse(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) round(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) split(sourceValue *ValueExpression, delimiterValue *ValueExpression, indexValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) sqrt(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) startsWith(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) strCount(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) strCmp(leftValue *ValueExpression, rightValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) subStr(sourceValue *ValueExpression, indexValue *ValueExpression, lengthValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) trim(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) trimLeft(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) trimRight(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) upper(sourceValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) utcNow() (*ValueExpression, error) {
	return nil, nil
}

// Filter Expression Operator Implementations

func (et *ExpressionTree) multiplyOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) divideOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) modulusOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) addOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) subtractOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) bitShiftLeftOp(leftValue *ValueExpression, rightValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) bitShiftRightOp(leftValue *ValueExpression, rightValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) bitwiseAndOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) bitwiseOrOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) bitwiseXorOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) lessThanOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) lessThanOrEqualOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) greaterThanOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) greaterThanOrEqualOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) equalOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum, exactMatch bool) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) notEqualOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum, exactMatch bool) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isNullOp(leftValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) isNotNullOp(leftValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) likeOp(leftValue *ValueExpression, rightValue *ValueExpression, exactMatch bool) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) notLikeOp(leftValue *ValueExpression, rightValue *ValueExpression, exactMatch bool) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) andOp(leftValue *ValueExpression, rightValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) orOp(leftValue *ValueExpression, rightValue *ValueExpression) (*ValueExpression, error) {
	return nil, nil
}
