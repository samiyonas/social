package main

import (
    "net/http"
    "time"
    "github.com/go-chi/chi/v5"
    _ "github.com/go-chi/chi/v5/middleware"
    "github.com/samiyonas/social/internal/store"
)

type application struct {
    config config
    store store.Storage
}

type config struct {
    addr string
    db dbConfig
}

type dbConfig struct {
    addr string
    maxOpenConns int
    maxIdleConns int
    maxIdleTime string
}

func (app *application) multiplexer () http.Handler{
    r := chi.NewRouter() // implements http.Handler inteface (ServeHTTP method)

    r.Route("/v1", func (r chi.Router) {
        r.Get("/health", app.healthCheckHandler)
    })

    return r
}

func (app *application) run (mux http.Handler) error {
    srv := &http.Server{
        Addr: app.config.addr,
        Handler: mux,
        WriteTimeout: time.Second * 30, // timesout if our server fails to write in 30 sec
        ReadTimeout: time.Second * 10, // timeout if the client fails to read in 10 sec
        IdleTimeout: time.Minute,
    }

    return srv.ListenAndServe()
}
