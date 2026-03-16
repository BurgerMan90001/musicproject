package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"musicproject.com/internal/repository"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

func HandleSignup(jwtKey []byte, repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w, r)
			return
		}
		ctx := r.Context()

		signup, err := ReadJSON[model.SignupRequest](r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if err := auth.ValidateEmail(signup.Email); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteErrJSON(w, err, http.StatusBadRequest)
			return
		}
		if err := auth.ValidatePassword(signup.Password); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			WriteErrJSON(w, err, http.StatusBadRequest)
			return
		}

		if user, _ := repo.GetUserByEmail(ctx, signup.Email); user != nil {
			WriteErrJSON(w, ErrUserAlreadyExists,
				http.StatusConflict)
			return
		}

		id, err := repo.PutUser(ctx, signup.Email, signup.Password)
		if err != nil {
			log.Printf("repository put error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		user := &model.User{
			ID:    id,
			Email: signup.Email,
		}

		accessToken, err := auth.GenerateToken(jwtKey, id,
			auth.TokenAccess, auth.ExpiresInOneDay)

		refreshToken, err := auth.GenerateToken(jwtKey, id,
			auth.TokenRefresh, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		res := model.LoginResponse{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
			User:         user,
		}

		WriteJSON(w, StatusSucess, res, http.StatusOK)
	}
}

func HandleLogin(jwtKey []byte, repo repository.Repository) http.HandlerFunc {
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

		tokenString, err := auth.GenerateToken(jwtKey, user.ID,
			auth.TokenAccess, auth.ExpiresInOneDay)

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

func HandleRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		}
	}
}

func HandleLogout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		}
	}
}
func HandleEmailReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		}
	}
}
