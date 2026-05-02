package handler

import (
	"context"
	"net/http"
	"strings"

	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/rating"
	"songsled.com/internal/services/search"
	"songsled.com/internal/services/upload"
	"songsled.com/pkg/model"
)

type songRepo interface {
	GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error)
	PutSong(ctx context.Context, s *model.Song) (uuid.UUID, error)
}

func handleSongs(searchService search.Service, repo songRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonutil.NotImplemented(w)
	}
}
func handleGetSong(repo songRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			jsonutil.WriteError(w, &model.Error{
				Code:    http.StatusNotFound,
				Message: "Song not found",
			})
			return
		}
		// params := r.URL.Query()
		// if params != nil {
		// 	jsonutil.WriteNotFound(w, errors.New(""))
		// 	return
		// }
		ctx := r.Context()

		song, err := repo.GetSongByID(ctx, id)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}
		jsonutil.WriteJSON(w, song, http.StatusOK)

	}
}
func handlePutSongsMetadata(repo songRepo) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		_, err := repo.PutSong(ctx, nil)
		if err != nil {
			jsonutil.InternalServerError(w, err)
			return
		}
	}
}
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

// Takes metadata first
// Then returns url for the cloud file upload or service file handler
func handleSongUpload(songService *upload.Song) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// TODO context canceling for file uploads
		ctx := r.Context()
		contentType := r.Header.Get("Content-Type")

		switch {
		// Audio file upload request to local storage MAYBE
		case strings.HasPrefix(contentType, "audio/"):
			// TODO limit uploads
			file, header, err := r.FormFile("file")
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			defer file.Close()
			if err := songService.UploadFile(ctx, file, header); err != nil {
				jsonutil.WriteError(w, err)
				return
			}

		// Metadata upload request
		default:
			songRequest, err := jsonutil.ReadJson[*model.UploadSongRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			url, err := songService.UploadMetadata(ctx, songRequest)
			if err != nil {
				jsonutil.WriteError(w, err)
				return

			}
			// Set the location where the file is going to be uploaded
			w.Header().Set("Location", url)
			w.WriteHeader(http.StatusOK)
		}
	})
}

// func handleSongSearch() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 	}
// }
