package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os/signal"
	"path/filepath"
	"syscall"

	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/server"
)

func main() {
	var env string
	flag.StringVar(&env, "env", "dev", "environment type")
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("panic %v", r)
		}
	}()

	if err := run(ctx, env); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server shutdown")
}

func run(ctx context.Context, env string) error {

	configFile := "config.dev.yml"
	envFile := ".env.dev"
	if env == "prod" {
		configFile = "config.prod.yml"
		envFile = ".env.prod"
	}

	cfg, err := config.LoadConfig(filepath.Join("config", configFile))
	if err != nil {
		return fmt.Errorf("Config load: %v", err)
	}
	// load env
	err = secrets.LoadEnv(filepath.Join("config", envFile))
	if err != nil {
		return fmt.Errorf("New env secret: %v", err)
	}
	//create database connection
	db, err := postgres.NewDB(ctx, env)
	if err != nil {
		return fmt.Errorf("New postgres: %v", err)
	}
	// create server
	server, err := server.NewServer(cfg.API.Port)
	if err != nil {
		return fmt.Errorf("New server: %v", err)
	}

	// create handler
	handler, err := handler.NewMux(ctx, cfg, db)
	if err != nil {
		return fmt.Errorf("New handler: %w", err)
	}

	// start server
	return server.ServeHTTPHandler(ctx, handler)
}
