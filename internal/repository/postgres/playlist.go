package postgres

import (
	"context"

	"github.com/google/uuid"
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
	// TODO
	// r.q.GetPlaylistByID()
	return nil, nil
}
