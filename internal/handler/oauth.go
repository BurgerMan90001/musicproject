package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/model"
	"musicproject.com/pkg/util/fileutil"
)

func handleOauthGoogleLogin(cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func handleOauthGoogleRedirect(jwtKey []byte, cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		state := r.FormValue("state")

		if state != "state" {
			log.Println("state mismatch")
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
		tokenString, err := auth.GenerateToken(jwtKey, user, auth.TokenRefresh, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}
		fileutil.WriteJSON(w, &model.Login{
			AccessToken:  tokenString,
			RefreshToken: tokenString,
			User:         user,
		})
	}
}

func getUserInfoGoogle(ctx context.Context, cfg *oauth2.Config, token *oauth2.Token) (*model.GoogleUserInfo, error) {
	client := cfg.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get: %v", err)
	}
	defer resp.Body.Close()

	userInfo, err := fileutil.ReadJSON[model.GoogleUserInfo](resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %v", err)
	}

	return userInfo, nil
}
