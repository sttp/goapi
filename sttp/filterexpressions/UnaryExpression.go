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

import "errors"

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

func (ue *UnaryExpression) unaryBoolean(value bool) (*ValueExpression, error) {
	switch ue.unaryType {
	case ExpressionUnaryType.Not:
		value = !value
	case ExpressionUnaryType.Plus:
		return nil, errors.New("cannot apply unary \"+\" operator to \"Boolean\"")
	case ExpressionUnaryType.Minus:
		return nil, errors.New("cannot apply unary \"-\" operator to \"Boolean\"")
	default:
		return nil, errors.New("unexpected unary type encountered")
	}

	return newValueExpression(ExpressionValueType.Boolean, value), nil
}

func (ue *UnaryExpression) unaryInt32(value int32) (*ValueExpression, error) {
	switch ue.unaryType {
	case ExpressionUnaryType.Plus:
		value = +value
	case ExpressionUnaryType.Minus:
		value = -value
	case ExpressionUnaryType.Not:
		value = ^value
	default:
		return nil, errors.New("unexpected unary type encountered")
	}

	return newValueExpression(ExpressionValueType.Int32, value), nil
}

func (ue *UnaryExpression) unaryInt64(value int64) (*ValueExpression, error) {
	switch ue.unaryType {
	case ExpressionUnaryType.Plus:
		value = +value
	case ExpressionUnaryType.Minus:
		value = -value
	case ExpressionUnaryType.Not:
		value = ^value
	default:
		return nil, errors.New("unexpected unary type encountered")
	}

	return newValueExpression(ExpressionValueType.Int64, value), nil
}

func (ue *UnaryExpression) unaryDecimal(value float64) (*ValueExpression, error) {
	switch ue.unaryType {
	case ExpressionUnaryType.Plus:
		value = +value
	case ExpressionUnaryType.Minus:
		value = -value
	case ExpressionUnaryType.Not:
		return nil, errors.New("cannot apply unary \"~\" operator to \"Decimal\"")
	default:
		return nil, errors.New("unexpected unary type encountered")
	}

	return newValueExpression(ExpressionValueType.Decimal, value), nil
}

func (ue *UnaryExpression) unaryDouble(value float64) (*ValueExpression, error) {
	switch ue.unaryType {
	case ExpressionUnaryType.Plus:
		value = +value
	case ExpressionUnaryType.Minus:
		value = -value
	case ExpressionUnaryType.Not:
		return nil, errors.New("cannot apply unary \"~\" operator to \"Double\"")
	default:
		return nil, errors.New("unexpected unary type encountered")
	}

	return newValueExpression(ExpressionValueType.Double, value), nil
}
