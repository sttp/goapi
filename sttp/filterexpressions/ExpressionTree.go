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
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/araddon/dateparse"
	"github.com/sttp/goapi/sttp/data"
	"github.com/sttp/goapi/sttp/guid"
)

var whitespace string = " \t\n\v\f\r\x85\xA0"

// ExpressionTree represents a tree of expressions for evaluation.
type ExpressionTree struct {
	currentRow *data.DataRow
	table      *data.DataTable

	// TopLimit represents the parsed value associated with the "TOP" keyword.
	TopLimit int32

	// OrderByTerms represents the order by elements parsed from the "ORDER BY" keyword.
	OrderByTerms []OrderByTerm

	// Root is the starting Expression for evaluation of the expression tree, or nil if
	// there is not one. This is the root expression of the ExpressionTree. Assign root
	// value before calling Evaluate.
	Root Expression
}

// NewExpressionTree creates a new expression tree.
func NewExpressionTree(table *data.DataTable) *ExpressionTree {
	return &ExpressionTree{
		table:    table,
		TopLimit: -1,
	}
}

// Table gets reference to the data table associated with the ExpressionTree.
func (et *ExpressionTree) Table() *data.DataTable {
	return et.table
}

// Evaluate executes the filter expression parser for the specified row for the ExpressionTree.
// Root expression should be assigned before calling Evaluate.
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
		return nil, errors.New("failed while evaluating unary expression value: " + err.Error())
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
		return nil, errors.New("failed while getting column \"" + column.Name() + "\" " + column.Type().String() + " value for current row: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"IN\" expression source value: " + err.Error())
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
			return nil, errors.New("failed while evaluating \"IN\" expression argument " + strconv.Itoa(i) + ": " + err.Error())
		}

		if valueType, err = ExpressionOperatorType.Equal.deriveComparisonOperationValueType(inListValue.ValueType(), argumentValue.ValueType()); err != nil {
			return nil, errors.New("failed while deriving \"IN\" expresssion equality comparison operation value type: " + err.Error())
		}

		if result, err = et.equalOp(inListValue, argumentValue, valueType, exactMatch); err != nil {
			return nil, errors.New("failed while comparing \"IN\" expresssion source value to argument " + strconv.Itoa(i) + " for equality: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Abs\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Ceiling\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Convert\" function source value, first argument: " + err.Error())
	}

	if targetType, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"Convert\" function target type, second argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Contains\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"Contains\" function test value, first argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.contains(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"Contains\" function optional ignore case value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"DateAdd\" function source value, first argument: " + err.Error())
	}

	if addValue, err = et.evaluateAs(arguments[1], ExpressionValueType.Int32); err != nil {
		return nil, errors.New("failed while evaluating \"DateAdd\" function add value, second argument: " + err.Error())
	}

	if intervalType, err = et.evaluateAs(arguments[2], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"DateAdd\" function interval type, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"DateDiff\" function left value, first argument: " + err.Error())
	}

	if rightValue, err = et.evaluateAs(arguments[1], ExpressionValueType.DateTime); err != nil {
		return nil, errors.New("failed while evaluating \"DateDiff\" function right value, second argument: " + err.Error())
	}

	if intervalType, err = et.evaluateAs(arguments[2], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"DateDiff\" function interval type, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"DatePart\" function source value, first argument: " + err.Error())
	}

	if intervalType, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"DatePart\" function interval type, second argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"EndsWith\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"EndsWith\" function test value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.endsWith(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"EndsWith\" function optional ignore case value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Floor\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"IIf\" function test value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"IndexOf\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"IndexOf\" function test value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.indexOf(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"IndexOf\" function optional ignore case value, third argument: " + err.Error())
	}

	return et.indexOf(sourceValue, testValue, ignoreCase)
}

func (et *ExpressionTree) evaluateIsDate(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsDate\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var testValue *ValueExpression
	var err error

	if testValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, errors.New("failed while evaluating \"IsDate\" function test value, first argument: " + err.Error())
	}

	return et.isDate(testValue), nil
}

func (et *ExpressionTree) evaluateIsInteger(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsInteger\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var testValue *ValueExpression
	var err error

	if testValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, errors.New("failed while evaluating \"IsInteger\" function test value, first argument: " + err.Error())
	}

	return et.isInteger(testValue), nil
}

func (et *ExpressionTree) evaluateIsGuid(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsGuid\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var testValue *ValueExpression
	var err error

	if testValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, errors.New("failed while evaluating \"IsGuid\" function test value, first argument: " + err.Error())
	}

	return et.isGuid(testValue), nil
}

func (et *ExpressionTree) evaluateIsNull(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 2 {
		return nil, errors.New("\"IsNull\" function expects 2 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var testValue, defaultValue *ValueExpression
	var err error

	if testValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"IsNull\" function test value, first argument: " + err.Error())
	}

	if defaultValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"IsNull\" function default value, second argument: " + err.Error())
	}

	return et.isNull(testValue, defaultValue)
}

func (et *ExpressionTree) evaluateIsNumeric(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) != 1 {
		return nil, errors.New("\"IsNumeric\" function expects 1 argument, received " + strconv.Itoa(len(arguments)))
	}

	var testValue *ValueExpression
	var err error

	if testValue, err = et.evaluate(arguments[0]); err != nil {
		return nil, errors.New("failed while evaluating \"IsNumeric\" function test value, first argument: " + err.Error())
	}

	return et.isNumeric(testValue), nil
}

