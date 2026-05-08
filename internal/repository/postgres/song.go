package postgres

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/google/uuid"
	"songsled.com/internal/repository"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Song struct {
	q *gensqlc.Queries
}

func NewSongRepo(q *gensqlc.Queries) *Song {
	return &Song{q}
}

func (r *Song) NewSong(ctx context.Context, s *model.Song) (uuid.UUID, error) {
	songId, err := r.q.NewSong(ctx, gensqlc.NewSongParams{
		AlbumID:      uuid.NullUUID{UUID: s.AlbumID, Valid: s.AlbumID != uuid.Nil},
		SongName:     s.Name,
		Duration:     int32(s.Duration),
		CreationDate: s.CreationDate,
		SongAudioUrl: s.Audio,
	})

	if err != nil {
		return songId, err
	}
	return songId, nil
}
func (r *Song) GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Song, error) {
	s, err := r.q.GetSongByID(ctx, songId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	

	return &model.Song{
		SongID:       s.SongID,
		AlbumID:      s.AlbumID.UUID,
		Name:         s.SongName,
		Genres:       strings.Split(string(s.Genres), ","),
		Artists:      strings.Split(string(s.Artists), ","),
		Duration:     int(s.Duration),
		CreationDate: s.SongCreationDate,
		Streams:      int(s.Streams),
		Cover:        s.SongCoverUrl.String,
		Audio:        s.SongAudioUrl,
	}, nil
}
func (r *Song) GetSongs(ctx context.Context, n int) ([]*model.Song, error) {
	d, err := r.q.GetSongs(ctx, int32(n))
	if err != nil {
		return nil, err
	}
	var songs []*model.Song
	for _, s := range d {
		songs = append(songs, &model.Song{
			SongID:       s.SongID,
			AlbumID:      s.AlbumID.UUID,
			Name:         s.SongName,
			Genres:       strings.Split(string(s.GenreList), ","),
			Artists:      strings.Split(string(s.ArtistList), ","),
			Duration:     int(s.Duration),
			CreationDate: s.SongCreationDate,
			Streams:      int(s.Streams),
			Cover:        s.SongCoverUrl.String,
			Audio:        s.SongAudioUrl,
		})
	}
	return songs, nil
}

// TODO
func (r *Song) GetSongsByGenre(ctx context.Context, genre string) ([]*model.Song, error) {
	// r.q.GetSongsByGenre()

	return nil, nil
}
func (r *Song) GetSongRandom(ctx context.Context, n int) ([]*model.Song, error) {
	// res, err := r.q.GetRandomSongs(ctx, int32(n))
	// if err != nil {
	// 	return nil, err
	// }

	// var songs []*model.Song
	// for _, s := range res {
	// 	songs = append(songs, toModelSong(s))
	// }
	return nil, nil
}
func (r *Song) PutArtistSong(ctx context.Context, songId, artistId uuid.UUID) error {
	return r.q.PutArtistSong(ctx, gensqlc.PutArtistSongParams{
		ArtistID: artistId,
		SongID:   songId,
	})
}
func (r *Song) PutSongGenre(ctx context.Context, songId, genreId uuid.UUID) error {
	return r.q.PutSongGenre(ctx, gensqlc.PutSongGenreParams{
		GenreID: genreId,
		SongID:  songId,
	})
}

// Update a song's cover image
func (r *Song) PutSongCover(ctx context.Context, songId uuid.UUID, songCoverUrl string) error {
	r.q.PutSongCover(ctx, gensqlc.PutSongCoverParams{
		SongID:       songId,
		SongCoverUrl: sql.NullString{String: songCoverUrl, Valid: songCoverUrl != ""},
	})
	return nil
}
func (r *Song) Search(ctx context.Context) ([]*model.Song, error) {

	return nil, nil
}

func (r *Song) DeleteSongByID(ctx context.Context, id uuid.UUID) error {

	return nil
}
