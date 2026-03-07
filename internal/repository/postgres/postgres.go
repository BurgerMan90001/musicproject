package postgres

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository struct {
	db *sql.DB
}

func New(url string) *Repository {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	db.SetMaxIdleConns(5)
	return &Repository{db}
}
