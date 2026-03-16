package model

import "encoding/json"

type Response struct {
	Status  string          `json:"status"`
	Data    json.RawMessage `json:"data,omitempty"`
	Message string          `json:"message,omitempty"`
}

// func Marshal(w http.ResponseWriter, v any) error {

// }
