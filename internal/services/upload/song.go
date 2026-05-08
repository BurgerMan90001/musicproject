package upload

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/google/uuid"
	"songsled.com/internal/config"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/encode"
	"songsled.com/internal/services/file"
	"songsled.com/pkg/crand"
	"songsled.com/pkg/model"
)

type Service struct {
	// Required
	bucket string
	// Optional
	caching bool
	// Should be 1 hour or less
	urlTtl time.Duration
	// Required
	// File storage
	store file.Blobstore
	// Optional
	// Encodes files before storing
	encoder encode.HLSEncoder

	songRepo *postgres.Song
}

// type songRepo interface {
// 	NewSong(ctx context.Context, s *model.Song) (uuid.UUID, error)
// }

func New(bucket string,
	encoding, caching bool,
	urlTtl time.Duration,
	store file.Blobstore,
	songRepo *postgres.Song,
) (*Service, error) {
	var encoder encode.HLSEncoder
	if encoding {
		encoder = encode.NewFFmpeg(config.Encoder{})
	}
	if bucket == "" {
		return nil, errors.New("Upload.NewSong: bucket is empty")
	}
	// if audioFolder == "" {
	// 	return nil, errors.New("Upload.NewSong:  audioFolder is empty")
	// }

	return &Service{bucket, caching, urlTtl,
		store, encoder, songRepo}, nil
}

// Uploads the song's metadata to the repository.
// Returns the url to upload the song file to.
func (s *Service) UploadSongMetadata(ctx context.Context,
	req *model.SongUploadRequest) (uuid.UUID, error) {

	if err := validateUploadRequest(req); err != nil {
		return uuid.Nil, err
	}

	//https://storage.googleapis.com/spume-musicproject/audio/24e692e5-mysong.mp3

	if req.Audio == "" {
		return uuid.Nil, &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Missing link to audio file",
		}
	}
	song := &model.Song{
		Name: req.Name,
		// Artists:      req.Artists,
		Duration: req.Duration,
		// Genres:       req.Genres,
		CreationDate: req.CreationDate,

		Audio: req.Audio,
	}
	if req.Cover != "" {
		song.Cover = req.Cover
	}

	// Put song metadata in repository
	songId, err := s.songRepo.NewSong(ctx, song)
	if err != nil {
		return uuid.Nil, err
	}
	// TODO OPTIMIZE DO WITHIN SINGLE QUERY
	// Add songs and genres
	for _, artistId := range req.Artists {
		err := s.songRepo.PutArtistSong(ctx, songId, artistId)
		if err != nil {
			return uuid.Nil, err
		}
	}
	for _, genreId := range req.Genres {
		err := s.songRepo.PutSongGenre(ctx, songId, genreId)
		if err != nil {
			return uuid.Nil, err
		}
	}

	return songId, nil
}

// Returns the image upload url for the audio directory
func (s *Service) UploadUrl(ctx context.Context, folder, filename, contentType string) (string, string, error) {

	// TODO 
	if filename == "" {
		filename = crand.NewShort()
	}
	
	// Add additional random characters to avoid collisions
	key := filepath.Join(folder, fmt.Sprintf("%s-%s", crand.NewShort(), filename))
	return s.store.CreateObjectUrl(ctx, s.bucket, key, s.caching, s.urlTtl, contentType)
}

// Returns the image upload url for the images directory
// func (s *Service) UploadImageUrl(ctx context.Context, folder, filename string) (string, string, error) {
// 	image := filepath.Join("images", folder, fmt.Sprintf("%s-%s", crand.NewShort(), filename))
// 	return s.store.CreateObjectUrl(ctx, s.bucket, image, s.caching, s.urlTtl, "images/*")
// }

// Local uploads of audio files
// Handles multipart file uploads if the service uses local uploads
// func (s *Song) UploadFile(ctx context.Context, file multipart.File,
// 	header *multipart.FileHeader) error {

// 	contents, err := io.ReadAll(file)
// 	if err != nil {
// 		return err
// 	}
// 	// Check if file is an audio file
// 	contentType := http.DetectContentType(contents)
// 	if !strings.HasPrefix(contentType, "audio/") {
// 		return fmt.Errorf("Incorrect mime type: %v", contentType)
// 	}
// 	return nil
// }
