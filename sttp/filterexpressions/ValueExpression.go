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
	"errors"
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
	if value != nil {
		// TODO: debug validation code - consider removing for production:
		switch valueType {
		case ExpressionValueType.Boolean:
			if _, ok := value.(bool); !ok {
				panic("cannot create Boolean value expression; value is not bool")
			}
		case ExpressionValueType.Int32:
			if _, ok := value.(int32); !ok {
				panic("cannot create Int32 value expression; value is not int32")
			}
		case ExpressionValueType.Int64:
			if _, ok := value.(int64); !ok {
				panic("cannot create Int64 value expression; value is not int64")
			}
		case ExpressionValueType.Decimal:
			if _, ok := value.(float64); !ok {
				panic("cannot create Decimal value expression; value is not float64")
			}
		case ExpressionValueType.Double:
			if _, ok := value.(float64); !ok {
				panic("cannot create Double value expression; value is not float64")
			}
		case ExpressionValueType.String:
			if _, ok := value.(string); !ok {
				panic("cannot create String value expression; value is not string")
			}
		case ExpressionValueType.Guid:
			if _, ok := value.(guid.Guid); !ok {
				panic("cannot create Guid value expression; value is not guid.Guid")
			}
		case ExpressionValueType.DateTime:
			if _, ok := value.(time.Time); !ok {
				panic("cannot create DateTime value expression; value is not time.Time")
			}
		default:
			panic("cannot create Boolean value expression; unexpected expression value type: 0x" + strconv.FormatInt(int64(valueType), 16))
		}
	}

	return &ValueExpression{
		value:     value,
		valueType: valueType,
	}
}

// newValueExpression takes parameter array to allow for single-lined, less cluttered, more readable code
func newValueExpression(valueType ExpressionValueTypeEnum, params []interface{}) (*ValueExpression, error) {
	if len(params) != 2 {
		return nil, errors.New("unexpected newValueExpression parameter count")
	}

	if params[1] != nil {
		err, ok := params[1].(error)

		if !ok {
			return nil, errors.New("second newValueExpression was not an error")
		}

		return nil, err
	}

	return NewValueExpression(valueType, params[0]), nil
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
		value, err := ve.BooleanValue()

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
	case ExpressionValueType.String:
		value, err := ve.StringValue()

		if err != nil {
			return ""
		}

		return value
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

func paramArray(value interface{}, err error) []interface{} {
	return []interface{}{value, err}
}

// BooleanValue gets the ValueExpression value cast as a bool.
// An error will be returned if value type is not ExpressionValueType.Boolean.
func (ve *ValueExpression) BooleanValue() (bool, error) {
	err := ve.validateValueType(ExpressionValueType.Boolean)

	if err != nil {
		return false, err
	}

	if ve.value == nil {
		return false, nil
	}

	return ve.value.(bool), nil
}

func (ve *ValueExpression) booleanValue() []interface{} {
	return paramArray(ve.BooleanValue())
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

func (ve *ValueExpression) int32Value() []interface{} {
	return paramArray(ve.Int32Value())
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

func (ve *ValueExpression) int64Value() []interface{} {
	return paramArray(ve.Int64Value())
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

func (ve *ValueExpression) decimalValue() []interface{} {
	return paramArray(ve.DecimalValue())
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

func (ve *ValueExpression) doubleValue() []interface{} {
	return paramArray(ve.DoubleValue())
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

func (ve *ValueExpression) stringValue() []interface{} {
	return paramArray(ve.StringValue())
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

func (ve *ValueExpression) guidValue() []interface{} {
	return paramArray(ve.GuidValue())
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

func (ve *ValueExpression) dateTimeValue() []interface{} {
	return paramArray(ve.DateTimeValue())
}

// True is a value expression of type boolean with a true value.
var True *ValueExpression = NewValueExpression(ExpressionValueType.Boolean, true)

// False is a value expression of type boolean with a false value.
var False *ValueExpression = NewValueExpression(ExpressionValueType.Boolean, false)

// EmptyString is a value expression of type string with a value of an empty string.
var EmptyString *ValueExpression = NewValueExpression(ExpressionValueType.String, "")

// NullValue gets the target expression value type with a value of nil.
func NullValue(targetValueType ExpressionValueTypeEnum) *ValueExpression {
	switch targetValueType {
	case ExpressionValueType.Boolean:
		return NewValueExpression(ExpressionValueType.Boolean, nil)
	case ExpressionValueType.Int32:
		return NewValueExpression(ExpressionValueType.Int32, nil)
	case ExpressionValueType.Int64:
		return NewValueExpression(ExpressionValueType.Int64, nil)
	case ExpressionValueType.Decimal:
		return NewValueExpression(ExpressionValueType.Decimal, nil)
	case ExpressionValueType.Double:
		return NewValueExpression(ExpressionValueType.Double, nil)
	case ExpressionValueType.String:
		return NewValueExpression(ExpressionValueType.String, nil)
	case ExpressionValueType.Guid:
		return NewValueExpression(ExpressionValueType.Guid, nil)
	case ExpressionValueType.DateTime:
		return NewValueExpression(ExpressionValueType.DateTime, nil)
	default:
		return NewValueExpression(ExpressionValueType.Undefined, nil)
	}
}
