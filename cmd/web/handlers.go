package main

import (
	"ecogoly/pkg/models"
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

// home page
func (app *Applicaton) Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.NotFound(w)
		return
	}

	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		app.ServeError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.ServeError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// interactive map page
func (app *Applicaton) MapPage(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	s, err := app.usersDB.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.ServeError(w, err)
		}
		return
	}

	fmt.Fprintf(w, "%v", s)
}

// user profile page
func (app *Applicaton) ProfilePage(w http.ResponseWriter, r *http.Request) {

}

func (app *Applicaton) RegisterUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		if _, err := app.usersDB.Insert(name, surname); err != nil {
			app.ServeError(w, err)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		ts, err := template.ParseFiles("./ui/html/login.html")
		if err != nil {
			app.ServeError(w, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		err = ts.Execute(w, nil)
		if err != nil {
			app.ServeError(w, err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (app *Applicaton) ViewProfile(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/p" {
		app.NotFound(w)
		return
	}
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		app.NotFound(w)
		return
	}

	s, err := app.usersDB.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrNoRecord) {
			app.NotFound(w)
		} else {
			app.ServeError(w, err)
		}
		return
	}

	ts, err := template.ParseFiles("./ui/html/profile.html")
	if err != nil {
		app.ServeError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, s)
	if err != nil {
		app.ServeError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Applicaton) mapPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/map" {
		app.NotFound(w)
		return
	}
	ts, err := template.ParseFiles("./ui/html/mapTest.html")
	if err != nil {
		app.ServeError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		app.ServeError(w, err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (app *Applicaton) SaveMarker(w http.ResponseWriter, r *http.Request) {
	//check method
	if r.Method == "POST" {
		name := r.FormValue("marker-name")
		desc := r.FormValue("marker-description")
		add := r.FormValue("marker-address")
		status := 1
		typ := 1
		if _, err := app.markersDB.Insert(name, desc, add, status, typ); err != nil {
			app.ServeError(w, err)
			return
		}
		http.Redirect(w, r, "/map", http.StatusSeeOther)

	} else {
		//return error404 if method wasn't POST
		app.NotFound(w)
		return
	}

}

func (app *Applicaton) getMarkers(w http.ResponseWriter, r *http.Request) {

}
