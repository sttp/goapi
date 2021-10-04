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

	"github.com/sttp/goapi/sttp/data"
	"github.com/sttp/goapi/sttp/guid"
)

// ValueExpression represents a value expression.
type ValueExpression struct {
	value     interface{}
	valueType ExpressionValueTypeEnum
}

// NewValueExpression creates a new value expression.
func NewValueExpression(valueType ExpressionValueTypeEnum, value interface{}) *ValueExpression {
	if value != nil {
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
			panic("cannot create new value expression; unexpected expression value type: 0x" + strconv.FormatInt(int64(valueType), 16))
		}
	}

	return newValueExpression(valueType, value)
}

func newValueExpression(valueType ExpressionValueTypeEnum, value interface{}) *ValueExpression {
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
		return strconv.FormatBool(ve.booleanValue())
	case ExpressionValueType.Int32:
		return strconv.FormatInt(int64(ve.int32Value()), 10)
	case ExpressionValueType.Int64:
		return strconv.FormatInt(ve.int64Value(), 10)
	case ExpressionValueType.Decimal:
		return strconv.FormatFloat(ve.decimalValue(), 'f', 6, 64)
	case ExpressionValueType.Double:
		return strconv.FormatFloat(ve.doubleValue(), 'f', 6, 64)
	case ExpressionValueType.String:
		return ve.stringValue()
	case ExpressionValueType.Guid:
		return ve.guidValue().String()
	case ExpressionValueType.DateTime:
		return ve.dateTimeValue().Format(data.DateTimeFormat)
	default:
		return ""
	}
}

// IsNull gets a flag that determines if the ValueExpression value is null.
func (ve *ValueExpression) IsNull() bool {
	return ve.value == nil
}

// True is a value expression of type boolean with a true value.
var True *ValueExpression = newValueExpression(ExpressionValueType.Boolean, true)

// False is a value expression of type boolean with a false value.
var False *ValueExpression = newValueExpression(ExpressionValueType.Boolean, false)

// EmptyString is a value expression of type string with a value of an empty string.
var EmptyString *ValueExpression = newValueExpression(ExpressionValueType.String, "")

// NullValue gets the target expression value type with a value of nil.
func NullValue(targetValueType ExpressionValueTypeEnum) *ValueExpression {
	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(ExpressionValueType.Boolean, nil)
	case ExpressionValueType.Int32:
		return newValueExpression(ExpressionValueType.Int32, nil)
	case ExpressionValueType.Int64:
		return newValueExpression(ExpressionValueType.Int64, nil)
	case ExpressionValueType.Decimal:
		return newValueExpression(ExpressionValueType.Decimal, nil)
	case ExpressionValueType.Double:
		return newValueExpression(ExpressionValueType.Double, nil)
	case ExpressionValueType.String:
		return newValueExpression(ExpressionValueType.String, nil)
	case ExpressionValueType.Guid:
		return newValueExpression(ExpressionValueType.Guid, nil)
	case ExpressionValueType.DateTime:
		return newValueExpression(ExpressionValueType.DateTime, nil)
	default:
		return newValueExpression(ExpressionValueType.Undefined, nil)
	}
}

func (ve *ValueExpression) validateValueType(valueType ExpressionValueTypeEnum) error {
	if valueType != ve.valueType {
		return fmt.Errorf("cannot read expression value as \"%s\", type is \"%s\"", valueType.String(), ve.valueType.String())
	}

	return nil
}

// BooleanValue gets the ValueExpression value cast as a bool.
// An error will be returned if value type is not ExpressionValueType.Boolean.
func (ve *ValueExpression) BooleanValue() (bool, error) {
	err := ve.validateValueType(ExpressionValueType.Boolean)

	if err != nil {
		return false, err
	}

	return ve.booleanValue(), nil
}

func (ve *ValueExpression) booleanValue() bool {
	if ve.value == nil {
		return false
	}

	return ve.value.(bool)
}

