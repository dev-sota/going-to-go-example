package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"

	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
)

func newGormError(err error, msg string) apperror.Error {
	errstr := err.Error()
	switch errstr {
	case gorm.ErrRecordNotFound.Error():
		return apperror.New(apperror.CodeNotFound, fmt.Errorf("%s: %s", msg, errstr))
	default:
		return apperror.New(apperror.CodeError, fmt.Errorf("%s: %s", msg, errstr))
	}
}
