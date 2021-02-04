package user

import (
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
)

type FindOutput struct {
	User model.User
}

type AddOutput struct {
	User model.User
}

type LoginOutput struct {
	Token string
}
