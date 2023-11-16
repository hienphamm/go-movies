package model

import "time"

type Movie struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	ReleaseDate time.Time `json:"release_date"`
	MPAARating  string    `json:"mpaa_rating"`
	Runtime     int       `json:"runtime"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"_"`
	UpdatedAt   time.Time `json:"_"`
}
