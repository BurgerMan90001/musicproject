package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	mpostgres "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/testcontainers/testcontainers-go"
	pgcontainer "github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
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

// Callers must handle cleanup with testcontainers.Terminate
func NewContainer(ctx context.Context) (*pgcontainer.PostgresContainer, error) {
	username := "musicproject"
	password := "admin"
	c, err := pgcontainer.Run(ctx, "postgres:alpine",
		// pgcontainer.WithInitScripts(schama),
		pgcontainer.WithUsername(username),
		pgcontainer.WithPassword(password),
		pgcontainer.WithDatabase(username),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).WithStartupTimeout(5*time.Second)))
	if err != nil {
		return nil, fmt.Errorf("Run postgres container: %w", err)
	}

	s, err := c.ConnectionString(ctx)
	if err != nil {
		return nil, err
	}
	if err := os.Setenv("PG_URI", s); err != nil {
		return nil, err
	}
	return c, nil
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

// Set forceVer to 0 to not force
func (r *Repo) Migrate(migrations, action string, forceVer int) error {
	driver, err := mpostgres.WithInstance(r.DB, &postgres.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithDatabaseInstance(
		migrations, "postgres", driver)
	if err != nil {
		return fmt.Errorf("Migrate NewWithDatabaseInstance: %w", err)
	}
	if forceVer > 1 {
		if err := m.Force(forceVer); err != nil {
			return fmt.Errorf("Migrate force version: %w", err)
		}
	}
	slog.Info("Migrating")
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
