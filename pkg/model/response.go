package model

import "time"

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message,omitempty"`
	Details []string `json:"details,omitempty"`
}

type HealthResponse struct {
	Message   string    `json:"message"`
	Timestamp time.Time `json:"timeStamp"`
}
