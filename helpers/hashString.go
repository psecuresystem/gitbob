package helpers

import (
	"crypto/sha1"
	"encoding/hex"
)

func HashString(val string) string {
	hasher := sha1.New()
	hasher.Write([]byte(val))
	return hex.EncodeToString(hasher.Sum(nil))
}