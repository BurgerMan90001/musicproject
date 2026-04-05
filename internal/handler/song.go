package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/rating"
	"musicproject.com/internal/services/song"
	"musicproject.com/pkg/model"
)

func HandleSongsMetadata(repo repository.Song) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		// claims, ok := contextClaims(ctx)

		switch r.Method {
		case http.MethodGet:
			song, err := repo.GetByID(ctx, id)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					http.Error(w, err.Error(), http.StatusNotFound)
					return
				}
				InternalServerError(w, err)
				return
			}
			jsonutil.WriteJSON(w, song, http.StatusOK)

		case http.MethodPut:
			_, err := repo.Put(ctx, nil)
			if err != nil {
				InternalServerError(w, err)
				return
			}
		default:
			MethodNotAllowedError(w)
		}
	}
}

func HandleSongRating(ratingService *rating.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		songId, err := uuid.Parse(r.PathValue("songId"))
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			aggregated, err := ratingService.GetAggregatedRating(ctx, songId)
			if err != nil {
				jsonutil.WriteError(w, err, http.StatusInternalServerError)
				return
			}
			jsonutil.WriteJSON(w, aggregated, http.StatusOK)
		case http.MethodPut:
			rating, err := jsonutil.ReadJSON[model.Rating](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err, http.StatusBadRequest)
				return
			}
			if err := ratingService.PutRating(ctx, songId, uuid.Nil, 0); err != nil {
				jsonutil.WriteError(w, err, http.StatusInternalServerError)
				return
			}

			jsonutil.WriteJSON(w, rating, http.StatusOK)

		default:
			MethodNotAllowedError(w)
		}
	}
}

func HandleSongUpload(songService *song.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		defer file.Close()

		songRequest, err := jsonutil.ReadJSON[model.UploadSongRequest](r.Body)

		ctx := r.Context()

		song, err := songService.UploadSong(ctx, file, handler, songRequest)
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusOK)
			return
		}
		jsonutil.WriteJSON(w, song, http.StatusCreated)
	}
}
