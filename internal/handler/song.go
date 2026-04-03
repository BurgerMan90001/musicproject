package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/file"
	"musicproject.com/internal/services/rating"
	"musicproject.com/pkg/model"
)

func HandleSongs(repo repository.Song) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
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
				WriteError(w, ErrInternalServerError, http.StatusInternalServerError)
				return
			}
			WriteJSON(w, song, http.StatusOK)

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
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			aggregated, err := ratingService.GetAggregatedRating(ctx, songId)
			if err != nil {
				WriteError(w, err, http.StatusInternalServerError)
				return
			}
			WriteJSON(w, aggregated, http.StatusOK)
		case http.MethodPut:
			rating, err := model.ReadJSON[model.Rating](r.Body)
			if err != nil {
				WriteError(w, err, http.StatusBadRequest)
				return
			}
			if err := ratingService.PutRating(ctx, songId, uuid.Nil, 0); err != nil {
				WriteError(w, err, http.StatusInternalServerError)
				return
			}

			WriteJSON(w, rating, http.StatusOK)

		default:
			MethodNotAllowedError(w)
		}
	}

}
func HandleSongUpload(fileService *file.Song) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		file, handler, err := r.FormFile("file")
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		defer file.Close()

		ctx := r.Context()

		if err := fileService.UploadSong(ctx, file, handler); err != nil {
			WriteError(w, err, http.StatusOK)
			return
		}
		
		// repo.PutSong(ctx, uuid.Nil, &model.Song{
		// 	ID:     uuid.Nil,
		// 	Source: "",
		// })

		// Save file
	}
}
