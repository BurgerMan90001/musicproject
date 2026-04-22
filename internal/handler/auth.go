package handler

import (
	"context"
	"log"
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

type authService interface {
	Signup(ctx context.Context, email, password string) (*model.User, *model.TokenPair, error)
	Login(ctx context.Context, email, password string) (*model.User, *model.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*model.TokenPair, error)
	Logout(ctx context.Context, refreshToken string) error
}
type emailService interface {
}

func handleSignup(authService authService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonutil.WriteMethodNotAllowed(w)
			return
		}
		ctx := r.Context()

		signup, err := jsonutil.ReadJson[model.SignupRequest](r.Body)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		user, tokenPair, err := authService.Signup(ctx, signup.Email, signup.Password)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		// Set cookies with max age of 24 hours
		setCookie(w, model.TokenAccess, tokenPair.AccessToken, 86400)
		setCookie(w, model.TokenRefresh, tokenPair.RefreshToken, 86400)

		jsonutil.WriteJSON(w, user, http.StatusCreated)
	}
}

func handleLogin(authService authService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonutil.WriteMethodNotAllowed(w)
			return
		}
		ctx := r.Context()

		login, err := jsonutil.ReadJson[model.LoginRequest](r.Body)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		user, tokenPair, err := authService.Login(ctx, login.Email, login.Password)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		// Set cookies with max age of 24 hours
		setCookie(w, model.TokenAccess, tokenPair.AccessToken, 86400)
		setCookie(w, model.TokenRefresh, tokenPair.RefreshToken, 86400)
		jsonutil.WriteJSON(w, user, http.StatusOK)
	}
}

func handleRefresh(authService authService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// Allowed methods
		case http.MethodPost, http.MethodGet:
		default:
			jsonutil.WriteMethodNotAllowed(w)
			return
		}

		// Try getting refresh token from cookie
		cookie, err := r.Cookie(string(model.TokenRefresh))
		if err != nil {
			jsonutil.WriteError(w, &model.Error{
				Code:    http.StatusUnauthorized,
				Message: "No refresh token",
				Details: []string{err.Error()},
			})
			return
		}
		ctx := r.Context()

		tokenPair, err := authService.Refresh(ctx, cookie.Value)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		setCookie(w, model.TokenRefresh, tokenPair.RefreshToken, 86400)
		jsonutil.WriteJSON(w, tokenPair.AccessToken, http.StatusOK)
	}
}

func handleLogout(authService authService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonutil.WriteMethodNotAllowed(w)
			return
		}
		ctx := r.Context()

		cookie, err := r.Cookie(string(model.TokenRefresh))
		if err != nil {
			jsonutil.WriteError(w, auth.ErrNoToken)
			return
		}

		if err := authService.Logout(ctx, cookie.Value); err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		// Clear the cookies
		clearCookie(w, model.TokenAccess)
		clearCookie(w, model.TokenRefresh)

		jsonutil.WriteJSON(w, nil, http.StatusNoContent)
	}
}

func handleEmailReset(emailService emailService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			jsonutil.WriteMethodNotAllowed(w)
			return
		}

	}
}

/* Oauth handler functions */
func handleOauthLogin(oauth auth.Oauth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oauth.RedirectURL(w)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func handleOauthRedirect(oauth auth.Oauth) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		code := r.FormValue("code")
		state := r.FormValue("state")

		stateCookie, err := r.Cookie("oauthState")
		if state != stateCookie.Value || err != nil {
			log.Println("invalid google oauth state")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		user, tokenPair, err := oauth.Login(ctx, code)
		if err != nil {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		// TODO set proper max ages
		setCookie(w, model.TokenAccess, tokenPair.AccessToken, 86400)
		setCookie(w, model.TokenRefresh, tokenPair.RefreshToken, 86400)

		jsonutil.WriteJSON(w, user, http.StatusOK)
	}
}

//	func requestRefreshToken(r *http.Request) (string, error) {
//		var refreshToken string
//		cookie, err := r.Cookie(string(model.TokenRefresh))
//		if err == nil {
//			refreshToken = cookie.Value
//		} else {
//			body, err := jsonutil.ReadJson[model.RefreshRequest](r.Body)
//			if err != nil {
//				return "", err
//			}
//			refreshToken = body.RefreshToken
//		}
//		return refreshToken, nil
//	}
