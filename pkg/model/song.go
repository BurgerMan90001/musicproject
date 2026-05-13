package model

import (
	"github.com/google/uuid"
)

type Artist struct {
	ArtistId uuid.UUID `json:"artistId"`
	Name     string    `json:"name"`
	Avatar   string    `json:"avatar"`
}

type Genre struct {
	GenreId uuid.UUID `json:"genreId"`
	Name    string    `json:"name"`
}

// A song's metadata
type Song struct {
	// Required
	SongId  uuid.UUID `json:"songId"`
	AlbumId uuid.UUID `json:"albumId"`
	Name    string    `json:"name"`

	// Genres and artists are strings
	Genres       []string `json:"genres"`
	Artists      []string `json:"artists"`
	Duration     int      `json:"duration"`
	CreationDate string   `json:"creationDate"`

	// Defaults to 0
	Streams int `json:"streams"`

	// Optional
	// Cover url
	Cover string `json:"cover,omitempty"`
	// Required
	// Song file location
	Audio string `json:"audio"`
}

// YYYYMMDD = "2006-01-02"
// type ReleaseDate struct {
// 	Year  int
// 	Month int
// 	Day   int
// }

// Upload song metadata
type SongUploadRequest struct {
	// Required
	Name    string   `json:"name"`
	Artists []string `json:"artists"`

	// Optional
	Genres []string `json:"genres"`

	Duration int `json:"duration"`

	// YYYY-MM-DD format
	CreationDate string `json:"creationDate"`

	// Optional
	AlbumID uuid.UUID `json:"albumId"`

	// Image and audio locations
	Cover string `json:"cover,omitempty"`
	Audio string `json:"audio"`
}

// type Link struct {
// 	Rel  string `json:"rel"`
// 	Href string `json:"href"`
// }

// type NewSong struct {
// 	Name     string      `json:"name"`
// 	Artists  []uuid.UUID `json:"artists"`
// 	Duration int         `json:"duration"`
// 	// Optional
// 	Genres []uuid.UUID `json:"genres"`
// 	// YYYY-MM-DD format
// 	CreationDate string `json:"creationDate"`

// 	// Optional
// 	AlbumID uuid.UUID `json:"albumId"`

// 	// Optional
// 	// Image string `json:"image,omitempty"`
// 	// Url   string `json:"url,omitempty"`
// }

// type newSong struct {
// 	name         string
// 	albumId      uuid.UUID
// 	duration     int
// 	creationDate string
// 	image        string
// 	url          string
// }

type Playlist struct {
	PlaylistID uuid.UUID `json:"playlistId"`
	UserID     uuid.UUID `json:"userId"`
	Name       string    `json:"name"`
	// Optional
	// Cover url for the playlist
	Cover string `json:"cover,omitempty"`
	Songs []Song `json:"songs"`
}
type NewPlaylistRequest struct {
	Name     string      `json:"name"`
	SongsIDs []uuid.UUID `json:"songIds"`
}
type Rating struct {
	SongID uuid.UUID `json:"songId"`
	UserID uuid.UUID `json:"userId"`
	Value  float64   `json:"value"`
}

// type SongPlayer struct {
// 	SongID    int     `json:"songId"`
// 	TimeStamp float64 `json:"timeStamp"`
// }
