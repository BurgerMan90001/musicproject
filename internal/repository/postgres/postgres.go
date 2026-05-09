package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"songsled.com/internal/config/secrets"
	"songsled.com/internal/repository/postgres/gensqlc"
)

type Repo struct {
	DB      *sql.DB
	Queries *gensqlc.Queries
}

func New(ctx context.Context, uriVar string) (*Repo, error) {
	uri, err := secrets.Getenv(uriVar)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, fmt.Errorf("Pg open: %v", err)
	}

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("Pg ping: %v, %v", uri, err)
	}
	return &Repo{db, gensqlc.New(db)}, nil
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
