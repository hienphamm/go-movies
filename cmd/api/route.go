package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (app *application) route() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(app.enableCORS)

	r.Get("/", Hello)
	r.Get("/movies", app.AllMovies)

	return r
}
