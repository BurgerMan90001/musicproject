package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"

	"musicproject.com/internal/config/secrets"
	"musicproject.com/schema"
)

func NewDB(ctx context.Context) (*sql.DB, error) {
	uri := os.Getenv("PG_URI")

	db, err := newDBFromUri(ctx, uri)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewTestDB(t *testing.T, ctx context.Context, image string) *sql.DB {
	t.Helper()
	pg := newPostgresContainer(t, ctx, image)

	db, err := newDBFromUri(ctx, pg.URI)
	require.NoError(t, err)

	err = schema.LoadTestData(ctx, db)
	require.NoError(t, err)

	return db
}

func newDBFromUri(ctx context.Context, uri string) (*sql.DB, error) {
	// Verify that credentials are not empty
	_, err := secrets.GetEnvMap("PG_USERNAME", "PG_PASSWORD", "PG_DATABASE", "PG_HOST")
	if err != nil {
		return nil, err
	}
	if uri == "" {
		return nil, fmt.Errorf("postgres uri empty")
	}
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("Pg open: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Pg ping: %v", err)
	}
	if err := schema.LoadSchema(ctx, db); err != nil {
		return nil, err
	}

	return db, nil
}

// func uri(username, password, database string) string {
// 	return fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", username, password, database)
// }
