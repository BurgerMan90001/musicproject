package postgres

import (
	"context"
	"strings"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Playlist struct {
	q *gensqlc.Queries
}

func NewPlaylistRepo(q *gensqlc.Queries) *Playlist {

	return &Playlist{q}
}
func (r *Playlist) NewPlaylist(ctx context.Context, name string, songIds []uuid.UUID) (uuid.UUID, error) {

	playlistId, err := r.q.NewPlaylist(ctx, gensqlc.NewPlaylistParams{
		UserID:       uuid.NullUUID{},
		PlaylistName: name,
	})
	if err != nil {
		return uuid.Nil, err
	}
	for _, id := range songIds {
		err := r.q.PutPlaylistSong(ctx, gensqlc.PutPlaylistSongParams{
			PlaylistID: playlistId,
			SongID:     id,
		})
		if err != nil {
			return uuid.Nil, err
		}
	}
	// r.q.getPlaylists

	return playlistId, nil
}

func (r *Playlist) GetPlaylists(ctx context.Context, n int32) ([]*model.Playlist, error) {
	l, err := r.q.GetPlaylists(ctx, n)
	if err != nil {
		return nil, err
	}
	var playlists []*model.Playlist
	for _, p := range l {
		playlists = append(playlists, &model.Playlist{
			PlaylistID: p.PlaylistID,
			UserID:     p.UserID.UUID,
			Name:       p.PlaylistName,
			Cover:      p.PlaylistCoverUrl.String,
		})
	}
	return playlists, nil
}

func (r *Playlist) GetPlaylistSongsById(ctx context.Context, playlistId uuid.UUID) ([]*model.Song, error) {
	l, err := r.q.GetPlaylistSongsById(ctx, gensqlc.GetPlaylistSongsByIdParams{
		PlaylistID: playlistId,
		Limit:      10,
	})
	if err != nil {
		return nil, err
	}

	var songs []*model.Song
	for _, s := range l {
		songs = append(songs, &model.Song{
			SongId:       s.SongID,
			AlbumId:      s.AlbumID.UUID,
			Name:         s.SongName,
			Genres:       strings.Split(string(s.GenreList), ","),
			Artists:      strings.Split(string(s.ArtistList), ","),
			CreationDate: s.CreationDate,
			Streams:      int(s.Streams),
			Cover:        s.SongCoverUrl.String,
			Audio:        s.SongAudioUrl,
		})
	}

	return songs, nil
}

func (r *Playlist) PutPlaylistSong(ctx context.Context, playlistId uuid.UUID,
	songId uuid.UUID) error {
	return r.q.PutPlaylistSong(ctx, gensqlc.PutPlaylistSongParams{
		PlaylistID: playlistId,
		SongID:     songId,
	})
}
