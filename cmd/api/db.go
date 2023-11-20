package api

import (
	"database/sql"
	_ "github.com/jackc/pgconn"
	_ "github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func openDB(dsn string) (*sql.DB, error) {
	conn, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = conn.Ping()
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (app *application) connectToDB() (*sql.DB, error) {
	connection, err := openDB(app.DNS)
	if err != nil {
		return nil, err
	}
	return connection, nil
}
