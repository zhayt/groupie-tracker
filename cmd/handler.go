package main

import (
	"github.com/zhayt/groupie-tracker/pkg/service"
	"html/template"
	"net/http"
	"strconv"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		app.notFound(w)
		return
	}

	tmp, err := template.ParseFiles("./web/templates/index.html")

	if app.hash == nil {
		app.hash, err = service.GetAll(app.api)
		if err != nil {
			app.serverError(w, err)
			return
		}
	}

	err = tmp.Execute(w, app.hash)
	if err != nil {
		app.serverError(w, err)
		return
	}
}

func (app *application) showArtist(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.notFound(w)
		return
	}

	res, err := service.GetById(app.api, strconv.Itoa(id))

	if err != nil {
		if err.Error() == "not found" {
			app.notFound(w)
			return
		}
		app.serverError(w, err)
		return
	}

	tmp, err := template.ParseFiles("./web/templates/artist.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tmp.Execute(w, res)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
