package password

import (
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password *string) (err error) {
	res, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	if err != nil {
		return
	}
	*password = string(res)
	return
}

func Authorize(hashedPassword string, password string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return
}
