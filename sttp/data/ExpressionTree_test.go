//******************************************************************************************************
//  ExpressionTree_test.go - Gbtc
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
	"fmt"
	"math"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/shopspring/decimal"
	"github.com/sttp/goapi/sttp/guid"
	"github.com/sttp/goapi/sttp/xml"
)

func testEvaluateBooleanLiteralExpression(t *testing.T, b bool) {
	result, err := EvaluateExpression(strconv.FormatBool(b), false)

	if err != nil {
		t.Fatal("TestEvaluateBooleanLiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateBooleanLiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestEvaluateBooleanLiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.BooleanValue()

	if err != nil {
		t.Fatal("TestEvaluateBooleanLiteralExpression: failed to retrieve value: " + err.Error())
	}

	if ve != b {
		t.Fatal("TestEvaluateBooleanLiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateBooleanLiteralExpression(t *testing.T) {
	testEvaluateBooleanLiteralExpression(t, false)
	testEvaluateBooleanLiteralExpression(t, true)
}

func testEvaluateInt32LiteralExpression(t *testing.T, i32 int32) {
	result, err := EvaluateExpression(strconv.FormatInt(int64(i32), 10), false)

	if err != nil {
		t.Fatal("TestEvaluateInt32LiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateInt32LiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestEvaluateInt32LiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.Int32Value()

	if err != nil {
		t.Fatal("TestEvaluateInt32LiteralExpression: failed to retrieve value: " + err.Error())
	}

	if ve != i32 {
		t.Fatal("TestEvaluateInt32LiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateInt32LiteralExpression(t *testing.T) {
	testEvaluateInt32LiteralExpression(t, math.MinInt32 + 1) // Min int32 value interpreted as int64
	testEvaluateInt32LiteralExpression(t, -1)
	testEvaluateInt32LiteralExpression(t, 0)
	testEvaluateInt32LiteralExpression(t, 1)
	testEvaluateInt32LiteralExpression(t, math.MaxInt32)
}

func testEvaluateInt64LiteralExpression(t *testing.T, i64 int64) {
	result, err := EvaluateExpression(strconv.FormatInt(i64, 10), false)

	if err != nil {
		t.Fatal("TestEvaluateInt64LiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateInt64LiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Int64 {
		t.Fatal("TestEvaluateInt64LiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.Int64Value()

	if err != nil {
		t.Fatal("TestEvaluateInt64LiteralExpression: failed to retrieve value: " + err.Error())
	}

	if ve != i64 {
		t.Fatal("TestEvaluateInt64LiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateInt64LiteralExpression(t *testing.T) {
	testEvaluateInt64LiteralExpression(t, math.MinInt64 + 1) // Min int64 value interpreted as Decimal
	testEvaluateInt64LiteralExpression(t, math.MaxInt64)
}

func TestEvaluateDecimalLiteralExpression(t *testing.T) {
	const dec string = "-9223372036854775809.87686876"
	var d decimal.Decimal
	d, _ = decimal.NewFromString(dec)

	result, err := EvaluateExpression(dec, false)

	if err != nil {
		t.Fatal("TestEvaluateDecimalLiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateDecimalLiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Decimal {
		t.Fatal("TestEvaluateDecimalLiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.DecimalValue()

	if err != nil {
		t.Fatal("TestEvaluateDecimalLiteralExpression: failed to retrieve value: " + err.Error())
	}

	if !ve.Equal(d) {
		t.Fatal("TestEvaluateDecimalLiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateDoubleLiteralExpression(t *testing.T) {
	var d float64 = 123.456e-6

	result, err := EvaluateExpression("123.456E-6", false)

	if err != nil {
		t.Fatal("TestEvaluateDoubleLiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateDoubleLiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Double {
		t.Fatal("TestEvaluateDoubleLiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.DoubleValue()

	if err != nil {
		t.Fatal("TestEvaluateDoubleLiteralExpression: failed to retrieve value: " + err.Error())
	}

	if ve != d {
		t.Fatal("TestEvaluateDoubleLiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateStringLiteralExpression(t *testing.T) {
	s := "'Hello, literal string expression'"

	result, err := EvaluateExpression(s, false)

	if err != nil {
		t.Fatal("TestEvaluateStringLiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateStringLiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.String {
		t.Fatal("TestEvaluateStringLiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.StringValue()

	if err != nil {
		t.Fatal("TestEvaluateStringLiteralExpression: failed to retrieve value: " + err.Error())
	}

	if ve != s[1:len(s)-1] {
		t.Fatal("TestEvaluateStringLiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateGuidLiteralExpression(t *testing.T) {
	g := guid.New()

	result, err := EvaluateExpression(g.String(), false)

	if err != nil {
		t.Fatal("TestEvaluateGuidLiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateGuidLiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Guid {
		t.Fatal("TestEvaluateGuidLiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.GuidValue()

	if err != nil {
		t.Fatal("TestEvaluateGuidLiteralExpression: failed to retrieve value: " + err.Error())
	}

	if !ve.Equal(g) {
		t.Fatal("TestEvaluateGuidLiteralExpression: retrieved value does not match source")
	}
}

func TestEvaluateDateTimeLiteralExpression(t *testing.T) {
	sr := "2006-01-01 00:00:00"
	dt, _ := dateparse.ParseAny(sr)
	testDateTime(t, dt, sr)

	sr = "2019-01-1 00:00:59.999"
	dt, _ = dateparse.ParseAny(sr)
	testDateTime(t, dt, sr)

	dt = time.Now()
	testDateTime(t, dt, dt.Format(time.RFC3339Nano))

	dt = time.Now().UTC()
	testDateTime(t, dt, dt.Format(time.RFC3339Nano))
}

func testDateTime(t *testing.T, dt time.Time, sr string) {
	result, err := EvaluateExpression("#"+sr+"#", false)

	if err != nil {
		t.Fatal("TestEvaluateDateTimeLiteralExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateDateTimeLiteralExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestEvaluateDateTimeLiteralExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.DateTimeValue()

	if err != nil {
		t.Fatal("TestEvaluateDateTimeLiteralExpression: failed to retrieve value: " + err.Error())
	}

	if !ve.Equal(dt) {
		t.Fatal("TestEvaluateDateTimeLiteralExpression: retrieved value does not match source")
	}
}

//gocyclo: ignore
func TestSignalIDSetExpressions(t *testing.T) {
	dataSet, _, _, statID, freqID := createDataSet()

	idSet, err := SelectSignalIDSet(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'FREQ'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestSignalIDSetExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != freqID {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'STAT'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestSignalIDSetExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != statID {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, statID.String(), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestSignalIDSetExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != statID {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, ";;"+statID.String()+";;;", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestSignalIDSetExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != statID {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	freqUUID := freqID.String()
	idSet, err = SelectSignalIDSet(dataSet, "'"+freqUUID[1:len(freqUUID)-1]+"'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestSignalIDSetExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != freqID {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, fmt.Sprintf("%s;%s;%s", statID.String(), freqID.String(), statID.String()), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 2 {
		t.Fatal("TestSignalIDSetExpressions: expected 2 results, received: " + strconv.Itoa(len(idSet)))
	}

	if !idSet.Contains(statID) || !idSet.Contains(freqID) {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, fmt.Sprintf("%s;%s;%s;FILTER ActiveMeasurements WHERE True", statID.String(), freqID.String(), statID.String()), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 2 {
		t.Fatal("TestSignalIDSetExpressions: expected 2 results, received: " + strconv.Itoa(len(idSet)))
	}

	if !idSet.Contains(statID) || !idSet.Contains(freqID) {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, "FILTER ActiveMeasurements WHERE SignalID = '"+freqUUID[1:len(freqUUID)-1]+"'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestSignalIDSetExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != freqID {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, "FILTER ActiveMeasurements WHERE SignalID = '"+freqUUID[1:len(freqUUID)-1]+"' OR SignalID = "+statID.String(), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSignalIDSetExpressions: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 2 {
		t.Fatal("TestSignalIDSetExpressions: expected 2 results, received: " + strconv.Itoa(len(idSet)))
	}

	if !idSet.Contains(statID) || !idSet.Contains(freqID) {
		t.Fatal("TestSignalIDSetExpressions: retrieve Guid value does not match source")
	}

	_, err = SelectSignalIDSet(dataSet, "", "", nil, false)

	if err == nil {
		t.Fatal("TestSignalIDSetExpressions: error expected, received none")
	}

	_, err = SelectSignalIDSet(dataSet, "bad expression", "ActiveMeasurements", nil, false)

	if err == nil {
		t.Fatal("TestSignalIDSetExpressions: error expected, received none")
	}
}

//gocyclo: ignore
func TestSelectDataRowsExpressions(t *testing.T) {
	dataSet, signalIDField, signalTypeField, statID, freqID := createDataSet()

	rows, err := SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'FREQ'; FILTER ActiveMeasurements WHERE SignalType = 'STAT' ORDER BY SignalID", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	// FREQ should be before STAT because of multiple statement evaluation order
	if !getRowGuid(t, rows[0], signalIDField).Equal(freqID) || !getRowGuid(t, rows[1], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'FREQ' OR SignalType = 'STAT'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	// Row with stat comes before row with freq (single expression statement)
	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'FREQ' OR SignalType = 'STAT' ORDER BY BINARY SignalType", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	// FREQ should sort before STAT with order by
	if !getRowGuid(t, rows[0], signalIDField).Equal(freqID) || !getRowGuid(t, rows[1], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'STAT' OR SignalType = 'FREQ' ORDER BY SignalType DESC", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	// Now descending
	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	freqUUID := freqID.String()
	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalID = "+statID.String()+" OR SignalID = '"+freqUUID[1:len(freqUUID)-1]+"' ORDER BY SignalType", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	// FREQ should sort before STAT with order by
	if !getRowGuid(t, rows[0], signalIDField).Equal(freqID) || !getRowGuid(t, rows[1], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalID = "+statID.String()+" OR SignalID = '"+freqUUID[1:len(freqUUID)-1]+"' ORDER BY SignalType;"+statID.String(), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	// Because expression includes Guid statID as a literal (at the end), it will parse first
	// regardless of order by in filter statement
	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE True", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE IsNull(NULL, False) OR Coalesce(Null, true)", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE IIf(IsNull(NULL, False) OR Coalesce(Null, true), Len(SignalType) == 4, false)", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType IS !NULL", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE Len(SubStr(Coalesce(Trim(SignalType), 'OTHER'), 0, 0X2)) = 2", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE LEN(SignalTYPE) > 3.5", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE Len(SignalType) & 0x4 == 4", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE RegExVal('ST.+', SignalType) == 'STAT'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 1 {
		t.Fatal("TestSelectDataRowsExpressions: expected 1 result, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE RegExMatch('FR.+', SignalType)", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 1 {
		t.Fatal("TestSelectDataRowsExpressions: expected 1 result, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType IN ('FREQ', 'STAT') ORDER BY SignalType", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(freqID) || !getRowGuid(t, rows[1], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalID IN ("+statID.String()+", "+freqID.String()+")", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType LIKE 'ST%'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 1 {
		t.Fatal("TestSelectDataRowsExpressions: expected 1 result, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType LIKE '*EQ'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 1 {
		t.Fatal("TestSelectDataRowsExpressions: expected 1 result, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType LIKE '*TA%'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 1 {
		t.Fatal("TestSelectDataRowsExpressions: expected 1 result, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE -Len(SignalType) <= 0", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	// number converted to string and compared
	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType == 0", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 0 {
		t.Fatal("TestSelectDataRowsExpressions: expected 0 results, received: " + strconv.Itoa(len(rows)))
	}

	// number converted to string and compared
	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE SignalType > 99", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	rows, err = SelectDataRows(dataSet, "FILTER ActiveMeasurements WHERE Len(SignalType) / 0x2 = 2", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestSelectDataRowsExpressions: error executing SelectDataRows: " + err.Error())
	}

	if len(rows) != 2 {
		t.Fatal("TestSelectDataRowsExpressions: expected 2 results, received: " + strconv.Itoa(len(rows)))
	}

	if !getRowGuid(t, rows[0], signalIDField).Equal(statID) || !getRowGuid(t, rows[1], signalIDField).Equal(freqID) {
		t.Fatal("TestSelectDataRowsExpressions: retrieve Guid value or order does not match source")
	}

	if getRowString(t, rows[0], signalTypeField) != "STAT" || getRowString(t, rows[1], signalTypeField) != "FREQ" {
		t.Fatal("TestSelectDataRowsExpressions: retrieve string value or order does not match source")
	}
}

func getRowGuid(t *testing.T, row *DataRow, columnIndex int) guid.Guid {
	id, null, err := row.GuidValue(columnIndex)

	if null || err != nil {
		t.Fatal("TestSelectDataRowsExpressions: failed to retrieve Guid value")
	}

	return id
}

func getRowString(t *testing.T, row *DataRow, columnIndex int) string {
	value, null, err := row.StringValue(columnIndex)

	if null || err != nil {
		t.Fatal("TestSelectDataRowsExpressions: failed to retrieve string value")
	}

	return value
}

//gocyclo: ignore
func TestMetadataExpressions(t *testing.T) {
	// Two sample metadata files exist, test both
	for i := 0; i < 2; i++ {
		fileName := fmt.Sprintf("../../test/MetadataSample%d.xml", i+1)

		var doc xml.XmlDocument
		err := doc.LoadXmlFromFile(fileName)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error loading XML document: " + err.Error())
		}

		dataSet := NewDataSet()
		err = dataSet.ParseXmlDocument(&doc)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error loading DataSet from XML document: " + err.Error())
		}

		if dataSet.TableCount() != 4 {
			t.Fatal("TestMetadataExpressions: expected 4 results, received: " + strconv.Itoa(dataSet.TableCount()))
		}

		table := dataSet.Table("MeasurementDetail")

		if table == nil {
			t.Fatal("TestMetadataExpressions: table not found in DataSet")
		}

		if table.ColumnCount() != 11 {
			t.Fatal("TestMetadataExpressions: expected table column count: " + strconv.Itoa(table.ColumnCount()))
		}

		if table.ColumnByName("ID") == nil {
			t.Fatal("TestMetadataExpressions: missing expected table column")
		}

		if table.ColumnByName("id").Type() != DataType.String {
			t.Fatal("TestMetadataExpressions: unexpected table column type")
		}

		if table.ColumnByName("SignalID") == nil {
			t.Fatal("TestMetadataExpressions: missing expected table column")
		}

		if table.ColumnByName("signalID").Type() != DataType.Guid {
			t.Fatal("TestMetadataExpressions: unexpected table column type")
		}

		if table.RowCount() == 0 {
			t.Fatal("TestMetadataExpressions: unexpected empty table")
		}

		table = dataSet.Table("DeviceDetail")

		if table == nil {
			t.Fatal("TestMetadataExpressions: table not found in DataSet")
		}

		if table.ColumnCount() != 19+i { // Second test adds a computed column
			t.Fatal("TestMetadataExpressions: expected table column count: " + strconv.Itoa(table.ColumnCount()))
		}

		if table.ColumnByName("ACRONYM") == nil {
			t.Fatal("TestMetadataExpressions: missing expected table column")
		}

		if table.ColumnByName("Acronym").Type() != DataType.String {
			t.Fatal("TestMetadataExpressions: unexpected table column type")
		}

		if table.ColumnByName("Name") == nil {
			t.Fatal("TestMetadataExpressions: missing expected table column")
		}

		if table.ColumnByName("name").Type() != DataType.String {
			t.Fatal("TestMetadataExpressions: unexpected table column type")
		}

		if table.RowCount() != 1 {
			t.Fatal("TestMetadataExpressions: expected table row count: " + strconv.Itoa(table.RowCount()))
		}

		dataRow := table.Row(0)

		acronym, null, err := dataRow.StringValueByName("Acronym")

		if null || err != nil {
			t.Fatal("TestMetadataExpressions: unexpected NULL column value in row")
		}

		name, null, err := dataRow.StringValueByName("Name")

		if null || err != nil {
			t.Fatal("TestMetadataExpressions: unexpected NULL column value in row")
		}

		if !strings.EqualFold(acronym, name) {
			t.Fatal("TestMetadataExpressions: unexpected column values in row")
		}

		// In test data set, DeviceDetail.OriginalSource is null
		if _, null, _ = dataRow.StringValueByName("OriginalSource"); !null {
			t.Fatal("TestMetadataExpressions: unexpected column value in row")
		}

		// In test data set, DeviceDetail.ParentAcronym is not null, but is an empty string
		if parentAcronym, null, _ := dataRow.StringValueByName("ParentAcronym"); len(parentAcronym) > 0 || null {
			t.Fatal("TestMetadataExpressions: unexpected column value in row")
		}

		idSet, err := SelectSignalIDSet(dataSet, "FILTER MeasurementDetail WHERE SignalAcronym = 'FREQ'", "MeasurementDetail", nil, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER TOP 8 MeasurementDetail WHERE SignalAcronym = 'STAT'", "MeasurementDetail", nil, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 8 {
			t.Fatal("TestMetadataExpressions: expected 8 results, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER TOP 0 MeasurementDetail WHERE SignalAcronym = 'STAT'", "MeasurementDetail", nil, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 0 {
			t.Fatal("TestMetadataExpressions: expected 0 results, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER TOP -1 MeasurementDetail WHERE SignalAcronym = 'STAT'", "MeasurementDetail", nil, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) == 0 {
			t.Fatal("TestMetadataExpressions: expected non-zero result set, received: " + strconv.Itoa(len(idSet)))
		}

		deviceDetailIDFields := new(TableIDFields)
		deviceDetailIDFields.SignalIDFieldName = "UniqueID"
		deviceDetailIDFields.MeasurementKeyFieldName = "Name"
		deviceDetailIDFields.PointTagFieldName = "Acronym"

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Convert(Longitude, 'System.Int32') = -89", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Convert(latitude, 'int16') = 35", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Convert(latitude, 'Int32') = 35", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Convert(Convert(Latitude, 'Int32'), 'String') = 35", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Convert(Latitude, 'Single') > 35", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		// test decimal comparison
		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Longitude < 0.0", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Acronym IN ('Test', 'Shelby')", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE Acronym not IN ('Test', 'Apple')", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE NOT (Acronym IN ('Test', 'Apple'))", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		idSet, err = SelectSignalIDSet(dataSet, "FILTER DeviceDetail WHERE NOT Acronym !IN ('Shelby', 'Apple')", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectSignalIDSet: " + err.Error())
		}

		if len(idSet) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(idSet)))
		}

		rows, err := SelectDataRows(dataSet, "Acronym LIKE 'Shel%'", "DeviceDetail", deviceDetailIDFields, false)

		if err != nil {
			t.Fatal("TestMetadataExpressions: error executing SelectDataRows: " + err.Error())
		}

		if len(rows) != 1 {
			t.Fatal("TestMetadataExpressions: expected 1 result, received: " + strconv.Itoa(len(rows)))
		}
	}
}

//gocyclo: ignore
func TestBasicExpressions(t *testing.T) {
	var doc xml.XmlDocument
	err := doc.LoadXmlFromFile("../../test/MetadataSample2.xml")

	if err != nil {
		t.Fatal("TestBasicExpressions: error loading XML document: " + err.Error())
	}

	dataSet := NewDataSet()
	err = dataSet.ParseXmlDocument(&doc)

	if err != nil {
		t.Fatal("TestBasicExpressions: error loading DataSet from XML document: " + err.Error())
	}

	dataRows, err := dataSet.Table("MeasurementDetail").Select("SignalAcronym = 'STAT'", "", -1)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during table Select: " + err.Error())
	}

	if len(dataRows) != 116 {
		t.Fatal("TestBasicExpressions: expected 116 results, received: " + strconv.Itoa(len(dataRows)))
	}

	dataRows, err = dataSet.Table("PhasorDetail").Select("Type = 'V'", "", -1)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during table Select: " + err.Error())
	}

	if len(dataRows) != 2 {
		t.Fatal("TestBasicExpressions: expected 2 results, received: " + strconv.Itoa(len(dataRows)))
	}

	valueExpression, err := EvaluateDataRowExpression(dataSet.Table("SchemaVersion").Row(0), "VersionNumber > 0", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err := valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	dataRow := dataSet.Table("DeviceDetail").Row(0)

	valueExpression, err = EvaluateDataRowExpression(dataRow, "AccessID % 2 = 0 AND FramesPerSecond % 4 <> 2 OR AccessID % 1 = 0", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "AccessID % 2 = 0 AND (FramesPerSecond % 4 <> 2 OR -AccessID % 1 = 0)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "AccessID % 2 = 0 AND (FramesPerSecond % 4 <> 2 AND AccessID % 1 = 0)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "AccessID % 2 >= 0 || (FramesPerSecond % 4 <> 2 AND AccessID % 1 = 0)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "AccessID % 2 = 0 OR FramesPerSecond % 4 != 2 && AccessID % 1 == 0", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "!AccessID % 2 = 0 || FramesPerSecond % 4 = 0x2 && AccessID % 1 == 0", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "NOT AccessID % 2 = 0 OR FramesPerSecond % 4 >> 0x1 = 1 && AccessID % 1 == 0x0", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "!AccessID % 2 = 0 OR FramesPerSecond % 4 >> 1 = 1 && AccessID % 3 << 1 & 4 >= 4", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "OriginalSource IS NULL", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "ParentAcronym IS NOT NULL", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "NOT ParentAcronym IS NULL && Len(parentAcronym) == 0", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "-FramesPerSecond", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	value, err := valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if value != -30 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "~FramesPerSecond", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	value, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if value != -31 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "~FramesPerSecond * -1 - 1 << -2", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	value, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if value != -2147483648 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "NOT True", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "!True", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "~True", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Len(IsNull(OriginalSource, 'A')) = 1", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExMatch('SH', Acronym)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExMatch('SH', Name)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExMatch('S[hH]', Name) && RegExMatch('S[hH]', Acronym)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExVal('Sh\\w+', Name)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.String {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	s, err := valueExpression.StringValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if s != "Shelby" {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "SubStr(RegExVal('Sh\\w+', Name), 2) == 'ElbY'", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "SubStr(RegExVal('Sh\\w+', Name), 3, 2) == 'lB'", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExVal('Sh\\w+', Name) IN ('NT', Acronym, 'NT')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExVal('Sh\\w+', Name) IN ===('NT', Acronym, 'NT')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "RegExVal('Sh\\w+', Name) IN BINARY ('NT', Acronym, 3.05)", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Name IN===(0x9F, Acronym, 'Shelby')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Acronym LIKE === 'Sh*'", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "name LiKe binaRY 'SH%'", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Name === 'Shelby' && Name== 'SHelBy' && Name !=='SHelBy'", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Now()", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err := valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	now := time.Now()

	if dt != now && !dt.Before(now) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "UtcNow()", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	now = time.Now().UTC()

	if dt != now && !dt.Before(now) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "#2019-02-04T03:00:52.73-05:00#", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if dt.Month() != 2 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "#2019-02-04T03:00:52.73-05:00#", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if dt.Day() != 4 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04T03:00:52.73-05:00#, 'Year')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err := valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 2019 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019/02/04 03:00:52.73-05:00#, 'Month')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 2 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04 03:00:52.73-05:00#, 'Day')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 4 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04 03:00#, 'Hour')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 3 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04 03:00:52.73-05:00#, 'Hour')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 8 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2/4/2019 3:21:55#, 'Minute')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 21 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#02/04/2019 03:21:55.33#, 'Second')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 55 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#02/04/2019 03:21:5.033#, 'Millisecond')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 33 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(DateAdd('2019-02-04 03:00:52.73-05:00', 1, 'Year'), 'year')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 2020 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd('2019-02-04', 2, 'Month')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 4, 4, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd(#1/31/2019#, 1, 'Day')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 2, 1, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd(#2019-01-31#, 2, 'Week')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 2, 14, 0, 0, 0, 0, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd(#2019-01-31#, 25, 'Hour')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 2, 1, 1, 0, 0, 0, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd(#2018-12-31 23:58#, 3, 'Minute')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 1, 1, 0, 1, 0, 0, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd('2019-01-1 00:59', 61, 'Second')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 1, 1, 1, 0, 1, 0, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd('2019-01-1 00:00:59.999', 2, 'Millisecond')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 1, 1, 0, 1, 0, 1000000, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateAdd(#1/1/2019 0:0:1.029#, -FramesPerSecond, 'Millisecond')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	dt, err = valueExpression.DateTimeValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !dt.Equal(time.Date(2019, 1, 1, 0, 0, 0, 999000000, time.UTC)) {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'Year')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 2 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'month')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 35 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'DAY')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 1095 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'Week')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 156 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'WeekDay')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 1095 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'Hour')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 26280 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'Minute')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 1576800 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2006-01-01 00:00:00#, #2008-12-31 00:00:00#, 'Second')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 94608000 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DateDiff(#2008-12-30 00:02:50.546#, '2008-12-31', 'Millisecond')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 86229454 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04 03:00:52.73-05:00#, 'DayOfyear')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 35 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04 03:00:52.73-05:00#, 'Week')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 6 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "DatePart(#2019-02-04 03:00:52.73-05:00#, 'WeekDay')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if i32 != 2 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "IsDate(#2019-02-04 03:00:52.73-05:00#) AND IsDate('2/4/2019') ANd isdate(updatedon) && !ISDATE(2.5) && !IsDate('ImNotADate')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "IsInteger(32768) AND IsInteger('1024') and ISinTegeR(FaLsE) And isinteger(accessID) && !ISINTEGER(2.5) && !IsInteger('ImNotAnInteger')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "IsGuid({9448a8b5-35c1-4dc7-8c42-8712153ac08a}) AND IsGuid('9448a8b5-35c1-4dc7-8c42-8712153ac08a') anD isGuid(9448a8b5-35c1-4dc7-8c42-8712153ac08a) AND IsGuid(Convert(9448a8b5-35c1-4dc7-8c42-8712153ac08a, 'string')) aND isguid(nodeID) && !ISGUID(2.5) && !IsGuid('ImNotAGuid')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "IsNumeric(32768) && isNumeric(123.456e-67) AND IsNumeric(3.14159265) and ISnumeric(true) AND IsNumeric('1024' ) and IsNumeric(2.5) aNd isnumeric(longitude) && !ISNUMERIC(9448a8b5-35c1-4dc7-8c42-8712153ac08a) && !IsNumeric('ImNotNumeric')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestBasicExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	result, err = valueExpression.BooleanValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if !result {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Convert(maxof(12, '99.9', 99.99), 'Double')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	f64, err := valueExpression.DoubleValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if f64 != 99.99 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}

	valueExpression, err = EvaluateDataRowExpression(dataRow, "Convert(minof(12, '99.9', 99.99), 'double')", false)

	if err != nil {
		t.Fatal("TestBasicExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	f64, err = valueExpression.DoubleValue()

	if err != nil {
		t.Fatal("TestBasicExpressions: error getting value: " + err.Error())
	}

	if f64 != 12.0 {
		t.Fatal("TestBasicExpressions: unexpected value expression result")
	}
}

func TestNegativeExpressions(t *testing.T) {
	_, err := EvaluateExpression("Convert(123, 'unknown')", false)

	if err == nil {
		t.Fatal("TestNegativeExpressions: expected error during EvaluateDataRowExpression")
	}

	_, err = EvaluateExpression("I am a bad expression", false)

	if err == nil {
		t.Fatal("TestNegativeExpressions: expected error during EvaluateDataRowExpression")
	}
}

//gocyclo: ignore
func TestMiscExpressions(t *testing.T) {
	var doc xml.XmlDocument
	err := doc.LoadXmlFromFile("../../test/MetadataSample2.xml")

	if err != nil {
		t.Fatal("TestMiscExpressions: error loading XML document: " + err.Error())
	}

	dataSet := NewDataSet()
	err = dataSet.ParseXmlDocument(&doc)

	if err != nil {
		t.Fatal("TestMiscExpressions: error loading DataSet from XML document: " + err.Error())
	}

	dataRow := dataSet.Table("DeviceDetail").Row(0)

	valueExpression, err := EvaluateDataRowExpression(dataRow, "AccessID ^ 2 + FramesPerSecond XOR 4", false)

	if err != nil {
		t.Fatal("TestMiscExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestMiscExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err := valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestMiscExpressions: error getting value: " + err.Error())
	}

	if i32 != 38 {
		t.Fatal("TestMiscExpressions: unexpected value expression result")
	}

	// test computed column with expression defined in schema
	g := guid.New()
	valueExpression, err = EvaluateDataRowExpression(dataRow, g.String(), false)

	if err != nil {
		t.Fatal("TestMiscExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Guid {
		t.Fatal("TestMiscExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	ge, err := valueExpression.GuidValue()

	if err != nil {
		t.Fatal("TestMiscExpressions: error getting value: " + err.Error())
	}

	if !g.Equal(ge) {
		t.Fatal("TestMiscExpressions: unexpected value expression result")
	}

	// test edge case of evaluating standalone Guid not used as a row identifier
	valueExpression, err = EvaluateDataRowExpression(dataRow, "ComputedCol", false)

	if err != nil {
		t.Fatal("TestMiscExpressions: error during EvaluateDataRowExpression: " + err.Error())
	}

	if valueExpression.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestMiscExpressions: unexpected value expression type: " + valueExpression.Type().String())
	}

	i32, err = valueExpression.Int32Value()

	if err != nil {
		t.Fatal("TestMiscExpressions: error getting value: " + err.Error())
	}

	if i32 != 32 {
		t.Fatal("TestMiscExpressions: unexpected value expression result")
	}
}

func TestFilterExpressionStatementCount(t *testing.T) {
	dataSet, _, _, statID, freqID := createDataSet()

	//                                                          Statements:  1  2  3  4
	parser, err := NewFilterExpressionParserForDataSet(dataSet, fmt.Sprintf("%s;%s;%s;FILTER ActiveMeasurements WHERE True", statID.String(), freqID.String(), statID.String()), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterExpressionStatementCount: unexpected NewFilterExpressionParserForDataSet error: " + err.Error())
	}

	parser.TrackFilteredRows = false
	parser.TrackFilteredSignalIDs = true

	if err := parser.Evaluate(true, false); err != nil {
		t.Fatal("TestFilterExpressionStatementCount: unexpected Evaluate error: " + err.Error())
	}

	idSet := parser.FilteredSignalIDSet()

	if len(idSet) != 2 {
		t.Fatal("TestFilterExpressionStatementCount: expected 2 results, received: " + strconv.Itoa(len(idSet)))
	}

	if !idSet.Contains(statID) || !idSet.Contains(freqID) {
		t.Fatal("TestFilterExpressionStatementCount: retrieve Guid value does not match source")
	}

	if parser.FilterExpressionStatementCount() != 4 {
		t.Fatal("TestFilterExpressionStatementCount: expected 4 results, received: " + strconv.Itoa(parser.FilterExpressionStatementCount()))
	}
}
