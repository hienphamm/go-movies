package main

import (
	"errors"
	"fmt"
	"net/http"
	"os"
)

const port = 8080

type application struct {
	Domain string
}

func main() {
	// set application config
	var app application

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
