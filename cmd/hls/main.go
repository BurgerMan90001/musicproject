package main

import (
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func main() {
	run()
}

func run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// ServeFile handles Range headers, 206 responses, ETags, and If-Modified-Since.
		// The path after /video/ maps to the file system.
		//path := "audio/" + r.URL.Path[len("/audio/"):]
		fs := http.FileServer(http.Dir("audio"))

		// Use CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Range")


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

	log.Println("serving on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