func (et *ExpressionTree) evaluateLastIndexOf(arguments []Expression) (*ValueExpression, error) {
	if len(arguments) < 2 || len(arguments) > 3 {
		return nil, errors.New("\"LastIndexOf\" function expects 2 or 3 arguments, received " + strconv.Itoa(len(arguments)))
	}

	var sourceValue, testValue *ValueExpression
	var err error

	if sourceValue, err = et.evaluateAs(arguments[0], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"LastIndexOf\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"LastIndexOf\" function test value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.lastIndexOf(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"LastIndexOf\" function optional ignore case value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Len\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Lower\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"NthIndexOf\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"NthIndexOf\" function test value, second argument: " + err.Error())
	}

	if indexValue, err = et.evaluateAs(arguments[2], ExpressionValueType.Int32); err != nil {
		return nil, errors.New("failed while evaluating \"NthIndexOf\" function index value, third argument: " + err.Error())
	}

	if len(arguments) == 3 {
		return et.nthIndexOf(sourceValue, testValue, indexValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[3], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"NthIndexOf\" function optional ignore case value, fourth argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Power\" function source value, first argument: " + err.Error())
	}

	if exponentValue, err = et.evaluateAs(arguments[1], ExpressionValueType.Int32); err != nil {
		return nil, errors.New("failed while evaluating \"Power\" function exponent value, second argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"RegExMatch\" function expression value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"RegExMatch\" function test value, second argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"RegExVal\" function expression value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"RegExVal\" function test value, second argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Replace\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"Replace\" function test value, second argument: " + err.Error())
	}

	if replaceValue, err = et.evaluateAs(arguments[2], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"Replace\" function replace value, third argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.replace(sourceValue, testValue, replaceValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[3], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"Replace\" function optional ignore case value, fourth argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Reverse\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Round\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Split\" function source value, first argument: " + err.Error())
	}

	if delimeterValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"Split\" function delimiter value, second argument: " + err.Error())
	}

	if indexValue, err = et.evaluateAs(arguments[2], ExpressionValueType.Int32); err != nil {
		return nil, errors.New("failed while evaluating \"Split\" function index value, third argument: " + err.Error())
	}

	if len(arguments) == 3 {
		return et.split(sourceValue, delimeterValue, indexValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[3], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"Split\" function optional ignore case value, fourth argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Sqrt\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"StartsWith\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"StartsWith\" function test value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.startsWith(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"StartsWith\" function optional ignore case value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"StrCount\" function source value, first argument: " + err.Error())
	}

	if testValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"StrCount\" function test value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.strCount(sourceValue, testValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"StrCount\" function optional ignore case value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"StrCmp\" function left value, first argument: " + err.Error())
	}

	if rightValue, err = et.evaluateAs(arguments[1], ExpressionValueType.String); err != nil {
		return nil, errors.New("failed while evaluating \"StrCmp\" function right value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.strCmp(leftValue, rightValue, NullValue(ExpressionValueType.Boolean))
	}

	var ignoreCase *ValueExpression

	if ignoreCase, err = et.evaluateAs(arguments[2], ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while evaluating \"StrCmp\" function optional ignore case value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"SubStr\" function source value, first argument: " + err.Error())
	}

	if indexValue, err = et.evaluateAs(arguments[1], ExpressionValueType.Int32); err != nil {
		return nil, errors.New("failed while evaluating \"SubStr\" function index value, second argument: " + err.Error())
	}

	if len(arguments) == 2 {
		return et.subStr(sourceValue, indexValue, NullValue(ExpressionValueType.Int32))
	}

	if lengthValue, err = et.evaluateAs(arguments[2], ExpressionValueType.Int32); err != nil {
		return nil, errors.New("failed while evaluating \"SubStr\" function optional length value, third argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Trim\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"TrimLeft\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"TrimRight\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"Upper\" function source value, first argument: " + err.Error())
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
		return nil, errors.New("failed while evaluating \"" + operatorExpression.OperatorType().String() + "\" operator left operand: " + err.Error())
	}

	var rightValue *ValueExpression

	if rightValue, err = et.evaluate(operatorExpression.RightValue()); err != nil {
		return nil, errors.New("failed while evaluating \"" + operatorExpression.OperatorType().String() + "\" operator right operand: " + err.Error())
	}

	var valueType ExpressionValueTypeEnum

	if valueType, err = operatorExpression.OperatorType().deriveOperationValueType(leftValue.ValueType(), rightValue.ValueType()); err != nil {
		return nil, errors.New("failed while deriving \"" + operatorExpression.OperatorType().String() + "\" operator value type: " + err.Error())
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
		return nil, errors.New("\"Abs\" function source value, first argument, must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, sourceValue.booleanValue()), nil
	case ExpressionValueType.Int32:
		abs := func(value int32) int32 {
			if value < 0 {
				return -value
			}
			return value
		}

		return newValueExpression(ExpressionValueType.Int32, abs(sourceValue.int32Value())), nil
	case ExpressionValueType.Int64:
		abs := func(value int64) int64 {
			if value < 0 {
				return -value
			}
			return value
		}

		return newValueExpression(ExpressionValueType.Int64, abs(sourceValue.int64Value())), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Abs(sourceValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Abs(sourceValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) ceiling(sourceValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Ceiling\" function source value, first argument, must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	if sourceValue.ValueType().IsIntegerType() {
		return sourceValue, nil
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Ceil(sourceValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Ceil(sourceValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) coalesce(arguments []Expression) (*ValueExpression, error) {
	testValue, err := et.evaluate(arguments[0])

	if err != nil {
		return nil, errors.New("failed while evaluating \"Coalesce\" function argument 0: " + err.Error())
	}

	if !testValue.IsNull() {
		return testValue, nil
	}

	for i := 1; i < len(arguments); i++ {
		listValue, err := et.evaluate(arguments[i])

		if err != nil {
			return nil, errors.New("failed while evaluating \"Coalesce\" function argument " + strconv.Itoa(i) + ": " + err.Error())
		}

		if !listValue.IsNull() {
			return listValue, nil
		}
	}

	return testValue, nil
}

func (et *ExpressionTree) convert(sourceValue *ValueExpression, targetType *ValueExpression) (*ValueExpression, error) {
	if targetType.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Convert\" function target type, second argument, must be a \"String\"")
	}

	if targetType.IsNull() {
		return nil, errors.New("\"Convert\" function target type, second argument, is null")
	}

	targetTypeName := strings.ToUpper(targetType.stringValue())

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
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Contains\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Contains\" function test value, second argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Boolean), nil
	}

	if testValue.IsNull() {
		return False, nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"Contains\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Boolean, strings.Contains(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(testValue.stringValue()))), nil
	}

	return newValueExpression(ExpressionValueType.Boolean, strings.Contains(sourceValue.stringValue(), testValue.stringValue())), nil
}

func (et *ExpressionTree) dateAdd(sourceValue *ValueExpression, addValue *ValueExpression, intervalType *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.DateTime && sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DateAdd\" function source value, first argument, must be a \"DateTime\" or a \"String\"")
	}

	if !addValue.ValueType().IsIntegerType() {
		return nil, errors.New("\"DateAdd\" function add value, second argument, must be an integer type")
	}

	if intervalType.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DateAdd\" function interval type, third argument, must be a \"String\"")
	}

	if addValue.IsNull() {
		return nil, errors.New("\"DateAdd\" function add value, second argument, is null")
	}

	if intervalType.IsNull() {
		return nil, errors.New("\"DateAdd\" function interval type, third argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return sourceValue, nil
	}

	var err error

	// DateTime parameters should support strings as well as literals
	if sourceValue, err = sourceValue.Convert(ExpressionValueType.DateTime); err != nil {
		return nil, errors.New("failed while converting \"DateAdd\" function source value, first argument, to \"DateTime\": " + err.Error())
	}

	var interval TimeIntervalEnum

	if interval, err = ParseTimeInterval(intervalType.stringValue()); err != nil {
		return nil, errors.New("failed while parsing \"DateAdd\" function interval value, third argument, as a valid time interval: " + err.Error())
	}

	value := addValue.integerValue(0)

	switch interval {
	case TimeInterval.Year:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().AddDate(value, 0, 0)), nil
	case TimeInterval.Month:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().AddDate(0, value, 0)), nil
	case TimeInterval.DayOfYear:
		fallthrough
	case TimeInterval.Day:
		fallthrough
	case TimeInterval.WeekDay:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().AddDate(0, 0, value)), nil
	case TimeInterval.Week:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().AddDate(0, 0, value*7)), nil
	case TimeInterval.Hour:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().Add(time.Hour*time.Duration(value))), nil
	case TimeInterval.Minute:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().Add(time.Minute*time.Duration(value))), nil
	case TimeInterval.Second:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().Add(time.Second*time.Duration(value))), nil
	case TimeInterval.Millisecond:
		return newValueExpression(ExpressionValueType.DateTime, sourceValue.dateTimeValue().Add(time.Millisecond*time.Duration(value))), nil
	default:
		return nil, errors.New("unexpected time interval encountered")
	}
}

