//******************************************************************************************************
//  OperatorExpression.go - Gbtc
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

// OperatorExpression represents an operator expression.
type OperatorExpression struct {
	operatorType ExpressionOperatorTypeEnum
	leftValue    Expression
	rightValue   Expression
}

// NewOperatorExpression creates a new operator expression.
func NewOperatorExpression(operatorType ExpressionOperatorTypeEnum, leftValue, rightValue Expression) *OperatorExpression {
	return &OperatorExpression{
		operatorType: operatorType,
		leftValue:    leftValue,
		rightValue:   rightValue,
	}
}

// Type gets expression type of the OperatorExpression.
func (*OperatorExpression) Type() ExpressionTypeEnum {
	return ExpressionType.Operator
}

// OperatorType gets operator type of the OperatorExpression.
func (oe *OperatorExpression) OperatorType() ExpressionOperatorTypeEnum {
	return oe.operatorType
}

// LeftValue gets the left value expression of the OperatorExpression.
func (oe *OperatorExpression) LeftValue() Expression {
	return oe.leftValue
}

// RightValue gets the right value expression of the OperatorExpression.
func (oe *OperatorExpression) RightValue() Expression {
	return oe.rightValue
}
