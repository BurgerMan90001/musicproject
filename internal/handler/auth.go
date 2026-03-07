package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/auth"
	"musicproject.com/internal/controller/user"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

func handleSignup(jwtKey []byte, c *user.Controller) http.HandlerFunc {
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
			fmt.Fprintln(w, "empty email or password")
			return
		}
		id := uuid.New()
		err := c.PutUser(ctx, id, email, password)
		if err != nil {
			log.Printf("repository put error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		tokenString, err := auth.GenerateToken(jwtKey, &model.User{
			ID:    id,
			Email: email,
		}, auth.ExpiresInOneDay)

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

func handleLogin(jwtKey []byte, c *user.Controller) http.HandlerFunc {
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

		user, err := c.GetUserByEmail(ctx, email)

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
		}, auth.ExpiresInOneDay)

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