func (et *ExpressionTree) dateDiff(leftValue *ValueExpression, rightValue *ValueExpression, intervalType *ValueExpression) (*ValueExpression, error) {
	if leftValue.ValueType() != ExpressionValueType.DateTime && leftValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DateDiff\" function left value, first argument, must be a \"DateTime\" or a \"String\"")
	}

	if rightValue.ValueType() != ExpressionValueType.DateTime && rightValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DateDiff\" function right value, second argument, must be a \"DateTime\" or a \"String\"")
	}

	if intervalType.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DateDiff\" function interval type, third argument, must be a \"String\"")
	}

	if intervalType.IsNull() {
		return nil, errors.New("\"DateDiff\" function interval type, third argument, is null")
	}

	// If either test value is Null, result is Null
	if leftValue.IsNull() || rightValue.IsNull() {
		return NullValue(ExpressionValueType.Int32), nil
	}

	var err error

	// DateTime parameters should support strings as well as literals
	if leftValue, err = leftValue.Convert(ExpressionValueType.DateTime); err != nil {
		return nil, errors.New("failed while converting \"DateDiff\" function left value, first argument, to \"DateTime\": " + err.Error())
	}

	if rightValue, err = rightValue.Convert(ExpressionValueType.DateTime); err != nil {
		return nil, errors.New("failed while converting \"DateDiff\" function right value, second argument, to \"DateTime\": " + err.Error())
	}

	var interval TimeIntervalEnum

	if interval, err = ParseTimeInterval(intervalType.stringValue()); err != nil {
		return nil, errors.New("failed while parsing \"DateDiff\" function interval value, third argument, as a valid time interval: " + err.Error())
	}

	rightDate := rightValue.dateTimeValue()
	leftDate := leftValue.dateTimeValue()

	if interval < TimeInterval.DayOfYear {
		switch interval {
		case TimeInterval.Year:
			return newValueExpression(ExpressionValueType.Int32, int32(rightDate.Year()-leftDate.Year())), nil
		case TimeInterval.Month:
			months := (rightDate.Year() - leftDate.Year()) * 12
			months += int(rightDate.Month() - leftDate.Month())
			return newValueExpression(ExpressionValueType.Int32, int32(months)), nil
		default:
			return nil, errors.New("unexpected time interval encountered")
		}
	}

	delta := rightDate.Sub(leftDate)

	switch interval {
	case TimeInterval.DayOfYear:
		fallthrough
	case TimeInterval.Day:
		fallthrough
	case TimeInterval.WeekDay:
		return newValueExpression(ExpressionValueType.Int32, int32(delta.Hours()/24.0)), nil
	case TimeInterval.Week:
		return newValueExpression(ExpressionValueType.Int32, int32(delta.Hours()/24.0/7.0)), nil
	case TimeInterval.Hour:
		return newValueExpression(ExpressionValueType.Int32, int32(delta.Hours())), nil
	case TimeInterval.Minute:
		return newValueExpression(ExpressionValueType.Int32, int32(delta.Minutes())), nil
	case TimeInterval.Second:
		return newValueExpression(ExpressionValueType.Int32, int32(delta.Seconds())), nil
	case TimeInterval.Millisecond:
		return newValueExpression(ExpressionValueType.Int32, int32(delta.Milliseconds())), nil
	default:
		return nil, errors.New("unexpected time interval encountered")
	}
}

