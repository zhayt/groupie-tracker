package main

import (
	"github.com/zhayt/groupie-tracker/pkg/service"
	"html/template"
	"net/http"
	"strconv"
	"strings"
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

	f := service.NewArtists(nil)
	f.AllArtists = app.hash
	f.SearchedArtist = app.hash
	err = tmp.Execute(w, f)
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

func (app *application) searchArtist(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	f := service.NewArtists(r.PostForm)

	f.AllArtists, err = service.GetAll(app.api)
	if err != nil {
		app.serverError(w, err)
		return
	}

	//err = f.Required("search")
	//if err != nil {
	//	app.clientError(w, http.StatusBadRequest)
	//	return
	//}

	toSearch, err := f.ParseValue(strings.TrimSpace(r.FormValue("search")))
	if err != nil {
		app.clientError(w, http.StatusBadRequest)
		return
	}

	f.Search(toSearch)

	tmp, err := template.ParseFiles("./web/templates/index.html")
	if err != nil {
		app.serverError(w, err)
		return
	}

	err = tmp.Execute(w, f)
	if err != nil {
		app.serverError(w, err)
		return
	}
}
