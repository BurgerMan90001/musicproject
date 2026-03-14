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

func New(cfg config.Config) *Repository {
	var url string

	switch cfg.Repository.Type {
	case "test":
		url = cfg.Repository.TestURL
	case "postgres":
		url = cfg.Repository.URL
	default:
		url = cfg.Repository.TestURL
	}

	ctx := context.Background()
	db, err := sql.Open("postgres", url)
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
