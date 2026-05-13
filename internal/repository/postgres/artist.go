package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Artist struct {
	q *gensqlc.Queries
}

func NewArtistRepo(q *gensqlc.Queries) *Artist {
	return &Artist{q}
}

func (r *Artist) NewArtist(ctx context.Context, artistName, coverUrl string) (uuid.UUID, error) {
	return r.q.NewArtist(ctx, gensqlc.NewArtistParams{
		ArtistName:      artistName,
		ArtistAvatarUrl: sql.NullString{String: coverUrl, Valid: coverUrl != ""},
	})
}

func (r *Artist) GetArtistSongs(ctx context.Context, artistId uuid.UUID, n int32) ([]*model.Song, error) {
	l, err := r.q.GetArtistSongs(ctx, gensqlc.GetArtistSongsParams{
		ArtistID: artistId,
		Limit:    n,
	})
	if err != nil {
		return nil, fmt.Errorf("Get artist songs: %w", err)
	}
	var songs []*model.Song
	for _, s := range l {
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

func (r *Artist) GetArtistById(ctx context.Context, artistId uuid.UUID) (*model.Artist, error) {
	a, err := r.q.GetArtistById(ctx, artistId)
	if err != nil {
		return nil, fmt.Errorf("Get artist by id: %w", err)
	}
	return &model.Artist{ArtistId: artistId, Name: a.ArtistName, Avatar: a.ArtistAvatarUrl.String}, nil
}
func (r *Artist) GetArtistByName(ctx context.Context, artistName string) (*model.Artist, error) {
	a, err := r.q.GetArtistByName(ctx, artistName)
	if err != nil {
		return nil, fmt.Errorf("Get artist by name: %w", err)
	}

	return &model.Artist{ArtistId: a.ArtistID, Name: artistName}, nil
}

// TODO
func (r *Artist) GetArtists() {
	// r.q.GetArt
}
