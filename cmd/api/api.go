package main

import (
	"log"
	"net/http"
	"time"

	"github.com/akhilr007/socials/internal/handler"
	appMiddleware "github.com/akhilr007/socials/internal/middleware"
	"github.com/akhilr007/socials/internal/store"

	"github.com/go-chi/chi/v5"
)

type application struct {
	config        config
	store         store.Storage
	postHandler   *handler.PostHandler
	appMiddleware AppMiddleware
}

type config struct {
	addr string
	db   dbConfig
	env  string
}

type dbConfig struct {
	addr         string
	maxOpenConns int
	maxIdleConns int
	maxIdleTime  string
}

type AppMiddleware struct {
	postMiddleware appMiddleware.PostMiddleware
}

func (app *application) run(mux *chi.Mux) error {

	srv := &http.Server{
		Addr:         app.config.addr,
		Handler:      mux,
		WriteTimeout: time.Second * 30,
		ReadTimeout:  time.Second * 10,
		IdleTimeout:  time.Minute,
	}

	log.Printf("server has started at %s", app.config.addr)
	return srv.ListenAndServe()
}
