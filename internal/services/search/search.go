package search

import (
	"context"

	"musicproject.com/pkg/model"
)

type Service interface {
	SearchSongs(ctx context.Context, query string) ([]model.Song, error)
	Filter(ctx context.Context) error
}

// func New(repo repository.Song) *Service {

// 	return &Service{repo: repo}
// }

// func (s *Service) Filter(ctx context.Context, query ...string) {

// }
