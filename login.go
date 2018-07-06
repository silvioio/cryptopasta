// Package cryptopasta provides a key derivation function for login systems.
//
// The password hash is generated with scrypt and parameters defined in salt.go.
package cryptopasta

import "golang.org/x/crypto/scrypt"

var (
	scryptSalt []byte
	scryptN    int
	scryptR    int
	scryptP    int
)

// LoginKey returns a 32 byte scrypt derived key.
func LoginKey(password []byte) *[32]byte {
	key := [32]byte{}
	derKey, err := scrypt.Key([]byte(password), scryptSalt, scryptN, scryptR, scryptP, 32)
	if err != nil {
		panic(err)
	}
	copy(key[:], derKey)
	return &key
}
