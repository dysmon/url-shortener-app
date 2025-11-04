package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	router := mux.NewRouter()

	router.NotFoundHandler = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowedHandler = http.HandlerFunc(app.methodNotAllowedResponse)

	router.Use(app.logRequest)

	router.HandleFunc("/v1/healthcheck", app.healthcheckHandler).Methods(http.MethodGet)

	router.HandleFunc("/v1/shorten", app.shortenURLHandler).Methods(http.MethodPost)
	router.HandleFunc("/v1/{shortCode}", app.redirectHandler).Methods(http.MethodGet)

	return router
}
