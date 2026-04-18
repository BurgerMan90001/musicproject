package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"path/filepath"
	"strings"

	"musicproject.com/internal/config"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/encode"
	"musicproject.com/internal/services/file"
	"musicproject.com/pkg/model"
)

type Song struct {
	parent  string
	store   file.Blobstore
	encoder encode.HLSEncoder
	repo    repository.Song
}

func New(parent string, encoding bool, store file.Blobstore, repo repository.Song) *Song {
	var encoder encode.HLSEncoder
	if encoding {
		encoder = encode.NewFFmpeg(config.Encoder{})
	}
	// TODO Use local filesystem if the bucket is empty
	if parent == "" {
		parent = ""
	}
	return &Song{parent, store, encoder, repo}
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
	
	url, err := s.store.CreateObjectUrl(ctx, s.parent, filepath.Join("files/audio", songRequest.Filename), true)
	if err != nil {
		return "", err
	}
	return url, nil
}

// Local uploads of audio files
// Handles file uploads if the url is here
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
