package main

import (
	"musicproject.com/config"
	"musicproject.com/internal/handler"
)

func main() {
	cfg := config.LoadConfig()

	//repo := postgres.New(cfg)
	server := handler.NewServer(cfg, nil)

	// start server
	server.Listen()
}
