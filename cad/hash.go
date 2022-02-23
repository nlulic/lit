package cad

import (
	"crypto/sha1"
	"fmt"
)

// Hash will return the sha1 hash of the passed bytes and will also append
// a header containing the passed type and the size **before** hashing
func Hash(in []byte, objectType string) string {
	hasher := sha1.New()
	hasher.Write(withHeader(in, objectType))

	return fmt.Sprintf("%x", hasher.Sum(nil))
}

func withHeader(in []byte, objectType string) []byte {
	header := fmt.Sprintf("%s %d\u0000", objectType, len(in))

	return append([]byte(header), in...)
}
