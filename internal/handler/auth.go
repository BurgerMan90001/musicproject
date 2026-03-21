package handler

import (
	"errors"
	"log"
	"net/http"

	"musicproject.com/internal/repository"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

func HandleSignup(authService *auth.Service, repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			//MethodNotAllowedError(w, r)
			WriteError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		signup, err := model.ReadJSON[model.SignupRequest](r.Body)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		session, err := authService.Signup(ctx, signup.Email, signup.Password)
		if err != nil {
			switch {
			case errors.Is(err, auth.ErrInvalidEmail),
				errors.Is(err, auth.ErrInvalidPassword):
				WriteError(w, err, http.StatusBadRequest)
				return
			default:
				WriteError(w, err, http.StatusInternalServerError)
				return
			}
		}
		WriteJSON(w, session, http.StatusOK)
	}
}

func HandleLogin(authService *auth.Service, repo repository.Repository) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}
		login, err := model.ReadJSON[model.LoginRequest](r.Body)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
		}
		ctx := r.Context()

		session, err := authService.Login(ctx, login.Email, login.Password)
		if err != nil {
			switch {
			//case errors.Is(err, auth.):
			default:
				WriteError(w, err, http.StatusInternalServerError)
			}
			return
		}
		WriteJSON(w, session, http.StatusOK)

		// if err != nil {
		// 	if errors.Is(err, repository.ErrNotFound) ||
		// 		!auth.ComparePassword(password, user.PasswordHash) {
		// 		w.WriteHeader(http.StatusUnauthorized)
		// 		fmt.Println("invalid email or password")
		// 		return
		// 	}

		// 	w.WriteHeader(http.StatusInternalServerError)
		// 	return
		// }

		// tokenString, err := auth.GenerateToken(jwtKey, user.ID,
		// 	auth.TokenAccess, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		WriteJSON(w, nil, http.StatusOK)
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

/* Oauth handler functions */
func HandleOauthLogin(authService *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := authService.Google.RedirectURL(w)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func HandleOauthGoogleRedirect(authService *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//code := r.FormValue("code")
		state := r.FormValue("state")

		stateCookie, err := r.Cookie("oauthState")

		if state != stateCookie.Value || err != nil {
			log.Println("invalid google oauth state")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		//authService
		//ctx := context.Background()
		//token, err := cfg.Exchange(ctx, code)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("unable to extchange token: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		// user := &model.User{
		// 	ID:    id,
		// 	Email: userInfo.Email,
		//

		// userInfo, err := authService.GetUserInfoGoogle(ctx, token)
		// if err != nil {
		// 	log.Printf("get user info error: %v", err)
		// 	http.Redirect(w, r, "/", http.StatusFound)
		// 	return
		// }

		//tokenString, err := auth.GenerateToken(jwtKey, id, auth.TokenRefresh, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		//authService.GetUserInfoGoogle(ctx, )

	}
}
