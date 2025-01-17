package main

import (
	"net/http"

	"jpech.dev/views"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /static/", http.FileServerFS(views.Files))

	mux.HandleFunc("GET /", app.getIndex)
    mux.HandleFunc("GET /about", app.getAbout)
    mux.HandleFunc("GET /projects", app.getProjects)
    mux.HandleFunc("GET /links", app.getLinks)

	return mux
}
