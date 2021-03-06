package dao

import (
	"fmt"

	"gorm.io/gorm"

	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
)

func newGormError(err error, msg string) apperror.Error {
	switch err {
	case gorm.ErrRecordNotFound:
		return apperror.New(apperror.CodeNotFound, fmt.Errorf("%s: %s", msg, err.Error()))
	default:
		return apperror.New(apperror.CodeError, fmt.Errorf("%s: %s", msg, err.Error()))
	}
}
