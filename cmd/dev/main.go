package main

import (
	"context"

	"musicproject.com/config"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/pkg/testutil"
)

func main() {
	cfg := config.LoadConfig()

	ctx := context.Background()
	pg, err := testutil.NewPostgresContainer(ctx, cfg.Repository.Postgres)
	if err != nil {
		panic(err)
	}
	defer pg.Terminate(ctx)

	repo := postgres.New(pg.URL)
	server := handler.NewServer(cfg, repo)

	// start server
	server.Listen()
}
