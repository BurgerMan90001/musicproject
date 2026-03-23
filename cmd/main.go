package main

import (
	"log"
	"net/http"

	"musicproject.com/config"
	"musicproject.com/internal/handler"
)

func main() {
	cfg := config.ReadConfigFile(config.TypeDev)

	mux := http.NewServeMux()

	handler.New(mux, cfg)


	// start server
	log.Printf("Server listening at %s", cfg.URL())
	if err := http.ListenAndServe(cfg.URL(), mux); err != nil {
		panic(err)
	}
}
