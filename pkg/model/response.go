package model

type Error struct {
	Code    int      `json:"code"`
	Message string   `json:"message,omitempty"`
	Details []string `json:"details,omitempty"`
}

type Health struct {
	Message string `json:"message,omitempty"`
}
