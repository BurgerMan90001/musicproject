package search

import (
	"context"
	"database/sql"

	"musicproject.com/pkg/model"
)

type Postgres struct {
	db *sql.DB
}

func NewPostgres(db *sql.DB) *Postgres {
	return &Postgres{db}
}

func (s *Postgres) SearchSongs(ctx context.Context) ([]model.Song, error) {
	query := ""
	rows, err := s.db.QueryContext(ctx, query)
	if err != nil {

		return nil, err
	}
	for rows.Next() {
		var ()
		err := rows.Scan()
		if err != nil {
			return nil, err
		}

	}
	return nil, nil
}
