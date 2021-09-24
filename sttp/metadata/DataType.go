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

package metadata

// DataTypeEnum defines the type for the DataType enumeration.
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

// Name gets the DataType enumeration name as a string.
func (dte DataTypeEnum) Name() string {
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
