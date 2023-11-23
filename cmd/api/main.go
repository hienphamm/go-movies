package api

import (
	"context"
	"errors"
	"flag"
	"fmt"
	"github.com/hienphamm/go-movies/internal/repository"
	"github.com/hienphamm/go-movies/internal/repository/dbrepo"
	"log"
	"net/http"
	"os"
	"time"
)

const port = 8080

type application struct {
	DNS          string
	Domain       string
	DB           repository.DatabaseRepo
	auth         Auth
	JWTSecret    string
	JWTIssuer    string
	JWTAudience  string
	JWTDomain    string
	CookieDomain string
}

func Run() {
	// set application config
	var app application

	//read from command line
	flag.StringVar(&app.DNS, "DNS", "host=localhost port=5432 user=hienphamm password=secret dbname=movies sslmode=disable timezone=UTC", "Postgres connection string")
	flag.StringVar(&app.JWTSecret, "jwt-secret", "verysecret", "signing secret")
	flag.StringVar(&app.JWTIssuer, "jwt-issuer", "example.com", "signing issuer")
	flag.StringVar(&app.JWTAudience, "jwt-audience", "example.com", "signing audience")
	flag.StringVar(&app.CookieDomain, "cookie-domain", "localhost", "cookie domain")
	flag.StringVar(&app.Domain, "domain", "example.com", "domain")
	flag.Parse()

	//connect to database
	conn, err := app.connectToDB()
	if err != nil {
		log.Fatal("got error", err)
	}

	app.DB = &dbrepo.PostgresDBRepo{DB: conn}
	defer app.DB.Connection().Close(context.Background())

	app.auth = Auth{
		Issuer:        app.JWTIssuer,
		Audience:      app.JWTAudience,
		Secret:        app.JWTSecret,
		TokenExpiry:   time.Minute * 15,
		RefreshExpiry: time.Hour * 24,
		CookieDomain:  "/",
		CookiePath:    "_Host-refresh_token",
		CookieName:    app.CookieDomain,
	}

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), app.route())
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}
}
