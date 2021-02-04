package password

import (
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password *string) apperror.Error {
	res, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	if err != nil {
		return apperror.New(apperror.CodeError, err)
	}
	*password = string(res)
	return nil
}

func Authorize(hashedPassword string, password string) apperror.Error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return apperror.New(apperror.CodeError, err)
	}
	return nil
}
