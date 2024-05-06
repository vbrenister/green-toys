package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type config struct {
	Port string
}

type application struct {
	config config
	logger *slog.Logger
}

func main() {

	var config config

	flag.StringVar(&config.Port, "port", ":4000", "Port to run the server on")
	flag.Parse()

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		config: config,
		logger: logger,
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	logger.Info("Starting server", "port", config.Port)

	err := srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}
