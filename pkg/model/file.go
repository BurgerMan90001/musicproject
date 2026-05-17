package model

type FileUploadResponse struct {
	Location string `json:"location"`
	Links    []Link `json:"links"`
}

type SongDownloadResponse struct {
	Song  *Song  `json:"song"`
	Links []Link `json:"links"`
}
