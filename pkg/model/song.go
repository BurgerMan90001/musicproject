package model

import (
	"github.com/google/uuid"
)

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
	// Required
	SongID  uuid.UUID `json:"id"`
	AlbumID uuid.UUID `json:"albumId"`
	UserID  uuid.UUID `json:"userId"`
	Name    string    `json:"name"`
	Genre   string    `json:"genre"`

	Duration    int    `json:"duration"`
	ReleaseDate string `json:"releaseDate"`
	// Defaults to 0
	Streams int `json:"streams"`

	// Optional
	// Image url
	Image string `json:"image,omitempty"`
	// Required
	// Song file location
	URL string `json:"url"`
}

// YYYYMMDD = "2006-01-02"
// type ReleaseDate struct {
// 	Year  int
// 	Month int
// 	Day   int
// }

// Upload song metadata
type UploadSongRequest struct {
	// Required
	Name     string `json:"name"`
	Genre    string `json:"genre"`
	Filename string `json:"filename"`
	// Optional
	Image string `json:"image,omitempty"`
}

type Playlist struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Songs []Song    `json:"songs"`
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
