package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
	"time"

	"url-shortener/internal/jsonlog"
)

type application struct {
	urlMap   map[string]string
	mapMutex *sync.RWMutex
	config   config
	logger   *jsonlog.Logger
}

func main() {
	cfg := loadConfigFromEnv()
	flag.Parse()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	app := &application{
		urlMap:   make(map[string]string),
		mapMutex: &sync.RWMutex{},
		config:   *cfg,
		logger:   logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.routes(),
		ErrorLog:     log.New(logger, "", 0),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.PrintInfo("starting server", map[string]string{
		"addr": srv.Addr,
		"env":  cfg.env,
	})

	err := srv.ListenAndServe()
	logger.PrintFatal(err, nil)
}
