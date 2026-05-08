package handler

import (
	"net/http"

	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/rating"
	"songsled.com/pkg/model"
)

func handleSongRating(ratingService *rating.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		songId, err := uuid.Parse(r.PathValue("songId"))
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			aggregated, err := ratingService.GetAggregatedRating(ctx, songId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, aggregated, http.StatusOK)
		case http.MethodPut:
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

		default:
		}
	}
}
