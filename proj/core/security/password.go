package security

import (
	"crypto/sha256"
)

func EncodePassword(password string) string {
	h := sha256.New()
	h.Write([]byte(password))

	return string(h.Sum(nil))
}

func IsPasswordValid(password string, hash string) bool {
	return EncodePassword(password) == hash
}
