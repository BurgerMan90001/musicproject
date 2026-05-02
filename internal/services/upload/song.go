package upload

import (
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/google/uuid"
	"songsled.com/internal/config"
	"songsled.com/internal/services/encode"
	"songsled.com/internal/services/file"
	"songsled.com/pkg/crand"
	"songsled.com/pkg/model"
)

type Song struct {
	// Required
	bucket string
	// Prefix for files
	prefix string

	caching bool
	// Should be 1 hour or less
	urlTtl time.Duration
	// Required
	// File storage
	store file.Blobstore
	// Optional
	// Encodes files before storing
	encoder encode.HLSEncoder
	// Required
	// Repository to store song metadata
	repo songRepo
}

type songRepo interface {
	PutSong(ctx context.Context, s *model.Song) (uuid.UUID, error)
}

func NewSong(bucket, prefix string,
	encoding, caching bool,
	urlTtl time.Duration,
	store file.Blobstore, repo songRepo) (*Song, error) {
	var encoder encode.HLSEncoder
	if encoding {
		encoder = encode.NewFFmpeg(config.Encoder{})
	}
	if bucket == "" {
		return nil, errors.New("Upload.NewSong: bucket is empty")
	}
	if prefix == "" {
		return nil, errors.New("Upload.NewSong: prefix is empty")
	}

	return &Song{bucket,
		prefix, caching, urlTtl,
		store, encoder,
		repo}, nil
}

// Uploads the song's metadata to the repository.
// Returns the url to upload the song file to.
func (s *Song) UploadMetadata(ctx context.Context,
	songRequest *model.UploadSongRequest) (string, error) {

	if err := validateUploadRequest(songRequest); err != nil {
		return "", err
	}
	//https://storage.googleapis.com/spume-musicproject/audio/24e692e5-mysong.mp3

	// Add additional random characters to avoid collisions
	filename := fmt.Sprintf("%s-%s", crand.NewShort(), songRequest.Filename)
	key := filepath.Join(s.prefix, filename)
	presignUrl, objectUrl, err := s.store.CreateObjectUrl(ctx, s.bucket, key, s.caching, s.urlTtl)
	if err != nil {
		return "", err
	}

	// Put song metadata in repository
	if _, err := s.repo.PutSong(ctx, &model.Song{
		Name:  songRequest.Name,
		Genre: songRequest.Genre,
		Image: songRequest.Image,

		URL: objectUrl,
	}); err != nil {
		return "", err
	}

	return presignUrl, nil
}

// Local uploads of audio files
// Handles multipart file uploads if the service uses local uploads
func (s *Song) UploadFile(ctx context.Context, file multipart.File,
	header *multipart.FileHeader) error {

	contents, err := io.ReadAll(file)
	if err != nil {
		return err
	}
	// Check if file is an audio file
	contentType := http.DetectContentType(contents)
	if !strings.HasPrefix(contentType, "audio/") {
		return fmt.Errorf("Incorrect mime type: %v", contentType)
	}
	return nil
}
