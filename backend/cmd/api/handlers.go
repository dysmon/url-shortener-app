package main

import (
	"net/http"
	"net/url"

	"url-shortener/internal/models"

	"github.com/gorilla/mux"
)

func (app *application) shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	var urlRequest *models.UrlRequest
	err := app.readJson(w, r, &urlRequest)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if _, err := url.ParseRequestURI(urlRequest.URL); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	shortCode := generateShortCode(urlRequest.URL)

	app.mapMutex.Lock()
	app.urlMap[shortCode] = urlRequest.URL
	app.mapMutex.Unlock()

	err = app.writeJson(w, http.StatusOK, envelope{
		"OriginalURL": urlRequest.URL,
		"ShortURL":    "http://" + r.Host + "/v1/" + shortCode,
	})

	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) redirectHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	shortCode := vars["shortCode"]

	app.mapMutex.RLock()
	originalUrl, exists := app.urlMap[shortCode]
	app.mapMutex.RUnlock()

	if !exists {
		app.notFoundResponse(w, r)
		return
	}

	http.Redirect(w, r, originalUrl, http.StatusMovedPermanently)
}