func (et *ExpressionTree) datePart(sourceValue *ValueExpression, intervalType *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.DateTime && sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DatePart\" function source value, first argument, must be a \"DateTime\" or a \"String\"")
	}

	if intervalType.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"DatePart\" function interval type, second argument, must be a \"String\"")
	}

	if intervalType.IsNull() {
		return nil, errors.New("\"DatePart\" function interval type, second argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Int32), nil
	}

	var err error

	// DateTime parameters should support strings as well as literals
	if sourceValue, err = sourceValue.Convert(ExpressionValueType.DateTime); err != nil {
		return nil, errors.New("failed while converting \"DatePart\" function source value, first argument, to \"DateTime\": " + err.Error())
	}

	var interval TimeIntervalEnum

	if interval, err = ParseTimeInterval(intervalType.stringValue()); err != nil {
		return nil, errors.New("failed while parsing \"DatePart\" function interval value, second argument, as a valid time interval: " + err.Error())
	}

	switch interval {
	case TimeInterval.Year:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Year())), nil
	case TimeInterval.Month:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Month())), nil
	case TimeInterval.DayOfYear:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().YearDay())), nil
	case TimeInterval.Day:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Day())), nil
	case TimeInterval.WeekDay:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Weekday()+1)), nil
	case TimeInterval.Week:
		_, week := sourceValue.dateTimeValue().ISOWeek()
		return newValueExpression(ExpressionValueType.Int32, int32(week)), nil
	case TimeInterval.Hour:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Hour())), nil
	case TimeInterval.Minute:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Minute())), nil
	case TimeInterval.Second:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Second())), nil
	case TimeInterval.Millisecond:
		return newValueExpression(ExpressionValueType.Int32, int32(sourceValue.dateTimeValue().Nanosecond()/1e6)), nil
	default:
		return nil, errors.New("unexpected time interval encountered")
	}
}

func (et *ExpressionTree) endsWith(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"EndsWith\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"EndsWith\" function test value, second argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Boolean), nil
	}

	if testValue.IsNull() {
		return False, nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"EndsWith\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Boolean, strings.HasSuffix(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(testValue.stringValue()))), nil
	}

	return newValueExpression(ExpressionValueType.Boolean, strings.HasSuffix(sourceValue.stringValue(), testValue.stringValue())), nil
}

func (et *ExpressionTree) floor(sourceValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Floor\" function source value, first argument, must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	if sourceValue.ValueType().IsIntegerType() {
		return sourceValue, nil
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Floor(sourceValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Floor(sourceValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) iif(testValue *ValueExpression, leftResultValue Expression, rightResultValue Expression) (*ValueExpression, error) {
	if testValue.ValueType() != ExpressionValueType.Boolean {
		return nil, errors.New("\"IIf\" function test value, first argument, must be a \"Boolean\"")
	}

	var result *ValueExpression
	var err error

	// Null test expression evaluates to false, that is, right expression
	if testValue.booleanValue() {
		result, err = et.evaluate(leftResultValue)

		if err != nil {
			return nil, errors.New("failed while evaluating \"IIf\" function left result value, second argument: " + err.Error())
		}

		return result, nil
	}

	result, err = et.evaluate(rightResultValue)

	if err != nil {
		return nil, errors.New("failed while evaluating \"IIf\" function right result value, third argument: " + err.Error())
	}

	return result, nil
}

func (et *ExpressionTree) indexOf(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"IndexOf\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"IndexOf\" function test value, second argument, must be a \"String\"")
	}

	if testValue.IsNull() {
		return nil, errors.New("\"IndexOf\" function test value, second argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Int32), nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"IndexOf\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Int32, int32(strings.Index(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(testValue.stringValue())))), nil
	}

	return newValueExpression(ExpressionValueType.Int32, int32(strings.Index(sourceValue.stringValue(), testValue.stringValue()))), nil
}

func (et *ExpressionTree) isDate(testValue *ValueExpression) *ValueExpression {
	if testValue.IsNull() {
		return False
	}

	if testValue.ValueType() == ExpressionValueType.DateTime {
		return True
	}

	if testValue.ValueType() == ExpressionValueType.String {
		if _, err := dateparse.ParseAny(testValue.stringValue()); err == nil {
			return True
		}
	}

	return False
}

func (et *ExpressionTree) isInteger(testValue *ValueExpression) *ValueExpression {
	if testValue.IsNull() {
		return False
	}

	if testValue.ValueType().IsIntegerType() {
		return True
	}

	if testValue.ValueType() == ExpressionValueType.String {
		if _, err := strconv.Atoi(testValue.stringValue()); err == nil {
			return True
		}
	}

	return False
}

func (et *ExpressionTree) isGuid(testValue *ValueExpression) *ValueExpression {
	if testValue.IsNull() {
		return False
	}

	if testValue.ValueType() == ExpressionValueType.Guid {
		return True
	}

	if testValue.ValueType() == ExpressionValueType.String {
		if _, err := guid.Parse(testValue.stringValue()); err == nil {
			return True
		}
	}

	return False
}

func (et *ExpressionTree) isNull(testValue *ValueExpression, defaultValue *ValueExpression) (*ValueExpression, error) {
	if defaultValue.IsNull() {
		return nil, errors.New("\"IsNull\" default value, second argument, is null")
	}

	if testValue.IsNull() {
		return defaultValue, nil
	}

	return testValue, nil
}

