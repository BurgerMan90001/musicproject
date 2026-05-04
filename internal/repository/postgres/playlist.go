package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"songsled.com/internal/repository"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Playlist struct {
	q *gensqlc.Queries
}

func NewPlaylist(q *gensqlc.Queries) *Playlist {
	return &Playlist{q}
}

func (r *Playlist) GetPlaylistByID(ctx context.Context, playlistId uuid.UUID) (*model.Playlist, error) {

	_, err := r.q.GetPlaylistByID(ctx, playlistId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
	}
	// l, err := r.q.GetPlaylistSongs(ctx, playlistId)
	// if err != nil {
	// 	if !errors.Is(err, sql.ErrNoRows) {
	// 		return nil, err
	// 	}
	// }
	// for _, s := range l {
	// 	// s.
	// }

	// r.q.GetPlaylistByID()
	return &model.Playlist{
		ID: playlistId,
		// Name: p.PlaylistName,
	}, nil
}

func (r *Playlist) PutPlaylist(ctx context.Context, p *model.Playlist) (uuid.UUID, error) {
	return uuid.Nil, nil
	// return r.q.PutPlaylist(ctx, p)
}
