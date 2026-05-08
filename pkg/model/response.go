package model

import "time"



type HealthResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timeStamp"`
}


type Link struct {
	Rel string `json:"rel"`
	Href string `json:"href"`
} 