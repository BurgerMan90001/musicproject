package handler

import (
	"net/http"
	"net/url"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/upload"
	"songsled.com/pkg/model"
)

//	type songRepo interface {
//		GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error)
//		PutSong(ctx context.Context, s *model.Song) (uuid.UUID, error)
//	}
func handleSongs(songRepo *postgres.Song, genreRepo *postgres.Genre, uploadService *upload.Service) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()

			// var n int32 = 10
			// sn := r.URL.Query().Get("n")
			// if sn != "" {
			// 	n, err : = strconv.ParseInt(sn, 10, 32)
			// 	if err != nil {
			// 		jsonutil.WriteError(w, &model.Error{
			// 			Code:    http.StatusBadRequest,
			// 			Message: "Missing query param n",
			// 		})
			// 		return
			// 	}
			// }

			genre := r.URL.Query().Get("genre")
			if genre != "" {
				songs, err := songRepo.GetSongsByGenre(ctx, genre)
				if err != nil {
					jsonutil.WriteError(w, err)
					return
				}
				jsonutil.WriteJSON(w, songs, http.StatusOK)
				return
			}

			songs, err := songRepo.GetSongs(ctx, 10)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, songs, http.StatusOK)
		})
		// Song upload
		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			songRequest, err := jsonutil.ReadJson[*model.SongUploadRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			ctx := r.Context()
			songId, err := uploadService.UploadSongMetadata(ctx, songRequest)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			l, err := url.JoinPath(os.Getenv("API_URL"), "v1", "songs", songId.String())
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", l)
			w.WriteHeader(http.StatusCreated)

		})
		r.Get("/{songId}", func(w http.ResponseWriter, r *http.Request) {
			songId, err := uuid.Parse(r.PathValue("songId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			ctx := r.Context()
			song, err := songRepo.GetSongByID(ctx, songId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			jsonutil.WriteJSON(w, song, http.StatusOK)
		})
		r.Get("/{songId}/genres", func(w http.ResponseWriter, r *http.Request) {
			songId, err := uuid.Parse(r.PathValue("songId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			ctx := r.Context()
			genres, err := songRepo.GetSongGenres(ctx, songId)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, genres, http.StatusOK)
		})
		r.Put("/{songId}/genres", func(w http.ResponseWriter, r *http.Request) {
			songId, err := uuid.Parse(r.PathValue("songId"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "Song not found",
				})
				return
			}
			type request struct {
				GenreIds []uuid.UUID `json:"genreIds"`
			}
			b, err := jsonutil.ReadJson[request](r.Body)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			ctx := r.Context()

			if err := songRepo.PutSongGenres(ctx, songId, b.GenreIds); err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, nil, http.StatusNoContent)
		})

	}
}
