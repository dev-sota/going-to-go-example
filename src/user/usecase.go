package user

import (
	"fmt"

	"github.com/ispec-inc/going-to-go-example/pkg/apperror"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/repository"
	"github.com/ispec-inc/going-to-go-example/pkg/password"
	"github.com/ispec-inc/going-to-go-example/pkg/registry"
	"github.com/ispec-inc/going-to-go-example/pkg/value"
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
	_, aerr = u.user.FindByEmail(inp.User.Email)
	if aerr == nil || aerr.Code() != apperror.CodeNotFound {
		aerr = apperror.New(apperror.CodeInvalid, fmt.Errorf("Email address already exists"))
		return
	}

	err := password.Encrypt(&inp.User.Password)
	if err != nil {
		aerr = apperror.New(apperror.CodeError, err)
		return
	}

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

func (u Usecase) Login(inp LoginInput) (out LoginOutput, aerr apperror.Error) {
	user, aerr := u.user.FindByEmail(inp.Email)
	if aerr != nil {
		return
	}

	err := password.Authorize(user.Password, inp.Password)
	if err != nil {
		aerr = apperror.New(apperror.CodeError, err)
		return
	}

	claims := value.NewClaims(user.ID)
	tkn, aerr := value.NewToken(claims)
	if aerr != nil {
		return
	}
	out.Token = tkn

	return out, nil
}
