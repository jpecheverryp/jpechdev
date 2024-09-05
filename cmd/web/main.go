package main

import (
	"log"
	"net/http"
	"html/template"
)

type application struct {
	templateCache map[string]*template.Template
}

func main() {
	port := ":5173"

	templateCache, err := newTemplateCache()
	if err != nil {
		log.Fatal(err)
	}

	app := &application{
		templateCache: templateCache,
	}

	log.Printf("starting server in port %s", port)
	log.Fatal(http.ListenAndServe(port, app.routes()))
}
