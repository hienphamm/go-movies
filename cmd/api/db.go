package api

import (
	"context"
	_ "github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "postgres://hienphamm:secret@localhost:5432/movies")
	if err != nil {
		return nil, err
	}
	err = conn.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (app *application) connectToDB() (*pgx.Conn, error) {
	connection, err := openDB(app.DNS)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
