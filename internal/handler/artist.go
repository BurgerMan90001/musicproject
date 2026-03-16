package handler

import "net/http"

func HandleArtists() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//id := r.FormValue("id")
		switch r.Method {
		case http.MethodGet:
			//artist, err
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)

		}
	}
}
