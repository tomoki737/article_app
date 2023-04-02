package utils

import (
	"crypto/sha512"
	"encoding/hex"
)

func HashString(str string) string {
	b := []byte(str)
	sha512 := sha512.Sum512(b)
	hashed := hex.EncodeToString(sha512[:])
	return hashed
}
