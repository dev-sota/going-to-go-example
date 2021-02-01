//go:generate mockgen -package mock -source=user.go -destination=../mock/user.go

package repository

import (
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
)

type User interface {
	Find(id int64) (model.User, apperror.Error)
	FindByUserID(uid int64) (model.User, apperror.Error)
	Create(minv model.User) apperror.Error
}
