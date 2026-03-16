package main

import (
	"log"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/server"
)

func main() {
	cfg := config.ReadConfigFile()

	mux := http.NewServeMux()

	repo := postgres.New(cfg)
	server := server.New(mux, repo, cfg)

	server.Handle()

	// start server
	log.Printf("Server listening at %s", cfg.ApiUrl())
	if err := http.ListenAndServe(cfg.ApiUrl(), mux); err != nil {
		panic(err)
	}
}
