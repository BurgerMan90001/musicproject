package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"songsled.com/internal/config/secrets"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatal("No action set (up, down)")
	}

	fmt.Println(args)
	if err := run(args); err != nil {
		log.Fatal(err.Error())
	}
}

func run(args []string) error {

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

	m, err := migrate.NewWithDatabaseInstance(
		"file://database/migrate", "postgres", driver)
	if err != nil {
		return err
	}

	var action string = args[0]
	if len(args) > 1 && args[1] == "force" {
		if len(args) <= 2 {
			return errors.New("No version")
		}
		v, err := strconv.Atoi(args[2])
		if err != nil {
			return fmt.Errorf("Parse force version: %w", err)
		}
		if err := m.Force(v); err != nil {
			return err
		}
	}

	switch action {
	case "up":
		return m.Up()
	case "down":
		return m.Down()
	case "drop":
		return m.Drop()
	default:

		return fmt.Errorf("Invalid action: %s", action)
	}
}
