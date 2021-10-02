//******************************************************************************************************
//  FunctionExpression.go - Gbtc
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

// FunctionExpression represents a function expression.
type FunctionExpression struct {
	functionType ExpressionFunctionTypeEnum
	arguments    []Expression
}

// NewFunctionExpression creates a new function expression.
func NewFunctionExpression(functionType ExpressionFunctionTypeEnum, arguments []Expression) *FunctionExpression {
	return &FunctionExpression{
		functionType: functionType,
		arguments:    arguments,
	}
}

// Type gets expression type of the FunctionExpression.
func (*FunctionExpression) Type() ExpressionTypeEnum {
	return ExpressionType.Function
}

// FunctionType gets function type of the FunctionExpression.
func (fe *FunctionExpression) FunctionType() ExpressionFunctionTypeEnum {
	return fe.functionType
}

// Arguments gets the expression arguments of the FunctionExpression.
func (fe *FunctionExpression) Arguments() []Expression {
	return fe.arguments
}
