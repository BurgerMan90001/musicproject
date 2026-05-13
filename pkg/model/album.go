package model

import "github.com/google/uuid"

type Album struct {
	AlbumId  uuid.UUID `json:"albumId"`
	Songs    []Song    `json:"songs"`
	ArtistID int       `json:"artistId"`
}
