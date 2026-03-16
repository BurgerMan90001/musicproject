package handler

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
)

func HandleUserID(repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			user, err := repo.GetUserByID(ctx, id)
			if err != nil {
				if errors.Is(err, repository.ErrNotFound) {
					WriteErrJSON(w, ErrUserNotFound, http.StatusNotFound)
					return
				}
				InternalServerError(w, r, err)
				return
			}

			WriteJSON(w, "success", user, http.StatusOK)
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

			err := repo.DeleteUserByID(ctx, id)
			if err != nil && errors.Is(err, repository.ErrNotFound) {

				WriteErrJSON(w, err, http.StatusNotFound)
				return
			}

		default:
			MethodNotAllowedError(w, r)
		}
	}
}
