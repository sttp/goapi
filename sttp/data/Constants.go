//******************************************************************************************************
//  Constants.go - Gbtc
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
package data

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ExpressionTypeEnum defines the type of the ExpressionType enumeration.
type ExpressionTypeEnum int

// ExpressionType is an enumeration of possible expression types.
var ExpressionType = struct {
	// Value defines a value expression type.
	Value ExpressionTypeEnum
	// Unary defines an unary expression type.
	Unary ExpressionTypeEnum
	// Column defines a column expression type.
	Column ExpressionTypeEnum
	// InList defines an in-list expression type.
	InList ExpressionTypeEnum
	// Function defines a function expression type.
	Function ExpressionTypeEnum
	// Operator defines an operator expression type.
	Operator ExpressionTypeEnum
}{
	Value:    0,
	Unary:    1,
	Column:   2,
	InList:   3,
	Function: 4,
	Operator: 5,
}

// String gets the ExpressionType enumeration value as a string.
func (ete ExpressionTypeEnum) String() string {
	switch ete {
	case ExpressionType.Value:
		return "Value"
	case ExpressionType.Unary:
		return "Unary"
	case ExpressionType.Column:
		return "Column"
	case ExpressionType.InList:
		return "InList"
	case ExpressionType.Function:
		return "Function"
	case ExpressionType.Operator:
		return "Operator"
	default:
		return "0x" + strconv.FormatInt(int64(ete), 16)
	}
}

// ExpressionValueTypeEnum defines the type of the ExpressionValueType enumeration.
type ExpressionValueTypeEnum int

// ExpressionValueType is an enumeration of possible expression value data types.
// These expression value data types are reduced to a reasonable set of possible types
// that can be represented in a filter expression. All data table column values will be
// mapped to these types.
var ExpressionValueType = struct {
	// Boolean defines a bool value type for an expression.
	Boolean ExpressionValueTypeEnum
	// Int32 defines an int32 value type for an expression.
	Int32 ExpressionValueTypeEnum
	// Int64 defines an int64 value type for an expression.
	Int64 ExpressionValueTypeEnum
	// Decimal defines a decimal.Decimal value type for an expression.
	Decimal ExpressionValueTypeEnum
	// Double defines a float64 value type for an expression.
	Double ExpressionValueTypeEnum
	// String defines a string value type for an expression.
	String ExpressionValueTypeEnum
	// Guid defines a guid.Guid value type for an expression.
	Guid ExpressionValueTypeEnum
	// DateTime defines a time.Time value type for an expression.
	DateTime ExpressionValueTypeEnum
	// Undefined defines a nil value type for an expression.
	Undefined ExpressionValueTypeEnum // make sure value is always last in enum
}{
	Boolean:   0,
	Int32:     1,
	Int64:     2,
	Decimal:   3,
	Double:    4,
	String:    5,
	Guid:      6,
	DateTime:  7,
	Undefined: 8,
}

// ZeroExpressionValueType defines the zero value for the ExpressionValueType enumeration
var ZeroExpressionValueType ExpressionValueTypeEnum = 0

// ExpressionValueTypeLen gets the number of elements in the ExpressionValueType enumeration.
func ExpressionValueTypeLen() int {
	return int(ExpressionValueType.Undefined) + 1
}

// String gets the ExpressionValueType enumeration value as a string.
func (evte ExpressionValueTypeEnum) String() string {
	switch evte {
	case ExpressionValueType.Boolean:
		return "Boolean"
	case ExpressionValueType.Int32:
		return "Int32"
	case ExpressionValueType.Int64:
		return "Int64"
	case ExpressionValueType.Decimal:
		return "Decimal"
	case ExpressionValueType.Double:
		return "Double"
	case ExpressionValueType.String:
		return "String"
	case ExpressionValueType.Guid:
		return "Guid"
	case ExpressionValueType.DateTime:
		return "DateTime"
	case ExpressionValueType.Undefined:
		return "Undefined"
	default:
		return "0x" + strconv.FormatInt(int64(evte), 16)
	}
}

// IsIntegerType gets a flag that determines if the ExpressionValueType enumeration value represents an integer type.
func (evte ExpressionValueTypeEnum) IsIntegerType() bool {
	switch evte {
	case ExpressionValueType.Boolean:
		return true
	case ExpressionValueType.Int32:
		return true
	case ExpressionValueType.Int64:
		return true
	default:
		return false
	}
}

