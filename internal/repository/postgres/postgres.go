package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"musicproject.com/config"
	"musicproject.com/pkg/util/fileutil"
)

type Repository struct {
	db *sql.DB
}

func New(cfg *config.Config) *Repository {
	ctx := context.Background()
	db, err := sql.Open("postgres", cfg.Repository.URL)
	if err != nil {
		panic(err)
	}
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	if err := fileutil.ExecSql(ctx, db, "schema/schema.sql"); err != nil {
		panic(err)
	}

	return &Repository{db}
}
