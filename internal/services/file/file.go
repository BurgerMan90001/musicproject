package file

import (
	"context"

	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

type Service struct {
	repo repository.Repository
}

func New() *Service {
	return &Service{}
}

func (s *Service) UploadSong(ctx context.Context, song *model.Song) error {
	return nil
}
func (s *Service) Backup(ctx context.Context) error {
	return nil

}
