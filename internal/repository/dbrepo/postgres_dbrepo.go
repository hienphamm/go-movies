package dbrepo

import (
	"context"
	"github.com/hienphamm/go-movies/internal/model"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *pgx.Conn
}

const dbTimeout = time.Second * 3

func (p *PostgresDBRepo) Connection() *pgx.Conn {
	return p.DB
}

func (p *PostgresDBRepo) AllMovies() ([]*model.Movie, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
	select
		id, title, release_date, runtime,
		mpaa_rating, description, coalesce(image,''),
		created_at, updated_at
	from
		movies 
	order by
		title
	`

	rows, err := p.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var movies []*model.Movie

	for rows.Next() {
		var movie model.Movie
		if err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.ReleaseDate,
			&movie.Runtime,
			&movie.MPAARating,
			&movie.Description,
			&movie.Image,
			&movie.CreatedAt,
			&movie.UpdatedAt,
		); err != nil {
			log.Fatal(err)
		}

		movies = append(movies, &movie)
	}

	return movies, nil
}

func (p *PostgresDBRepo) GetUserByEmail(email string) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `select id, email, first_name, last_name, created_at, updated_at from users where email = $1`
	row, err := p.DB.Query(ctx, query, email)

	if err != nil {
		return nil, err
	}
	defer row.Close()

	var user model.User
	err = row.Scan(
		&user.ID,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)

	return &user, nil
}
