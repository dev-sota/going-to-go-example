package user

import "github.com/ispec-inc/going-to-go-example/pkg/view"

type GetUserResponse struct {
	User view.User `json:"user"`
}

type AddUserResponse struct {
	User view.User `json:"user"`
}
