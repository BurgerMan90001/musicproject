package handler

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
)

func HandleUserID(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			if repo == nil {
				InternalServerError(w, ErrNilRepo)
				return
			}
			user, err := repo.GetUserByID(ctx, id)
			if err != nil {
				switch err {
				case repository.ErrNotFound:
					WriteError(w, err, http.StatusNotFound)
				default:
					InternalServerError(w, err)
				}
				return
			}
			WriteJSON(w, user, http.StatusOK)

		case http.MethodDelete:
			/*
				claims, err := auth.JWTParseToken(jwtKey, r)

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintln(w, "Invalid token:", err)
					return
				}
			*/

			err := repo.DeleteUserByID(ctx, id)
			if err != nil {
				switch err {
				case repository.ErrNotFound:
					WriteError(w, err, http.StatusNotFound)

				default:
					InternalServerError(w, err)
				}
				return
			}

		default:
			MethodNotAllowedError(w)
		}
	}
}
