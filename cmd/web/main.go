package main

import (
	"database/sql"
	"flag"
	"github.com/vbrenister/green-toys/internal/models"
	"log/slog"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

type config struct {
	port    string
	logging struct {
		requests bool
	}
	db struct {
		dsn string
	}
}

type application struct {
	config config
	logger *slog.Logger
	toys   models.ToyModel
}

func main() {
	var config config

	flag.StringVar(&config.port, "port", ":4000", "Port to run the server on")
	flag.BoolVar(&config.logging.requests, "request-logging", true, "Enable request logging")
	flag.StringVar(&config.db.dsn, "dsn", "postgres://green_toys_user:password@localhost/green_toys_db?sslmode=disable", "PostgreSQL DSN")

	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := openDB(config.db.dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := &application{
		config: config,
		logger: logger,
		toys:   models.NewToyModel(db),
	}

	srv := &http.Server{
		Addr:    config.port,
		Handler: app.routes(),
	}

	logger.Info("Starting server", "port", config.port)

	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
