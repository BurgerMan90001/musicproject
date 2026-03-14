package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
	"musicproject.com/pkg/util/handleutil"
)

func handleSignup(jwtKey []byte, repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		id, err := uuid.Parse(r.FormValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			//log.Printf("internal server error: %v", err)
			return
		}
		if id == uuid.Nil {
			id = uuid.New()
		}
		email := r.FormValue("email")
		password := r.FormValue("password")

		passwordHash, err := auth.HashPassword(password)
		if err != nil {
			if errors.Is(err, auth.ErrInvalidPassword) {
				http.Error(w, auth.ErrInvalidPassword.Error(), http.StatusBadRequest)
				return
			}
			handleutil.InternalServerError(w, r, err)
			return
		}
		user := &model.User{
			ID:           id,
			Email:        email,
			PasswordHash: passwordHash,
		}
		if err := repo.PutUser(ctx, id, user); err != nil {
			log.Printf("repository put error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tokenString, err := auth.GenerateToken(jwtKey, &model.User{
			ID:    id,
			Email: email,
		}, auth.TokenAccess, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/jwt")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, tokenString)
	}
}

func handleLogin(jwtKey []byte, repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		email := r.FormValue("email")
		password := r.FormValue("password")

		if email == "" || password == "" {
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintln(w, "empty username or password")
			return
		}

		user, err := repo.GetUserByEmail(ctx, email)

		if err != nil {
			if errors.Is(err, repository.ErrNotFound) ||
				!auth.ComparePassword(password, user.PasswordHash) {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Println("invalid email or password")
				return
			}

			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tokenString, err := auth.GenerateToken(jwtKey, &model.User{
			ID:    user.ID,
			Email: user.Email,
		}, auth.TokenAccess, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/jwt")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, tokenString)
	}
}

func handleRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		}
	}
}
func handleEmailReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		}
	}
}
