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
	"testing"
	"time"

	"github.com/araddon/dateparse"
	"github.com/shopspring/decimal"
	"github.com/sttp/goapi/sttp/guid"
)

func TestEvaluateBooleanExpression(t *testing.T) {
	b := true

	result, err := EvaluateExpression(strconv.FormatBool(b), false)

	if err != nil {
		t.Fatal("TestEvaluateBooleanExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateBooleanExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Boolean {
		t.Fatal("TestEvaluateBooleanExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.BooleanValue()

	if err != nil {
		t.Fatal("TestEvaluateBooleanExpression: failed to retrieve value: " + err.Error())
	}

	if ve != b {
		t.Fatal("TestEvaluateBooleanExpression: retrieved value does not match source")
	}
}

func TestEvaluateInt32Expression(t *testing.T) {
	var i32 int32 = math.MaxInt32

	result, err := EvaluateExpression(strconv.FormatInt(int64(i32), 10), false)

	if err != nil {
		t.Fatal("TestEvaluateInt32Expression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateInt32Expression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Int32 {
		t.Fatal("TestEvaluateInt32Expression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.Int32Value()

	if err != nil {
		t.Fatal("TestEvaluateInt32Expression: failed to retrieve value: " + err.Error())
	}

	if ve != i32 {
		t.Fatal("TestEvaluateInt32Expression: retrieved value does not match source")
	}
}

func TestEvaluateInt64Expression(t *testing.T) {
	var i64 int64 = math.MaxInt64

	result, err := EvaluateExpression(strconv.FormatInt(i64, 10), false)

	if err != nil {
		t.Fatal("TestEvaluateInt64Expression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateInt64Expression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Int64 {
		t.Fatal("TestEvaluateInt64Expression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.Int64Value()

	if err != nil {
		t.Fatal("TestEvaluateInt64Expression: failed to retrieve value: " + err.Error())
	}

	if ve != i64 {
		t.Fatal("TestEvaluateInt64Expression: retrieved value does not match source")
	}
}

func TestEvaluateDecimalExpression(t *testing.T) {
	const dec string = "-9223372036854775809.87686876"
	var d decimal.Decimal
	d, _ = decimal.NewFromString(dec)

	result, err := EvaluateExpression(dec, false)

	if err != nil {
		t.Fatal("TestEvaluateDecimalExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateDecimalExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Decimal {
		t.Fatal("TestEvaluateDecimalExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.DecimalValue()

	if err != nil {
		t.Fatal("TestEvaluateDecimalExpression: failed to retrieve value: " + err.Error())
	}

	if !ve.Equal(d) {
		t.Fatal("TestEvaluateDecimalExpression: retrieved value does not match source")
	}
}

func TestEvaluateDoubleExpression(t *testing.T) {
	var d float64 = 123.456e-6

	result, err := EvaluateExpression("123.456E-6", false)

	if err != nil {
		t.Fatal("TestEvaluateDoubleExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateDoubleExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Double {
		t.Fatal("TestEvaluateDoubleExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.DoubleValue()

	if err != nil {
		t.Fatal("TestEvaluateDoubleExpression: failed to retrieve value: " + err.Error())
	}

	if ve != d {
		t.Fatal("TestEvaluateDoubleExpression: retrieved value does not match source")
	}
}

func TestEvaluateStringExpression(t *testing.T) {
	s := "'Hello, literal string expression'"

	result, err := EvaluateExpression(s, false)

	if err != nil {
		t.Fatal("TestEvaluateStringExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateStringExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.String {
		t.Fatal("TestEvaluateStringExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.StringValue()

	if err != nil {
		t.Fatal("TestEvaluateStringExpression: failed to retrieve value: " + err.Error())
	}

	if ve != s[1:len(s)-1] {
		t.Fatal("TestEvaluateStringExpression: retrieved value does not match source")
	}
}

func TestEvaluateGuidExpression(t *testing.T) {
	g := guid.New()

	result, err := EvaluateExpression(g.String(), false)

	if err != nil {
		t.Fatal("TestEvaluateGuidExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateGuidExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.Guid {
		t.Fatal("TestEvaluateGuidExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.GuidValue()

	if err != nil {
		t.Fatal("TestEvaluateGuidExpression: failed to retrieve value: " + err.Error())
	}

	if !ve.Equal(g) {
		t.Fatal("TestEvaluateGuidExpression: retrieved value does not match source")
	}
}

func TestEvaluateDateTimeExpression(t *testing.T) {
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
		t.Fatal("TestEvaluateDateTimeExpression: error parsing expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestEvaluateDateTimeExpression: received no result")
	}

	if result.ValueType() != ExpressionValueType.DateTime {
		t.Fatal("TestEvaluateDateTimeExpression: received unexpected type: " + result.ValueType().String())
	}

	ve, err := result.DateTimeValue()

	if err != nil {
		t.Fatal("TestEvaluateDateTimeExpression: failed to retrieve value: " + err.Error())
	}

	if !ve.Equal(dt) {
		t.Fatal("TestEvaluateDateTimeExpression: retrieved value does not match source")
	}
}

func TestFilterSignalIDs(t *testing.T) {
	dataSet, _, _, statID, freqID := createDataSet()

	idSet, err := SelectSignalIDSet(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'FREQ'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterSignalIDs: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestFilterSignalIDs: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != freqID {
		t.Fatal("TestFilterSignalIDs: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, "FILTER ActiveMeasurements WHERE SignalType = 'STAT'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterSignalIDs: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestFilterSignalIDs: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != statID {
		t.Fatal("TestFilterSignalIDs: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, statID.String(), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterSignalIDs: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestFilterSignalIDs: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != statID {
		t.Fatal("TestFilterSignalIDs: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, ";;"+statID.String()+";;;", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterSignalIDs: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestFilterSignalIDs: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != statID {
		t.Fatal("TestFilterSignalIDs: retrieve Guid value does not match source")
	}

	freqUUID := freqID.String()
	idSet, err = SelectSignalIDSet(dataSet, "'"+freqUUID[1:len(freqUUID)-1]+"'", "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterSignalIDs: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 1 {
		t.Fatal("TestFilterSignalIDs: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if idSet.Keys()[0] != freqID {
		t.Fatal("TestFilterSignalIDs: retrieve Guid value does not match source")
	}

	idSet, err = SelectSignalIDSet(dataSet, fmt.Sprintf("%s;%s;%s", statID.String(), freqID.String(), statID.String()), "ActiveMeasurements", nil, false)

	if err != nil {
		t.Fatal("TestFilterSignalIDs: error executing SelectSignalIDSet: " + err.Error())
	}

	if len(idSet) != 2 {
		t.Fatal("TestFilterSignalIDs: expected 1 result, received: " + strconv.Itoa(len(idSet)))
	}

	if !idSet.Contains(statID) || !idSet.Contains(freqID) {
		t.Fatal("TestFilterSignalIDs: retrieve Guid value does not match source")
	}
}

//func TestFilterSignalIDs(t *testing.T) {
//}
