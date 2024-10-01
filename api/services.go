package api_user

import (
	"crypto/sha256"
	"encoding/hex"
)

func HashPassword(password string) string {
	hasher := sha256.New()
	hasher.Write([]byte(password))
	return hex.EncodeToString(hasher.Sum(nil))
}

func ToString(value *string) string {
	if value == nil {
		return ""
	}
	return *value
}
