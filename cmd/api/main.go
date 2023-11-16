package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
)

const port = 8080

type application struct {
	DNS    string
	Domain string
}

func main() {
	// set application config
	var app application
	flag.StringVar(&app.DNS, "DNS", "host=localhost port=5432 user=hienphamm password=secret db=movies sslmode=disable timezone=UTC", "Postgres connection string")
	flag.Parse()
	//read from command line

	//connect to database
	app.Domain = "example.com"

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.route())
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
