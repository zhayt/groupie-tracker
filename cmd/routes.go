package main

import "net/http"

func (app *application) routes() http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(app.home))
	mux.Handle("/artist", http.HandlerFunc(app.showArtist))
	mux.HandleFunc("/search", app.search)

	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./web/static"))))
	return app.logRequest(app.onlyGetMethod(mux))
}
