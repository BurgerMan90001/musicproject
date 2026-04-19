package auth

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"time"

	"golang.org/x/oauth2"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/pkg/model"
)

type GoogleOauth struct {
	cfg *oauth2.Config
}

func NewOauth(redirectUrl string, scopes []string, endpoint oauth2.Endpoint) (*GoogleOauth, error) {
	secretList, err := secrets.GetEnvMap("GOOGLE_OAUTH_CLIENT_ID", "GOOGLE_OAUTH_CLIENT_SECRET")
	if err != nil {
		return nil, err
	}
	return &GoogleOauth{&oauth2.Config{
		ClientID:     secretList["GOOGLE_OAUTH_CLIENT_ID"],
		ClientSecret: secretList["GOOGLE_OAUTH_CLIENT_SECRET"],
		RedirectURL:  redirectUrl,
		Scopes:       scopes,
		Endpoint:     endpoint,
	}}, nil
}
func (s *GoogleOauth) Login(ctx context.Context, code string) (*model.User, *model.TokenPair, error) {
	userInfo, err := s.getUserInfo(ctx, &oauth2.Token{})
	if err != nil {
		return nil, nil, err
	}
	user := &model.User{
		Email: userInfo.Email,
	}

	return user, nil, nil
}

// Generates state cookie and returns redirect url
func (s *GoogleOauth) RedirectURL(w http.ResponseWriter) string {
	state := generateStateCookie(w)
	url := s.cfg.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return url
}

func (s *GoogleOauth) getUserInfo(ctx context.Context, token *oauth2.Token) (*model.OauthUserInfo, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
	defer cancel()

	client := s.cfg.Client(ctx, token)

	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	userInfo, err := jsonutil.ReadJson[*model.OauthUserInfo](resp.Body)
	if err != nil {
		return nil, err
	}

	return userInfo, nil
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

// func (s *GoogleOauth) validateStateCookie(r *http.Request) error {
// 	//code := r.FormValue("code")
// 	state := r.FormValue("state")
// 	stateCookie, err := r.Cookie("oauthState")
// 	if state != stateCookie.Value || err != nil {
// 		return errors.New("invalid google oauth state")

// 	}
// 	return nil
// }
