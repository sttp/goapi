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

	"github.com/araddon/dateparse"
	"github.com/shopspring/decimal"
	"github.com/sttp/goapi/sttp/data"
	"github.com/sttp/goapi/sttp/guid"
)

// ValueExpression represents a value expression.
type ValueExpression struct {
	value     interface{}
	valueType ExpressionValueTypeEnum
}

const debug bool = true

// NewValueExpression creates a new value expression.
//gocyclo:ignore
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
			if _, ok := value.(decimal.Decimal); !ok {
				panic("cannot create Decimal value expression; value is not decimal.Decimal")
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

	if debug {
		return &ValueExpression{
			value:     value,
			valueType: valueType,
		}
	}

	return newValueExpression(valueType, value)
}

func newValueExpression(valueType ExpressionValueTypeEnum, value interface{}) *ValueExpression {
	if debug {
		return NewValueExpression(valueType, value)
	}

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
		return ve.decimalValue().String()
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

func (ve *ValueExpression) integerValue(defaultValue int) int {
	switch ve.ValueType() {
	case ExpressionValueType.Boolean:
		return ve.booleanValueAsInt()
	case ExpressionValueType.Int32:
		return int(ve.int32Value())
	case ExpressionValueType.Int64:
		return int(ve.int64Value())
	default:
		return defaultValue
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

func (ve *ValueExpression) booleanValueAsInt() int {
	if ve.booleanValue() {
		return 1
	}

	return 0
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

// DecimalValue gets the ValueExpression value cast as a decimal.Decimal.
// An error will be returned if value type is not ExpressionValueType.Decimal.
func (ve *ValueExpression) DecimalValue() (decimal.Decimal, error) {
	err := ve.validateValueType(ExpressionValueType.Decimal)

	if err != nil {
		return decimal.Zero, err
	}

	return ve.decimalValue(), nil
}

func (ve *ValueExpression) decimalValue() decimal.Decimal {
	if ve.value == nil {
		return decimal.Zero
	}

	return ve.value.(decimal.Decimal)
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

	switch ve.ValueType() {
	case ExpressionValueType.Boolean:
		return ve.convertFromBoolean(targetValueType)
	case ExpressionValueType.Int32:
		return ve.convertFromInt32(targetValueType)
	case ExpressionValueType.Int64:
		return ve.convertFromInt64(targetValueType)
	case ExpressionValueType.Decimal:
		return ve.convertFromDecimal(targetValueType)
	case ExpressionValueType.Double:
		return ve.convertFromDouble(targetValueType)
	case ExpressionValueType.String:
		return ve.convertFromString(targetValueType)
	case ExpressionValueType.Guid:
		return ve.convertFromGuid(targetValueType)
	case ExpressionValueType.DateTime:
		return ve.convertFromDateTime(targetValueType)
	case ExpressionValueType.Undefined:
		// Change Undefined values to Nullable of target type
		return NullValue(targetValueType), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromBoolean(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	value := ve.booleanValueAsInt()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(targetValueType, value != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.Int64:
		return newValueExpression(targetValueType, int64(value)), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(targetValueType, decimal.NewFromInt(int64(value))), nil
	case ExpressionValueType.Double:
		return newValueExpression(targetValueType, float64(value)), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return nil, errors.New("cannot convert \"Boolean\" to \"" + targetValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromInt32(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	value := ve.int32Value()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(targetValueType, value != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.Int64:
		return newValueExpression(targetValueType, int64(value)), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(targetValueType, decimal.NewFromInt(int64(value))), nil
	case ExpressionValueType.Double:
		return newValueExpression(targetValueType, float64(value)), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return nil, errors.New("cannot convert \"Int32\" to \"" + targetValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromInt64(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	value := ve.int64Value()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(targetValueType, value != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(targetValueType, int32(value)), nil
	case ExpressionValueType.Int64:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(targetValueType, decimal.NewFromInt(value)), nil
	case ExpressionValueType.Double:
		return newValueExpression(targetValueType, float64(value)), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return nil, errors.New("cannot convert \"Int64\" to \"" + targetValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromDecimal(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	value := ve.decimalValue()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(targetValueType, !value.Equal(decimal.Zero)), nil
	case ExpressionValueType.Int32:
		return newValueExpression(targetValueType, int32(value.IntPart())), nil
	case ExpressionValueType.Int64:
		return newValueExpression(targetValueType, value.IntPart()), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.Double:
		f64, _ := value.Float64()
		return newValueExpression(targetValueType, f64), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return nil, errors.New("cannot convert \"Decimal\" to \"" + targetValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromDouble(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	value := ve.doubleValue()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(targetValueType, value != 0.0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(targetValueType, int32(value)), nil
	case ExpressionValueType.Int64:
		return newValueExpression(targetValueType, int64(value)), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(targetValueType, decimal.NewFromFloat(value)), nil
	case ExpressionValueType.Double:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return nil, errors.New("cannot convert \"Double\" to \"" + targetValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromString(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	value := ve.stringValue()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		targetValue, _ := strconv.ParseBool(value)
		return newValueExpression(targetValueType, targetValue), nil
	case ExpressionValueType.Int32:
		targetValue, _ := strconv.ParseInt(value, 10, 32)
		return newValueExpression(targetValueType, int32(targetValue)), nil
	case ExpressionValueType.Int64:
		targetValue, _ := strconv.ParseInt(value, 10, 64)
		return newValueExpression(targetValueType, targetValue), nil
	case ExpressionValueType.Decimal:
		targetValue, _ := decimal.NewFromString(value)
		return newValueExpression(targetValueType, targetValue), nil
	case ExpressionValueType.Double:
		targetValue, _ := strconv.ParseFloat(value, 64)
		return newValueExpression(targetValueType, targetValue), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.Guid:
		targetValue, _ := guid.Parse(value)
		return newValueExpression(targetValueType, targetValue), nil
	case ExpressionValueType.DateTime:
		targetValue, _ := dateparse.ParseAny(value)
		return newValueExpression(targetValueType, targetValue), nil
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}

func (ve *ValueExpression) convertFromGuid(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	switch targetValueType {
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.Guid:
		return newValueExpression(targetValueType, ve.guidValue()), nil
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
}

func (ve *ValueExpression) convertFromDateTime(targetValueType ExpressionValueTypeEnum) (*ValueExpression, error) {
	result := ve.dateTimeValue()
	value := result.Unix()

	switch targetValueType {
	case ExpressionValueType.Boolean:
		return newValueExpression(targetValueType, value != 0), nil
	case ExpressionValueType.Int32:
		return newValueExpression(targetValueType, int32(value)), nil
	case ExpressionValueType.Int64:
		return newValueExpression(targetValueType, value), nil
	case ExpressionValueType.Decimal:
		return newValueExpression(targetValueType, decimal.NewFromInt(value)), nil
	case ExpressionValueType.Double:
		return newValueExpression(targetValueType, float64(value)), nil
	case ExpressionValueType.String:
		return newValueExpression(targetValueType, ve.String()), nil
	case ExpressionValueType.DateTime:
		return newValueExpression(targetValueType, result), nil
	case ExpressionValueType.Guid:
		return nil, errors.New("cannot convert \"DateTime\" to \"" + targetValueType.String() + "\"")
	default:
		return nil, errors.New("unexpected expression value type encountered")
	}
}
