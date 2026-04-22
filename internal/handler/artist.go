package handler

import (
	"net/http"

	"musicproject.com/internal/jsonutil"
)

func HandleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id := r.FormValue("id")
		switch r.Method {
		case http.MethodGet:
			//artist, err
		default:
			jsonutil.WriteMethodNotAllowed(w)
		}
	}
}
