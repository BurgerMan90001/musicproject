package handler

import (
	"log"
	"net/http"

	"musicproject.com/internal/repository"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

const (
	AccessCookie  = "accessKey"
	RefreshCookie = "refreshKey"
)

func HandleSignup(authService *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			WriteError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		signup, err := model.ReadJSON[model.SignupRequest](r.Body)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}

		tokenPair, err := authService.Signup(ctx, signup.Email, signup.Password)
		if err != nil {
			switch err {
			case auth.ErrInvalidEmail,
				auth.ErrInvalidPassword:
				WriteError(w, err, http.StatusBadRequest)
				return
			case auth.ErrUserAlreadyExists:
				WriteError(w, err, http.StatusConflict)
			default:
				WriteError(w, err, http.StatusInternalServerError)
				return
			}
		}

		WriteJSON(w, tokenPair, http.StatusOK)
	}
}

func HandleLogin(authService *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()

		login, err := model.ReadJSON[model.LoginRequest](r.Body)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}

		tokenPair, err := authService.Login(ctx, login.Email, login.Password)
		if err != nil {
			switch err {
			case auth.ErrInvalidPassword,
				auth.ErrInvalidEmail:
				WriteError(w, err, http.StatusUnauthorized)
			case repository.ErrNotFound:
				WriteError(w, err, http.StatusNotFound)
			default:
				WriteError(w, err, http.StatusInternalServerError)
			}
			return
		}
		WriteJSON(w, tokenPair, http.StatusOK)

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

		// if err != nil {
		// 	log.Printf("generate token error: %v", err)
		// 	http.Redirect(w, r, "/", http.StatusFound)
		// 	return
		// }
		//authService.GetUserInfoGoogle(ctx, )

	}
}
