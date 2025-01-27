package main

import "net/http"

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "index.html")
}

func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {
    app.render(w, http.StatusOK, "index.html")
}

func (app *application) getAbout(w http.ResponseWriter, r *http.Request) {
    app.render(w, http.StatusOK, "index.html")
}

func (app *application) getContact(w http.ResponseWriter, r *http.Request) {
    app.render(w, http.StatusOK, "contact.html")
}
