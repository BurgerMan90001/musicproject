package main

import (
	"fmt"
	"log"
	"net/http"

	"okapi.com/config"
	"okapi.com/internal/handler"
)

func main() {
	cfg := config.ReadConfigFile()

	port := cfg.APIConfig.Port
	host := cfg.APIConfig.Host

	mux := http.NewServeMux()

	handler := handler.New(mux, cfg)

	handler.Register("")

	// start server
	log.Printf("Server listening at %v:%d", host, port)
	if err := http.ListenAndServe(fmt.Sprintf("%v:%d", host, port), mux); err != nil {
		panic(err)
	}

}
