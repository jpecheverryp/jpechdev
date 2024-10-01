package main

import (
	"html/template"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger        *slog.Logger
	templateCache map[string]*template.Template
}

func main() {
	port := ":5173"

	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	templateCache, err := newTemplateCache()
	//if err != nil {
		//logger.Error(err.Error())
		//os.Exit(1)
	//}

	app := &application{
		logger:        logger,
		templateCache: templateCache,
	}

	logger.Info("starting server", "port", port)
	err = http.ListenAndServe(port, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