// Int32Value gets the ValueExpression value cast as an int32.
// An error will be returned if value type is not ExpressionValueType.Int32.
func (ve *ValueExpression) Int32Value() (int32, error) {
	err := ve.validateValueType(ExpressionValueType.Int32)

	if err != nil {
		return 0, err
	}

	return ve.int32Value(), nil
}

func (ve *ValueExpression) int32Value() int32 {
	if ve.value == nil {
		return 0
	}

	return ve.value.(int32)
}

// Int64Value gets the ValueExpression value cast as an int64.
// An error will be returned if value type is not ExpressionValueType.Int64.
func (ve *ValueExpression) Int64Value() (int64, error) {
	err := ve.validateValueType(ExpressionValueType.Int64)

	if err != nil {
		return 0, err
	}

	return ve.int64Value(), nil
}

func (ve *ValueExpression) int64Value() int64 {
	if ve.value == nil {
		return 0
	}

	return ve.value.(int64)
}

// DecimalValue gets the ValueExpression value cast as a float64.
// An error will be returned if value type is not ExpressionValueType.Decimal.
func (ve *ValueExpression) DecimalValue() (float64, error) {
	err := ve.validateValueType(ExpressionValueType.Decimal)

	if err != nil {
		return 0.0, err
	}

	return ve.decimalValue(), nil
}

func (ve *ValueExpression) decimalValue() float64 {
	if ve.value == nil {
		return 0.0
	}

	return ve.value.(float64)
}

// DoubleValue gets the ValueExpression value cast as a float64.
// An error will be returned if value type is not ExpressionValueType.Double.
func (ve *ValueExpression) DoubleValue() (float64, error) {
	err := ve.validateValueType(ExpressionValueType.Double)

	if err != nil {
		return 0.0, err
	}

	return ve.doubleValue(), nil
}

func (ve *ValueExpression) doubleValue() float64 {
	if ve.value == nil {
		return 0.0
	}

	return ve.value.(float64)
}

// StringValue gets the ValueExpression value cast as a string.
// An error will be returned if value type is not ExpressionValueType.String.
func (ve *ValueExpression) StringValue() (string, error) {
	err := ve.validateValueType(ExpressionValueType.String)

	if err != nil {
		return "", err
	}

	return ve.stringValue(), nil
}

func (ve *ValueExpression) stringValue() string {
	if ve.value == nil {
		return ""
	}

	return ve.value.(string)
}

// GuidValue gets the ValueExpression value cast as a guid.Guid.
// An error will be returned if value type is not ExpressionValueType.Guid.
func (ve *ValueExpression) GuidValue() (guid.Guid, error) {
	err := ve.validateValueType(ExpressionValueType.Guid)

	if err != nil {
		return guid.Guid{}, err
	}

	return ve.guidValue(), nil
}

func (ve *ValueExpression) guidValue() guid.Guid {
	if ve.value == nil {
		return guid.Guid{}
	}

	return ve.value.(guid.Guid)
}

// DateTimeValue gets the ValueExpression value cast as a time.Time.
// An error will be returned if value type is not ExpressionValueType.DateTime.
func (ve *ValueExpression) DateTimeValue() (time.Time, error) {
	err := ve.validateValueType(ExpressionValueType.DateTime)

	if err != nil {
		return time.Time{}, err
	}

	return ve.dateTimeValue(), nil
}

func (ve *ValueExpression) dateTimeValue() time.Time {
	if ve.value == nil {
		return time.Time{}
	}

	return ve.value.(time.Time)
}

