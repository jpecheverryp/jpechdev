package main

import "net/http"

func (app *application) getIndex(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "index.html")
}

func (app *application) getProjects(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "index.html")
}

func (app *application) getAbout(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "about.html")
}

func (app *application) getContact(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "contact.html")
}

func (app *application) getGitHub(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://github.com/jpecheverryp", http.StatusSeeOther)
}

func (app *application) getLinkedIn(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "https://www.linkedin.com/in/jpechdev/", http.StatusSeeOther)
}
