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

	"songsled.com/internal/config"
	"songsled.com/internal/handler"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/server"
	"songsled.com/internal/services/file"
)

func main() {
	var (
		env        string
		configFile string
	)

	flag.StringVar(&env, "env", "dev", "environment")
	flag.StringVar(&configFile, "config", filepath.Join("config.yml"), "specifies the location of the config file")
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("panic %v", r)
		}
	}()

	if err := run(ctx, configFile, env); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server shutdown")
}

func run(ctx context.Context, configFile, env string) error {
	// if env == "dev" {
	// 	if err := secrets.ReadEnvFile(filepath.Join(".env.dev")); err != nil {
	// 		return fmt.Errorf("Load env file: %v", err)
	// 	}
	// }

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return fmt.Errorf("Load config file: %v", err)
	}

	//create database connection
	repo, err := postgres.New(ctx, "PG_URI")
	if err != nil {
		return fmt.Errorf("New postgres connection: %v", err)
	}
	defer repo.DB.Close()

	// create server
	server, err := server.NewServer(cfg.API.Port)
	if err != nil {
		return fmt.Errorf("New server: %v", err)
	}

	store, err := file.New(ctx, &file.AWSS3{}, cfg.File)
	if err != nil {
		return err
	}

	// rdb := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// defer rdb.Close()

	// create handler
	handler, err := handler.New(ctx, cfg, store, repo, nil)
	if err != nil {
		return fmt.Errorf("New mux handler: %w", err)
	}

	// start server
	return server.ServeHTTPHandler(ctx, handler)
}