// IsNumericType gets a flag that determines if the ExpressionValueType enumeration value represents a numeric type.
func (evte ExpressionValueTypeEnum) IsNumericType() bool {
	switch evte {
	case ExpressionValueType.Boolean:
		return true
	case ExpressionValueType.Int32:
		return true
	case ExpressionValueType.Int64:
		return true
	case ExpressionValueType.Decimal:
		return true
	case ExpressionValueType.Double:
		return true
	default:
		return false
	}
}

// ExpressionUnaryTypeEnum defines the type of the ExpressionUnaryType enumeration.
type ExpressionUnaryTypeEnum int

// ExpressionUnaryType is an enumeration of the possible expression unary types.
var ExpressionUnaryType = struct {
	// Plus defines the "+" unary operator.
	Plus ExpressionUnaryTypeEnum
	// Minus defines the "-" unary operator.
	Minus ExpressionUnaryTypeEnum
	// Not defines the "~" or "!" unary operator.
	Not ExpressionUnaryTypeEnum
}{
	Plus:  0,
	Minus: 1,
	Not:   2,
}

// String gets the ExpressionUnaryType enumeration value as a string.
func (eute ExpressionUnaryTypeEnum) String() string {
	switch eute {
	case ExpressionUnaryType.Plus:
		return "+"
	case ExpressionUnaryType.Minus:
		return "-"
	case ExpressionUnaryType.Not:
		return "~"
	default:
		return "0x" + strconv.FormatInt(int64(eute), 16)
	}
}

// ExpressionFunctionTypeEnum defines the type of the ExpressionFunctionType enumeration.
type ExpressionFunctionTypeEnum int

