// Package cryptopasta provides a key derivation function for login systems.
//
// The password hash is generated with scrypt and parameters defined in salt.go.
package cryptopasta

import "golang.org/x/crypto/scrypt"

// LoginKey returns a 32 byte scrypt derived key.
func LoginKey(password []byte) *[32]byte {
	key := [32]byte{}
	derKey, err := scrypt.Key([]byte(password), []byte{0xd0, 0xdf, 0x83, 0xc0, 0x5d, 0x53, 0x87, 0xb9}, 1<<15, 8, 1, 32)
	if err != nil {
		panic(err)
	}
	copy(key[:], derKey)
	return &key
}
