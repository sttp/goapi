package OperationalEncoding

type OperationalEncoding uint32

const (
	// UTF-16, little endian (deprecated)
	UTF16LE OperationalEncoding = 0x00000000
	// UTF-16, big endian (deprecated)
	UTF16BE OperationalEncoding = 0x00000100
	// UTF-8
	UTF8 OperationalEncoding = 0x00000200
)
