package main

import (
	"log"
	"os"
	"strconv"
)

type config struct {
	port int
	env  string
}

func loadConfigFromEnv() *config {
	cfg := &config{}

	portStr := os.Getenv("PORT")
	if portStr == "" {
		log.Println("PORT not set, defaulting to 8000")
		cfg.port = 8000
	} else {
		port, err := strconv.Atoi(portStr)
		if err != nil {
			log.Printf("Invalid PORT value %q, defaulting to 8000\n", portStr)
			cfg.port = 8000
		} else {
			cfg.port = port
		}
	}

	cfg.env = os.Getenv("ENV")
	if cfg.env == "" {
		log.Println("ENV not set, defaulting to 'development'")
		cfg.env = "development"
	}

	return cfg
}
