package song

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/file"
	"musicproject.com/pkg/model"
)

type Service struct {
	store  file.Blobstore
	repo   repository.Song
	parent string
}

func NewSong(store file.Blobstore, repo repository.Song) *Service {
	return &Service{store, repo, "/test"}
}

func (s *Service) UploadSong(ctx context.Context, file multipart.File,
	header *multipart.FileHeader, songRequest model.UploadSongRequest) (*model.Song, error) {

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}
	// Check if file is an audio file
	mimeType := http.DetectContentType(contents)
	if !isAudioType(mimeType) {
		return nil, fmt.Errorf("Upload incorrect mime type: %v", mimeType)
	}

	if err := s.store.CreateObject(ctx, "", header.Filename, contents, true, mimeType); err != nil {
		return nil, err
	}

	song := &model.Song{
		Name:  songRequest.Name,
		Genre: songRequest.Genre,
		Image: songRequest.Image,
	}

	songId, err := s.repo.Put(ctx, song)
	if err != nil {
		return nil, err
	}

	song.ID = songId

	return song, nil
}
func (s *Service) Backup(ctx context.Context) error {
	return nil
}
func (s *Service) DownloadByID(ctx context.Context, songId uuid.UUID) error {
	return nil
}
func (s *Service) DeleteByID(ctx context.Context, songId uuid.UUID) error {
	return nil
}

func (s *Service) ListSongs() {

}

func isAudioType(mimeType string) bool {
	return strings.HasPrefix(mimeType, "audio/")
}
