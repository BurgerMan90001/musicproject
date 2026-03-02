package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"okapi.com/internal/auth"
	"okapi.com/pkg/model"
	"okapi.com/pkg/util/fileutil"
)

func handleOathGoogleLogin(cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		url := cfg.AuthCodeURL("state", oauth2.AccessTypeOffline)
		http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	}
}
func handleOathGoogleRedirect(jwtKey []byte, cfg *oauth2.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		state := r.FormValue("state")

		if state != "state" {
			log.Println("state mismatch")
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		ctx := context.Background()
		token, err := cfg.Exchange(ctx, code)

		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Printf("unable to extchange token: %v", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		userInfo, err := getUserInfoGoogle(ctx, cfg, token)
		if err != nil {
			log.Printf("get user info error: %v", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}

		tokenString, err := auth.GenerateToken(jwtKey, &model.User{
			ID:    userInfo.ID,
			Email: userInfo.Email,
		}, auth.ExpiresInOneDay)

		if err != nil {
			log.Printf("generate token error: %v", err)
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			return
		}
		fileutil.WriteJSON(w, &model.Login{
			AccessToken: tokenString,
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
