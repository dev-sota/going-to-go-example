package user

import (
	"github.com/ispec-inc/going-to-go-example/pkg/value"
	"github.com/ispec-inc/going-to-go-example/pkg/view"
)

type Response struct {
	User view.User `json:"user"`
}

type LoginResponse struct {
	Token value.Token `json:"token"`
}
