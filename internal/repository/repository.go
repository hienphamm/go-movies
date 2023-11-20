package repository

import (
	"database/sql"
	"github.com/hienphamm/go-movies/internal/model"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	AllMovies() ([]*model.Movie, error)
}
