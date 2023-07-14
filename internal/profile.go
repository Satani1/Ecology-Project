package internal

import (
	"html/template"
	"net/http"
	"time"
)

func (app *Applicaton) ViewProfile(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("name")

	//get user from DB
	user, err := app.Repository.GetUserByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html")
	//find template
	ts, err := template.ParseFiles("./public/html/profile.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//exec template
	if err := ts.Execute(w, user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (app *Applicaton) QuitProfile(w http.ResponseWriter, r *http.Request) {
	//delete cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
	})

	w.WriteHeader(http.StatusOK)
}
