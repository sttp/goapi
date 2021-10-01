//******************************************************************************************************
//  Expression.go - Gbtc
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

import (
	"errors"
)

// Expression represents the fundamental type for all expression types
type Expression struct {
	expressionType ExpressionTypeEnum
}

func newExpression(expressionType ExpressionTypeEnum) Expression {
	return Expression{expressionType: expressionType}
}

// Type gets the type of the expression, e.g., ExpressionType.Unary or ExpressionType.Operator.
func (e *Expression) Type() ExpressionTypeEnum {
	return e.expressionType
}

// ValueExpression gets the expression cast to a ValueExpression.
func GetValueExpression(expression interface{}) (*ValueExpression, error) {
	if ve, ok := expression.(*ValueExpression); ok {
		return ve, nil
	}

	return nil, errors.New("expression is not a ValueExpression")
}

// UnaryExpression gets the expression cast to a UnaryExpression.
func GetUnaryExpression(expression interface{}) (*UnaryExpression, error) {
	if ue, ok := expression.(*UnaryExpression); ok {
		return ue, nil
	}

	return nil, errors.New("expression is not a UnaryExpression")
}

// ColumnExpression gets the expression cast to a ColumnExpression.
func GetColumnExpression(expression interface{}) (*ColumnExpression, error) {
	if ce, ok := expression.(*ColumnExpression); ok {
		return ce, nil
	}

	return nil, errors.New("expression is not a ColumnExpression")
}

// InListExpression gets the expression cast to a InListExpression.
func GetInListExpression(expression interface{}) (*InListExpression, error) {
	if ine, ok := expression.(*InListExpression); ok {
		return ine, nil
	}

	return nil, errors.New("expression is not a InListExpression")
}

// FunctionExpression gets the expression cast to a FunctionExpression.
func GetFunctionExpression(expression interface{}) (*FunctionExpression, error) {
	if fe, ok := expression.(*FunctionExpression); ok {
		return fe, nil
	}

	return nil, errors.New("expression is not a FunctionExpression")
}

// OperatorExpression gets the expression cast to a OperatorExpression.
func GetOperatorExpression(expression interface{}) (*OperatorExpression, error) {
	if oe, ok := expression.(*OperatorExpression); ok {
		return oe, nil
	}

	return nil, errors.New("expression is not a OperatorExpression")
}
