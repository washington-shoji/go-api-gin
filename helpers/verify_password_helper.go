package helpers

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func VerifyPassword(encryptedPw string, password string) error {
	fmt.Println("encryptedPw", encryptedPw)
	fmt.Println("password", password)
	return bcrypt.CompareHashAndPassword([]byte(encryptedPw), []byte(password))
}
