package handler

import (
	"errors"
	"log"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/controller/user"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/util/fileutil"
	"musicproject.com/pkg/util/handleutil"
)

func handleUser(c *user.Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.FormValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if id == uuid.Nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			user, err := c.GetUserByID(ctx, id)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					w.WriteHeader(http.StatusNotFound)
					return
				}
				log.Printf("repository get error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fileutil.WriteJSON(w, user)
		case http.MethodPut:

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
			handleutil.ErrMethodNotAllowed(w, r)
		}
	}
}
