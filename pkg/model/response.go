package model

import (
	"time"
)

type HealthResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timeStamp"`
}

type Link struct {
	Rel  string `json:"rel"`
	Href string `json:"href"`
}

type FileUploadResponse struct {
	Href  string `json:"href"`
	Links []Link `json:"links"`
}
