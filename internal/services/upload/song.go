package upload

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/google/uuid"
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

	// Required
	songRepo   *postgres.Song
	genreRepo  *postgres.Genre
	artistRepo *postgres.Artist
}

// type songRepo interface {
// 	NewSong(ctx context.Context, s *model.Song) (uuid.UUID, error)
// }

func New(bucket string,
	// encoding,
	caching bool,
	urlTtl time.Duration,
	store file.Blobstore,
	songRepo *postgres.Song,
	genreRepo *postgres.Genre,
	artistRepo *postgres.Artist,
) (*Service, error) {
	var encoder encode.HLSEncoder
	// if encoding {
	// 	encoder = encode.NewFFmpeg(config.Encoder{})
	// }
	if bucket == "" {
		return nil, errors.New("Upload.NewSong: bucket is empty")
	}
	// if audioFolder == "" {
	// 	return nil, errors.New("Upload.NewSong:  audioFolder is empty")
	// }

	return &Service{bucket, caching, urlTtl,
		store, encoder, songRepo, genreRepo, artistRepo}, nil
}

// Uploads the song's metadata to the repository.
// Returns the url to upload the song file to.
func (s *Service) UploadSongMetadata(ctx context.Context,
	req *model.SongUploadRequest) (uuid.UUID, error) {

	if err := validateUploadRequest(req); err != nil {
		return uuid.Nil, err
	}

	if req.Audio == "" {
		return uuid.Nil, &model.Error{
			Code:    http.StatusBadRequest,
			Message: "Missing link to audio file",
		}
	}
	// Put song metadata in repository
	songId, err := s.songRepo.NewSong(ctx, req)
	if err != nil {
		return uuid.Nil, err
	}
	// TODO OPTIMIZE DO WITHIN SINGLE QUERY OR SOMETHING
	// Add songs and genres
	var artistIds []uuid.UUID
	for _, name := range req.Artists {
		artist, err := s.artistRepo.GetArtistByName(ctx, name)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			id, err := s.artistRepo.NewArtist(ctx, name, "")
			if err != nil {
				return uuid.Nil, err
			}
			artistIds = append(artistIds, id)
		} else if err != nil {
			return uuid.Nil, err
		} else {
			artistIds = append(artistIds, artist.ArtistId)
		}

	}
	if err := s.songRepo.PutSongArtists(ctx, songId, artistIds); err != nil {
		return uuid.Nil, err
	}

	var genreIds []uuid.UUID
	for _, name := range req.Genres {
		genreId, err := s.genreRepo.GetGenreByName(ctx, name)
		if err != nil && errors.Is(err, sql.ErrNoRows) {
			id, err := s.genreRepo.NewGenre(ctx, name)
			if err != nil {
				return uuid.Nil, err
			}
			genreIds = append(genreIds, id)

		} else if err != nil {
			return uuid.Nil, err
		} else {
			genreIds = append(genreIds, genreId)
		}

	}
	if err := s.songRepo.PutSongGenres(ctx, songId, genreIds); err != nil {
		return uuid.Nil, err
	}
	if req.Cover != "" {
		if err := s.songRepo.PutSongCover(ctx, songId, req.Cover); err != nil {
			return uuid.Nil, err
		}
	}

	return songId, nil
}

// Returns the image upload url for the audio directory
func (s *Service) UploadUrl(ctx context.Context, folder, filename, contentType string) (string, string, error) {

	if filename == "" {
		filename = crand.NewShort()
	}

	// Add additional random characters to avoid collisions
	key := filepath.Join(folder, fmt.Sprintf("%s-%s", crand.NewShort(), filename))
	return s.store.CreateObjectUrl(ctx, s.bucket, key, s.caching, s.urlTtl, contentType)
}
