//******************************************************************************************************
//  DataSet_test.go - Gbtc
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
//  10/11/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package data

import (
	"strconv"
	"testing"

	"github.com/sttp/goapi/sttp/guid"
)

func createDataColumn(dataTable *DataTable, columnName string, dataType DataTypeEnum) int {
	dataColumn := dataTable.CreateColumn(columnName, dataType, "")
	dataTable.AddColumn(dataColumn)
	return dataTable.ColumnByName(columnName).Index()
}

func createDataSet() (*DataSet, int, int, guid.Guid, guid.Guid) {
	dataSet := NewDataSet()
	dataTable := dataSet.CreateTable("ActiveMeasurements")
	var dataRow *DataRow

	signalIDField := createDataColumn(dataTable, "SignalID", DataType.Guid)
	signalTypeField := createDataColumn(dataTable, "SignalType", DataType.String)

	statID := guid.New()
	dataRow = dataTable.CreateRow()
	dataRow.SetValue(signalIDField, statID)
	dataRow.SetValue(signalTypeField, "STAT")
	dataTable.AddRow(dataRow)

	freqID := guid.New()
	dataRow = dataTable.CreateRow()
	dataRow.SetValue(signalIDField, freqID)
	dataRow.SetValue(signalTypeField, "FREQ")
	dataTable.AddRow(dataRow)

	dataSet.AddTable(dataTable)

	return dataSet, signalIDField, signalTypeField, statID, freqID
}

func TestCreateDataSet(t *testing.T) {
	dataSet, _, _, _, _ := createDataSet()

	if dataSet.TableCount() != 1 {
		t.Fatal("TestCreateDataSet: expected table count of 1, received: " + strconv.Itoa(dataSet.TableCount()))
	}

	dataTable := dataSet.Tables()[0]

	if dataTable.RowCount() != 2 {
		t.Fatal("TestCreateDataSet: expected row count of 2, received: " + strconv.Itoa(dataTable.RowCount()))
	}
}
