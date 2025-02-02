package main

import (
	"net/http"

	"github.com/justinas/alice"
	"jpech.dev/views"
)

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", http.FileServerFS(views.Files))

	mux.HandleFunc("GET /", app.getIndex)
	mux.HandleFunc("GET /about", app.getAbout)
	mux.HandleFunc("GET /projects", app.getProjects)
	mux.HandleFunc("GET /contact", app.getContact)

	standard := alice.New(app.logRequest)

	return standard.Then(mux)
}
