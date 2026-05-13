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

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			p, err := playlistRepo.GetPlaylists(ctx, 20)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, p, http.StatusOK)
		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
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
			playlistId, err := playlistRepo.NewPlaylist(ctx, res.Name, res.SongsIDs)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			l, err := apiJoinUrl("v1", "playlists", playlistId.String())
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", l)
			w.WriteHeader(http.StatusCreated)
		})
		r.Get("/{playlistId}", func(w http.ResponseWriter, r *http.Request) {
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
			songs, err := playlistRepo.GetPlaylistSongsById(ctx, playlistId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, songs, http.StatusOK)
		})
		r.Put("/{playlistId}", func(w http.ResponseWriter, r *http.Request) {
			_, err := uuid.Parse(r.PathValue("id"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusBadRequest,
					Message: "Playlist not found",
					Details: err.Error(),
				})
				return
			}
			jsonutil.NotImplemented(w)
		})
		r.Put("/{playlistId}/songs", func(w http.ResponseWriter, r *http.Request) {
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

			type request struct {
				SongId uuid.UUID `json:"songId"`
			}
			b, err := jsonutil.ReadJson[request](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			if err := playlistRepo.PutPlaylistSong(ctx, playlistId, b.SongId); err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			l, err := apiJoinUrl("v1", "playlists", playlistId.String(), "songs")
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", l)
			w.WriteHeader(http.StatusCreated)
		})
	}
}
