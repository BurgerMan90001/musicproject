package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/rating"
	"musicproject.com/internal/services/upload"
	"musicproject.com/pkg/model"
)

func handleGetSongsMetadata(repo repository.Song) http.HandlerFunc {
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
			rating, err := jsonutil.ReadJson[*model.Rating](r.Body)
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

// Takes metadata first
// Then returns url for the cloud file upload or service file handler
func handleSongUpload(songService *upload.Song) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonutil.MethodNotAllowedError(w)
			return
		}
		// TODO context canceling for file uploads
		ctx := r.Context()
		contentType := r.Header.Get("Content-Type")

		switch {
		// Audio file upload request to local storage MAYBE
		case strings.HasPrefix(contentType, "audio/"):
			// TODO limit uploads
			file, header, err := r.FormFile("file")
			if err != nil {
				jsonutil.WriteError(w, err, http.StatusBadRequest)
				return
			}
			defer file.Close()
			if err := songService.UploadFile(ctx, file, header); err != nil {
				jsonutil.WriteError(w, err, http.StatusBadRequest)
				return
			}

		// Metadata upload request
		default:
			songRequest, err := jsonutil.ReadJson[*model.UploadSongRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err, http.StatusBadRequest)
				return
			}

			url, err := songService.UploadMetadata(ctx, songRequest)
			if err != nil {
				jsonutil.WriteError(w, err, http.StatusBadRequest, err)
				return
			}
			// Set the location where the file is going to be uploaded
			w.Header().Set("Location", url)
			w.WriteHeader(http.StatusOK)
		}
	}
}
