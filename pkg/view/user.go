package view

import (
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
)

type User struct {
	ID    int64  `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

func NewUser(m model.User) User {
	return User{
		ID:    m.ID,
		Email: m.Email,
		Name:  m.Name,
		Age:   m.Age,
	}
}