func (et *ExpressionTree) isNumeric(testValue *ValueExpression) *ValueExpression {
	if testValue.IsNull() {
		return False
	}

	if testValue.ValueType().IsNumericType() {
		return True
	}

	if testValue.ValueType() == ExpressionValueType.String {
		if _, err := strconv.ParseFloat(testValue.stringValue(), 64); err == nil {
			return True
		}
	}

	return False
}

func (et *ExpressionTree) lastIndexOf(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"LastIndexOf\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"LastIndexOf\" function test value, second argument, must be a \"String\"")
	}

	if testValue.IsNull() {
		return nil, errors.New("\"LastIndexOf\" function test value, second argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Int32), nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"LastIndexOf\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Int32, int32(strings.LastIndex(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(testValue.stringValue())))), nil
	}

	return newValueExpression(ExpressionValueType.Int32, int32(strings.LastIndex(sourceValue.stringValue(), testValue.stringValue()))), nil
}

func (et *ExpressionTree) len(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Len\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Int32), nil
	}

	return newValueExpression(ExpressionValueType.Int32, int32(len(sourceValue.stringValue()))), nil
}

func (et *ExpressionTree) lower(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Lower\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.String), nil
	}

	return newValueExpression(ExpressionValueType.String, strings.ToLower(sourceValue.stringValue())), nil
}

func (et *ExpressionTree) maxOf(arguments []Expression) (*ValueExpression, error) {
	testValue, err := et.evaluate(arguments[0])

	if err != nil {
		return nil, errors.New("failed while evaluating \"MaxOf\" function argument 0: " + err.Error())
	}

	for i := 1; i < len(arguments); i++ {
		nextValue, err := et.evaluate(arguments[i])

		if err != nil {
			return nil, errors.New("failed while evaluating \"MaxOf\" function argument " + strconv.Itoa(i) + ": " + err.Error())
		}

		valueType, err := ExpressionOperatorType.GreaterThan.deriveComparisonOperationValueType(testValue.ValueType(), nextValue.ValueType())

		if err != nil {
			return nil, errors.New("failed while deriving \"MaxOf\" function greater than comparison operation value type: " + err.Error())
		}

		result, err := et.greaterThanOp(nextValue, testValue, valueType)

		if err != nil {
			return nil, errors.New("failed while executing \">\" comparison operation in \"MaxOf\" function: " + err.Error())
		}

		if result.booleanValue() || (testValue.IsNull() && !nextValue.IsNull()) {
			testValue = nextValue
		}
	}

	return testValue, nil
}

func (et *ExpressionTree) minOf(arguments []Expression) (*ValueExpression, error) {
	testValue, err := et.evaluate(arguments[0])

	if err != nil {
		return nil, errors.New("failed while evaluating \"MinOf\" function argument 0: " + err.Error())
	}

	for i := 1; i < len(arguments); i++ {
		nextValue, err := et.evaluate(arguments[i])

		if err != nil {
			return nil, errors.New("failed while evaluating \"MinOf\" function argument " + strconv.Itoa(i) + ": " + err.Error())
		}

		valueType, err := ExpressionOperatorType.LessThan.deriveComparisonOperationValueType(testValue.ValueType(), nextValue.ValueType())

		if err != nil {
			return nil, errors.New("failed while deriving \"MinOf\" function less than comparison operation value type: " + err.Error())
		}

		result, err := et.greaterThanOp(nextValue, testValue, valueType)

		if err != nil {
			return nil, errors.New("failed while executing \"<\" comparison operation in \"MinOf\" function: " + err.Error())
		}

		if result.booleanValue() || (testValue.IsNull() && !nextValue.IsNull()) {
			testValue = nextValue
		}
	}

	return testValue, nil
}

func (et *ExpressionTree) now() (*ValueExpression, error) {
	return newValueExpression(ExpressionValueType.DateTime, time.Now()), nil
}

func (et *ExpressionTree) nthIndexOf(sourceValue *ValueExpression, testValue *ValueExpression, indexValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"NthIndexOf\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"NthIndexOf\" function test value, second argument, must be a \"String\"")
	}

	if !indexValue.ValueType().IsIntegerType() {
		return nil, errors.New("\"NthIndexOf\" function index value, third argument, must be an integer type")
	}

	if testValue.IsNull() {
		return nil, errors.New("\"NthIndexOf\" function test value, second argument, is null")
	}

	if indexValue.IsNull() {
		return nil, errors.New("\"NthIndexOf\" function index value, third argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Int32), nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"NthIndexOf\" function optional ignore case value, fourth argument, to \"Boolean\": " + err.Error())
	}

	var source, test string

	if ignoreCase.booleanValue() {
		source = strings.ToUpper(sourceValue.stringValue())
		test = strings.ToUpper(testValue.stringValue())
	} else {
		source = sourceValue.stringValue()
		test = testValue.stringValue()
	}

	return newValueExpression(ExpressionValueType.Int32, int32(findNthIndex(source, test, indexValue.integerValue(-1)))), nil
}

