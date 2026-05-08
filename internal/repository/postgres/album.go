package postgres

import (
	"context"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
)

type Album struct {
	q *gensqlc.Queries
}

func NewAlbumRepo(q *gensqlc.Queries) *Album {
	return &Album{q}
}
func (r *Album) NewAlbum(ctx context.Context, albumName, creationDate string) (uuid.UUID, error) {
	return r.q.NewAlbum(ctx, gensqlc.NewAlbumParams{
		AlbumName:    albumName,
		CreationDate: creationDate,
	})
}
