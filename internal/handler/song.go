package handler

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/upload"
	"songsled.com/pkg/model"
)

//	type songRepo interface {
//		GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error)
//		PutSong(ctx context.Context, s *model.Song) (uuid.UUID, error)
//	}
func handleSongs(repo *postgres.Song, uploadService *upload.Service) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			n, err := strconv.Atoi(r.URL.Query().Get("n"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusBadRequest,
					Message: "Missing query param n",
				})
				return
			}
			genre := r.URL.Query().Get("genre")
			if genre != "" {
				songs, err := repo.GetSongsByGenre(ctx, genre)
				if err != nil {
					jsonutil.WriteError(w, err)
					return
				}
				jsonutil.WriteJSON(w, songs, http.StatusOK)
				return
			}

			songs, err := repo.GetSongs(ctx, n)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, songs, http.StatusOK)
		})
		// Song upload
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			songRequest, err := jsonutil.ReadJson[*model.SongUploadRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			ctx := r.Context()
			if _, err := uploadService.UploadSongMetadata(ctx, songRequest); err != nil {
				jsonutil.WriteError(w, err)
				return
			}

		})
		r.Get("/{songId}", func(w http.ResponseWriter, r *http.Request) {
			id, err := uuid.Parse(r.PathValue("songId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			ctx := r.Context()
			song, err := repo.GetSongByID(ctx, id)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, song, http.StatusOK)
		})

		// Update song metadata
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {

		})

	}
}
