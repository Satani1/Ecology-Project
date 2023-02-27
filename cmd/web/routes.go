package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *Applicaton) routes() *mux.Router {
	rMux := mux.NewRouter()

	rMux.HandleFunc("/", app.home).Methods("GET")
	rMux.HandleFunc("/map", app.mapPage)
	rMux.HandleFunc("/profile", app.profilePage)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	rMux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return rMux
}
