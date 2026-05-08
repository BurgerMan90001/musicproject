package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/pkg/model"
)

func handleAlbums() func(r chi.Router) {
	return func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/{albumId}", func(w http.ResponseWriter, r *http.Request) {
			_, err := uuid.Parse(r.PathValue("albumId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}

		})
	}
}
