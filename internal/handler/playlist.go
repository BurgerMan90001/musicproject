package handler

import (
	"context"
	"net/http"

	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/pkg/model"
)

type playlistRepo interface {
	GetPlaylistByID(ctx context.Context, playlistId uuid.UUID) (*model.Playlist, error)
}

func handlePlaylistsID(playlistRepo playlistRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
func handlePlaylists(playlistRepo playlistRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			playlistId, err := uuid.Parse(r.PathValue("id"))
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			ctx := r.Context()

			playlist, err := playlistRepo.GetPlaylistByID(ctx, playlistId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, playlist, http.StatusOK)
		case http.MethodPut:
		case http.MethodDelete:

		default:
		}

	}
}
