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
	switch cfg.Repository.Type {
	case "test":
		db, err := sql.Open("postgres", cfg.Repository.TestURL)
		if err != nil {
			panic(err)
		}
		ctx := context.TODO()
		if err := fileutil.ExecSql(ctx, db, "schema/schema.sql"); err != nil {
			panic(err)
		}
		return &Repository{db}
	case "postgres":
		db, err := sql.Open("postgres", cfg.Repository.URL)
		if err != nil {
			panic(err)
		}
		return &Repository{db}
	default:
		db, err := sql.Open("postgres", cfg.Repository.URL)
		if err != nil {
			panic(err)
		}
		return &Repository{db}
	}

}
