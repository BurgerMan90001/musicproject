package search

import (
	"context"

	"songsled.com/pkg/model"
)

type Service interface {
	SearchSongs(ctx context.Context, query string) ([]model.Song, error)
	Filter(ctx context.Context) error
}


