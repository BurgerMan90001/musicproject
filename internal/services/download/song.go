package download

import (
	"time"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/encode"
	"musicproject.com/internal/services/file"
)

type Song struct {
	parent  string
	store   file.Blobstore
	encoder encode.HLSEncoder
	repo    repository.Song
	// Should be around 5-15 minutes
	urlTtl time.Duration
}

func NewSong(urlTtl time.Duration) *Song {

	return &Song{urlTtl: urlTtl}
}

func (s *Song) DownloadByID(songId uuid.UUID) {

}
