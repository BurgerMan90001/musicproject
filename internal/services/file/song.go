package file

import (
	"context"
	"mime/multipart"

	"github.com/google/uuid"
)

type Song struct {
	store Blobstore
}

func NewSong(store Blobstore) *Song {
	return &Song{store}
}

func (s *Song) UploadSong(ctx context.Context, file multipart.File,
	header *multipart.FileHeader) error {

	// object := &Object{
	// 	name: song.Name,
	// }
	if err := s.store.CreateObject(ctx, "", "", nil, false, ""); err != nil {
		return nil
	}
	return nil
}
func (s *Song) Backup(ctx context.Context) error {
	return nil
}
func (s *Song) DownloadSongByID(ctx context.Context, songId uuid.UUID) error {
	return nil
}
func (s *Song) DeleteSongByID(ctx context.Context, songId uuid.UUID) error {
	return nil
}
