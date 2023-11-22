package repository

import (
	"github.com/hienphamm/go-movies/internal/model"
	"github.com/jackc/pgx/v4"
)

type DatabaseRepo interface {
	Connection() *pgx.Conn
	AllMovies() ([]*model.Movie, error)
	GetUserByEmail(email string) (*model.User, error)
}
