package cryptopasta

import (
	"strings"
	"testing"
)

func TestEncodeB64(t *testing.T) {
	encTests := []struct {
		testStr    string
		encodedStr string
	}{
		{
			testStr:    "The quick brown fox jumps over the lazy dog",
			encodedStr: "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw==",
		},
	}

	for _, tt := range encTests {
		actualStr := EncodeB64([]byte(tt.testStr))
		if strings.Compare(actualStr, tt.encodedStr) != 0 {
			t.Errorf("encoded string does not match the expected string")
		}
	}
}

func TestDecodeB64(t *testing.T) {
	decTests := []struct {
		testStr    string
		decodedStr string
	}{
		{
			testStr:    "VGhlIHF1aWNrIGJyb3duIGZveCBqdW1wcyBvdmVyIHRoZSBsYXp5IGRvZw==",
			decodedStr: "The quick brown fox jumps over the lazy dog",
		},
	}
	for _, tt := range decTests {
		actual, err := DecodeB64(tt.testStr)
		if err != nil {
			t.Fatal(err)
		}
		actualStr := string(actual)
		if strings.Compare(actualStr, tt.decodedStr) != 0 {
			t.Errorf("encoded string does not match the expected string")
		}
	}
}

func TestEncodeHex(t *testing.T) {
	encTests := []struct {
		testStr    string
		encodedStr string
	}{
		{
			testStr:    "The quick brown fox jumps over the lazy dog",
			encodedStr: "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
		},
	}

	for _, tt := range encTests {
		actualStr := EncodeHex([]byte(tt.testStr))
		if strings.Compare(actualStr, tt.encodedStr) != 0 {
			t.Errorf("encoded string does not match the expected string")
		}
	}
}

func TestDecodeHex(t *testing.T) {
	decTests := []struct {
		testStr    string
		decodedStr string
	}{
		{
			testStr:    "54686520717569636b2062726f776e20666f78206a756d7073206f76657220746865206c617a7920646f67",
			decodedStr: "The quick brown fox jumps over the lazy dog",
		},
	}
	for _, tt := range decTests {
		actual, err := DecodeHex(tt.testStr)
		if err != nil {
			t.Fatal(err)
		}
		actualStr := string(actual)
		if strings.Compare(actualStr, tt.decodedStr) != 0 {
			t.Errorf("encoded string does not match the expected string")
		}
	}
}
