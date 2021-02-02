package user

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/go-playground/validator"
	"github.com/ispec-inc/going-to-go-example/pkg/domain/model"
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

func (h handler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	inp := user.FindInput{
		ID: int64(id),
	}
	out, aerr := h.usecase.Find(inp)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	usrres := view.NewUser(out.User)
	res := Response{
		User: usrres,
	}
	presenter.Response(w, res)
}

func (h handler) Add(w http.ResponseWriter, r *http.Request) {
	var request addRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}
	validate := validator.New()
	if err := validate.Struct(request); err != nil {
		presenter.BadRequestError(w, err)
		return
	}

	usr := model.User{
		Name: request.Name,
		Age:  request.Age,
	}
	inp := user.AddInput{
		User: usr,
	}
	out, aerr := h.usecase.Add(inp)
	if aerr != nil {
		presenter.ApplicationException(w, aerr)
		return
	}

	usrres := view.NewUser(out.User)
	res := Response{
		User: usrres,
	}
	presenter.Response(w, res)
}
