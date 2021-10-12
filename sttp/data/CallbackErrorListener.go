//******************************************************************************************************
//  CallbackErrorListener.go - Gbtc
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
//  10/07/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package data

import "github.com/antlr/antlr4/runtime/Go/antlr"

// CallbackErrorListener defines a implementation of an ANTLR error listener
// that reports any parsing exceptions to a user defined callback.
type CallbackErrorListener struct {
	*antlr.DefaultErrorListener

	// ParsingExceptionCallback defines a callback for reporting ANTLR parsing exceptions.
	ParsingExceptionCallback func(message string)
}

// NewCallbackErrorListener creates a new NewCallbackErrorListener.
func NewCallbackErrorListener() *CallbackErrorListener {
	return &CallbackErrorListener{
		DefaultErrorListener: antlr.NewDefaultErrorListener(),
	}
}

// SyntaxError is called when ANTLR parser encounters a syntax error.
func (cel *CallbackErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	if cel.ParsingExceptionCallback == nil {
		return
	}

	cel.ParsingExceptionCallback(msg)
}
