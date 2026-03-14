package main

import (
	"fmt"
	"log"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
)

func main() {
	cfg := config.ReadConfigFile("config/base.yml")

	port := cfg.API.Port
	host := cfg.API.Host

	mux := http.NewServeMux()

	repo := postgres.New(cfg)
	handler := handler.New(mux, repo, cfg)

	handler.Register("")

	// start server
	log.Printf("Server listening at %v:%d", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), mux); err != nil {
		panic(err)
	}
}
