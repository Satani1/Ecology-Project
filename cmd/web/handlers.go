package main

import (
	"html/template"
	"net/http"
)

// home page
func (app *Applicaton) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		app.serveError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.serveError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// interactive map page
func (app *Applicaton) mapPage(w http.ResponseWriter, r *http.Request) {

}

// user profile page
func (app *Applicaton) profilePage(w http.ResponseWriter, r *http.Request) {

}