// ExpressionFunctionType is an enumeration of possible expression function types.
var ExpressionFunctionType = struct {
	// Abs defines a function type that returns the absolute value of the specified numeric expression.
	Abs ExpressionFunctionTypeEnum
	// Ceiling defines a function type that returns the smallest integer that is greater than, or equal to, the specified numeric expression.
	Ceiling ExpressionFunctionTypeEnum
	// Coalesce defines a function type that returns the first non-null value in expression list.
	Coalesce ExpressionFunctionTypeEnum
	// Convert defines a function type that converts expression to the specified type.
	Convert ExpressionFunctionTypeEnum
	// Contains defines a function type that returns flag that determines if source string contains test string.
	Contains ExpressionFunctionTypeEnum
	// DateAdd defines a function type that adds value at specified interval to source date and then returns the date.
	DateAdd ExpressionFunctionTypeEnum
	// DateDiff defines a function type that returns the difference between left and right value at specified interval.
	DateDiff ExpressionFunctionTypeEnum
	// DatePart defines a function type that returns specified interval of source date.
	DatePart ExpressionFunctionTypeEnum
	// EndsWith defines a function type that returns flag that determines if source string ends with test string.
	EndsWith ExpressionFunctionTypeEnum
	// Floor defines a function type that returns the largest integer value that is smaller than, or equal to, the specified numeric expression.
	Floor ExpressionFunctionTypeEnum
	// IIf defines a function type that returns leftValue if result of expression is true, else returns rightValue.
	IIf ExpressionFunctionTypeEnum
	// IndexOf defines a function type that returns zero-based index of first occurrence of test in source, or -1 if not found.
	IndexOf ExpressionFunctionTypeEnum
	// IsDate defines a function type that returns flag that determines if expression is a DateTime value or can be parsed as one.
	IsDate ExpressionFunctionTypeEnum
	// IsInteger defines a function type that returns flag that determines if expression is an integer value or can be parsed as one.
	IsInteger ExpressionFunctionTypeEnum
	// IsGuid defines a function type that returns flag that determines if expression is a Guid value or can be parsed as one.
	IsGuid ExpressionFunctionTypeEnum
	// IsNull defines a function type that returns the specified defaultValue if expression is null, otherwise returns the expression.
	IsNull ExpressionFunctionTypeEnum
	// IsNumeric defines a function type that returns flag that determines if expression is a numeric value or can be parsed as one.
	IsNumeric ExpressionFunctionTypeEnum
	// LastIndexOf defines a function type that returns zero-based index of last occurrence of test in source, or -1 if not found.
	LastIndexOf ExpressionFunctionTypeEnum
	// Len defines a function type that returns length of expression interpreted as a string.
	Len ExpressionFunctionTypeEnum
	// Lower defines a function type that returns lower-case representation of expression interpreted as a string.
	Lower ExpressionFunctionTypeEnum
	// MaxOf defines a function type that returns value in expression list with maximum value.
	MaxOf ExpressionFunctionTypeEnum
	// MinOf defines a function type that returns value in expression list with minimum value.
	MinOf ExpressionFunctionTypeEnum
	// Now defines a function type that returns a DateTime value representing the current local system time.
	Now ExpressionFunctionTypeEnum
	// NthIndexOf defines a function type that returns zero-based index of the Nth, represented by index value, occurrence of test in source, or -1 if not found.
	NthIndexOf ExpressionFunctionTypeEnum
	// Power defines a function type that returns the value of specified numeric expression raised to the power of specified numeric exponent.
	Power ExpressionFunctionTypeEnum
	// RegExMatch defines a function type that returns flag that determines if test, interpreted as a string, is a match for specified regex string-based regular expression.
	RegExMatch ExpressionFunctionTypeEnum
	// RegExVal defines a function type that returns value from test, interpreted as a string, that is matched by specified regex string-based regular expression.
	RegExVal ExpressionFunctionTypeEnum
	// Replace defines a function type that returns a string where all instances of test found in source are replaced with replace value - all parameters interpreted as strings.
	Replace ExpressionFunctionTypeEnum
	// Reverse defines a function type that returns string where all characters in expression interpreted as a string are reversed.
	Reverse ExpressionFunctionTypeEnum
	// Round defines a function type that returns the nearest integer value to the specified numeric expression
	Round ExpressionFunctionTypeEnum
	// Split defines a function type that returns zero-based Nth, represented by index, value in source split by delimiter, or null if out of range.
	Split ExpressionFunctionTypeEnum
	// Sqrt defines a function type that returns the square root of the specified numeric expression
	Sqrt ExpressionFunctionTypeEnum
	// StartsWith defines a function type that returns flag that determines if source string starts with test string.
	StartsWith ExpressionFunctionTypeEnum
	// StrCount defines a function type that returns count of occurrences of test in source.
	StrCount ExpressionFunctionTypeEnum
	// StrCmp defines a function type that returns -1 if left is less-than right, 1 if left is greater-than right, or 0 if left equals right.
	StrCmp ExpressionFunctionTypeEnum
	// SubStr defines a function type that returns portion of source interpreted as a string starting at index.
	SubStr ExpressionFunctionTypeEnum
	// Trim defines a function type that removes white-space from the beginning and end of expression interpreted as a string.
	Trim ExpressionFunctionTypeEnum
	// TrimLeft defines a function type that removes white-space from the beginning of expression interpreted as a string.
	TrimLeft ExpressionFunctionTypeEnum
	// TrimRight defines a function type that removes white-space from the end of expression interpreted as a string.
	TrimRight ExpressionFunctionTypeEnum
	// Upper defines a function type that returns upper-case representation of expression interpreted as a string.
	Upper ExpressionFunctionTypeEnum
	// UtcNow defines a function type that returns a DateTime value representing the current UTC system time.
	UtcNow ExpressionFunctionTypeEnum
}{
	Abs:         0,
	Ceiling:     1,
	Coalesce:    2,
	Convert:     3,
	Contains:    4,
	DateAdd:     5,
	DateDiff:    6,
	DatePart:    7,
	EndsWith:    8,
	Floor:       9,
	IIf:         10,
	IndexOf:     11,
	IsDate:      12,
	IsInteger:   13,
	IsGuid:      14,
	IsNull:      15,
	IsNumeric:   16,
	LastIndexOf: 17,
	Len:         18,
	Lower:       19,
	MaxOf:       20,
	MinOf:       21,
	Now:         22,
	NthIndexOf:  23,
	Power:       24,
	RegExMatch:  25,
	RegExVal:    26,
	Replace:     27,
	Reverse:     28,
	Round:       29,
	Split:       30,
	Sqrt:        31,
	StartsWith:  32,
	StrCount:    33,
	StrCmp:      34,
	SubStr:      35,
	Trim:        36,
	TrimLeft:    37,
	TrimRight:   38,
	Upper:       39,
	UtcNow:      40,
}

