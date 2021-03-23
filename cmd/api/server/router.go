package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ispec-inc/going-to-go-example/cmd/api/adapter/invitation"
	"github.com/ispec-inc/going-to-go-example/cmd/api/adapter/user"
	"github.com/ispec-inc/going-to-go-example/cmd/api/middleware"
	"github.com/ispec-inc/going-to-go-example/pkg/presenter"
	"github.com/ispec-inc/going-to-go-example/pkg/registry"
)

func NewRouter(repo registry.Repository) http.Handler {
	r := chi.NewRouter()

	authmiddleware := middleware.NewAuth(repo)
	invitationHandler := invitation.NewHandler(repo)
	userHandler := user.NewHandler(repo)

	r = commonMiddleware(r)

	r.Mount("/invitations", invitation.NewRouter(invitationHandler))
	r.Mount("/users", user.NewRouter(userHandler, authmiddleware))
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		presenter.Success(w)
	})

	return r
}
