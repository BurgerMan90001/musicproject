package download

import (
	"context"
	"time"

	"github.com/google/uuid"
	"songsled.com/internal/services/file"
	"songsled.com/pkg/model"
)

type songRepo interface {
	GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Error, error)
}

type Song struct {
	bucket string
	prefix string
	store  file.Blobstore
	repo   songRepo
	// Should be around 5-15 minutes
	urlTtl time.Duration
}

func NewSong(urlTtl time.Duration) *Song {

	return &Song{urlTtl: urlTtl}
}

func (s *Song) DownloadByID(ctx context.Context, songId uuid.UUID) error {
	s.repo.GetSongByID(ctx, songId)

	return nil
	// key := filepath.Join(s.prefix)
	// url, err := s.store.GetObjectUrl(ctx, s.bucket)
}
