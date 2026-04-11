package handler

import (
	"context"
	"log"
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services/auth"
	"musicproject.com/internal/services/email"
	"musicproject.com/pkg/model"
)

type authService interface {
	Signup(ctx context.Context, email, password string) (*model.User, *model.TokenPair, error)
	Login(ctx context.Context, email, password string) (*model.User, *model.TokenPair, error)
	Refresh(ctx context.Context, refreshToken string) (*model.TokenPair, error)
	Logout(ctx context.Context, accessToken, refreshToken string) error
}

type authHandler struct {
	authService authService
}

func (h *authHandler) handleSignup() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()

		signup, err := jsonutil.ReadJSON[model.SignupRequest](r.Body)
		if err != nil {
			jsonutil.WriteError(w, ErrInvalidRequestBody, http.StatusBadRequest)
			return
		}

		user, tokenPair, err := h.authService.Signup(ctx, signup.Email, signup.Password)
		if err != nil {
			switch err {
			case auth.ErrInvalidEmail,
				auth.ErrInvalidPassword:
				jsonutil.WriteError(w, err, http.StatusBadRequest)
			case auth.ErrUserAlreadyExists:
				jsonutil.WriteError(w, err, http.StatusConflict)
			default:
				InternalServerError(w, err)
			}
			return
		}

		http.SetCookie(w, accessCookie(tokenPair.AccessToken, 1))
		http.SetCookie(w, refreshCookie(tokenPair.RefreshToken, 1))

		jsonutil.WriteJSON(w, user, http.StatusCreated)
	}
}

func (h *authHandler) handleLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()

		login, err := jsonutil.ReadJSON[model.LoginRequest](r.Body)
		if err != nil {
			jsonutil.WriteError(w, ErrInvalidRequestBody, http.StatusBadRequest)
			return
		}

		user, tokenPair, err := h.authService.Login(ctx, login.Email, login.Password)
		if err != nil {
			switch err {
			case auth.ErrIncorrectLogin:
				jsonutil.WriteError(w, err, http.StatusUnauthorized)
			case repository.ErrNotFound:
				jsonutil.WriteError(w, err, http.StatusNotFound)
			default:
				InternalServerError(w, err)
			}
			return
		}

		http.SetCookie(w, accessCookie(tokenPair.AccessToken, 1))
		http.SetCookie(w, refreshCookie(tokenPair.RefreshToken, 1))
		jsonutil.WriteJSON(w, user, http.StatusOK)
	}
}

func (h *authHandler) handleRefresh() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		// Allowed methods
		case http.MethodPost, http.MethodGet:
		default:
			MethodNotAllowedError(w)
			return
		}

		var refreshToken string

		cookie, err := r.Cookie(auth.TokenRefresh)
		if err == nil {
			refreshToken = cookie.Value
		} else {
			body, err := jsonutil.ReadJSON[model.RefreshRequest](r.Body)
			if err != nil {
				jsonutil.WriteError(w, auth.ErrNoRefeshToken, http.StatusUnauthorized)
				return
			}
			refreshToken = body.RefreshToken
		}

		if refreshToken == "" {
			jsonutil.WriteError(w, auth.ErrNoRefeshToken, http.StatusUnauthorized)
			return
		}

		ctx := r.Context()

		tokenPair, err := h.authService.Refresh(ctx, refreshToken)
		if err != nil {
			switch err {
			case auth.ErrInvalidTokenType:
				jsonutil.WriteError(w, auth.ErrNoRefeshToken, http.StatusUnauthorized)
			}

			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}
		jsonutil.WriteJSON(w, tokenPair, http.StatusOK)
	}
}

func HandleLogout(authService authService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			MethodNotAllowedError(w)
			return
		}
		ctx := r.Context()
		access, err := r.Cookie(auth.TokenAccess)
		if err != nil {
			jsonutil.WriteError(w, err, http.StatusBadRequest)
			return
		}

		if err := authService.Logout(ctx, access.Value, access.Value); err != nil {
			jsonutil.WriteError(w, err, http.StatusInternalServerError)
			return
		}

		// Clear the cookie
		http.SetCookie(w, accessCookie("", -1))
		http.SetCookie(w, refreshCookie("", -1))

		jsonutil.WriteJSON(w, nil, http.StatusNoContent)
	}
}
func (h *authHandler) handleEmailReset(emailService *email.Service) http.HandlerFunc {
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
func HandleOauthRedirect(oauth auth.Oauth) http.HandlerFunc {
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
		http.SetCookie(w, accessCookie(tokenPair.AccessToken, 1))
		http.SetCookie(w, refreshCookie(tokenPair.RefreshToken, 1))

		jsonutil.WriteJSON(w, user, http.StatusOK)
	}
}
