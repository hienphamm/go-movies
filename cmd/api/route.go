package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (a *application) route() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)

	r.Get("/", Hello)

	return r
}
