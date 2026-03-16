package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
)

func HandleSongs(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			song, err := repo.GetSongByID(ctx, id)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				InternalServerError(w, r, err)
				return
			}
			WriteJSON(w, "success", song, http.StatusOK)

		case http.MethodPut:
			_, err := repo.PutSong(ctx, id, nil)
			if err != nil {
				InternalServerError(w, r, err)
			}
		default:
			MethodNotAllowedError(w, r)
		}
	}
}
