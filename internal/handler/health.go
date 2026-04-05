package handler

import (
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/pkg/model"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	jsonutil.WriteJSON(w, model.Health{
		Message: "alive",
	}, http.StatusOK)
}
