package handler

import (
	"net/http"
	"time"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/pkg/model"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	jsonutil.WriteJSON(w, model.HealthResponse{
		Message:   "alive",
		Timestamp: time.Now().UTC(),
	}, http.StatusOK)
}
