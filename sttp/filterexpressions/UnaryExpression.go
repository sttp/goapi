//******************************************************************************************************
//  UnaryExpression.go - Gbtc
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

// UnaryExpression represents a unary expression.
type UnaryExpression struct {
	value     Expression
	unaryType ExpressionUnaryTypeEnum
}

// NewUnaryExpression creates a new unary expression.
func NewUnaryExpression(unaryType ExpressionUnaryTypeEnum, value Expression) *UnaryExpression {
	return &UnaryExpression{
		value:     value,
		unaryType: unaryType,
	}
}

// Type gets expression type of the UnaryExpression.
func (*UnaryExpression) Type() ExpressionTypeEnum {
	return ExpressionType.Unary
}

// Value gets the expression value of the UnaryExpression.
func (ue *UnaryExpression) Value() Expression {
	return ue.value
}

// UnaryType gets unary type of the UnaryExpression.
func (ue *UnaryExpression) UnaryType() ExpressionUnaryTypeEnum {
	return ue.unaryType
}
