package main

import (
	"musicproject.com/config"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
)

func main() {
	cfg := config.LoadConfig()

	repo := postgres.New("")
	server := handler.NewServer(cfg, repo)

	// start server
	server.Listen()
}
