package dbrepo

import (
	"context"
	"database/sql"
	"github.com/hienphamm/go-movies/cmd/internal/model"
	"log"
	"time"
)

type PostgresDBRepo struct {
	DB *sql.DB
}

const dbTimeout = time.Second * 3

func (p *PostgresDBRepo) Connection() *sql.DB {
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

	rows, err := p.DB.QueryContext(ctx, query)
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
