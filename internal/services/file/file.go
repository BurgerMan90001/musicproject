package file

import (
	"context"

	"musicproject.com/pkg/model"
)

type Service struct {
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
