package cryptopasta

import (
	"bytes"
	"fmt"
	"testing"
)

func TestLoginKey(t *testing.T) {
	scryptTests := []struct {
		password []byte
		hash     []byte
	}{
		{
			password: []byte("password"),
			hash:     []byte("mfVw9ru1gA0JL9Xo5Mw7nlCUIhRuEWwLVvHUKlqqI58="),
		},
	}

	scryptSalt = []byte{0, 1, 2, 3, 4, 5, 6, 7}

	for _, tt := range scryptTests {
		login := LoginKey(tt.password)
		fmt.Println(EncodeB64(login[:]))
		if bytes.Compare([]byte(EncodeB64(login[:])), tt.hash) != 0 {
			t.Errorf("password hash does not match")
		}
	}
}

func BenchmarkScrypt(b *testing.B) {
	scryptSalt = []byte{0, 1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < b.N; i++ {
		_ = LoginKey([]byte("thisisareallybadpassword"))

	}
}
