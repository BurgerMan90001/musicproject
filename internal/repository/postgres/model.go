package postgres

import (
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"songsled.com/pkg/model"
)

func toModelSong(
	songId uuid.UUID,
	albumId uuid.NullUUID,
	songName string,
	genreList []byte,
	artistList []byte,
	songCreationDate string,
	streams int32,
	songCoverUrl sql.NullString,
	songAudioUrl string,
) *model.Song {
	return &model.Song{
		SongId:       songId,
		Name:         songName,
		AlbumId:      albumId.UUID,
		Genres:       strings.Split(string(genreList), ","),
		Artists:      strings.Split(string(artistList), ","),
		CreationDate: songCreationDate,
		Streams:      int(streams),
		Cover:        songCoverUrl.String,
		Audio:        songAudioUrl,
	}
}

func toModelAlbum() *model.Album {
	return &model.Album{}
}
