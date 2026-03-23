package postgres

import (
	"context"
	"database/sql"

	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/util/fileutil"
)

type Repository struct {
	User   repository.User
	Song   repository.Song
	Rating repository.Rating
}

func New(cfg config.Config) Repository {

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

	return Repository{&User{db}, &Song{db}, &Rating{db}}
}
