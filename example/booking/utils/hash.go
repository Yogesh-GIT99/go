package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(Password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(Password), 16)

	return string(bytes), err
}