// Convert attempts to convert the ValueExpression to the specified targetValueType.
func (ve *ValueExpression) Convert(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	// If source value is Null, result is Null, regardless of target type
	if ve.IsNull() {
		return NullValue(targetValueType), nil
	}

	var targetValue interface{}

	switch ve.ValueType() {
	case ExpressionValueType.Boolean:
		result := ve.booleanValue()
		var value int

		if result {
			value = 1
		}

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue = result
		case ExpressionValueType.Int32:
			targetValue = value
		case ExpressionValueType.Int64:
			targetValue = int64(value)
		case ExpressionValueType.Decimal:
			targetValue = float64(value)
		case ExpressionValueType.Double:
			targetValue = float64(value)
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return nil, errors.New("cannot convert \"Boolean\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int32:
		value := ve.int32Value()

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue = value == 0
		case ExpressionValueType.Int32:
			targetValue = value
		case ExpressionValueType.Int64:
			targetValue = int64(value)
		case ExpressionValueType.Decimal:
			targetValue = float64(value)
		case ExpressionValueType.Double:
			targetValue = float64(value)
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return nil, errors.New("cannot convert \"Int32\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Int64:
		value := ve.int64Value()

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue = value == 0
		case ExpressionValueType.Int32:
			targetValue = int32(value)
		case ExpressionValueType.Int64:
			targetValue = value
		case ExpressionValueType.Decimal:
			targetValue = float64(value)
		case ExpressionValueType.Double:
			targetValue = float64(value)
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return nil, errors.New("cannot convert \"Int64\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Decimal:
		value := ve.decimalValue()

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue = value == 0.0
		case ExpressionValueType.Int32:
			targetValue = int32(value)
		case ExpressionValueType.Int64:
			targetValue = int64(value)
		case ExpressionValueType.Decimal:
			targetValue = value
		case ExpressionValueType.Double:
			targetValue = float64(value)
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return nil, errors.New("cannot convert \"Decimal\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Double:
		value := ve.doubleValue()

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue = value == 0.0
		case ExpressionValueType.Int32:
			targetValue = int32(value)
		case ExpressionValueType.Int64:
			targetValue = int64(value)
		case ExpressionValueType.Decimal:
			targetValue = float64(value)
		case ExpressionValueType.Double:
			targetValue = value
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.Guid:
			fallthrough
		case ExpressionValueType.DateTime:
			return nil, errors.New("cannot convert \"Double\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.String:
		value := ve.stringValue()

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue, _ = strconv.ParseBool(value)
		case ExpressionValueType.Int32:
			i, _ := strconv.ParseInt(value, 10, 32)
			targetValue = int32(i)
		case ExpressionValueType.Int64:
			targetValue, _ = strconv.ParseInt(value, 10, 64)
		case ExpressionValueType.Decimal:
			targetValue, _ = strconv.ParseFloat(value, 64)
		case ExpressionValueType.Double:
			targetValue, _ = strconv.ParseFloat(value, 64)
		case ExpressionValueType.String:
			targetValue = value
		case ExpressionValueType.Guid:
			targetValue, _ = guid.Parse(value)
		case ExpressionValueType.DateTime:
			targetValue, _ = time.Parse(data.DateTimeFormat, value)
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Guid:
		switch targetValueType {
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.Guid:
			targetValue, _ = ve.GuidValue()
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
			return nil, errors.New("cannot convert \"Guid\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.DateTime:
		result := ve.dateTimeValue()
		value := result.Unix()

		switch targetValueType {
		case ExpressionValueType.Boolean:
			targetValue = value == 0
		case ExpressionValueType.Int32:
			targetValue = int32(value)
		case ExpressionValueType.Int64:
			targetValue = int64(value)
		case ExpressionValueType.Decimal:
			targetValue = float64(value)
		case ExpressionValueType.Double:
			targetValue = float64(value)
		case ExpressionValueType.String:
			targetValue = ve.String()
		case ExpressionValueType.DateTime:
			targetValue = result
		case ExpressionValueType.Guid:
			return nil, errors.New("cannot convert \"DateTime\" to \"" + targetValueType.String() + "\"")
		default:
			return nil, errors.New("unexpected expression value type encountered")
		}
	case ExpressionValueType.Undefined:
		// Change Undefined values to Nullable of target type
		return NullValue(targetValueType), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}

	return NewValueExpression(targetValueType, targetValue), nil
}
