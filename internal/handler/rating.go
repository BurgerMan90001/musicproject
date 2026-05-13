package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/rating"
	"songsled.com/pkg/model"
)

func handleRatings(ratingService *rating.Service) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			songId, err := uuid.Parse(r.URL.Query().Get("songId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			ctx := r.Context()
			aggregated, err := ratingService.GetAggregatedRating(ctx, songId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, aggregated, http.StatusOK)

		})
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			rating, err := jsonutil.ReadJson[*model.Rating](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			if err := ratingService.Put(ctx, rating); err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, rating, http.StatusOK)
		})
	}
}
