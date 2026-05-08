package postgres

import (
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

func toModelSong(s gensqlc.Song) *model.Song {
	return &model.Song{
		SongID: s.SongID,
		AlbumID: s.AlbumID.UUID,
		
	}
}

func toModelAlbum(a gensqlc.Album) *model.Album {
	return &model.Album{}
}
