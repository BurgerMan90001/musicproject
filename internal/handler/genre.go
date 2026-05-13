package handler

import (
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/pkg/model"
)

func handleGenres(repo *postgres.Genre) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			genreName := r.URL.Query().Get("genreName")
			if genreName != "" {

			}
			genres, err := repo.GetGenres(ctx, 20)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, genres, http.StatusOK)
		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			type request struct {
				Name string `json:"genreName"`
			}
			b, err := jsonutil.ReadJson[request](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			genreId, err := repo.NewGenre(ctx, b.Name)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			location, err := url.JoinPath(os.Getenv("API_URL"), "v1", "genres", genreId.String())
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", location)
			w.WriteHeader(http.StatusSeeOther)
		})
		r.Get("/{genreId}", func(w http.ResponseWriter, r *http.Request) {
			genreId, err := uuid.Parse(r.URL.Query().Get("genreId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			ctx := r.Context()
			genreName, err := repo.GetGenreById(ctx, genreId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, model.Genre{
				GenreId: genreId,
				Name:    genreName,
			}, http.StatusOK)
		})
	}
}