// String gets the ExpressionFunctionType enumeration value as a string.
//gocyclo:ignore
func (efte ExpressionFunctionTypeEnum) String() string {
	switch efte {
	case ExpressionFunctionType.Abs:
		return "Abs"
	case ExpressionFunctionType.Ceiling:
		return "Ceiling"
	case ExpressionFunctionType.Coalesce:
		return "Coalesce"
	case ExpressionFunctionType.Convert:
		return "Convert"
	case ExpressionFunctionType.Contains:
		return "Contains"
	case ExpressionFunctionType.DateAdd:
		return "DateAdd"
	case ExpressionFunctionType.DateDiff:
		return "DateDiff"
	case ExpressionFunctionType.DatePart:
		return "DatePart"
	case ExpressionFunctionType.EndsWith:
		return "EndsWith"
	case ExpressionFunctionType.Floor:
		return "Floor"
	case ExpressionFunctionType.IIf:
		return "IIf"
	case ExpressionFunctionType.IndexOf:
		return "IndexOf"
	case ExpressionFunctionType.IsDate:
		return "IsDate"
	case ExpressionFunctionType.IsInteger:
		return "IsInteger"
	case ExpressionFunctionType.IsGuid:
		return "IsGuid"
	case ExpressionFunctionType.IsNull:
		return "IsNull"
	case ExpressionFunctionType.IsNumeric:
		return "IsNumeric"
	case ExpressionFunctionType.LastIndexOf:
		return "LastIndexOf"
	case ExpressionFunctionType.Len:
		return "Len"
	case ExpressionFunctionType.Lower:
		return "Lower"
	case ExpressionFunctionType.MaxOf:
		return "MaxOf"
	case ExpressionFunctionType.MinOf:
		return "MinOf"
	case ExpressionFunctionType.Now:
		return "Now"
	case ExpressionFunctionType.NthIndexOf:
		return "NthIndexOf"
	case ExpressionFunctionType.Power:
		return "Power"
	case ExpressionFunctionType.RegExMatch:
		return "RegExMatch"
	case ExpressionFunctionType.RegExVal:
		return "RegExVal"
	case ExpressionFunctionType.Replace:
		return "Replace"
	case ExpressionFunctionType.Reverse:
		return "Reverse"
	case ExpressionFunctionType.Round:
		return "Round"
	case ExpressionFunctionType.Split:
		return "Split"
	case ExpressionFunctionType.Sqrt:
		return "Sqrt"
	case ExpressionFunctionType.StartsWith:
		return "StartsWith"
	case ExpressionFunctionType.StrCount:
		return "StrCount"
	case ExpressionFunctionType.StrCmp:
		return "StrCmp"
	case ExpressionFunctionType.SubStr:
		return "SubStr"
	case ExpressionFunctionType.Trim:
		return "Trim"
	case ExpressionFunctionType.TrimLeft:
		return "TrimLeft"
	case ExpressionFunctionType.TrimRight:
		return "TrimRight"
	case ExpressionFunctionType.Upper:
		return "Upper"
	case ExpressionFunctionType.UtcNow:
		return "UtcNow"
	default:
		return "0x" + strconv.FormatInt(int64(efte), 16)
	}
}

// ExpressionOperatorTypeEnum defines the type of the ExpressionOperatorType enumeration.
type ExpressionOperatorTypeEnum int

