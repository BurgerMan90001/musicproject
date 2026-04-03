package handler

import (
	"log"
	"net/http"

	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

const (
	AccessCookie  = "accessToken"
	RefreshCookie = "refreshToken"
)

func setTokenPair(w http.ResponseWriter, tokenPair *model.TokenPair) {

	accessCookie := &http.Cookie{
		Name:  AccessCookie,
		Value: tokenPair.AccessToken,

		Path:     "/v1/auth/refresh",
		HttpOnly: true,
		Secure:   true,
	}
	refreshCookie := &http.Cookie{
		Name:  RefreshCookie,
		Value: tokenPair.RefreshToken,

		Path:     "/",
		HttpOnly: true,
		Secure:   true,
	}
	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
}

func handleSignup(authService *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			WriteError(w, ErrInvalidMethod, http.StatusMethodNotAllowed)
			return
		}
		ctx := r.Context()

		signup, err := model.ReadJSON[model.SignupRequest](r.Body)
		if err != nil {
			WriteError(w, ErrInvalidRequestBody, http.StatusBadRequest)
			return
		}

		user, tokenPair, err := authService.Signup(ctx, signup.Email, signup.Password)
		if err != nil {
			switch err {
			case auth.ErrInvalidEmail,
				auth.ErrInvalidPassword:
				WriteError(w, err, http.StatusBadRequest)
			case auth.ErrUserAlreadyExists:
				WriteError(w, err, http.StatusConflict)
			default:
				InternalServerError(w, err)
			}
			return
		}

		setTokenPair(w, tokenPair)
		WriteJSON(w, user, http.StatusCreated)
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
			WriteError(w, ErrInvalidRequestBody, http.StatusBadRequest)
			return
		}

		user, tokenPair, err := authService.Login(ctx, login.Email, login.Password)
		if err != nil {
			switch err {
			case auth.ErrIncorrectLogin:
				WriteError(w, err, http.StatusUnauthorized)
			case repository.ErrNotFound:
				WriteError(w, err, http.StatusNotFound)
			default:
				InternalServerError(w, err)
			}
			return
		}

		setTokenPair(w, tokenPair)
		WriteJSON(w, user, http.StatusOK)
	}
}

func HandleRefresh(authService *auth.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}

		var refreshToken string

		cookie, err := r.Cookie(RefreshCookie)
		if err == nil {
			refreshToken = cookie.Value
		} else {
			body, err := model.ReadJSON[model.RefreshRequest](r.Body)
			if err != nil {
				WriteError(w, auth.ErrNoRefeshToken, http.StatusBadRequest)
			}
			refreshToken = body.RefreshToken
		}

		// if refreshToken == "" {
		// 	WriteError(w, auth.ErrNoRefeshToken, http.StatusBadRequest)
		// 	return
		// }

		ctx := r.Context()

		tokenPair, err := authService.Refresh(ctx, refreshToken)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		WriteJSON(w, tokenPair, http.StatusOK)
	}
}

func HandleLogout(authService auth.JWTService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()
		cookie, err := r.Cookie(AccessCookie)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}

		if err := authService.RevokeToken(ctx, cookie.Value); err != nil {
			WriteError(w, err, http.StatusInternalServerError)
			return
		}

		//authService.Logout(ctx)

		// Clear the cookie
		http.SetCookie(w, &http.Cookie{
			Name:     RefreshCookie,
			Value:    "",
			HttpOnly: true,
			Secure:   true,
			SameSite: http.SameSiteStrictMode,
			Path:     "/v1/auth/refresh",
			MaxAge:   -1,
		})

		WriteJSON(w, nil, http.StatusNoContent)
	}
}
func HandleEmailReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}

	}
}

/* Oauth handler functions */
func HandleOauthLogin(oauth auth.Oauth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oauth.RedirectURL(w)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func HandleOauthGoogleRedirect(oauth auth.Oauth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		state := r.FormValue("state")

		stateCookie, err := r.Cookie("oauthState")
		if state != stateCookie.Value || err != nil {
			log.Println("invalid google oauth state")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		ctx := r.Context()

		user, tokenPair, err := oauth.Login(ctx, code)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		setTokenPair(w, tokenPair)
		WriteJSON(w, user, http.StatusOK)
		// userInfo, err := oauth.GetUserInfo(ctx, token)
		// if err != nil {
		// 	log.Printf("get user info error: %v", err)
		// 	http.Redirect(w, r, "/", http.StatusFound)
		// 	return
		// }

		// tokenPair, err := authService.GenerateTokenPair()

		// if err != nil {
		// 	log.Printf("generate token error: %v", err)
		// 	http.Redirect(w, r, "/", http.StatusFound)
		// 	return
		// }

	}
}
