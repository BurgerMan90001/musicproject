package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/pkg/model"
)

func handleArtists(repo *postgres.Artist) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {

			type NewArtistRequest struct {
				Name      string `json:"name"`
				AvatarUrl string `json:"avatarUrl"`
			}

			b, err := jsonutil.ReadJson[NewArtistRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			ctx := r.Context()
			artistId, err := repo.NewArtist(ctx, b.Name, b.AvatarUrl)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			l, err := apiJoinUrl("v1", "artists", artistId.String())
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", l)
			w.WriteHeader(http.StatusCreated)
		})

		r.Get("/{artistId}", func(w http.ResponseWriter, r *http.Request) {
			artistId, err := uuid.Parse(r.URL.Query().Get("artistId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			ctx := r.Context()
			a, err := repo.GetArtistById(ctx, artistId)
			if err != nil {
				jsonutil.WriteError(w, err)
			}

			jsonutil.WriteJSON(w, a, http.StatusOK)
		})
	}
}
