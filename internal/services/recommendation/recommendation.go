package recommendation

import (
	"context"

	"songsled.com/pkg/model"
)

type songRepo interface {
	GetRandom(ctx context.Context, n int) ([]model.Song, error)
}

type Service struct {
	songRepo songRepo
}

func New(songRepo songRepo) *Service {
	return &Service{songRepo: songRepo}
}

func (s *Service) GetRandom(ctx context.Context) ([]model.Song, error) {
	return s.songRepo.GetRandom(ctx, 10)
}
