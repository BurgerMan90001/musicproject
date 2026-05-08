package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/upload"
)

// images
func handleImages(uploadService *upload.Service) func(r chi.Router) {
	return func(r chi.Router) {

		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			jsonutil.NotImplemented(w)
		})

		// r.Get("/covers", func(w http.ResponseWriter, r *http.Request) {
		// 	jsonutil.NotImplemented(w)
		// })
		// Image upload
		r.HandleFunc("/covers", func(w http.ResponseWriter, r *http.Request) {

			ctx := r.Context()

			// type body struct {
			// 	songId uuid.UUID `json:"songId"`
			// }
			// b, err := jsonutil.ReadJson[body](r.Body)
			// if err != nil {
			// 	jsonutil.WriteError(w, err)
			// 	return
			// }
			location, _, err := uploadService.UploadImageUrl(ctx, "covers")

			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			w.Header().Set("Location", location)
			w.WriteHeader(http.StatusTemporaryRedirect)
		})

	}
}
