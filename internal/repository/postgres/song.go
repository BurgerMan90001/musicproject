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

func (r *Song) NewSong(ctx context.Context, s *model.SongUploadRequest) (uuid.UUID, error) {
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
	s, err := r.q.GetSongById(ctx, songId)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}

	return &model.Song{
		SongId:       s.SongID,
		AlbumId:      s.AlbumID.UUID,
		Name:         s.SongName,
		Genres:       strings.Split(string(s.GenreList), ","),
		Artists:      strings.Split(string(s.ArtistList), ","),
		Duration:     int(s.Duration),
		CreationDate: s.CreationDate,
		Streams:      int(s.Streams),
		Cover:        s.SongCoverUrl.String,
		Audio:        s.SongAudioUrl,
	}, nil
}
func (r *Song) GetSongs(ctx context.Context, n int32) ([]*model.Song, error) {
	d, err := r.q.GetSongs(ctx, int32(n))
	if err != nil {
		return nil, err
	}
	var songs []*model.Song
	for _, s := range d {
		songs = append(songs, &model.Song{
			SongId:       s.SongID,
			AlbumId:      s.AlbumID.UUID,
			Name:         s.SongName,
			Genres:       strings.Split(string(s.GenreList), ","),
			Artists:      strings.Split(string(s.ArtistList), ","),
			Duration:     int(s.Duration),
			CreationDate: s.CreationDate,
			Streams:      int(s.Streams),
			Cover:        s.SongCoverUrl.String,
			Audio:        s.SongAudioUrl,
		})
	}
	return songs, nil
}

func (r *Song) GetSongsByGenre(ctx context.Context, genreName string) ([]*model.Song, error) {
	l, err := r.q.GetSongsByGenre(ctx, genreName)
	if err != nil {
		return nil, err
	}

	var songs []*model.Song
	for _, s := range l {
		songs = append(songs, toModelSong(
			s.SongID,
			s.AlbumID,
			s.SongName,
			s.GenreList,
			s.ArtistList,
			s.Duration,
			s.CreationDate,
			s.Streams,
			s.SongCoverUrl,
			s.SongAudioUrl,
		))
	}
	return songs, nil
}

func (r *Song) GetSongGenres(ctx context.Context, songId uuid.UUID) ([]*model.Genre, error) {

	l, err := r.q.GetSongGenres(ctx, songId)
	if err != nil {
		return nil, err
	}
	var genres []*model.Genre
	for _, g := range l {
		genres = append(genres, &model.Genre{
			GenreId: g.GenreID,
			Name:    g.GenreName,
		})
	}
	return genres, nil
}

// TODO
func (r *Song) GetSongRandom(ctx context.Context, n int) ([]*model.Song, error) {
	// res, err := r.q.
	// if err != nil {
	// 	return nil, err
	// }

	// var songs []*model.Song
	// for _, s := range res {
	// 	songs = append(songs, toModelSong(s))
	// }
	return nil, nil
}
func (r *Song) PutSongArtists(ctx context.Context, songId uuid.UUID, artistIds []uuid.UUID) error {
	return r.q.PutSongArtists(ctx, gensqlc.PutSongArtistsParams{
		ArtistIds: artistIds,
		SongID:    songId,
	})
}
func (r *Song) PutSongGenres(ctx context.Context, songId uuid.UUID, genreIds []uuid.UUID) error {
	return r.q.PutSongGenres(ctx, gensqlc.PutSongGenresParams{
		GenreIds: genreIds,
		SongID:   songId,
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
