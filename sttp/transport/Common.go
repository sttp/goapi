//******************************************************************************************************
//  Common.go - Gbtc
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
//  09/16/2021 - J. Ritchie Carroll
//       Generated original version of source code.
//
//******************************************************************************************************

package transport

import (
	"bytes"
	"compress/gzip"
	"crypto/aes"
	"crypto/cipher"
	"io"
	"net"
	"strings"
)

func decipherAES(key, iv, data []byte) ([]byte, error) {
	var block cipher.Block
	var err error

	if block, err = aes.NewCipher(key); err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	out := make([]byte, len(data))
	mode.CryptBlocks(out, data)

	return out, nil
}

func decompressGZip(data []byte) ([]byte, error) {
	var reader *gzip.Reader
	var err error

	if reader, err = gzip.NewReader(bytes.NewReader(data)); err != nil {
		return nil, err
	}

	defer reader.Close()

	if data, err = io.ReadAll(reader); err != nil {
		return nil, err
	}

	return data, nil
}

func resolveDNSName(addr string) string {
	if strings.Contains(addr, ":") {
		host, _, err := net.SplitHostPort(addr)

		if err == nil {
			addr = host
		}
	}

	hostNames, err := net.LookupAddr(addr)

	if err == nil && len(hostNames) > 0 {
		return hostNames[0] + " (" + addr + ")"
	}

	return addr
}