// ExpressionOperatorType is an enumeration of possible expression operator types.
var ExpressionOperatorType = struct {
	// Multiply defines a "*" operator type.
	Multiply ExpressionOperatorTypeEnum
	// Divide defines a "/" operator type.
	Divide ExpressionOperatorTypeEnum
	// Modulus defines a "%" operator type.
	Modulus ExpressionOperatorTypeEnum
	// Add defines an "+" operator type.
	Add ExpressionOperatorTypeEnum
	// Subtract defines a "-" operator type.
	Subtract ExpressionOperatorTypeEnum
	// BitShiftLeft defines a "<<" operator type.
	BitShiftLeft ExpressionOperatorTypeEnum
	// BitShiftRight defines a ">>" operator type.
	BitShiftRight ExpressionOperatorTypeEnum
	// BitwiseAnd defines a "&" operator type.
	BitwiseAnd ExpressionOperatorTypeEnum
	// BitwiseOr defines a "|" operator type.
	BitwiseOr ExpressionOperatorTypeEnum
	// BitwiseXor defines a "^" operator type.
	BitwiseXor ExpressionOperatorTypeEnum
	// LessThan defines a "<" operator type.
	LessThan ExpressionOperatorTypeEnum
	// LessThanOrEqual defines a "<=" operator type.
	LessThanOrEqual ExpressionOperatorTypeEnum
	// GreaterThan defines a ">" operator type.
	GreaterThan ExpressionOperatorTypeEnum
	// GreaterThanOrEqual defines a ">=" operator type.
	GreaterThanOrEqual ExpressionOperatorTypeEnum
	// Equal defines an "=" or "==" operator type.
	Equal ExpressionOperatorTypeEnum
	// EqualExactMatch defines an "===" operator type.
	EqualExactMatch ExpressionOperatorTypeEnum
	// NotEqual defines a "<>" or "!=" operator type.
	NotEqual ExpressionOperatorTypeEnum
	// NotEqualExactMatch defines a "!==" operator type.
	NotEqualExactMatch ExpressionOperatorTypeEnum
	// IsNull defines a "IS NULL" operator type.
	IsNull ExpressionOperatorTypeEnum
	// IsNotNull defines a "IS NOT NULL" operator type.
	IsNotNull ExpressionOperatorTypeEnum
	// Like defines a "LIKE" operator type.
	Like ExpressionOperatorTypeEnum
	// LikeExactMatch defines a "LIKE BINARY" or "LIKE ===" operator type.
	LikeExactMatch ExpressionOperatorTypeEnum
	// NotLike defines a "NOT LIKE" operator type.
	NotLike ExpressionOperatorTypeEnum
	// NotLikeExactMatch defines a "NOT LIKE BINARY" or "NOT LIKE ===" operator type.
	NotLikeExactMatch ExpressionOperatorTypeEnum
	// And defines an "AND" or "&&" operator type.
	And ExpressionOperatorTypeEnum
	// Or defines an "OR" or "||" operator type.
	Or ExpressionOperatorTypeEnum
}{
	Multiply:           0,
	Divide:             1,
	Modulus:            2,
	Add:                3,
	Subtract:           4,
	BitShiftLeft:       5,
	BitShiftRight:      6,
	BitwiseAnd:         7,
	BitwiseOr:          8,
	BitwiseXor:         9,
	LessThan:           10,
	LessThanOrEqual:    11,
	GreaterThan:        12,
	GreaterThanOrEqual: 13,
	Equal:              14,
	EqualExactMatch:    15,
	NotEqual:           16,
	NotEqualExactMatch: 17,
	IsNull:             18,
	IsNotNull:          19,
	Like:               20,
	LikeExactMatch:     21,
	NotLike:            22,
	NotLikeExactMatch:  23,
	And:                24,
	Or:                 25,
}

// String gets the ExpressionOperatorType enumeration value as a string.
//gocyclo:ignore
func (eote ExpressionOperatorTypeEnum) String() string {
	switch eote {
	case ExpressionOperatorType.Multiply:
		return "*"
	case ExpressionOperatorType.Divide:
		return "/"
	case ExpressionOperatorType.Modulus:
		return "%"
	case ExpressionOperatorType.Add:
		return "+"
	case ExpressionOperatorType.Subtract:
		return "-"
	case ExpressionOperatorType.BitShiftLeft:
		return "<<"
	case ExpressionOperatorType.BitShiftRight:
		return ">>"
	case ExpressionOperatorType.BitwiseAnd:
		return "&"
	case ExpressionOperatorType.BitwiseOr:
		return "|"
	case ExpressionOperatorType.BitwiseXor:
		return "^"
	case ExpressionOperatorType.LessThan:
		return "<"
	case ExpressionOperatorType.LessThanOrEqual:
		return "<="
	case ExpressionOperatorType.GreaterThan:
		return ">"
	case ExpressionOperatorType.GreaterThanOrEqual:
		return ">="
	case ExpressionOperatorType.Equal:
		return "="
	case ExpressionOperatorType.EqualExactMatch:
		return "==="
	case ExpressionOperatorType.NotEqual:
		return "<>"
	case ExpressionOperatorType.NotEqualExactMatch:
		return "!=="
	case ExpressionOperatorType.IsNull:
		return "IS NULL"
	case ExpressionOperatorType.IsNotNull:
		return "IS NOT NULL"
	case ExpressionOperatorType.Like:
		return "LIKE"
	case ExpressionOperatorType.LikeExactMatch:
		return "LIKE BINARY"
	case ExpressionOperatorType.NotLike:
		return "NOT LIKE"
	case ExpressionOperatorType.NotLikeExactMatch:
		return "NOT LIKE BINARY"
	case ExpressionOperatorType.And:
		return "AND"
	case ExpressionOperatorType.Or:
		return "OR"
	default:
		return "0x" + strconv.FormatInt(int64(eote), 16)
	}
}

