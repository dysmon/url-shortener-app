package main

import (
	"net/http"
)

func (app *application) logRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		app.logger.PrintInfo("request", map[string]string{
			"remote_addr": r.RemoteAddr,
			"proto":       r.Proto,
			"method":      r.Method,
			"url":         r.URL.String(),
		})

		next.ServeHTTP(w, r)
	})
}
