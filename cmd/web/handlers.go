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

	ts, err := template.ParseFiles("./public/html/index.html")
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
		ts, err := template.ParseFiles("./public/html/login.html")
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

	ts, err := template.ParseFiles("./public/html/profile.html")
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
	ts, err := template.ParseFiles("./public/html/mapTest.html")
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

func (app *Applicaton) SaveMark(w http.ResponseWriter, r *http.Request) {
	//check method
	if r.Method == "POST" {
		name := r.FormValue("marker-name")
		desc := r.FormValue("marker-description")
		add := r.FormValue("marker-address")

		//get handler for filename, size and headers
		file, handler, err := r.FormFile("marker-photo")
		if err != nil {
			app.ServeError(w, err)
			return
		}
		defer file.Close()

		//filepath
		var dstPath = "public/photoDB/" + handler.Filename
		//create a file
		dst, err := os.Create(dstPath)
		if err != nil {
			app.ServeError(w, err)
			return
		}
		defer dst.Close()

		//saving a copy of file
		if _, err := io.Copy(dst, file); err != nil {
			app.ServeError(w, err)
			return
		}
		//insert text data and path to photo in the db markerks
		_, err = app.markersDB.Insert(name, desc, add, dstPath)
		if err != nil {
			app.ServeError(w, err)
			return
		}

		//redirect back to the /map page
		http.Redirect(w, r, "/map", http.StatusSeeOther)
	} else {
		app.ServeError(w, errors.New("Error with http method"))
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
func (app *Applicaton) updateMarkerToWork(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			app.ServeError(w, err)
			return
		}
		err = app.markersDB.UpdateMarkerToWork(id)
		if err != nil {
			app.ServeError(w, err)
			return
		}
		http.Redirect(w, r, "/map", http.StatusSeeOther)
	} else {
		app.NotFound(w)
		return
	}

}

func (app *Applicaton) closeMarker(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id, err := strconv.Atoi(r.URL.Query().Get("id"))
		if err != nil {
			app.ServeError(w, err)
			return
		}

		err = app.markersDB.Delete(id)
		if err != nil {
			app.ServeError(w, err)
			return
		}

		http.Redirect(w, r, "/map", http.StatusSeeOther)
	} else {
		app.NotFound(w)
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
		var dstPath = "public/html/photoDB/" + handler.Filename
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

func (app *Applicaton) photoPathToHTML(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, `<html><body><h1>ХУЙ</h1></body></html>`)
}
