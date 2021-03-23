package user

import (
	"github.com/ispec-inc/going-to-go-example/pkg/view"
)

type Response struct {
	User view.User `json:"user"`
}

type LoginResponse struct {
	Token string `json:"token"`
}
