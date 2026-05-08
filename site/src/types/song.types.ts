export interface Song {
  id: string;
  albumId?: string;
  name: string;
  genres: string;
  artists: string;
  streams: number;
  duration: number;
  creationDate: string;
  image?: string;
  audio: string;
}
// SongID  uuid.UUID `json:"id"`
// 	AlbumID uuid.UUID `json:"albumId"`
// 	Name    string    `json:"name"`
// 	// Genres and artists are strings
// 	Genres       string `json:"genres"`
// 	Artists      string `json:"artists"`
// 	Duration     int    `json:"duration"`
// 	CreationDate string `json:"creationDate"`

// 	// Defaults to 0
// 	Streams int `json:"streams"`

// 	// Optional
// 	// Image url
// 	Image string `json:"image,omitempty"`
// 	// Required
// 	// Song file location
// 	Url string `json:"url"`
