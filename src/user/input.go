package user

import "github.com/ispec-inc/going-to-go-example/pkg/domain/model"

type FindUserInput struct {
	ID int64
}

type AddUserInput struct {
	User model.User
}
