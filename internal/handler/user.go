package handler

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
)

type userHandler struct {
	repo repository.User
}

func (h *userHandler) handleUsers() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			MethodNotAllowedError(w)
			return
		}
	}
}
func (h *userHandler) handleUsersID() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		ctx := r.Context()

		switch r.Method {
		case http.MethodGet:
			// if repo == nil {
			// 	InternalServerError(w, ErrNilRepo)
			// 	return
			// }
			user, err := h.repo.GetByID(ctx, id)

			if err != nil {
				switch err {
				case repository.ErrNotFound:
					jsonutil.WriteError(w, err, http.StatusNotFound)
				default:
					InternalServerError(w, err)
				}
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

			err := h.repo.DeleteByID(ctx, id)
			if err != nil {
				switch err {
				case repository.ErrNotFound:
					jsonutil.WriteError(w, err, http.StatusNotFound)

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
