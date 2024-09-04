package main

import (
	"log"
	"net/http"
	"text/template"
)

func getIndex(w http.ResponseWriter, r *http.Request) {
	files := []string{
		"./views/index.html",
	}

	ts, err := template.ParseFiles(files...)

	if err != nil {
        log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = ts.ExecuteTemplate(w, "index", nil)
	if err != nil {
        log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func main() {
	port := ":5173"
	mux := http.NewServeMux()

    fileServer := http.FileServer(http.Dir("./static/"))
    mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))

	mux.HandleFunc("/", getIndex)

	log.Printf("starting server in port %s", port)
	log.Fatal(http.ListenAndServe(port, mux))
}
