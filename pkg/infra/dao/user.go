package dao

import (
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
	"github.com/ispec-inc/going-to-go-example/pkg/infra/entity"
	"github.com/ispec-inc/going-to-go-example/pkg/transaction"
	"gorm.io/gorm"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) User {
	return User{db}
}

func (repo User) Find(id int64) (model.User, apperror.Error) {
	var user entity.User
	if err := repo.db.First(&user, id).Error; err != nil {
		return model.User{}, newGormError(
			err, "error searching user in database",
		)
	}
	return user.ToModel(), nil
}

func (repo User) Create(mu model.User) apperror.Error {
	f := func(tx *gorm.DB) apperror.Error {
		user := entity.NewUserFromModel(mu)
		if err := tx.Create(&user).Error; err != nil {
			return newGormError(err, "error inserting user in database")
		}

		return nil
	}

	if aerr := transaction.Run(repo.db, f); aerr != nil {
		return aerr
	}

	return nil
}
