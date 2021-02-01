package user

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/going-to-go-example/pkg/presenter"
	"github.com/ispec-inc/going-to-go-example/pkg/registry"
	"github.com/ispec-inc/going-to-go-example/pkg/view"
	"github.com/ispec-inc/going-to-go-example/src/user"
)

type handler struct {
	usecase user.Usecase
}

func NewHandler(repo registry.Repository) handler {
	usecase := user.NewUsecase(repo)
	return handler{usecase}
}

func (h handler) GetUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	inp := user.FindUserInput{
		ID: int64(id),
	}
	out, aerr := h.usecase.FindUser(inp)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	usrres := view.NewUser(out.User)
	res := GetUserResponse{
		User: usrres,
	}
	presenter.Response(w, res)
}
