package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	"songsled.com/internal/config"
	"songsled.com/internal/config/secrets"
	"songsled.com/internal/handler"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/server"
	"songsled.com/internal/services/file"
)

func main() {
	var (
		envFile    string
		configFile string
		schemaFile string
	)

	flag.StringVar(&envFile, "envFile", filepath.Join(".env.dev"), "specifies the location of the env file")
	flag.StringVar(&configFile, "config", filepath.Join("config.dev.yml"), "specifies the location of the config file")
	flag.StringVar(&schemaFile, "schema", filepath.Join("database", "schema.sql"), "specifies the location of the schema file")
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("panic %v", r)
		}
	}()

	if err := run(ctx, configFile, envFile, schemaFile); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server shutdown")
}

func run(ctx context.Context, configFile, envFile, schemaFile string) error {
	if err := secrets.ReadEnvFile(envFile); err != nil {
		return fmt.Errorf("Load env file: %v", err)
	}

	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return fmt.Errorf("Load config file: %v", err)
	}

	//create database connection
	repo, err := postgres.New(ctx)
	if err != nil {
		return fmt.Errorf("New postgres connection: %v", err)
	}
	defer repo.DB.Close()
	if err := repo.ExecFile(ctx, schemaFile); err != nil {
		return err
	}
	if os.Getenv("LOAD_TESTDATA") == "true" {
		err = repo.ExecFile(ctx, filepath.Join("test", "integration", "testdata", "testdata.sql"))
	}
	// create server
	server, err := server.NewServer(cfg.API.Port)
	if err != nil {
		return fmt.Errorf("New server: %v", err)
	}

	store, err := file.New(ctx, &file.AWSS3{}, cfg.File.Region, cfg.File.Endpoint, cfg.File.Endpoint)
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
