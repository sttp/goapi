//******************************************************************************************************
//  InListExpression.go - Gbtc
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

// InListExpression represents an in-list expression.
type InListExpression struct {
	value         Expression
	arguments     []Expression
	hasNotkeyWord bool
	exactMatch    bool
}

// NewInListExpression creates a new in-list expression.
func NewInListExpression(value Expression, arguments []Expression, hasNotkeyWord, exactMatch bool) *InListExpression {
	return &InListExpression{
		value:         value,
		arguments:     arguments,
		hasNotkeyWord: hasNotkeyWord,
		exactMatch:    exactMatch,
	}
}

// Type gets expression type of the InListExpression.
func (*InListExpression) Type() ExpressionTypeEnum {
	return ExpressionType.InList
}

// Value gets the expression value of the InListExpression.
func (ile *InListExpression) Value() Expression {
	return ile.value
}

// Arguments gets the expression arguments of the InListExpression.
func (ile *InListExpression) Arguments() []Expression {
	return ile.arguments
}

// HasNotKeyword gets a flag that determines if the InListExpression has the "NOT" keyword.
func (ile *InListExpression) HasNotKeyword() bool {
	return ile.hasNotkeyWord
}

// ExtactMatch gets a flags that determines if the InListExpression has the "BINARY" or "===" keyword.
func (ile *InListExpression) ExtactMatch() bool {
	return ile.exactMatch
}
