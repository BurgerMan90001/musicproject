package handler

import (
	"log"
	"net/http"

	"musicproject.com/internal/repository"
	"musicproject.com/internal/services"
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

		Path: "/",
		HttpOnly: true,
		Secure: true,
	}
	refreshCookie := &http.Cookie{
		Name:  RefreshCookie,
		Value: tokenPair.RefreshToken,

		Path: "/",
		HttpOnly: true,
		Secure: true,
	}
	http.SetCookie(w, accessCookie)
	http.SetCookie(w, refreshCookie)
}

func HandleSignup(authService services.Auth) http.HandlerFunc {
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

func HandleLogin(authService services.Auth) http.HandlerFunc {
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

func HandleRefresh(jwtService services.JWT) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}

		cookie, err := r.Cookie(RefreshCookie)
		if err != nil {
			switch err {
			case http.ErrNoCookie:
				WriteError(w, auth.ErrNoRefeshToken, http.StatusBadRequest)
			default:
				WriteError(w, err, http.StatusBadRequest)
			}
			return
		}
		ctx := r.Context()

		tokenPair, err := jwtService.RefreshTokens(ctx, cookie.Value)
		if err != nil {
			WriteError(w, err, http.StatusBadRequest)
			return
		}
		WriteJSON(w, tokenPair, http.StatusOK)
	}
}

func HandleLogout(authService services.Auth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()
		authService.Logout(ctx)
		WriteJSON(w, nil, http.StatusOK)
	}
}
func HandleEmailReset() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
		}
	}
}

/* Oauth handler functions */
func HandleOauthLogin(oauth services.Oauth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oauth.RedirectURL(w)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func HandleOauthGoogleRedirect(oauth services.Oauth) http.HandlerFunc {
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
			//w.WriteHeader(http.StatusInternalServerError)
			//log.Printf("unable to extchange token: %v", err)
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
