package search

import (
	"context"

	"songsled.com/pkg/model"
)

var _ Service = (*Postgres)(nil)

// Postgres full text searching with pgvector
type Postgres struct {
}

func NewPostgres() *Postgres {
	return &Postgres{}
}

func (s *Postgres) SearchSongs(ctx context.Context, query string) ([]model.Song, error) {
	return nil, nil
}

func (s *Postgres) Filter(ctx context.Context) error {
	return nil
}
