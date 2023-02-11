package main

import (
	"flag"
	"github.com/zhayt/groupie-tracker/pkg/models"
	"log"
	"net/http"
	"os"
	"time"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	api      string
	hash     []*models.Artist
}

func main() {
	addr := flag.String("addr", ":9000", "HTTP network address")
	api_addr := flag.String("api_addr", "https://groupietrackers.herokuapp.com/api/artists", "API url")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := application{
		errorLog: errorLog,
		infoLog:  infoLog,
		api:      *api_addr,
	}

	srv := &http.Server{
		Addr:         *addr,
		ErrorLog:     errorLog,
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 25 * time.Second,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	if err != nil {
		errorLog.Fatalf(err.Error())
	}
}
