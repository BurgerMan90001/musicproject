package handler

import (
	"net/http"
	"time"

	"songsled.com/internal/jsonutil"
	"songsled.com/pkg/model"
)

func handleHealth(w http.ResponseWriter, _ *http.Request) {
	jsonutil.WriteJSON(w, model.HealthResponse{
		Message:   "alive",
		Timestamp: time.Now().UTC(),
	}, http.StatusOK)
}
