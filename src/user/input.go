package user

import "github.com/ispec-inc/going-to-go-example/pkg/domain/model"

type FindInput struct {
	ID int64
}

type AddInput struct {
	User model.User
}
