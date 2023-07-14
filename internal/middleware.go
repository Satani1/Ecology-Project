package internal

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt"
	"log"
	"net/http"
	"time"
)

func (app *Applicaton) RequireAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.InfoLog.Println("Middleware in")

		//get cookie off request
		tokenString, err := r.Cookie("Authorization")
		if err != nil {
			http.Error(w, "Need authorization", http.StatusBadRequest)
			return
		}

		//decode/validate
		token, err := jwt.Parse(tokenString.Value, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("invalid signing method")
			}
			return []byte(app.Secret), nil
		})

		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			//check exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				http.Error(w, "token expired", http.StatusUnauthorized)
				return
			}

			// Find the user with token sub
			sub := fmt.Sprintf("%v", claims["sub"])
			if user, err := app.Repository.GetUserByName(sub); err == nil {
				// Attach to request
				r.Header.Set("user", user.UID)
				r.Header.Set("name", user.Name)

				// Continue
				next.ServeHTTP(w, r)
			} else {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
		} else {
			http.Error(w, "token invalid", http.StatusUnauthorized)
			return
		}
		log.Println("Middleware quit")
	})
}
