package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", app.onlyGetMethod(http.HandlerFunc(app.home)))
	mux.Handle("/artist", app.onlyGetMethod(http.HandlerFunc(app.showArtist)))

	return app.logRequest(mux)
}