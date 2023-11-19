package main

import (
	"errors"
	"flag"
	"fmt"
	"github.com/hienphamm/go-movies/cmd/internal/repository"
	"github.com/hienphamm/go-movies/cmd/internal/repository/dbrepo"
	"log"
	"net/http"
	"os"
)

const port = 8080

type application struct {
	DNS    string
	Domain string
	DB     repository.DatabaseRepo
}

func main() {
	// set application config
	var app application

	//read from command line
	flag.StringVar(&app.DNS, "DNS", "host=localhost port=5432 user=hienphamm password=secret database=movies sslmode=disable timezone=UTC", "Postgres connection string")
	flag.Parse()

	//connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal(err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close()

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.route())
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
