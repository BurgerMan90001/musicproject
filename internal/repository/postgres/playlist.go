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

	return playlistId, nil
}

func (r *Playlist) GetPlaylistSongsByID(ctx context.Context, playlistId uuid.UUID) ([]model.Song, error) {
	playlistSongs, err := r.q.GetPlaylistSongsByID(ctx, gensqlc.GetPlaylistSongsByIDParams{
		PlaylistID: playlistId,
		Limit:      10,
	})
	if err != nil {
		return nil, err
	}

	var songs []model.Song
	for _, s := range playlistSongs {
		songs = append(songs, model.Song{
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
