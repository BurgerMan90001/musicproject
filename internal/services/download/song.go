package download

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/file"
	"songsled.com/pkg/model"
)

// type songRepo interface {
// 	GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Error, error)
// }

type Service struct {
	bucket string
	prefix string
	store  file.Blobstore
	// repo   songRepo
	// Should be around 5-15 minutes
	urlTtl time.Duration

	songRepo *postgres.Song
}

func NewSong(
	urlTtl time.Duration,

	songRepo *postgres.Song,
) *Service {

	return &Service{
		urlTtl:   urlTtl,
		songRepo: songRepo,
	}
}

// TODO
func (s *Service) DownloadSongByID(ctx context.Context, songId uuid.UUID) error {
	song, err := s.songRepo.GetSongByID(ctx, songId)
	if err != nil {
		return fmt.Errorf(": %w", err)
	}

	if song.Audio == "" {
		return errors.New("Could not find song's audio file")
	}

	return nil
	// key := filepath.Join(s.prefix)
	// url, err := s.store.GetObjectUrl(ctx, s.bucket)
}

func (s *Service) DownloadUrl(ctx context.Context, songId uuid.UUID) (*model.SongDownloadResponse, error) {
	song, err := s.songRepo.GetSongByID(ctx, songId)
	if err != nil {
		return nil, fmt.Errorf(": %w", err)
	}

	if song.Audio == "" {
		return nil, errors.New("Could not find song's audio file")
	}

	presign, err := s.store.GetObjectUrl(ctx, s.bucket, song.Audio, s.urlTtl)
	if err != nil {
		return nil, fmt.Errorf("Download service, GetObjectUrl: %w", err)
	}

	return &model.SongDownloadResponse{
		Song: song,
		Links: []model.Link{
			{
				Rel:  "download",
				Href: presign,
			},
		},
	}, nil
}
