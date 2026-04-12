package handler

import (
	"net/http"
	"path/filepath"

	"musicproject.com/internal/jsonutil"
)

// func HandleAudio(store file.Blobstore) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		//ctx := r.Context()

// 		// song, err := store.GetObject(ctx, "", "")
// 		// if err != nil {
// 		// 	jsonutil.WriteError(w, err, http.StatusBadRequest)
// 		// 	return
// 		// }
// 	}
// }

// TODO
func HandleAudioEncode() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func HandleAudio() http.HandlerFunc {
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
func HandleAudioUpload() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
