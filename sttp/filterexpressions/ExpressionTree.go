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
		var valueExpression *ValueExpression
		var err error

		if valueExpression, err = GetValueExpression(expression); err != nil {
			return nil, err
		}

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
	var unaryExpression *UnaryExpression
	var err error

	if unaryExpression, err = GetUnaryExpression(expression); err != nil {
		return nil, err
	}

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
		return unaryExpression.unaryBoolean(unaryValue.BooleanValue())
	case ExpressionValueType.Int32:
		return unaryExpression.unaryInt32(unaryValue.Int32Value())
	case ExpressionValueType.Int64:
		return unaryExpression.unaryInt64(unaryValue.Int64Value())
	case ExpressionValueType.Decimal:
		return unaryExpression.unaryDecimal(unaryValue.DecimalValue())
	case ExpressionValueType.Double:
		return unaryExpression.unaryDouble(unaryValue.DoubleValue())
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
	var columnExpression *ColumnExpression
	var err error

	if columnExpression, err = GetColumnExpression(expression); err != nil {
		return nil, err
	}

	var column *data.DataColumn

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
		return NewValueExpression(valueType, nil), nil
	}

	return NewValueExpression(valueType, value), nil
}

func (et *ExpressionTree) evaluateInList(expression Expression) (*ValueExpression, error) {
	var inListExpression *InListExpression
	var err error

	if inListExpression, err = GetInListExpression(expression); err != nil {
		return nil, err
	}

	var inListValue *ValueExpression

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
	var resultVal bool

	for i := 0; i < len(arguments); i++ {
		if argumentValue, err = et.evaluate(arguments[i]); err != nil {
			return nil, err
		}

		if valueType, err = et.deriveComparisonOperationValueType(ExpressionOperatorType.Equal, inListValue.ValueType(), argumentValue.ValueType()); err != nil {
			return nil, err
		}

		if result, err = et.equalOp(inListValue, argumentValue, valueType, exactMatch); err != nil {
			return nil, err
		}

		if resultVal, err = result.BooleanValue(); err != nil {
			return nil, err
		}

		if resultVal {
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
	return nil, nil
}

func (et *ExpressionTree) evaluateOperator(expression Expression) (*ValueExpression, error) {
	var operatorExpression *OperatorExpression
	var err error

	if operatorExpression, err = GetOperatorExpression(expression); err != nil {
		return nil, err
	}

	var leftValue *ValueExpression

	if leftValue, err = et.evaluate(operatorExpression.LeftValue()); err != nil {
		return nil, err
	}

	var rightValue *ValueExpression

	if rightValue, err = et.evaluate(operatorExpression.RightValue()); err != nil {
		return nil, err
	}

	var valueType ExpressionValueTypeEnum

	if valueType, err = et.deriveOperationValueType(ExpressionOperatorType.Equal, leftValue.ValueType(), rightValue.ValueType()); err != nil {
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

// Operation Value Type Selectors

func (et *ExpressionTree) deriveOperationValueType(operationType ExpressionOperatorTypeEnum, leftValueType ExpressionValueTypeEnum, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch operationType {
	case ExpressionOperatorType.Multiply:
		fallthrough
	case ExpressionOperatorType.Divide:
		fallthrough
	case ExpressionOperatorType.Add:
		fallthrough
	case ExpressionOperatorType.Subtract:
		return et.deriveArithmeticOperationValueType(operationType, leftValueType, rightValueType)
	case ExpressionOperatorType.Modulus:
		fallthrough
	case ExpressionOperatorType.BitwiseAnd:
		fallthrough
	case ExpressionOperatorType.BitwiseOr:
		fallthrough
	case ExpressionOperatorType.BitwiseXor:
		return et.deriveIntegerOperationValueType(operationType, leftValueType, rightValueType)
	case ExpressionOperatorType.LessThan:
		fallthrough
	case ExpressionOperatorType.LessThanOrEqual:
		fallthrough
	case ExpressionOperatorType.GreaterThan:
		fallthrough
	case ExpressionOperatorType.GreaterThanOrEqual:
		fallthrough
	case ExpressionOperatorType.Equal:
		fallthrough
	case ExpressionOperatorType.EqualExactMatch:
		fallthrough
	case ExpressionOperatorType.NotEqual:
		fallthrough
	case ExpressionOperatorType.NotEqualExactMatch:
		return et.deriveComparisonOperationValueType(operationType, leftValueType, rightValueType)
	case ExpressionOperatorType.And:
		fallthrough
	case ExpressionOperatorType.Or:
		return et.deriveBooleanOperationValueType(operationType, leftValueType, rightValueType)
	case ExpressionOperatorType.BitShiftLeft:
		fallthrough
	case ExpressionOperatorType.BitShiftRight:
		fallthrough
	case ExpressionOperatorType.IsNull:
		fallthrough
	case ExpressionOperatorType.IsNotNull:
		fallthrough
	case ExpressionOperatorType.Like:
		fallthrough
	case ExpressionOperatorType.LikeExactMatch:
		fallthrough
	case ExpressionOperatorType.NotLike:
		fallthrough
	case ExpressionOperatorType.NotLikeExactMatch:
		return leftValueType, nil
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression operator type encountered")
	}
}

func (et *ExpressionTree) deriveArithmeticOperationValueType(operationType ExpressionOperatorTypeEnum, leftValueType ExpressionValueTypeEnum, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch leftValueType {
	case ExpressionValueType.Boolean:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			return ExpressionValueType.Boolean, nil
		case ExpressionValueType.Int32:
			return ExpressionValueType.Int32, nil
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.String:
			if operationType == ExpressionOperatorType.Add {
				return ExpressionValueType.String, nil
			}
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Boolean\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int32:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			return ExpressionValueType.Int32, nil
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.String:
			if operationType == ExpressionOperatorType.Add {
				return ExpressionValueType.String, nil
			}
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Int32\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int64:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.String:
			if operationType == ExpressionOperatorType.Add {
				return ExpressionValueType.String, nil
			}
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Int64\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Decimal:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.String:
			if operationType == ExpressionOperatorType.Add {
				return ExpressionValueType.String, nil
			}
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Decimal\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Double:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.String:
			if operationType == ExpressionOperatorType.Add {
				return ExpressionValueType.String, nil
			}
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Double\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.String:
		if operationType == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"" + leftValueType.String() + "\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) deriveIntegerOperationValueType(operationType ExpressionOperatorTypeEnum, leftValueType ExpressionValueTypeEnum, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch leftValueType {
	case ExpressionValueType.Boolean:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			return ExpressionValueType.Boolean, nil
		case ExpressionValueType.Int32:
			return ExpressionValueType.Int32, nil
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			fallthrough
		case ExpressionValueType.String:
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Boolean\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int32:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			return ExpressionValueType.Int32, nil
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			fallthrough
		case ExpressionValueType.String:
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Int32\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int64:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			fallthrough
		case ExpressionValueType.String:
			fallthrough
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Int64\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Decimal:
		fallthrough
	case ExpressionValueType.Double:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"" + leftValueType.String() + "\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) deriveComparisonOperationValueType(operationType ExpressionOperatorTypeEnum, leftValueType ExpressionValueTypeEnum, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch leftValueType {
	case ExpressionValueType.Boolean:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.Boolean, nil
		case ExpressionValueType.Int32:
			return ExpressionValueType.Int32, nil
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Boolean\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int32:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.Int32, nil
		case ExpressionValueType.Int64:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Int32\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int64:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.Int64, nil
		case ExpressionValueType.Decimal:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Int64\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Decimal:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.Decimal, nil
		case ExpressionValueType.Double:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Decimal\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Double:
		switch rightValueType {
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.Double, nil
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Double\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.String:
		return leftValueType, nil
	case ExpressionValueType.Guid:
		switch rightValueType {
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.Guid, nil
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			fallthrough
		case ExpressionValueType.DateTime:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"Guid\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.DateTime:
		switch rightValueType {
		case ExpressionValueType.DateTime:
			fallthrough
		case ExpressionValueType.String:
			return ExpressionValueType.DateTime, nil
		case ExpressionValueType.Boolean:
			fallthrough
		case ExpressionValueType.Int32:
			fallthrough
		case ExpressionValueType.Int64:
			fallthrough
		case ExpressionValueType.Decimal:
			fallthrough
		case ExpressionValueType.Double:
			fallthrough
		case ExpressionValueType.Guid:
			return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"DateTime\" and \"" + rightValueType.String() + "\"")
		default:
			return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
		}
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (et *ExpressionTree) deriveBooleanOperationValueType(operationType ExpressionOperatorTypeEnum, leftValueType ExpressionValueTypeEnum, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	if leftValueType == ExpressionValueType.Boolean && rightValueType == ExpressionValueType.Boolean {
		return ExpressionValueType.Boolean, nil
	}

	return ZeroExpressionValueType, errors.New("cannot perform \"" + operationType.String() + "\" operation on \"" + leftValueType.String() + "\" and \"" + rightValueType.String() + "\"")
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

	var err error

	switch sourceValue.ValueType() {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, sourceValue.booleanValue())
	case ExpressionValueType.Int32:
		var i32 int32
		if i32, err = sourceValue.Int32Value(); err != nil {
			return nil, err
		}
		return NewValueExpression(ExpressionValueType.Int32, abs32(i32)), nil
	case ExpressionValueType.Int64:
		var i64 int64
		if i64, err = sourceValue.Int64Value(); err != nil {
			return nil, err
		}
		return NewValueExpression(ExpressionValueType.Int64, abs64(i64)), nil
	case ExpressionValueType.Decimal:
		var f64 float64
		if f64, err = sourceValue.DecimalValue(); err != nil {
			return nil, err
		}
		return NewValueExpression(ExpressionValueType.Decimal, math.Abs(f64)), nil
	case ExpressionValueType.Double:
		var f64 float64
		if f64, err = sourceValue.DoubleValue(); err != nil {
			return nil, err
		}
		return NewValueExpression(ExpressionValueType.Double, math.Abs(f64)), nil
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
	return nil, nil
}

func (et *ExpressionTree) coalesce(arguments []*ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) convert(sourceValue *ValueExpression, targetType *ValueExpression) (*ValueExpression, error) {
	return nil, nil
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

func (et *ExpressionTree) iIf(testValue *ValueExpression, leftResultValue Expression, rightResultValue Expression) (*ValueExpression, error) {
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

func (et *ExpressionTree) maxOf(arguments []*ValueExpression) (*ValueExpression, error) {
	return nil, nil
}

func (et *ExpressionTree) minOf(arguments []*ValueExpression) (*ValueExpression, error) {
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
