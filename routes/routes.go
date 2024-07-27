package routes

import (
	"github.com/go-chi/chi/v5"
)

func APIRoutes() *chi.Mux {

	r := chi.NewRouter()
	r.Mount("/posts", PostRoutes())

	return r
}
