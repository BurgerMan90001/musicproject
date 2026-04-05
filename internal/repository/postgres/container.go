package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	URL string
}

func newPostgresContainer(t *testing.T, ctx context.Context, cfg config.Postgres, sm secrets.Manager) *PostgresContainer {
	t.Helper()

	s, err := secrets.GetSecrets(ctx, sm, "PG_USERNAME",
		"PG_PASSWORD", "PG_DATABASE",
	)
	require.NoError(t, err)

	pg, err := postgres.Run(ctx, cfg.Image,
		postgres.WithUsername(s[0]),
		postgres.WithPassword(s[1]),
		postgres.WithDatabase(s[2]),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, pg)

	url, err := pg.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	return &PostgresContainer{PostgresContainer: pg, URL: url}
}
