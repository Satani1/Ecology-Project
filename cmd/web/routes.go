package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *Applicaton) Routes() *mux.Router {
	rMux := mux.NewRouter()

	rMux.HandleFunc("/", app.Home).Methods("GET")
	rMux.HandleFunc("/profile", app.ProfilePage)
	rMux.HandleFunc("/register", app.RegisterUser)
	rMux.HandleFunc("/p", app.ViewProfile)

	rMux.HandleFunc("/map", app.mapPage)
	rMux.HandleFunc("/testDB", app.testDB)
	rMux.HandleFunc("/m", app.getMarkers).Methods("GET")
	rMux.HandleFunc("/savemarker", app.SaveMarker)
	rMux.HandleFunc("/photo", app.UploadPhoto)

	fileServer := http.FileServer(http.Dir("./ui/static/"))
	rMux.Handle("/static/", http.StripPrefix("/static", fileServer))

	return rMux
}
