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

	"github.com/testcontainers/testcontainers-go"

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
		port       int
	)

	flag.StringVar(&env, "env", "dev", "environment")
	flag.StringVar(&configFile, "config", filepath.Join("config.yml"), "specifies the location of the config file")
	flag.IntVar(&port, "port", 8081, "")
	flag.Parse()

	ctx, done := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer func() {
		done()
		if r := recover(); r != nil {
			log.Fatalf("panic %v", r)
		}
	}()

	if err := run(ctx, configFile, env, port); err != nil {
		log.Fatal(err)
	}

	slog.Info("Server shutdown")
}

func run(ctx context.Context,
	configFile,
	env string,
	port int,
) error {
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		return fmt.Errorf("Load config file: %v", err)
	}

	if err := os.Setenv("ENV", env); err != nil {
		return err
	}

	if env == "dev" {
		c, err := postgres.NewContainer(ctx)
		if err != nil {
			return fmt.Errorf("New postgres container: %w", err)
		}
		defer func() {
			err := testcontainers.TerminateContainer(c)
			if err != nil {
				slog.Error(err.Error())
			}
		}()

	}
	//create database connection
	repo, err := postgres.New(ctx, "PG_URI")
	if err != nil {
		return fmt.Errorf("New postgres connection: %v", err)
	}
	defer repo.DB.Close()

	if env == "dev" {
		if err := repo.Migrate("file://database/migrate", "up", 0); err != nil {
			return fmt.Errorf("Migrate: %w", err)
		}
	}

	// create server
	server, err := server.NewServer(port)
	if err != nil {
		return fmt.Errorf("New server: %v", err)
	}

	store, err := file.New(ctx, &file.AWSS3{}, cfg.File)
	if err != nil {
		return err
	}

	// create handler
	handler, err := handler.New(ctx, cfg, store, repo, nil)
	if err != nil {
		return fmt.Errorf("New mux handler: %w", err)
	}

	// start server
	return server.ServeHTTPHandler(ctx, handler)
}
