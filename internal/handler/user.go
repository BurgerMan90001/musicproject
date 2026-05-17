package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/internal/repository/postgres"
	"songsled.com/pkg/model"
)

// type userRepo interface {
// 	GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error)
// 	DeleteUserByID(ctx context.Context, id uuid.UUID) error
// }

func handleUsers(userRepo *postgres.User) func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {

		})
		r.Get("/{userId}", func(w http.ResponseWriter, r *http.Request) {
			_, err := uuid.Parse(r.PathValue("id"))
			if err != nil {
				jsonutil.WriteError(w, &model.Error{
					Code:    http.StatusNotFound,
					Message: "User not found",
				})
				return
			}
			// ctx := r.Context()
		})
	}
}
func handleGetUsersId(userRepo *postgres.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// user, err := userRepo.GetUserByID(ctx, id)
		// if err != nil {
		// 	jsonutil.WriteError(w, err)
		// 	return
		// }
		// user.PasswordHash = ""
		// jsonutil.WriteJSON(w, user, http.StatusOK)

	}
}

func handleDelteUsersId(userRepo *postgres.User) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := uuid.Parse(r.PathValue("id"))
		if err != nil {
			jsonutil.WriteError(w, &model.Error{
				Code:    http.StatusNotFound,
				Message: "User not found",
			})
			return
		}
		ctx := r.Context()

		if err := userRepo.DeleteUserByID(ctx, id); err != nil {
			jsonutil.WriteError(w, err)
			return
		}

	}
}

func handleUserHistory() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
