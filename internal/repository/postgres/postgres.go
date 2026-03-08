package postgres

import (
	"context"
	"database/sql"
	"errors"

	embeddedpostgres "github.com/fergusstrange/embedded-postgres"
	_ "github.com/lib/pq"
	"musicproject.com/config"
	"musicproject.com/pkg/util/fileutil"
)

type Repository struct {
	db    *sql.DB
	Embed *embeddedpostgres.EmbeddedPostgres
}

func New(cfg config.Config) *Repository {
	switch cfg.Repository.Type {
	case "test":
		cfg := embeddedpostgres.DefaultConfig().StartParameters(map[string]string{
			"shared_buffers":  "16MB",
			"max_connections": "10",
		})
		embedDb := embeddedpostgres.NewDatabase(cfg)
		if err := embedDb.Start(); err != nil {
			panic(err)
		}
		db, err := sql.Open("postgres", cfg.GetConnectionURL()+"?sslmode=disable")
		if err != nil {
			panic(err)
		}

		if err := fileutil.ExecSql(context.TODO(), db, "schema/schema.sql"); err != nil {
			panic(err)
		}
		return &Repository{db, embedDb}
	case "postgres":
		db, err := sql.Open("postgres", cfg.Repository.URL)
		if err != nil {
			panic(err)
		}
		return &Repository{db, nil}
	default:
		db, err := sql.Open("postgres", cfg.Repository.URL)
		if err != nil {
			panic(err)
		}
		return &Repository{db, nil}
	}

}

func (r *Repository) Stop() error {

	return errors.Join(r.db.Close(), r.Embed.Stop())
}
