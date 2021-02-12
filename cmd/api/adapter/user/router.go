package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/going-to-go-example/cmd/api/middleware"
)

func NewRouter(handler handler, middleware middleware.Auth) http.Handler {
	r := chi.NewRouter()
	r.Post("/", handler.Add)
	r.Post("/login", handler.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.VerifyToken)
		r.Get("/{id}", handler.Get)
	})

	return r
}
