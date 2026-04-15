package upload

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strings"

	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/encode"
	"musicproject.com/internal/services/file"
	"musicproject.com/pkg/model"
)

type Song struct {
	store   file.Blobstore
	encoder encode.HLSEncoder
	repo    repository.Song
}

func New(store file.Blobstore, encoder encode.HLSEncoder, repo repository.Song) *Song {
	return &Song{store, encoder, repo}
}

// Uploads the song's metadata to the repository.
// Returns the url to upload the song file to.
func (s *Song) UploadMetadata(ctx context.Context,
	songRequest *model.UploadSongRequest) (string, error) {
	// Put song metadata in repository
	_, err := s.repo.Put(ctx, &model.Song{
		Name:  songRequest.Name,
		Genre: songRequest.Genre,
		Image: songRequest.Image,
	})
	if err != nil {
		return "", err
	}
	parent := "files/audio"
	url, err := s.store.CreateObjectUrl(ctx, parent, songRequest.Filename, true)
	if err != nil {
		return "", nil
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

//	func (s *Song) DownloadByID(ctx context.Context, songId uuid.UUID) error {
//		return nil
//	}
// func (s *Song) DeleteByID(ctx context.Context, songId uuid.UUID) error {
// 	return nil
// }
