package handler

import (
	"net/http"
	"path/filepath"

	"musicproject.com/internal/services/song"
)

// For testing audio serving
func handleAudio(songService *song.Service) http.HandlerFunc {
	//fs := http.FileServer(fileService.R)
	return func(w http.ResponseWriter, r *http.Request) {

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
	}
}
