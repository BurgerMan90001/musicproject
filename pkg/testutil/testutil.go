package testutil

import (
	"context"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"musicproject.com/config"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	URL string
}

func NewPostgresContainer(ctx context.Context, cfg config.Postgres) (*PostgresContainer, error) {
	pg, err := postgres.Run(ctx, cfg.Image,
		postgres.WithDatabase(cfg.Database),
		postgres.WithUsername(cfg.Username),
		postgres.WithPassword(cfg.Password),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		return nil, err
	}
	url, err := pg.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		return nil, err
	}

	return &PostgresContainer{PostgresContainer: pg, URL: url}, nil
}
