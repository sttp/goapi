//******************************************************************************************************
//  DataColumn.go - Gbtc
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

import "fmt"

// DataColumn represents a column, i.e., a field, in a DataTable defining a name and a data type.
// Data columns can also be computed where its value would be derived from other columns and
// functions (https://sttp.github.io/documentation/filter-expressions/) defined in an expression.
type DataColumn struct {
	parent     *DataTable
	name       string
	dataType   DataTypeEnum
	expression string
	computed   bool
	index      int
}

func newDataColumn(parent *DataTable, name string, dataType DataTypeEnum, expression string) *DataColumn {
	return &DataColumn{
		parent:     parent,
		name:       name,
		dataType:   dataType,
		expression: expression,
		computed:   len(expression) > 0,
		index:      -1,
	}
}

// Parent gets the parent DataTable of the DataColumn.
func (dc *DataColumn) Parent() *DataTable {
	return dc.parent
}

// Name gets the column name of the DataColumn.
func (dc *DataColumn) Name() string {
	return dc.name
}

// Type gets the column DataType enumeration value of the DataColumn.
func (dc *DataColumn) Type() DataTypeEnum {
	return dc.dataType
}

// Expression gets the column expression value of the DataColumn, if any.
func (dc *DataColumn) Expression() string {
	return dc.expression
}

// Computed gets a flag that determines if the DataColumn is a computed value,
// i.e., has a defined expression.
func (dc *DataColumn) Computed() bool {
	return dc.computed
}

// Index gets the index of the DataColumn within its parent DataTable columns collection.
func (dc *DataColumn) Index() int {
	return dc.index
}

// String gets a representation of the DataColumn as a string.
func (dc *DataColumn) String() string {
	dataType := dc.dataType.String()

	if dc.computed {
		dataType = "Computed " + dataType
	}

	return fmt.Sprintf("%s (%s)", dc.name, dataType)
}
