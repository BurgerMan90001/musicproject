package main

import (
	"fmt"
	"log"
	"net/http"

	"okapi.com/config"
	"okapi.com/internal/handler"
	"okapi.com/internal/util/fileutil"
)

func main() {

	cfg, err := fileutil.ReadYAML[config.ServiceConfig]("base.yml")
	if err != nil {
		panic(err)
	}
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
