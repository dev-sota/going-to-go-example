package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password string) (string, error) {
	res, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(res), err
}