func (et *ExpressionTree) power(sourceValue *ValueExpression, exponentValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Power\" function source value, first argument, must be numeric")
	}

	if !exponentValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Power\" function exponent value, second argument, must be numeric")
	}

	// If source value or exponent value is Null, result is Null
	if sourceValue.IsNull() || exponentValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	valueType, err := ExpressionOperatorType.Multiply.deriveArithmeticOperationValueType(sourceValue.ValueType(), exponentValue.ValueType())

	if err != nil {
		return nil, errors.New("failed while deriving \"Power\" function multiplicative arithmetic operation value type: " + err.Error())
	}

	if sourceValue, err = sourceValue.Convert(valueType); err != nil {
		return nil, errors.New("failed while converting \"Power\" function source value, first argument, to \"" + valueType.String() + "\": " + err.Error())
	}

	if exponentValue, err = sourceValue.Convert(valueType); err != nil {
		return nil, errors.New("failed while converting \"Power\" function exponent value, second argument, to \"" + valueType.String() + "\": " + err.Error())
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, math.Pow(float64(sourceValue.booleanValueAsInt()), float64(exponentValue.booleanValueAsInt())) != 0.0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(math.Pow(float64(sourceValue.int32Value()), float64(exponentValue.int32Value())))), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(math.Pow(float64(sourceValue.int64Value()), float64(exponentValue.int64Value())))), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Pow(sourceValue.decimalValue(), exponentValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Pow(sourceValue.doubleValue(), exponentValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) regExMatch(regexValue *ValueExpression, testValue *ValueExpression) (*ValueExpression, error) {
	return et.evaluateRegEx("RegExMatch", regexValue, testValue, false)
}

func (et *ExpressionTree) regExVal(regexValue *ValueExpression, testValue *ValueExpression) (*ValueExpression, error) {
	return et.evaluateRegEx("RegExVal", regexValue, testValue, true)
}

func (et *ExpressionTree) evaluateRegEx(functionName string, regexValue *ValueExpression, testValue *ValueExpression, returnMatchedValue bool) (*ValueExpression, error) {
	if regexValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"" + functionName + "\" function expression value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"" + functionName + "\" function test value, second argument, must be a \"String\"")
	}

	// If regex value or test value is Null, result is Null
	if regexValue.IsNull() || testValue.IsNull() {
		if returnMatchedValue {
			return NullValue(ExpressionValueType.String), nil
		}

		return NullValue(ExpressionValueType.Boolean), nil
	}

	regex, err := regexp.Compile(regexValue.stringValue())

	if err != nil {
		return nil, errors.New("failed while compiling \"" + functionName + "\" function expression value, first argument: " + err.Error())
	}

	testText := testValue.stringValue()
	result := regex.FindStringIndex(testText)

	if returnMatchedValue {
		// RegExVal returns any left-most matched value, otherwise empty string
		if result == nil {
			return EmptyString, nil
		}

		return newValueExpression(ExpressionValueType.String, testText[result[0]:result[1]]), nil
	}

	// RegExMatch returns boolean result determining if there was a matched value
	if result == nil {
		return False, nil
	}

	return True, nil
}

func (et *ExpressionTree) replace(sourceValue *ValueExpression, testValue *ValueExpression, replaceValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Replace\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Replace\" function test value, second argument, must be a \"String\"")
	}

	if replaceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Replace\" function replace value, third argument, must be a \"String\"")
	}

	if testValue.IsNull() {
		return nil, errors.New("\"Replace\" function test value, second argument, is null")
	}

	if replaceValue.IsNull() {
		return nil, errors.New("\"Replace\" function replace value, third argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return sourceValue, nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"Replace\" function optional ignore case value, fourth argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		regex, err := regexp.Compile("(?i)" + regexp.QuoteMeta(testValue.stringValue()))

		if err != nil {
			return nil, errors.New("failed while compiling \"Replace\" function case-insensitive RegEx replace expression for test value, second argument: " + err.Error())
		}

		return newValueExpression(ExpressionValueType.String, regex.ReplaceAllString(sourceValue.stringValue(), replaceValue.stringValue())), nil
	}

	return newValueExpression(ExpressionValueType.String, strings.ReplaceAll(sourceValue.stringValue(), testValue.stringValue(), replaceValue.stringValue())), nil
}

func (et *ExpressionTree) reverse(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Reverse\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.String), nil
	}

	chars := []rune(sourceValue.stringValue())

	for i, j := 0, len(chars)-1; i < len(chars)/2; i, j = i+1, j-1 {
		chars[i], chars[j] = chars[j], chars[i]
	}

	return newValueExpression(ExpressionValueType.String, string(chars)), nil
}

func (et *ExpressionTree) round(sourceValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Round\" function source value, first argument, must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	if sourceValue.ValueType().IsIntegerType() {
		return sourceValue, nil
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Round(sourceValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Round(sourceValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) split(sourceValue *ValueExpression, delimiterValue *ValueExpression, indexValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Split\" function source value, first argument, must be a \"String\"")
	}

	if delimiterValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Split\" function delimeter value, second argument, must be a \"String\"")
	}

	if !indexValue.ValueType().IsIntegerType() {
		return nil, errors.New("\"Split\" function index value, third argument, must be an integer type")
	}

	if delimiterValue.IsNull() {
		return nil, errors.New("\"Split\" delimiter test value, second argument, is null")
	}

	if indexValue.IsNull() {
		return nil, errors.New("\"Split\" function index value, third argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return sourceValue, nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"Split\" function optional ignore case value, fourth argument, to \"Boolean\": " + err.Error())
	}

	index := indexValue.integerValue(-1)
	var result []int

	if ignoreCase.booleanValue() {
		result = splitNthIndex(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(delimiterValue.stringValue()), index)
	} else {
		result = splitNthIndex(sourceValue.stringValue(), delimiterValue.stringValue(), index)
	}

	if result == nil {
		return EmptyString, nil
	}

	return newValueExpression(ExpressionValueType.String, sourceValue.stringValue()[result[0]:result[1]]), nil
}

func (et *ExpressionTree) sqrt(sourceValue *ValueExpression) (*ValueExpression, error) {
	if !sourceValue.ValueType().IsNumericType() {
		return nil, errors.New("\"Sqrt\" function source value, first argument, must be numeric")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(sourceValue.ValueType()), nil
	}

	switch sourceValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Double, math.Sqrt(float64(sourceValue.booleanValueAsInt()))), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Double, math.Sqrt(float64(sourceValue.int32Value()))), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Double, math.Sqrt(float64(sourceValue.int64Value()))), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, math.Sqrt(sourceValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, math.Sqrt(sourceValue.doubleValue())), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) startsWith(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"StartsWith\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"StartsWith\" function test value, second argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.Boolean), nil
	}

	if testValue.IsNull() {
		return False, nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"StartsWith\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Boolean, strings.HasPrefix(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(testValue.stringValue()))), nil
	}

	return newValueExpression(ExpressionValueType.Boolean, strings.HasPrefix(sourceValue.stringValue(), testValue.stringValue())), nil
}

