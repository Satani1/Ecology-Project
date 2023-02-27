package main

import (
	"html/template"
	"log"
	"net/http"
)

//home page
func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ts, err := template.ParseFiles("./ui/html/index.html")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = ts.Execute(w, nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//interactive map page
func mapPage(w http.ResponseWriter, r *http.Request) {

}

//user profile page
func profilePage(w http.ResponseWriter, r *http.Request) {

}
