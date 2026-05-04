package handler

import (
	"net/http"

	"songsled.com/internal/jsonutil"
	"songsled.com/internal/services/auth"
)

// type authService interface {
// 	Signup(ctx context.Context, email, password string) (*model.User, *model.TokenPair, error)
// 	Login(ctx context.Context, email, password string) (*model.User, *model.TokenPair, error)
// 	Refresh(ctx context.Context, refreshToken string) (*model.TokenPair, error)
// 	Logout(ctx context.Context, refreshToken string) error
// }
// type emailService interface {
// }

// func handleSignup(authService authService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		ctx := r.Context()

// 		signup, err := jsonutil.ReadJson[model.SignupRequest](r.Body)
// 		if err != nil {
// 			jsonutil.WriteError(w, &model.Error{
// 				Code:    http.StatusBadRequest,
// 				Message: "Invalid signup credentials",
// 			})
// 			return
// 		}

// 		user, tokenPair, err := authService.Signup(ctx, signup.Email, signup.Password)
// 		if err != nil {
// 			jsonutil.WriteError(w, err)
// 			return
// 		}
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "accessToken",
// 			Value:    tokenPair.AccessToken,
// 			HttpOnly: true,
// 			Secure:   true,
// 			MaxAge:   86400,

// 			SameSite: http.SameSiteStrictMode,
// 		})
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "refreshToken",
// 			Value:    tokenPair.RefreshToken,
// 			HttpOnly: true,
// 			Secure:   true,
// 			MaxAge:   86400,

// 			SameSite: http.SameSiteStrictMode,
// 		})

// 		jsonutil.WriteJSON(w, user, http.StatusCreated)
// 	}
// }

// func handleLogin(authService authService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		ctx := r.Context()

// 		login, err := jsonutil.ReadJson[model.LoginRequest](r.Body)
// 		if err != nil {
// 			jsonutil.WriteError(w, err)
// 			return
// 		}

// 		user, tokenPair, err := authService.Login(ctx, login.Email, login.Password)
// 		if err != nil {
// 			jsonutil.WriteError(w, err)
// 			return
// 		}

// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "accessToken",
// 			Value:    tokenPair.AccessToken,
// 			HttpOnly: true,
// 			Secure:   true,
// 			MaxAge:   86400,

// 			SameSite: http.SameSiteStrictMode,
// 		})
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "refreshToken",
// 			Value:    tokenPair.RefreshToken,
// 			HttpOnly: true,
// 			Secure:   true,
// 			MaxAge:   86400,

// 			SameSite: http.SameSiteStrictMode,
// 		})
// 		jsonutil.WriteJSON(w, user, http.StatusOK)
// 	}
// }

// func handleRefresh(authService authService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// 		// Try getting refresh token from cookie
// 		cookie, err := r.Cookie(string(model.TokenRefresh))
// 		if err != nil {
// 			jsonutil.WriteError(w, &model.Error{
// 				Code:    http.StatusUnauthorized,
// 				Message: "No refresh token",
// 				Details: err.Error(),
// 			})
// 			return
// 		}
// 		ctx := r.Context()

// 		tokenPair, err := authService.Refresh(ctx, cookie.Value)
// 		if err != nil {
// 			jsonutil.WriteError(w, err)
// 			return
// 		}
// 		http.SetCookie(w, &http.Cookie{
// 			Name:     "accessToken",
// 			Value:    tokenPair.AccessToken,
// 			HttpOnly: true,
// 			Secure:   true,
// 			MaxAge:   86400,

// 			SameSite: http.SameSiteStrictMode,
// 		})
// 		jsonutil.WriteJSON(w, tokenPair.AccessToken, http.StatusOK)
// 	}
// }

// func handleLogout(authService authService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {

// ctx := r.Context()

// cookie, err := r.Cookie(string(model.TokenRefresh))
// if err != nil {
// 	jsonutil.WriteError(w, auth.ErrInvalidToken())
// 	return
// }

// if err := authService.Logout(ctx, cookie.Value); err != nil {
// 	jsonutil.WriteError(w, err)
// 	return
// }

// Clear the cookies
// clearCookie(w, model.TokenAccess)
// clearCookie(w, model.TokenRefresh)

// 		jsonutil.WriteJSON(w, nil, http.StatusNoContent)
// 	}
// }

// func handleEmailReset(emailService emailService) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		jsonutil.NotImplemented(w)

//		}
//	}
func handleOidcLogin(oidc auth.Oidc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := oidc.RedirectUrl(w)

		http.Redirect(w, r, url, http.StatusFound)
	}
}

func handleOidcRedirect(oidc auth.Oidc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		t, err := oidc.Login(ctx, w, r)
		if err != nil {
			jsonutil.WriteError(w, err)
			return
		}

		jsonutil.WriteJSON(w, t, http.StatusOK)
	}
}

// slog.Info(t)
// http.SetCookie(w, &http.Cookie{
// 	Name:     "refreshToken",
// 	Value:    tokenPair.RefreshToken,
// 	HttpOnly: true,
// 	Secure:   true,
// 	MaxAge:   86400,

// 	SameSite: http.SameSiteStrictMode,
// })
// http.SetCookie(w, &http.Cookie{
// 	Name:     "accessToken",
// 	Value:    tokenPair.AccessToken,
// 	HttpOnly: true,
// 	Secure:   true,
// 	MaxAge:   86400,

// 	SameSite: http.SameSiteStrictMode,
// })
// 		jsonutil.WriteJSON(w, t, http.StatusOK)
// 	}
// }

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
