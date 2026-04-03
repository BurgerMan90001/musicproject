package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"musicproject.com/config"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/server"
)

func main() {
	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("panic %v", r)
		}
	}()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("config load error: %v", err)
	}

	//db := postgres.NewDB(cfg.Repository.Postgres.URL())
	server, err := server.NewServer(cfg.API.Port)
	if err != nil {
		log.Fatalf("new server error: %v", err)
	}

	// start server
	handler, err := handler.NewMux(ctx, cfg, nil)
	server.ServeHTTPHandler(ctx, handler)

	log.Println("server shutdown")
}
