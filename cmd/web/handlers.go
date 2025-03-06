package main

import (
	"net/http"

	"jpech.dev/views/page"
)

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.renderTempl(w, r, http.StatusOK, page.Index())
}

func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {
	app.renderTempl(w, r, http.StatusOK, page.Index())
}

func (app *application) getAbout(w http.ResponseWriter, r *http.Request) {
	app.renderTempl(w, r, http.StatusOK, page.About())
}

func (app *application) getContact(w http.ResponseWriter, r *http.Request) {
	app.renderTempl(w, r, http.StatusOK, page.Contact())
}

func (app *application) getGitHub(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/jpecheverryp", http.StatusSeeOther)
}

func (app *application) getLinkedIn(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.linkedin.com/in/jpechdev/", http.StatusSeeOther)
}
