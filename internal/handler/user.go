package handler

import (
	"errors"
	"log"
	"net/http"

	"okapi.com/internal/controller/user"
	"okapi.com/internal/repository"
	"okapi.com/pkg/util/fileutil"
)

func handleUser(c *user.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id := r.FormValue("id")
		if id == "" {
			w.WriteHeader(http.StatusBadRequest)
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			user, err := c.GetUserByID(ctx, id)
			if err != nil && errors.Is(err, repository.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
				return
			} else if err != nil {
				log.Printf("repository get error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			fileutil.WriteJSON(w, user)
		case http.MethodDelete:
			/*
				claims, err := auth.JWTParseToken(jwtKey, r)

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintln(w, "Invalid token:", err)
					return
				}
			*/

			err := c.DeleteUserByID(ctx, id)
			if err != nil && errors.Is(err, repository.ErrNotFound) {
				w.WriteHeader(http.StatusNotFound)
				fileutil.WriteJSON(w, err)
				return
			}

		default:
			http.Error(w, "invalid request method", http.StatusMethodNotAllowed)
		}
	}
}
