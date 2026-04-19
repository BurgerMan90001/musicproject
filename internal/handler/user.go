package handler

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
)

func handleUsers(userRepo repository.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			jsonutil.MethodNotAllowedError(w)
			return
		}
	}
}
func handleUsersID(userRepo repository.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			jsonutil.WriteJSON(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:

			user, err := userRepo.GetByID(ctx, id)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}
			user.PasswordHash = ""
			jsonutil.WriteJSON(w, user, http.StatusOK)

		case http.MethodDelete:
			/*
				claims, err := auth.JWTParseToken(jwtKey, r)

				if err != nil {
					w.WriteHeader(http.StatusUnauthorized)
					fmt.Fprintln(w, "Invalid token:", err)
					return
				}
			*/

			err := userRepo.DeleteByID(ctx, id)
			if err != nil {
				jsonutil.WriteError(w, err)
				return
			}

		default:
			jsonutil.MethodNotAllowedError(w)
		}
	}
}
