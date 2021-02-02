package user

import (
	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/repository"
	"github.com/ispec-inc/going-to-go-example/pkg/registry"
)

type Usecase struct {
	user repository.User
}

func NewUsecase(repo registry.Repository) Usecase {
	return Usecase{
		user: repo.NewUser(),
	}
}

func (u Usecase) Find(inp FindInput) (out FindOutput, aerr apperror.Error) {
	user, aerr := u.user.Find(inp.ID)
	if aerr != nil {
		return
	}
	out.User = user
	return out, nil
}

func (u Usecase) Add(inp AddInput) (out AddOutput, aerr apperror.Error) {
	aerr = u.user.Create(&inp.User)
	if aerr != nil {
		return
	}

	user, aerr := u.user.Find(inp.User.ID)
	if aerr != nil {
		return
	}
	out.User = user

	return out, nil
}
