package postgres

import (
	"context"
	"database/sql"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"musicproject.com/config"
	"musicproject.com/schema"
)

func NewDB(ctx context.Context, url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, err
	}

	if err := schema.LoadSchema(ctx, db); err != nil {
		return nil, err
	}
	return db, nil
}
func NewTestDB(t *testing.T, ctx context.Context, cfg config.Postgres) (*sql.DB, *PostgresContainer) {
	t.Helper()
	pg := newPostgresContainer(t, ctx, cfg)

	db, err := NewDB(ctx, pg.URL)
	require.NoError(t, err)

	err = schema.LoadTestData(ctx, db)
	require.NoError(t, err)

	return db, pg
}

type PostgresContainer struct {
	*postgres.PostgresContainer
	URL string
}

func newPostgresContainer(t *testing.T, ctx context.Context, cfg config.Postgres) *PostgresContainer {
	t.Helper()

	pg, err := postgres.Run(ctx, cfg.Image,
		postgres.WithDatabase(cfg.Database),
		postgres.WithUsername(cfg.Username),
		postgres.WithPassword(cfg.Password),
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
