// Package cryptopasta provides basic encoding and decoding helper for B64, Hex, etc.
package cryptopasta

import (
	"encoding/base64"
	"encoding/hex"
)

// EncodeB64 encodes a slice of bytes into a base64 encoded string.
func EncodeB64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// DecodeB64 decodes a b64 string into a slice of bytes.
func DecodeB64(str string) (data []byte, err error) {
	data, err = base64.StdEncoding.DecodeString(str)
	return
}

// EncodeHex encodes a slice of bytes into a hex encoded string.
func EncodeHex(data []byte) string {
	return hex.EncodeToString(data)
}

// DecodeHex decodes a hex string into a slice of bytes.
func DecodeHex(str string) (data []byte, err error) {
	data, err = hex.DecodeString(str)
	return
}
