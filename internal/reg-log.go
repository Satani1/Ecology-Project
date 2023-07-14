package internal

import (
	"ecogoly/pkg/models"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func (app *Applicaton) SingUp(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")

		id, err := uuid.NewUUID()
		if err != nil {
			app.ServeError(w, err)
			return
		}
		app.InfoLog.Println(id, id.String())
		app.InfoLog.Println(name, password)

		//hash password
		hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
		if err != nil {
			app.ErrogLog.Fatalln(err)
		}

		//create user
		user := models.User{
			UID:      id.String(),
			Name:     name,
			Password: string(hash),
		}

		//insert into bd
		if err := app.Repository.InsertUser(user); err != nil {
			app.ErrogLog.Fatalln(err)
		}

		w.WriteHeader(http.StatusCreated)

		http.Redirect(w, r, "/", http.StatusSeeOther)
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
	if r.Method == "POST" {
		name := r.FormValue("name")
		password := r.FormValue("password")

		//get user by name from DB
		user, err := app.Repository.GetUserByName(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}

		//compare pass
		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
			http.Error(w, "Invalid password", http.StatusBadRequest)
			return
		}

		//generate token
		if err := app.generateToken(w, user.UID, user.Name); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)

	} else {
		ts, err := template.ParseFiles("./public/html/singin.html")
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

func (app *Applicaton) generateToken(w http.ResponseWriter, id, name string) error {
	log.Println("in")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  id,
		"sub": name,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return err
	}
	log.Println(token)
	//send token to cookie

	http.SetCookie(w, &http.Cookie{
		Name:     "Authorization",
		Value:    tokenString,
		Path:     "",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	})
	return nil
}
