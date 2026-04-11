package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/schema"
)

func NewDB(ctx context.Context, sm secrets.Manager) (*sql.DB, error) {
	uri, err := sm.Get(ctx, "PG_URI")
	if err != nil {
		return nil, err
	}
	db, err := newDBFromUri(ctx, uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewTestDB(t *testing.T, ctx context.Context, cfg config.Postgres, sm secrets.Manager) *sql.DB {
	t.Helper()
	pg := newPostgresContainer(t, ctx, cfg, sm)

	db, err := newDBFromUri(ctx, pg.URI)
	require.NoError(t, err)

	err = schema.LoadTestData(ctx, db)
	require.NoError(t, err)

	return db
}

func newDBFromUri(ctx context.Context, url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("Pg open: %v", err)
	}

	if err := schema.LoadSchema(ctx, db); err != nil {
		return nil, err
	}
	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Pg ping: %v", err)
	}

	return db, nil
}

func uri(username, password, database string) string {
	return fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", username, password, database)
}
