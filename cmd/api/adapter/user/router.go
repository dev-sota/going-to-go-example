package user

import (
	"net/http"

	"github.com/go-chi/chi"
)

func NewRouter(handler handler) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", handler.GetUser)
	r.Post("/", handler.AddUser)
	return r
}
