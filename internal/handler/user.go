package handler

import (
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
)

type userHandler struct {
	userRepo repository.User
}

func NewUser(userRepo repository.User) (*userHandler, error) {
	if userRepo == nil {
		return nil, fmt.Errorf("nil user repo")

	}
	return &userHandler{userRepo: userRepo}, nil
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

			user, err := h.userRepo.GetByID(ctx, id)

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

			err := h.userRepo.DeleteByID(ctx, id)
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
