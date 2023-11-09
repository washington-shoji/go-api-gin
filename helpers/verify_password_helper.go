package helpers

import (
	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(encryptedPw string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(encryptedPw), []byte(password))
}
