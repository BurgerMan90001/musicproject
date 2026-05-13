package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/pkg/model"
)

func handleAlbums(repo *postgres.Album) func(r chi.Router) {
	return func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {})
		r.Get("/{albumId}", func(w http.ResponseWriter, r *http.Request) {
			albumId, err := uuid.Parse(r.PathValue("albumId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Album not found",
				})
				return
			}

			ctx := r.Context()

			repo.GetAlbum(ctx, albumId)
		})
		r.Get("/{albumId}/songs", func(w http.ResponseWriter, r *http.Request) {
			albumId, err := uuid.Parse(r.PathValue("albumId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Album not found",
				})
				return
			}

			ctx := r.Context()
			songs, err := repo.GetAlbumSongs(ctx, albumId, 30)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, songs, http.StatusOK)
		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			type req struct {
				AlbumName    string `json:"albumName"`
				CreationDate string `json:"creationDate"`
				CoverUrl     string `json:"coverUrl"`
			}
			ctx := r.Context()
			b, err := jsonutil.ReadJson[req](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			albumId, err := repo.NewAlbum(ctx, b.AlbumName, b.CreationDate, b.CoverUrl)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			l, err := apiJoinUrl("v1", "albums", albumId.String())
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", l)
			w.WriteHeader(http.StatusCreated)
		})

		r.Get("/{albumId}/songs", func(w http.ResponseWriter, r *http.Request) {
			albumId, err := uuid.Parse(r.PathValue("albumId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Album not found",
				})
				return
			}
			ctx := r.Context()
			songs, err := repo.GetAlbumSongs(ctx, albumId, 50)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, songs, http.StatusOK)
		})

	}
}
