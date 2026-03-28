package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services"
	"musicproject.com/pkg/model"
)

func HandleSongs(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			song, err := repo.GetSongByID(ctx, id)
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
			_, err := repo.PutSong(ctx, id, nil)
			if err != nil {
				InternalServerError(w, err)
				return
			}
		default:
			MethodNotAllowedError(w)
		}
	}
}

func HandleSongRating(ratingService services.Rating) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		songId, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			ratingService.GetAggregatedRating(ctx, songId)
		case http.MethodPut:
			//ratingService.PutRating(ctx, songId, uuid.Nil)

		default:
			MethodNotAllowedError(w)
		}
	}

}
func HandleSongUpload(fileService services.File) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()
		file, _, err := r.FormFile("")
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		defer file.Close()

		song := &model.Song{
			ID: uuid.Nil,
		}
		//handler.Header

		if err := fileService.UploadSong(ctx, song); err != nil {
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
