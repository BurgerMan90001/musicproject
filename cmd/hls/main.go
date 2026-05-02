package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"

	"songsled.com/internal/config"
	"songsled.com/internal/middleware"
	"songsled.com/internal/services/encode"
)

func main() {
	run()
}
func addHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Range")
		next.ServeHTTP(w, r)
	})
}
func fs() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// ServeFile handles Range headers, 206  responses, ETags, and If-Modified-Since.
		// The path after /video/ maps to the file system.
		//path := "audio/" + r.URL.Path[len("/audio/"):]
		id := r.PathValue("id")

		log.Println(id)
		fs := http.FileServerFS(os.DirFS("audio"))

		ext := strings.ToLower(filepath.Ext(r.URL.Path))
		switch ext {
		case ".m3u8":
			w.Header().Set("Content-Type", "application/vnd.apple.mpegurl")
			w.Header().Set("Cache-Control", "no-cache")
		case ".ts":
			w.Header().Set("Content-Type", "video/mp2t")
			w.Header().Set("Cache-Control", "public, max-age=3600")
		default:

		}

		fs.ServeHTTP(w, r)
	})

}
func run() {
	r := chi.NewRouter()
	r.Handle("/segment", handleSegment())
	r.With(middleware.Logger(), addHeaders).Handle("/audio/", http.StripPrefix("/audio/", fs()))
	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleSegment() http.Handler {
	encoder := encode.NewFFmpeg(config.Encoder{
		Logging: true,
	})
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx := r.Context()
		id := uuid.New()
		outputDir := filepath.Join("audio", id.String())
		err := encoder.Segment(ctx, "", outputDir)

		if err != nil {
			fmt.Fprintln(w, err)
		}
	})
}
