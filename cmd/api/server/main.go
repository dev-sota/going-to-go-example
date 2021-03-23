package main

import (
	"net/http"

	"github.com/ispec-inc/going-to-go-example/pkg/config"
	"github.com/ispec-inc/going-to-go-example/pkg/registry"
)

func main() {
	config.Init()

	repo, cleanup := registry.NewRepository()
	defer cleanup() //nolint

	r := NewRouter(repo)
	_ = http.ListenAndServe(":9000", r)
}
