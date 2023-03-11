package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func (app *Applicaton) Routes() *mux.Router {
	rMux := mux.NewRouter()

	//pages
	rMux.HandleFunc("/", app.Home).Methods("GET")  //home page
	rMux.HandleFunc("/register", app.RegisterUser) //register page
	rMux.HandleFunc("/profile", app.ViewProfile)   //profile page
	rMux.HandleFunc("/map", app.mapPage)           //map page

	//work with markers on the map
	rMux.HandleFunc("/m", app.getMarkers).Methods("GET")
	rMux.HandleFunc("/towork", app.updateMarkerToWork).Methods("POST")
	rMux.HandleFunc("/savemarker", app.SaveMark)
	rMux.HandleFunc("/photo", app.photoPathToHTML)
	rMux.HandleFunc("/toreport", app.closeMarker)

	fileServer := http.FileServer(http.Dir("./public"))

	rMux.PathPrefix("/public/").Handler(http.StripPrefix("/public", fileServer))

	return rMux
}
