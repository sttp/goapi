//******************************************************************************************************
//  BufferBlock.go - Gbtc
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
//  09/15/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package transport

import "github.com/sttp/goapi/sttp/guid"

// BufferBlock defines an atomic unit of data, i.e., a binary buffer, for transport in STTP.
type BufferBlock struct {
	// Measurement's globally unique identifier.
	SignalID guid.Guid

	// Buffer is an atomic unit of data, i.e., a binary buffer. This buffer typically
	// represents a partial image of a larger whole.
	Buffer []byte
}
