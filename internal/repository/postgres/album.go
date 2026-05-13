package postgres

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Album struct {
	q *gensqlc.Queries
}

func NewAlbumRepo(q *gensqlc.Queries) *Album {
	return &Album{q}
}
func (r *Album) NewAlbum(ctx context.Context, albumName, creationDate, coverUrl string) (uuid.UUID, error) {
	return r.q.NewAlbum(ctx, gensqlc.NewAlbumParams{
		AlbumName:     albumName,
		CreationDate:  creationDate,
		AlbumCoverUrl: sql.NullString{String: coverUrl, Valid: coverUrl != ""},
	})
}
func (r *Album) GetAlbum(ctx context.Context, albumId uuid.UUID) {

}

func (r *Album) GetAlbumSongs(ctx context.Context, albumId uuid.UUID, n int32) ([]*model.Song, error) {
	l, err := r.q.GetAlbumSongs(ctx, gensqlc.GetAlbumSongsParams{
		AlbumID: uuid.NullUUID{UUID: albumId, Valid: albumId != uuid.Nil},
		Limit:   n,
	})
	if err != nil {
		return nil, err
	}
	var songs []*model.Song
	for _, s := range l {
		songs = append(songs, &model.Song{
			SongId:  s.SongID,
			AlbumId: s.AlbumID.UUID,
			Name:    s.SongName,
			Genres:  []string{},
		})
	}

	return songs, nil
}
func (r *Album) PutAlbumArtist(ctx context.Context, albumId, artistId uuid.UUID) error {
	return r.q.PutAlbumArtist(ctx, gensqlc.PutAlbumArtistParams{
		AlbumID:  albumId,
		ArtistID: artistId,
	})
}

func (r *Album) PutAlbumSong(ctx context.Context) {

}
