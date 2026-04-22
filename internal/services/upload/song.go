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

	"musicproject.com/internal/config"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/encode"
	"musicproject.com/internal/services/file"
	"musicproject.com/pkg/model"
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
	repo repository.Song
}

func NewSong(bucket, prefix string, encoding bool, caching bool, urlTtl time.Duration,
	store file.Blobstore, repo repository.Song) (*Song, error) {
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

	// Put song metadata in repository
	_, err := s.repo.Put(ctx, &model.Song{
		Name:  songRequest.Name,
		Genre: songRequest.Genre,
		Image: songRequest.Image,
	})
	if err != nil {
		return "", err
	}
	key := filepath.Join(s.prefix, songRequest.Filename)
	url, err := s.store.CreateObjectUrl(ctx, s.bucket, key, s.caching, s.urlTtl)
	if err != nil {
		return "", err
	}
	return url, nil
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
