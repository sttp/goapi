//******************************************************************************************************
//  ValueExpression.go - Gbtc
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
//  10/01/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package filterexpressions

import (
	"fmt"
	"strconv"
	"time"

	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/ticks"
)

// ValueExpression represents a value expression.
type ValueExpression struct {
	value     interface{}
	valueType ExpressionValueTypeEnum
}

// NewValueExpression creates a new value expression.
func NewValueExpression(valueType ExpressionValueTypeEnum, value interface{}) *ValueExpression {
	return &ValueExpression{
		value:     value,
		valueType: valueType,
	}
}

// Type gets expression type of the ValueExpression.
func (*ValueExpression) Type() ExpressionTypeEnum {
	return ExpressionType.Value
}

// Value gets the value of a ValueExpression.
func (ve *ValueExpression) Value() interface{} {
	return ve.value
}

// ValueType gets data type of the value of a ValueExpression.
func (ve *ValueExpression) ValueType() ExpressionValueTypeEnum {
	return ve.valueType
}

// String gets the ValueExpression value as a string.
func (ve *ValueExpression) String() string {
	switch ve.valueType {
	case ExpressionValueType.Boolean:
		value, err := ve.BoolValue()

		if err != nil {
			return ""
		}

		return strconv.FormatBool(value)
	case ExpressionValueType.Int32:
		value, err := ve.Int32Value()

		if err != nil {
			return ""
		}

		return strconv.FormatInt(int64(value), 10)
	case ExpressionValueType.Int64:
		value, err := ve.Int64Value()

		if err != nil {
			return ""
		}

		return strconv.FormatInt(value, 10)
	case ExpressionValueType.Decimal:
		value, err := ve.DecimalValue()

		if err != nil {
			return ""
		}

		return strconv.FormatFloat(value, 'f', 6, 64)
	case ExpressionValueType.Double:
		value, err := ve.DoubleValue()

		if err != nil {
			return ""
		}

		return strconv.FormatFloat(value, 'f', 6, 64)
	case ExpressionValueType.Guid:
		value, err := ve.GuidValue()

		if err != nil {
			return ""
		}

		return value.String()
	case ExpressionValueType.DateTime:
		value, err := ve.DateTimeValue()

		if err != nil {
			return ""
		}

		return value.Format(ticks.TimeFormat)
	default:
		return ""
	}
}

// IsNull gets a flag that determines if the ValueExpression value is null.
func (ve *ValueExpression) IsNull() bool {
	return ve.value == nil
}

func (ve *ValueExpression) validateValueType(valueType ExpressionValueTypeEnum) error {
	if valueType != ve.valueType {
		return fmt.Errorf("cannot read expression value as \"%s\", type is \"%s\"", valueType.String(), ve.valueType.String())
	}

	return nil
}

// BoolValue gets the ValueExpression value cast as a bool.
// An error will be returned if value type is not ExpressionValueType.Boolean.
func (ve *ValueExpression) BoolValue() (bool, error) {
	err := ve.validateValueType(ExpressionValueType.Boolean)

	if err != nil {
		return false, err
	}

	if ve.value == nil {
		return false, nil
	}

	return ve.value.(bool), nil
}

// Int32Value gets the ValueExpression value cast as an int32.
// An error will be returned if value type is not ExpressionValueType.Int32.
func (ve *ValueExpression) Int32Value() (int32, error) {
	err := ve.validateValueType(ExpressionValueType.Int32)

	if err != nil {
		return 0, err
	}

	if ve.value == nil {
		return 0, nil
	}

	return ve.value.(int32), nil
}

// Int64Value gets the ValueExpression value cast as an int64.
// An error will be returned if value type is not ExpressionValueType.Int64.
func (ve *ValueExpression) Int64Value() (int64, error) {
	err := ve.validateValueType(ExpressionValueType.Int64)

	if err != nil {
		return 0, err
	}

	if ve.value == nil {
		return 0, nil
	}

	return ve.value.(int64), nil
}

// DecimalValue gets the ValueExpression value cast as a float64.
// An error will be returned if value type is not ExpressionValueType.Decimal.
func (ve *ValueExpression) DecimalValue() (float64, error) {
	err := ve.validateValueType(ExpressionValueType.Decimal)

	if err != nil {
		return 0.0, err
	}

	if ve.value == nil {
		return 0.0, nil
	}

	return ve.value.(float64), nil
}

// DoubleValue gets the ValueExpression value cast as a float64.
// An error will be returned if value type is not ExpressionValueType.Double.
func (ve *ValueExpression) DoubleValue() (float64, error) {
	err := ve.validateValueType(ExpressionValueType.Double)

	if err != nil {
		return 0.0, err
	}

	if ve.value == nil {
		return 0.0, nil
	}

	return ve.value.(float64), nil
}

// StringValue gets the ValueExpression value cast as a string.
// An error will be returned if value type is not ExpressionValueType.String.
func (ve *ValueExpression) StringValue() (string, error) {
	err := ve.validateValueType(ExpressionValueType.String)

	if err != nil {
		return "", err
	}

	if ve.value == nil {
		return "", nil
	}

	return ve.value.(string), nil
}

// GuidValue gets the ValueExpression value cast as a guid.Guid.
// An error will be returned if value type is not ExpressionValueType.Guid.
func (ve *ValueExpression) GuidValue() (guid.Guid, error) {
	err := ve.validateValueType(ExpressionValueType.Guid)

	if err != nil {
		return guid.Guid{}, err
	}

	if ve.value == nil {
		return guid.Guid{}, nil
	}

	return ve.value.(guid.Guid), nil
}

// DateTimeValue gets the ValueExpression value cast as a time.Time.
// An error will be returned if value type is not ExpressionValueType.DateTime.
func (ve *ValueExpression) DateTimeValue() (time.Time, error) {
	err := ve.validateValueType(ExpressionValueType.DateTime)

	if err != nil {
		return time.Time{}, err
	}

	if ve.value == nil {
		return time.Time{}, nil
	}

	return ve.value.(time.Time), nil
}
