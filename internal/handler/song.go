package handler

import (
	"context"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/rating"
	"musicproject.com/pkg/model"
)

func handleSongsMetadata(repo repository.Song) http.HandlerFunc {
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
				jsonutil.InternalServerError(w, err)
				return
			}
			jsonutil.WriteJSON(w, song, http.StatusOK)

		case http.MethodPut:
			_, err := repo.Put(ctx, nil)
			if err != nil {
				jsonutil.InternalServerError(w, err)
				return
			}
		default:
			jsonutil.MethodNotAllowedError(w)
		}
	}
}

func handleSongRating(ratingService *rating.Service) http.HandlerFunc {
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
			rating, err := jsonutil.ReadJSON[*model.Rating](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err, http.StatusBadRequest)
				return
			}
			if err := ratingService.Put(ctx, rating); err != nil {
				jsonutil.WriteError(w, err, http.StatusInternalServerError)
				return
			}

			jsonutil.WriteJSON(w, rating, http.StatusOK)

		default:
			jsonutil.MethodNotAllowedError(w)
		}
	}
}

type songService interface {
	UploadSong(ctx context.Context,
		file multipart.File, handler *multipart.FileHeader,
		songRequest model.UploadSongRequest) (*model.Song, error)
}

func handleSongUpload(songService songService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonutil.MethodNotAllowedError(w)
			return
		}
		// TODO limit uploads
		file, handler, err := r.FormFile("file")
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		defer file.Close()

		songRequest, err := jsonutil.ReadJSON[model.UploadSongRequest](r.Body)
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		song, err := songService.UploadSong(ctx, file, handler, songRequest)
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusOK)
			return
		}
		jsonutil.WriteJSON(w, song, http.StatusCreated)
	}
}
