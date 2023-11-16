package main

import (
	"encoding/json"
	"fmt"
	"log"
	"movies/cmd/internal/model"
	"net/http"
	"time"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world !")
}

func (app *application) AllMovies(w http.ResponseWriter, r *http.Request) {
	var movies []model.Movie

	rd, _ := time.Parse("2006-01-02", "1986-03-07")

	highlander := model.Movie{
		ID:          0,
		Title:       "Highlander",
		ReleaseDate: rd,
		MPAARating:  "R",
		Runtime:     116,
		Description: "A very nice movie",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	movies = append(movies, highlander)

	out, err := json.Marshal(highlander)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "Application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
