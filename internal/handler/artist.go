package handler

import (
	"net/http"

	"songsled.com/internal/jsonutil"
)

func handleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		jsonutil.NotImplemented(w)
		//id := r.FormValue("id")
		switch r.Method {
		case http.MethodGet:
			//artist, err
		default:
		}
	}
}
