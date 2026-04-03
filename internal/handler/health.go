package handler

import (
	"net/http"

	"musicproject.com/pkg/model"
)

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, model.Health{
		Message: "alive",
	}, http.StatusOK)
}
