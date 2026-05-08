package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/pkg/model"
)

//	type playlistRepo interface {
//		GetPlaylistByID(ctx context.Context, playlistId uuid.UUID) (*model.Playlist, error)
//		PutPlaylist(ctx context.Context, p *model.Playlist) (uuid.UUID, error)
//	}
func handlePlaylists(playlistRepo *postgres.Playlist) func(r chi.Router) {
	return func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})

		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			res, err := jsonutil.ReadJson[*model.NewPlaylistRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusBadRequest,
					Message: "Playlist not found",
					Details: err.Error(),
				})
				return
			}
			if _, err := playlistRepo.NewPlaylist(ctx, res.Name, res.SongsIDs); err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", "")
			jsonutil.WriteJSON(w, nil, http.StatusCreated)
		})
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			playlistId, err := uuid.Parse(r.PathValue("id"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusBadRequest,
					Message: "Playlist not found",
					Details: err.Error(),
				})
				return
			}
			ctx := r.Context()
			songs, err := playlistRepo.GetPlaylistSongsByID(ctx, playlistId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, songs, http.StatusOK)
		})

	}
}
