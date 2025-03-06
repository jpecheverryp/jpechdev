package main

import (
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	port := ":5173"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	app := &application{
		logger: logger,
	}

	logger.Info("starting server", "port", port)
	err := http.ListenAndServe(port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
