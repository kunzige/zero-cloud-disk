package utils

import (
	"crypto/sha1"
	"encoding/hex"
)

func Encrypt(message string) string {

	h := sha1.New()
	h.Write([]byte(message))
	sha1_hash := hex.EncodeToString(h.Sum(nil))

	return sha1_hash
}
