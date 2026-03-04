package model

type Artist struct {
	UUID   `json:"id"`
	Albums []Album `json:"albums"`
}

type Album struct {
	ID       UUID   `json:"id"`
	Songs    []Song `json:"songs"`
	ArtistID int    `json:"artistId"`
}

type Song struct {
	ID       UUID   `json:"id"`
	Name     string `json:"name"`
	Duration int    `json:"duration"`
	Image    string `json:"image"`
}

type SongPlayer struct {
	SongID    int `json:"songId"`
	TimeStamp int `json:"timeStamp"`
}

type Item struct {
	Container string `json:"container"`
}
