//******************************************************************************************************
//  DataType.go - Gbtc
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

import "strings"

// DataTypeEnum defines the type of the DataType enumeration.
type DataTypeEnum int

// DataType is an enumeration of the possible data types for a DataColumn.
var DataType = struct {
	// String represents a Go string data type.
	String DataTypeEnum
	// Boolean represents a Go bool data type.
	Boolean DataTypeEnum
	// DateTime represents a Go time.Time data type.
	DateTime DataTypeEnum
	// Single represents a Go float32 data type.
	Single DataTypeEnum
	// Double represents a Go float64 data type.
	Double DataTypeEnum
	// Decimal represents a Go decimal.Decimal data type.
	// Type defined in github.com/shopspring/decimal.
	Decimal DataTypeEnum
	// Guid represents a Go guid.Guid data type.
	// Type defined in github.com/sttp/goapi/sttp/guid.
	Guid DataTypeEnum
	// Int8 represents a Go int8 data type.
	Int8 DataTypeEnum
	// Int16 represents a Go int16 data type.
	Int16 DataTypeEnum
	// Int32 represents a Go int32 data type.
	Int32 DataTypeEnum
	// Int64 represents a Go int64 data type.
	Int64 DataTypeEnum
	// UInt8 represents a Go uint8 data type.
	UInt8 DataTypeEnum
	// UInt16 represents a Go uint16 data type.
	UInt16 DataTypeEnum
	// UInt32 represents a Go uint32 data type.
	UInt32 DataTypeEnum
	// UInt64 represents a Go uint64 data type.
	UInt64 DataTypeEnum
}{
	String:   0,
	Boolean:  1,
	DateTime: 2,
	Single:   3,
	Double:   4,
	Decimal:  5,
	Guid:     6,
	Int8:     7,
	Int16:    8,
	Int32:    9,
	Int64:    10,
	UInt8:    11,
	UInt16:   12,
	UInt32:   13,
	UInt64:   14,
}

// String gets the DataType enumeration name as a string.
//gocyclo:ignore
func (dte DataTypeEnum) String() string {
	switch dte {
	case DataType.String:
		return "String"
	case DataType.Boolean:
		return "Boolean"
	case DataType.DateTime:
		return "DateTime"
	case DataType.Single:
		return "Single"
	case DataType.Double:
		return "Double"
	case DataType.Decimal:
		return "Decimal"
	case DataType.Guid:
		return "Guid"
	case DataType.Int8:
		return "Int8"
	case DataType.Int16:
		return "Int16"
	case DataType.Int32:
		return "Int32"
	case DataType.Int64:
		return "Int64"
	case DataType.UInt8:
		return "UInt8"
	case DataType.UInt16:
		return "UInt16"
	case DataType.UInt32:
		return "UInt32"
	case DataType.UInt64:
		return "UInt64"
	default:
		return "Undefined"
	}
}

// ParseXsdDataType gets the DataType from the provided XSD data type. Return tuple includes
// boolean value that determines if parse was successful. See XML Schema Language Datatypes
// for possible xsdTypeName values: https://www.w3.org/TR/xmlschema-2/
//gocyclo:ignore
func ParseXsdDataType(xsdTypeName, extDataType string) (DataTypeEnum, bool) {
	switch xsdTypeName {
	case "string":
		if strings.HasPrefix(extDataType, "System.Guid") {
			return DataType.Guid, true
		}
		return DataType.String, true
	case "boolean":
		return DataType.Boolean, true
	case "dateTime":
		return DataType.DateTime, true
	case "float":
		return DataType.Single, true
	case "double":
		return DataType.Double, true
	case "decimal":
		return DataType.Decimal, true
	case "byte": // XSD defines byte as signed 8-bit int
		return DataType.Int8, true
	case "short":
		return DataType.Int16, true
	case "int":
		return DataType.Int32, true
	case "long":
		return DataType.Int64, true
	case "unsignedByte":
		return DataType.UInt8, true
	case "unsignedShort":
		return DataType.UInt16, true
	case "unsignedInt":
		return DataType.UInt32, true
	case "unsignedLong":
		return DataType.UInt64, true
	default:
		return DataType.String, false
	}
}
