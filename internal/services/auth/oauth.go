package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"musicproject.com/config"
	"musicproject.com/pkg/model"
)

type GoogleOauth struct {
	cfg *oauth2.Config
}

func NewGoogle(cfg config.Google) *GoogleOauth {
	return &GoogleOauth{&oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       cfg.Scopes,
		Endpoint:     google.Endpoint,
	}}
}
func (s *GoogleOauth) Login(ctx context.Context, code string) (*model.User, *model.TokenPair, error) {

	return nil, nil, nil
}

// Generates state cookie and returns redirect url
func (s *GoogleOauth) RedirectURL(w http.ResponseWriter) string {
	state := generateStateCookie(w)
	url := s.cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url
}

// func (s *GoogleOauth) exchange(ctx context.Context, code string) (*oauth2.Token, error) {
// 	return s.cfg.Exchange(ctx, code)
// }

func (s *GoogleOauth) getUserInfo(ctx context.Context, token *oauth2.Token) (*model.OauthUserInfo, error) {
	client := s.cfg.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var userInfo model.OauthUserInfo
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, err
	}

	return &userInfo, nil
}

func generateStateCookie(w http.ResponseWriter) string {
	b := make([]byte, 128)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	http.SetCookie(w, &http.Cookie{
		Name:  "oauthState",
		Value: state,

		HttpOnly: true,
		Secure:   true,
	})
	return state
}
func (s *GoogleOauth) validateStateCookie(r *http.Request) error {
	//code := r.FormValue("code")
	state := r.FormValue("state")
	stateCookie, err := r.Cookie("oauthState")
	if state != stateCookie.Value || err != nil {
		return errors.New("invalid google oauth state")
		// log.Println("invalid google oauth state")
		// http.Redirect(w, r, "/", http.StatusFound)
		// return
	}
	return nil
}
