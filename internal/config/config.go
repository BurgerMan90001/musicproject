package config

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"songsled.com/internal/fileutil"
)

const (
	DevConfig  = "config.dev.yml"
	ProdConfig = "config.prod.yml"
)


func LoadConfig(filename string) (*Config, error) {

	cfg, err := fileutil.ReadYaml[Config](filename)
	if err != nil {
		return nil, fmt.Errorf("Load config: %w", err)
	}

	return cfg, nil
}

func LoadSchema(ctx context.Context, filename string, db *sql.DB) error {
	schema, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("Postgres load schema: %w", err)
	}
	if _, err := db.ExecContext(ctx, string(schema)); err != nil {
		return fmt.Errorf("Postgres load schema: %w", err)
	}
	return nil
}
