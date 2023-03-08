package main

import (
	"ecogoly/pkg/models"
	"encoding/json"
	"errors"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
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
		if _, err := app.markersDB.Insert(name, desc, add); err != nil {
			app.ServeError(w, err)
			return
		}
		fmt.Fprintf(w, "name: %v\ndesc: %v\nadd: %v\n", name, desc, add)
		http.Redirect(w, r, "/map", http.StatusSeeOther)

	} else {
		//return error404 if method wasn't POST
		app.NotFound(w)
		return
	}

}

func (app *Applicaton) getMarkers(w http.ResponseWriter, r *http.Request) {
	markers, err := app.markersDB.GetAll()
	if err != nil {
		app.ServeError(w, err)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(markers)
}

// testing markersDB
func (app *Applicaton) testDB(w http.ResponseWriter, r *http.Request) {
	name := "Test"
	desc := "Тест описания"
	add := "Москва, варшавское шоссе, 10"
	if _, err := app.markersDB.Insert(name, desc, add); err != nil {
		app.ServeError(w, err)
		return
	}

}

const MAX_UPLOAD_SIZE = 5 << 20

func (app *Applicaton) UploadPhoto(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//max upload size
		r.ParseMultipartForm(MAX_UPLOAD_SIZE)

		//get handler for filename, size and headers
		file, handler, err := r.FormFile("marker-photo")
		if err != nil {
			app.ServeError(w, err)
			return
		}
		defer file.Close()

		fmt.Printf("Uploaded File: %+v\n", handler.Filename)
		fmt.Printf("File Size: %+v\n", handler.Size)
		fmt.Printf("MIME Header: %+v\n", handler.Header)
		//filepath
		var dstPath = "ui/html/photoDB" + handler.Filename
		//create a file
		dst, err := os.Create(dstPath)
		if err != nil {
			app.ServeError(w, err)
			return
		}
		defer dst.Close()

		//saving a copy of file
		if _, err = io.Copy(dst, file); err != nil {
			app.ServeError(w, err)
			return
		}

		fmt.Fprintf(w, "Successfully Uploaded File\n")
	} else {
		app.ServeError(w, errors.New("Error with http method"))
		return
	}
}
