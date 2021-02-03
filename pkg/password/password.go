package password

import (
	"fmt"

	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"golang.org/x/crypto/bcrypt"
)

func Encrypt(password *string) apperror.Error {
	res, err := bcrypt.GenerateFromPassword([]byte(*password), 10)
	if err != nil {
		return apperror.New(apperror.CodeError, fmt.Errorf("error password: %v", err))
	}
	*password = string(res)
	return nil
}
