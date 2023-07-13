package internal

import (
	"github.com/google/uuid"
	"html/template"
	"net/http"
)

func (app *Applicaton) SingUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		surname := r.FormValue("surname")
		password := r.FormValue("password")

		id, err := uuid.NewUUID()
		if err != nil {
			app.ServeError(w, err)
			return
		}
		app.InfoLog.Println(id, id.String())
		app.InfoLog.Println(name, surname, password)

		//id, err := app.UsersDB.Insert(name, surname, email)
		//if err != nil {
		//	app.ServeError(w, err)
		//	return
		//}
		//strID := strconv.Itoa(id)
		//http.Redirect(w, r, "/profile?id="+strID, http.StatusSeeOther)
	} else {
		ts, err := template.ParseFiles("./public/html/singup.html")
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

func (app *Applicaton) SingIn(w http.ResponseWriter, r *http.Request) {

}