func (et *ExpressionTree) strCount(sourceValue *ValueExpression, testValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"StrCount\" function source value, first argument, must be a \"String\"")
	}

	if testValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"StrCount\" function test value, second argument, must be a \"String\"")
	}

	if sourceValue.IsNull() || testValue.IsNull() {
		return newValueExpression(ExpressionValueType.Int32, int32(0)), nil
	}

	findValue := testValue.stringValue()

	if len(findValue) == 0 {
		return newValueExpression(ExpressionValueType.Int32, int32(0)), nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"StrCount\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Int32, int32(strings.Count(strings.ToUpper(sourceValue.stringValue()), strings.ToUpper(testValue.stringValue())))), nil
	}

	return newValueExpression(ExpressionValueType.Int32, int32(strings.Count(sourceValue.stringValue(), testValue.stringValue()))), nil
}

func (et *ExpressionTree) strCmp(leftValue *ValueExpression, rightValue *ValueExpression, ignoreCase *ValueExpression) (*ValueExpression, error) {
	if leftValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"StrCmp\" function left value, first argument, must be a \"String\"")
	}

	if rightValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"StrCmp\" function right value, second argument, must be a \"String\"")
	}

	// If both left and right values are null, values are considered equal
	if leftValue.IsNull() && rightValue.IsNull() {
		return newValueExpression(ExpressionValueType.Int32, int32(0)), nil
	}

	// If left value is null, right non-null value will be considered greater
	if leftValue.IsNull() {
		return newValueExpression(ExpressionValueType.Int32, int32(1)), nil
	}

	// If right value is null, left non-null value will be considered greater
	if rightValue.IsNull() {
		return newValueExpression(ExpressionValueType.Int32, int32(-1)), nil
	}

	var err error

	if ignoreCase, err = ignoreCase.Convert(ExpressionValueType.Boolean); err != nil {
		return nil, errors.New("failed while converting \"StrCmp\" function optional ignore case value, third argument, to \"Boolean\": " + err.Error())
	}

	if ignoreCase.booleanValue() {
		return newValueExpression(ExpressionValueType.Int32, int32(strings.Compare(strings.ToUpper(leftValue.stringValue()), strings.ToUpper(rightValue.stringValue())))), nil
	}

	return newValueExpression(ExpressionValueType.Int32, int32(strings.Compare(leftValue.stringValue(), rightValue.stringValue()))), nil
}

func (et *ExpressionTree) subStr(sourceValue *ValueExpression, indexValue *ValueExpression, lengthValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"SubStr\" function source value, first argument, must be a \"String\"")
	}

	if !indexValue.ValueType().IsIntegerType() {
		return nil, errors.New("\"SubStr\" function index value, second argument, must be an integer type")
	}

	if !lengthValue.ValueType().IsIntegerType() {
		return nil, errors.New("\"SubStr\" function length value, third argument, must be an integer type")
	}

	if indexValue.IsNull() {
		return nil, errors.New("\"SubStr\" function index value, second argument, is null")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return sourceValue, nil
	}

	sourceText := sourceValue.stringValue()
	index := indexValue.integerValue(0)

	if index < 0 || index >= len(sourceText) {
		return EmptyString, nil
	}

	if !lengthValue.IsNull() {
		length := lengthValue.integerValue(0)

		if length <= 0 {
			return EmptyString, nil
		}

		if index+length < len(sourceText) {
			return newValueExpression(ExpressionValueType.String, sourceText[index:index+length]), nil
		}
	}

	return newValueExpression(ExpressionValueType.String, sourceText[index:]), nil
}

func (et *ExpressionTree) trim(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Trim\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.String), nil
	}

	return newValueExpression(ExpressionValueType.String, strings.TrimSpace(sourceValue.stringValue())), nil
}

func (et *ExpressionTree) trimLeft(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"TrimLeft\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.String), nil
	}

	return newValueExpression(ExpressionValueType.String, strings.TrimLeft(sourceValue.stringValue(), whitespace)), nil
}

func (et *ExpressionTree) trimRight(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"TrimRight\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.String), nil
	}

	return newValueExpression(ExpressionValueType.String, strings.TrimRight(sourceValue.stringValue(), whitespace)), nil
}

func (et *ExpressionTree) upper(sourceValue *ValueExpression) (*ValueExpression, error) {
	if sourceValue.ValueType() != ExpressionValueType.String {
		return nil, errors.New("\"Upper\" function source value, first argument, must be a \"String\"")
	}

	// If source value is Null, result is Null
	if sourceValue.IsNull() {
		return NullValue(ExpressionValueType.String), nil
	}

	return newValueExpression(ExpressionValueType.String, strings.ToUpper(sourceValue.stringValue())), nil
}

func (et *ExpressionTree) utcNow() (*ValueExpression, error) {
	return newValueExpression(ExpressionValueType.DateTime, time.Now().UTC()), nil
}

// Filter Expression Operator Implementations

func convertOperands(leftValue **ValueExpression, rightValue **ValueExpression, valueType ExpressionValueTypeEnum) error {
	var err error

	if *leftValue, err = (*leftValue).Convert(valueType); err != nil {
		return errors.New("failed while converting left operand, to \"" + valueType.String() + "\": " + err.Error())
	}

	if *rightValue, err = (*rightValue).Convert(valueType); err != nil {
		return errors.New("failed while converting right operand, to \"" + valueType.String() + "\": " + err.Error())
	}

	return nil
}

