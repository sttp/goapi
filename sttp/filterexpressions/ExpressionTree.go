//******************************************************************************************************
//  ExpressionTree.go - Gbtc
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

package filterexpressions

import "github.com/sttp/goapi/sttp/data"

// ExpressionTree represents an evaluated tree of expressions
type ExpressionTree struct {
	currentRow *data.DataRow
	table      *data.DataTable
}

func evaluate(expression *Expression, targetValueType ExpressionValueTypeEnum) *ValueExpression {
	return nil
}

func evaluateUnary(expression *Expression) *ValueExpression {
	return nil
}

func evaluateColumn(expression *Expression) *ValueExpression {
	return nil
}

func evaluateInList(expression *Expression) *ValueExpression {
	return nil
}

func evaluateFunction(expression *Expression) *ValueExpression {
	return nil
}

func evaluateOperator(expression *Expression) *ValueExpression {
	return nil
}