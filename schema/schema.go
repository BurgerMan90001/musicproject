package schema

import (
	"context"
	"database/sql"
	"embed"
)

//go:embed schema.sql testdata.sql
var schemaFS embed.FS

func LoadSchema(ctx context.Context, db *sql.DB) error {
	schema, err := schemaFS.ReadFile("schema.sql")
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, string(schema)); err != nil {
		return err
	}
	return nil
}

func LoadTestData(ctx context.Context, db *sql.DB) error {
	data, err := schemaFS.ReadFile("testdata.sql")
	if err != nil {
		return err
	}
	if _, err := db.ExecContext(ctx, string(data)); err != nil {
		return err
	}
	return nil
}
