package main

import (
	"fmt"
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.infoLog.Println(fmt.Sprintf("%s - %s %s %s", r.RemoteAddr, r.Proto, r.Method, r.URL.String()))

		next.ServeHTTP(w, r)
	})
}

func (app *application) onlyGetMethod(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			app.clientError(w, http.StatusMethodNotAllowed)
			return
		}

		next.ServeHTTP(w, r)
	})
}
