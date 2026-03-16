package handler

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
)

func HandleOauthGoogleLogin(cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		state := generateStateCookie(w)
		url := cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func HandleOauthGoogleRedirect(jwtKey []byte, cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		state := r.FormValue("state")

		stateCookie, err := r.Cookie("oauthState")

		if state != stateCookie.Value || err != nil {
			log.Println("invalid google oauth state")
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		ctx := context.Background()
		token, err := cfg.Exchange(ctx, code)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("unable to extchange token: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		userInfo, err := getUserInfoGoogle(ctx, cfg, token)
		if err != nil {
			log.Printf("get user info error: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		id, err := uuid.NewV7()
		if err != nil {
			log.Printf("generate uuid error: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
		}
		user := &model.User{
			ID:    id,
			Email: userInfo.Email,
		}
		tokenString, err := auth.GenerateToken(jwtKey, id, auth.TokenRefresh, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		WriteJSON(w, StatusSucess, model.LoginResponse{
			AccessToken:  tokenString,
			RefreshToken: tokenString,
			User:         user,
		}, http.StatusOK)
	}
}
func generateStateCookie(w http.ResponseWriter) string {
	b := make([]byte, 128)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	http.SetCookie(w, &http.Cookie{
		Name:  "oauthState",
		Value: state,
	})
	return state
}
func getUserInfoGoogle(ctx context.Context, cfg *oauth2.Config, token *oauth2.Token) (*model.GoogleUserInfo, error) {
	client := cfg.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get: %v", err)
	}
	defer resp.Body.Close()

	userInfo, err := ReadJSON[model.GoogleUserInfo](resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %v", err)
	}

	return &userInfo, nil
}
