package main

import (
	"context"
	"fmt"
	"log"
	"os/signal"
	"syscall"

	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
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

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}

	log.Println("server shutdown")
}

func run(ctx context.Context) error {
	log.Println("Starting server")
	// load config
	cfg, err := config.LoadConfig()
	if err != nil {
		return fmt.Errorf("Config load: %v", err)
	}
	// create secret manager
	sm, err := secrets.NewEnv()
	if err != nil {
		return fmt.Errorf("New env secret: %v", err)
	}
	//create database connection
	db, err := postgres.NewDB(ctx, sm)
	if err != nil {
		return fmt.Errorf("New postgres: %v", err)
	}
	// db = nil
	// create server
	server, err := server.NewServer(cfg.API.Port)
	if err != nil {
		return fmt.Errorf("New server: %v", err)
	}

	// create handler
	handler, err := handler.NewMux(ctx, cfg, db, sm)
	if err != nil {
		return fmt.Errorf("New handler: %w", err)
	}

	// start server
	return server.ServeHTTPHandler(ctx, handler)
}
