package main

import (
	"net/http"

	"github.com/a-h/templ"
)

func (app *application) renderTempl(w http.ResponseWriter, r *http.Request, status int, component templ.Component) {
	w.WriteHeader(status)
	err := templ.Component.Render(component, r.Context(), w)
	if err != nil {
		app.logger.Error(err.Error())
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}
}
