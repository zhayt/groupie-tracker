package main

import (
	"github.com/zhayt/groupie-tracker/pkg/service"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		app.notFound(w)
		return
	}
	_, err := service.GetAllArtists()
	if err != nil {
		app.serverError(w, err)
		return
	}

	w.Write([]byte("Home page"))
}

func (app *application) showArtist(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Show Artist"))
}
