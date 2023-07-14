package internal

import (
	"ecogoly/pkg/repository/postgres"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Applicaton struct {
	ErrogLog   *log.Logger
	InfoLog    *log.Logger
	Repository *postgres.PostgresRepository
	Secret     string
}

func (app *Applicaton) Routes() *mux.Router {
	rMux := mux.NewRouter()

	//pages
	rMux.HandleFunc("/", app.Home).Methods("GET") //home page
	rMux.HandleFunc("/register", app.SingUp)      //register page
	rMux.HandleFunc("/login", app.SingIn)
	rMux.Handle("/profile", app.RequireAuth(http.HandlerFunc(app.ViewProfile))) //profile page
	rMux.Handle("/quit", app.RequireAuth(http.HandlerFunc(app.QuitProfile)))
	rMux.HandleFunc("/map", app.mapPage) //map page

	//work with markers on the map
	//rMux.HandleFunc("/m", app.getMarkers).Methods("GET")
	//rMux.HandleFunc("/towork", app.updateMarkerToWork).Methods("POST")
	//rMux.HandleFunc("/savemarker", app.SaveMark)
	//rMux.HandleFunc("/photo", app.photoPathToHTML)
	//rMux.HandleFunc("/toreport", app.closeMarker)

	fileServer := http.FileServer(http.Dir("./public"))

	rMux.PathPrefix("/public/").Handler(http.StripPrefix("/public", fileServer))

	return rMux
}