func (et *ExpressionTree) multiplyOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	// If left or right value is Null, result is Null
	if leftValue.IsNull() || rightValue.IsNull() {
		return NullValue(valueType), nil
	}

	if err := convertOperands(&leftValue, &rightValue, valueType); err != nil {
		return nil, errors.New("multiplication \"*\" operator " + err.Error())
	}

	switch valueType {
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()*rightValue.int32Value())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()*rightValue.int64Value())), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, float64(leftValue.decimalValue()*rightValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, float64(leftValue.doubleValue()*rightValue.doubleValue())), nil
	case ExpressionValueType.Boolean:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply multiplication \"*\" operator to \"" + valueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) divideOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	// If left or right value is Null, result is Null
	if leftValue.IsNull() || rightValue.IsNull() {
		return NullValue(valueType), nil
	}

	if err := convertOperands(&leftValue, &rightValue, valueType); err != nil {
		return nil, errors.New("division \"/\" operator " + err.Error())
	}

	switch valueType {
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()/rightValue.int32Value())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()/rightValue.int64Value())), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, float64(leftValue.decimalValue()/rightValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, float64(leftValue.doubleValue()/rightValue.doubleValue())), nil
	case ExpressionValueType.Boolean:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply division \"/\" operator to \"" + valueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) modulusOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	// If left or right value is Null, result is Null
	if leftValue.IsNull() || rightValue.IsNull() {
		return NullValue(valueType), nil
	}

	if err := convertOperands(&leftValue, &rightValue, valueType); err != nil {
		return nil, errors.New("modulus \"%\" operator " + err.Error())
	}

	switch valueType {
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()%rightValue.int32Value())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()%rightValue.int64Value())), nil
	case ExpressionValueType.Boolean:
		fallthrough
	case ExpressionValueType.Decimal:
		fallthrough
	case ExpressionValueType.Double:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply modulus \"%\" operator to \"" + valueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) addOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	// If left or right value is Null, result is Null
	if leftValue.IsNull() || rightValue.IsNull() {
		return NullValue(valueType), nil
	}

	if err := convertOperands(&leftValue, &rightValue, valueType); err != nil {
		return nil, errors.New("addition \"+\" operator " + err.Error())
	}

	switch valueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, leftValue.booleanValueAsInt()+rightValue.booleanValueAsInt() != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()+rightValue.int32Value())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()+rightValue.int64Value())), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, float64(leftValue.decimalValue()+rightValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, float64(leftValue.doubleValue()+rightValue.doubleValue())), nil
	case ExpressionValueType.String:
		return newValueExpression(ExpressionValueType.String, leftValue.stringValue()+rightValue.stringValue()), nil
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply addition \"+\" operator to \"" + valueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) subtractOp(leftValue *ValueExpression, rightValue *ValueExpression, valueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	// If left or right value is Null, result is Null
	if leftValue.IsNull() || rightValue.IsNull() {
		return NullValue(valueType), nil
	}

	if err := convertOperands(&leftValue, &rightValue, valueType); err != nil {
		return nil, errors.New("subtraction \"-\" operator " + err.Error())
	}

	switch valueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, leftValue.booleanValueAsInt()-rightValue.booleanValueAsInt() != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()-rightValue.int32Value())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()-rightValue.int64Value())), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, float64(leftValue.decimalValue()-rightValue.decimalValue())), nil
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, float64(leftValue.doubleValue()-rightValue.doubleValue())), nil
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply subtraction \"-\" operator to \"" + valueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) bitShiftLeftOp(leftValue *ValueExpression, rightValue *ValueExpression) (*ValueExpression, error) {
	// If left is Null, result is Null
	if leftValue.IsNull() {
		return leftValue, nil
	}

	if !rightValue.ValueType().IsIntegerType() {
		return nil, errors.New("BitShift operation shift value must be an integer")
	}

	if rightValue.IsNull() {
		return nil, errors.New("BitShift operation shift value is null")
	}

	switch leftValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, leftValue.booleanValueAsInt()<<rightValue.integerValue(0) != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()<<rightValue.integerValue(0))), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()<<rightValue.integerValue(0))), nil
	case ExpressionValueType.Decimal:
		fallthrough
	case ExpressionValueType.Double:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply left bit-shift \"<<\" operator to \"" + leftValue.ValueType().String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) bitShiftRightOp(leftValue *ValueExpression, rightValue *ValueExpression) (*ValueExpression, error) {
	// If left is Null, result is Null
	if leftValue.IsNull() {
		return leftValue, nil
	}

	if !rightValue.ValueType().IsIntegerType() {
		return nil, errors.New("BitShift operation shift value must be an integer")
	}

	if rightValue.IsNull() {
		return nil, errors.New("BitShift operation shift value is null")
	}

	switch leftValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, leftValue.booleanValueAsInt()>>rightValue.integerValue(0) != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, int32(leftValue.int32Value()>>rightValue.integerValue(0))), nil
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, int64(leftValue.int64Value()>>rightValue.integerValue(0))), nil
	case ExpressionValueType.Decimal:
		fallthrough
	case ExpressionValueType.Double:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		fallthrough
	case ExpressionValueType.Undefined:
		return nil, errors.New("cannot apply right bit-shift \">>\" operator to \"" + leftValue.ValueType().String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
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

// https://play.golang.org/p/-zlKH7mfboa
func findNthIndex(source, test string, index int) int {
	var result int

	for i := 0; i < index+1; i++ {
		location := strings.Index(source, test)

		if location == -1 {
			result = 0
			break
		}

		location++
		result += location
		source = source[location:]
	}

	return result - 1
}

// https://play.golang.org/p/_7nc06CqjKi
func splitNthIndex(source, test string, index int) []int {
	firstIndex := findNthIndex(source, test, index-1)
	secondIndex := findNthIndex(source, test, index)

	if firstIndex <= 0 && secondIndex <= 0 {
		return nil
	}

	if firstIndex <= 0 {
		return []int{0, secondIndex}
	}

	if secondIndex <= 0 {
		return []int{firstIndex + len(test), len(source)}
	}

	return []int{firstIndex + len(test), secondIndex}
}
