package main

import (
	"database/sql"
	"log"
	"path/filepath"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"songsled.com/internal/config/secrets"
)

// TODO
func main() {

	if err := run(); err != nil {
		log.Fatal(err.Error())
	}

}

func run() error {
	uri, err := secrets.Getenv("PG_URI")
	if err != nil {
		return err
	}
	db, err := sql.Open("postgres", uri)
	if err != nil {
		return err
	}
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(filepath.Join("database", "migrate"), "postgres", driver)
	if err != nil {
		return err
	}
	return m.Up()
}
