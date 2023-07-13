package internal

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *Applicaton) ServeError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrogLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *Applicaton) ClientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *Applicaton) NotFound(w http.ResponseWriter) {
	app.ClientError(w, http.StatusNotFound)
}
