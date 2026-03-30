package postgres

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
	"musicproject.com/schema"
)

type Repository struct {
	db *sql.DB
}

func New(url string) *Repository {
	ctx := context.Background()
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	if err := db.PingContext(ctx); err != nil {
		panic(err)
	}

	if err := schema.LoadSchema(ctx, db); err != nil {
		panic(err)
	}
	if err := schema.LoadTestData(ctx, db); err != nil {
		panic(err)
	}

	return &Repository{db}
}

func (r *Repository) Close() {
	r.db.Close()
}
