package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
	"songsled.com/internal/config/secrets"
	"songsled.com/internal/repository/postgres/gensqlc"
)

type Repo struct {
	DB      *sql.DB
	Queries *gensqlc.Queries
}

func New(ctx context.Context) (*Repo, error) {
	uri := os.Getenv("PG_URI")

	db, err := newDBFromUri(ctx, uri)
	if err != nil {
		return nil, err
	}

	return &Repo{db, gensqlc.New(db)}, nil
}

func NewTest(t *testing.T, ctx context.Context, image string) *Repo {
	t.Helper()
	pg := newPostgresContainer(t, ctx, image)

	db, err := newDBFromUri(ctx, pg.URI)
	require.NoError(t, err)

	return &Repo{db, gensqlc.New(db)}
}

func newDBFromUri(ctx context.Context, uri string) (*sql.DB, error) {
	// Verify that credentials are not empty
	_, err := secrets.GetenvMap("PG_USERNAME", "PG_PASSWORD")
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
		return nil, fmt.Errorf("Pg ping: %v, %v", uri, err)
	}

	return db, nil
}

// Reads and executes a .sql file
func (r *Repo) ExecFile(ctx context.Context, filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("postgres.ExecFile: %s, %w", filename, err)
	}

	if _, err := r.DB.ExecContext(ctx, string(data)); err != nil {
		return fmt.Errorf("postgres.ExecFile: %s, %w", filename, err)
	}
	return nil
}

type PostgresContainer struct {
	*postgres.PostgresContainer
	URI string
}

func newPostgresContainer(t *testing.T, ctx context.Context, image string) *PostgresContainer {
	t.Helper()

	secretList, err := secrets.GetenvMap("PG_USERNAME", "PG_PASSWORD")
	require.NoError(t, err)

	pg, err := postgres.Run(ctx, image,
		postgres.WithUsername(secretList["PG_USERNAME"]),
		postgres.WithPassword(secretList["PG_PASSWORD"]),
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
