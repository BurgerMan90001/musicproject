package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/upload"
	"songsled.com/pkg/model"
)

// images
func handleImages(uploadService *upload.Service) func(r chi.Router) {
	return func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			jsonutil.NotImplemented(w)
		})

		r.Get("/covers", func(w http.ResponseWriter, r *http.Request) {
			jsonutil.NotImplemented(w)
		})
		// Image upload
		r.Put("/covers", func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()
			filename := r.URL.Query().Get("filename")

			uploadLocation, objectLocation, err := uploadService.UploadUrl(ctx, "images/covers", filename, "image/*")
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