// TimeIntervalEnum defines the type of the TimeInterval enumeration.
type TimeIntervalEnum int

// TimeInterval is an enumeration of possible DateTime intervals.
var TimeInterval = struct {
	// Year represents the year part of a DateTime.
	Year TimeIntervalEnum
	// Month represents the month part (1-12) of a DateTime.
	Month TimeIntervalEnum
	// DayOfYear represents the day of the year (1-366) of a DateTime.
	DayOfYear TimeIntervalEnum
	// Day represents the of of the month (1-31) of a DateTime.
	Day TimeIntervalEnum
	// Week represents the week of the year (1-53) of a DateTime.
	Week TimeIntervalEnum
	// WeekDay represents the day of the week (1-7) of a DateTime.
	WeekDay TimeIntervalEnum
	// Hour represents the hour of the day (0-23) of a DateTime.
	Hour TimeIntervalEnum
	// Minute represents the minute of the hour (0-59) of a DateTime.
	Minute TimeIntervalEnum
	// Second represents the second of the minute (0-59) of a DateTime.
	Second TimeIntervalEnum
	// Millisecond represents the millisecond of the second (0-999) of a DateTime.
	Millisecond TimeIntervalEnum
}{
	Year:        0,
	Month:       1,
	DayOfYear:   2,
	Day:         3,
	Week:        4,
	WeekDay:     5,
	Hour:        6,
	Minute:      7,
	Second:      8,
	Millisecond: 9,
}

// ParseTimeInterval gets the TimeInterval parsed from the specified name. Case insensitive.
func ParseTimeInterval(name string) (TimeIntervalEnum, error) {
	name = strings.ToUpper(strings.TrimSpace(name))

	switch name {
	case "YEAR":
		return TimeInterval.Year, nil
	case "MONTH":
		return TimeInterval.Month, nil
	case "DAYOFYEAR":
		return TimeInterval.DayOfYear, nil
	case "DAY":
		return TimeInterval.Day, nil
	case "WEEK":
		return TimeInterval.Week, nil
	case "WEEKDAY":
		return TimeInterval.WeekDay, nil
	case "HOUR":
		return TimeInterval.Hour, nil
	case "MINUTE":
		return TimeInterval.Minute, nil
	case "SECOND":
		return TimeInterval.Second, nil
	case "MILLISECOND":
		return TimeInterval.Millisecond, nil
	default:
		return TimeInterval.Year, fmt.Errorf("specified time interval \"%s\" is unrecognized", name)
	}
}

// Operation Value Type Selectors

