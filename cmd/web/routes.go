package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *Applicaton) Routes() *mux.Router {
	rMux := mux.NewRouter()

	rMux.HandleFunc("/", app.Home).Methods("GET")
	rMux.HandleFunc("/map", app.MapPage)
	rMux.HandleFunc("/profile", app.ProfilePage)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	rMux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return rMux
}
