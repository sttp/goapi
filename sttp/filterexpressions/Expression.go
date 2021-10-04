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

// Expression is the interface that can represent all expression types
type Expression interface {
	// Type gets the type of the expression.
	Type() ExpressionTypeEnum
}

// GetValueExpression gets the expression cast to a ValueExpression.
// An error will be returned if expression is nil or not ExpressionType.Value.
func GetValueExpression(expression Expression) (*ValueExpression, error) {
	if expression == nil {
		return nil, errors.New("cannot get ValueExpression, expression is nil")
	}

	if expression.Type() != ExpressionType.Value {
		return nil, errors.New("expression is not a ValueExpression")
	}

	return expression.(*ValueExpression), nil
}

// GetUnaryExpression gets the expression cast to a UnaryExpression.
// An error will be returned if expression is nil or not ExpressionType.Unary.
func GetUnaryExpression(expression Expression) (*UnaryExpression, error) {
	if expression == nil {
		return nil, errors.New("cannot get UnaryExpression, expression is nil")
	}

	if expression.Type() != ExpressionType.Unary {
		return nil, errors.New("expression is not a UnaryExpression")
	}

	return expression.(*UnaryExpression), nil
}

// GetColumnExpression gets the expression cast to a ColumnExpression.
// An error will be returned if expression is nil or not ExpressionType.Column.
func GetColumnExpression(expression Expression) (*ColumnExpression, error) {
	if expression == nil {
		return nil, errors.New("cannot get ColumnExpression, expression is nil")
	}

	if expression.Type() != ExpressionType.Column {
		return nil, errors.New("expression is not a ColumnExpression")
	}

	return expression.(*ColumnExpression), nil
}

// GetInListExpression gets the expression cast to a InListExpression.
// An error will be returned if expression is nil or not ExpressionType.InList.
func GetInListExpression(expression Expression) (*InListExpression, error) {
	if expression == nil {
		return nil, errors.New("cannot get InListExpression, expression is nil")
	}

	if expression.Type() != ExpressionType.InList {
		return nil, errors.New("expression is not a InListExpression")
	}

	return expression.(*InListExpression), nil
}

// GetFunctionExpression gets the expression cast to a FunctionExpression.
// An error will be returned if expression is nil or not ExpressionType.Function.
func GetFunctionExpression(expression Expression) (*FunctionExpression, error) {
	if expression == nil {
		return nil, errors.New("cannot get FunctionExpression, expression is nil")
	}

	if expression.Type() != ExpressionType.Function {
		return nil, errors.New("expression is not a FunctionExpression")
	}

	return expression.(*FunctionExpression), nil
}

// GetOperatorExpression gets the expression cast to a OperatorExpression.
// An error will be returned if expression is nil or not ExpressionType.Operator.
func GetOperatorExpression(expression Expression) (*OperatorExpression, error) {
	if expression == nil {
		return nil, errors.New("cannot get OperatorExpression, expression is nil")
	}

	if expression.Type() != ExpressionType.Operator {
		return nil, errors.New("expression is not a OperatorExpression")
	}

	return expression.(*OperatorExpression), nil
}