//gocyclo:ignore
func (eote ExpressionOperatorTypeEnum) deriveOperationValueType(leftValueType, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch eote {
	case ExpressionOperatorType.Multiply:
		fallthrough
	case ExpressionOperatorType.Divide:
		fallthrough
	case ExpressionOperatorType.Add:
		fallthrough
	case ExpressionOperatorType.Subtract:
		return eote.deriveArithmeticOperationValueType(leftValueType, rightValueType)
	case ExpressionOperatorType.Modulus:
		fallthrough
	case ExpressionOperatorType.BitwiseAnd:
		fallthrough
	case ExpressionOperatorType.BitwiseOr:
		fallthrough
	case ExpressionOperatorType.BitwiseXor:
		return eote.deriveIntegerOperationValueType(leftValueType, rightValueType)
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
		return eote.deriveComparisonOperationValueType(leftValueType, rightValueType)
	case ExpressionOperatorType.And:
		fallthrough
	case ExpressionOperatorType.Or:
		return eote.deriveBooleanOperationValueType(leftValueType, rightValueType)
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

func (eote ExpressionOperatorTypeEnum) deriveArithmeticOperationValueType(leftValueType, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch leftValueType {
	case ExpressionValueType.Boolean:
		return eote.deriveArithmeticOperationValueTypeFromBoolean(rightValueType)
	case ExpressionValueType.Int32:
		return eote.deriveArithmeticOperationValueTypeFromInt32(rightValueType)
	case ExpressionValueType.Int64:
		return eote.deriveArithmeticOperationValueTypeFromInt64(rightValueType)
	case ExpressionValueType.Decimal:
		return eote.deriveArithmeticOperationValueTypeFromDecimal(rightValueType)
	case ExpressionValueType.Double:
		return eote.deriveArithmeticOperationValueTypeFromDouble(rightValueType)
	case ExpressionValueType.String:
		if eote == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"" + leftValueType.String() + "\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveArithmeticOperationValueTypeFromBoolean(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		if eote == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Boolean\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveArithmeticOperationValueTypeFromInt32(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		if eote == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Int32\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveArithmeticOperationValueTypeFromInt64(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		if eote == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Int64\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveArithmeticOperationValueTypeFromDecimal(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		if eote == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Decimal\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveArithmeticOperationValueTypeFromDouble(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		if eote == ExpressionOperatorType.Add {
			return ExpressionValueType.String, nil
		}
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Double\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveIntegerOperationValueType(leftValueType, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch leftValueType {
	case ExpressionValueType.Boolean:
		return eote.deriveIntegerOperationValueTypeFromBoolean(rightValueType)
	case ExpressionValueType.Int32:
		return eote.deriveIntegerOperationValueTypeFromInt32(rightValueType)
	case ExpressionValueType.Int64:
		return eote.deriveIntegerOperationValueTypeFromInt64(rightValueType)
	case ExpressionValueType.Decimal:
		fallthrough
	case ExpressionValueType.Double:
		fallthrough
	case ExpressionValueType.String:
		fallthrough
	case ExpressionValueType.Guid:
		fallthrough
	case ExpressionValueType.DateTime:
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"" + leftValueType.String() + "\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveIntegerOperationValueTypeFromBoolean(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Boolean\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveIntegerOperationValueTypeFromInt32(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Int32\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveIntegerOperationValueTypeFromInt64(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Int64\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueType(leftValueType, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	switch leftValueType {
	case ExpressionValueType.Boolean:
		return eote.deriveComparisonOperationValueTypeFromBoolean(rightValueType)
	case ExpressionValueType.Int32:
		return eote.deriveComparisonOperationValueTypeFromInt32(rightValueType)
	case ExpressionValueType.Int64:
		return eote.deriveComparisonOperationValueTypeFromInt64(rightValueType)
	case ExpressionValueType.Decimal:
		return eote.deriveComparisonOperationValueTypeFromDecimal(rightValueType)
	case ExpressionValueType.Double:
		return eote.deriveComparisonOperationValueTypeFromDouble(rightValueType)
	case ExpressionValueType.String:
		return leftValueType, nil
	case ExpressionValueType.Guid:
		return eote.deriveComparisonOperationValueTypeFromGuid(rightValueType)
	case ExpressionValueType.DateTime:
		return eote.deriveComparisonOperationValueTypeFromDateTime(rightValueType)
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromBoolean(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Boolean\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromInt32(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Int32\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromInt64(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Int64\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromDecimal(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Decimal\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromDouble(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Double\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromGuid(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"Guid\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveComparisonOperationValueTypeFromDateTime(rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
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
		return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"DateTime\" and \"" + rightValueType.String() + "\"")
	default:
		return ZeroExpressionValueType, errors.New("unexpected expression value type encountered")
	}
}

func (eote ExpressionOperatorTypeEnum) deriveBooleanOperationValueType(leftValueType, rightValueType ExpressionValueTypeEnum) (ExpressionValueTypeEnum, error) {
	if leftValueType == ExpressionValueType.Boolean && rightValueType == ExpressionValueType.Boolean {
		return ExpressionValueType.Boolean, nil
	}

	return ZeroExpressionValueType, errors.New("cannot perform \"" + eote.String() + "\" operation on \"" + leftValueType.String() + "\" and \"" + rightValueType.String() + "\"")
}
