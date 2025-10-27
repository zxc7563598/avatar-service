package utils

import (
	"crypto/sha256"
)

func HashToBytes(seed string) []byte {
	h := sha256.Sum256([]byte(seed))
	return h[:]
}

func BinaryHash(b []byte) int64 {
	if len(b) < 8 {
		return 0
	}
	var v int64
	for i := 0; i < 8; i++ {
		v = (v << 8) | int64(b[i])
	}
	return v
}
