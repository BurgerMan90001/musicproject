package postgres

import (
	"context"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Song struct {
	q *gensqlc.Queries
}

func toSong(s gensqlc.Song) *model.Song {
	return &model.Song{
		SongID: s.SongID,
		// AlbumID:  s.AlbumID,
		// UserID:   s.UserID,

		Name:     s.SongName,
		Streams:  int(s.Streams),
		Duration: int(s.Duration),
		Image:    s.SongImage.String,
		URL:      s.SongUrl,
	}
}

func NewSong(q *gensqlc.Queries) *Song {
	return &Song{q}
}

func (r *Song) GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Song, error) {
	s, err := r.q.GetSongByID(ctx, songId)
	
	if err != nil {
		return nil, err
	}
	s.SongID = songId
	return toSong(s), nil
}

// func (r *Song) GetSongsByGenre(ctx context.Context, genre string) ([]*model.Song, error) {

//		return nil, nil
//	}
func (r *Song) GetSongRandom(ctx context.Context, n int) ([]model.Song, error) {
	res, err := r.q.GetSongRandom(ctx, int32(n))
	if err != nil {
		return nil, err
	}
	var songs []model.Song
	for _, s := range res {
		songs = append(songs, *toSong(s))

	}
	return songs, nil
}
func (r *Song) PutSong(ctx context.Context, s *model.Song) (uuid.UUID, error) {
	return uuid.Max, nil
}
func (r *Song) Search(ctx context.Context) ([]model.Song, error) {
	return nil, nil
}

func (r *Song) DeleteSongByID(ctx context.Context, id uuid.UUID) error {

	return nil
}
