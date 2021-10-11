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
	"testing"
)

func TestExpressionTreeGeneration(t *testing.T) {
	result, err := Evaluate("{b4a26a66-a073-44a0-b03b-55d97badef74}", true)

	if err != nil {
		t.Fatal("TestExpressionTreeGeneration: failed to parse expression: " + err.Error())
	}

	if result == nil {
		t.Fatal("TestExpressionTreeGeneration: received no result")
	}

	if result.ValueType() != ExpressionValueType.Guid {
		t.Fatal("TestExpressionTreeGeneration: expected Guid value, received: " + result.ValueType().String())
	}
}
