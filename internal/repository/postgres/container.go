package postgres

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"musicproject.com/internal/config/secrets"
)

type PostgresContainer struct {
	*postgres.PostgresContainer
	URI string
}

func newPostgresContainer(t *testing.T, ctx context.Context, image string) *PostgresContainer {
	t.Helper()

	secretList, err := secrets.GetEnvMap("PG_USERNAME",
		"PG_PASSWORD", "PG_DATABASE",
	)
	require.NoError(t, err)

	pg, err := postgres.Run(ctx, image,
		postgres.WithUsername(secretList["PG_USERNAME"]),
		postgres.WithPassword(secretList["PG_PASSWORD"]),
		postgres.WithDatabase(secretList["PG_DATABASE"]),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	require.NoError(t, err)
	testcontainers.CleanupContainer(t, pg)

	uri, err := pg.ConnectionString(ctx, "sslmode=disable")
	require.NoError(t, err)

	return &PostgresContainer{PostgresContainer: pg, URI: uri}
}
