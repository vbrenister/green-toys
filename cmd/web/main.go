package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"

	"github.com/vbrenister/green-toys/internal/models"
)

type config struct {
	Port                  string
	RequestLoggingEnabled bool
}

type application struct {
	config config
	logger *slog.Logger
	toys   models.ToyModel
}

func main() {
	var config config

	flag.StringVar(&config.Port, "port", ":4000", "Port to run the server on")
	flag.BoolVar(&config.RequestLoggingEnabled, "request-logging", true, "Enable request logging")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: config,
		logger: logger,
		toys:   models.NewToyModel(),
	}

	srv := &http.Server{
		Addr:    config.Port,
		Handler: app.routes(),
	}

	logger.Info("Starting server", "port", config.Port)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
