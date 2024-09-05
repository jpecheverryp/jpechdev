package main

import "net/http"

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "index.html")
}
