package model

import "github.com/google/uuid"

type Artist struct {
	ID     uuid.UUID `json:"id"`
	Albums []Album   `json:"albums"`
}

type Album struct {
	ID       uuid.UUID `json:"id"`
	Songs    []Song    `json:"songs"`
	ArtistID int       `json:"artistId"`
}

// A song's metadata
type Song struct {
	ID       uuid.UUID `json:"id"`
	Name     string    `json:"name"`
	Genre    string    `json:"genre"`
	Streams  int       `json:"streams"`
	Duration int       `json:"duration"`
	Image    string    `json:"image"`
	Source   string    `json:"src"`
}

type SongFile struct {
}

type UploadSongRequest struct {
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Image    string `json:"image,omitempty"`
	Filename string `json:"filename,omitempty"`
}

type Rating struct {
	SongID uuid.UUID `json:"songId"`
	UserID uuid.UUID `json:"userId"`
	Value  float64   `json:"value"`
}

type SongPlayer struct {
	SongID    int     `json:"songId"`
	TimeStamp float64 `json:"timeStamp"`
}
