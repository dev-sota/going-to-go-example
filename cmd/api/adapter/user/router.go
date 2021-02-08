package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/going-to-go-example/cmd/api/middleware"
	"github.com/ispec-inc/going-to-go-example/pkg/presenter"
)

func NewRouter(handler handler, middleware middleware.Auth) http.Handler {
	r := chi.NewRouter()
	r.Get("/{id}", handler.Get)
	r.Post("/", handler.Add)
	r.Post("/login", handler.Login)

	r.Group(func(r chi.Router) {
		r.Use(middleware.VerifyToken)
		r.Get("/restrict", func(w http.ResponseWriter, r *http.Request) {
			presenter.Success(w)
		})
	})

	return r
}
