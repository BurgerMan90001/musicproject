package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/upload"
	"songsled.com/pkg/model"
)

func handleAudio(uploadService *upload.Service) func(r chi.Router) {
	return func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

			w.Header().Set("Location", "")
			w.WriteHeader(http.StatusFound)
		})

		r.Put("/", func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			filename := r.URL.Query().Get("filename")
			contentType := r.Header.Get("Content-Type")

			uploadLocation, objectLocation, err := uploadService.UploadUrl(ctx, "audio/", filename, contentType)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

			jsonutil.WriteJSON(w, &model.FileUploadResponse{
				Href: objectLocation,
				Links: []model.Link{
					{Rel: "upload", Href: uploadLocation},
				},
			}, http.StatusContinue)
		})
	}
}

// TODO
func handleAudioEncode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

/*
func handleAudio(songService *song.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//songId, err := uuid.Parse(r.PathValue("id"))
		// if err != nil {
		// 	jsonutil.WriteError(w, err, http.StatusBadRequest)
		// 	return
		// }
		//ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			//songService.GetByID(ctx, songId)
			//songService.
			w.Header().Set("Access-Control-Allow-Origin", "*")
			w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Range")

			ext := filepath.Ext(r.URL.Path)
			switch ext {
			case ".m3u8":
				w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
				w.Header().Set("Cache-Control", "no-cache")
			case ".ts":
				w.Header().Set("Content-Type", "video/mp2t")
				w.Header().Set("Cache-Control", "public, max-age=3600")
				//default:
			}
			//fs.ServeHTTP(w, r)
		case http.MethodPut:
		default:

			jsonutil.MethodNotAllowedError(w)
		}
	}
}
*/
