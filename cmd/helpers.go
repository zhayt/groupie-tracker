package main

import (
	"fmt"
	"html/template"
	"net/http"
	"runtime/debug"
)

type ErrorStatus struct {
	Code    int
	Message string
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	tmp, err := template.ParseFiles("./web/templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	err = tmp.Execute(w, ErrorStatus{http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError)})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	tmp, err := template.ParseFiles("./web/templates/error.html")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(status)
	err = tmp.Execute(w, ErrorStatus{status, http.StatusText(status)})
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
