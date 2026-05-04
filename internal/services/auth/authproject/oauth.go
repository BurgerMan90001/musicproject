package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	oidc "github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"songsled.com/internal/config/secrets"
	"songsled.com/pkg/crand"
	"songsled.com/pkg/model"
)

var _ Oauth = (*GoogleOauth)(nil)

type Oauth interface {
	Login(ctx context.Context, r *http.Request) (*oidc.IDToken, error)
	RedirectUrl(w http.ResponseWriter) string
}
type GoogleOauth struct {
	cfg *oauth2.Config

	stateCookie string
	ver         string
	verifier    *oidc.IDTokenVerifier
}

func NewOauth(ctx context.Context,
	issuer, redirect,
	clientIdVar, clientSecretVar string,
	scopes []string,

) (*GoogleOauth, error) {

	// // http://localhost:8080/auth/realms/songsled
	// // https://accounts.google.com

	provider, err := oidc.NewProvider(ctx, issuer)
	if err != nil {
		return nil, fmt.Errorf("New provider: %w", err)
	}
	clientID, err := secrets.Getenv(clientIdVar)
	if err != nil {
		return nil, err
	}
	clientSecret, err := secrets.Getenv(clientIdVar)
	if err != nil {
		return nil, err
	}

	cfg := &oauth2.Config{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		RedirectURL:  redirect,

		Endpoint: provider.Endpoint(),

		Scopes: scopes,
	}

	oidcConfig := &oidc.Config{
		ClientID: clientID,
	}
	verifier := provider.Verifier(oidcConfig)

	return &GoogleOauth{
		cfg: cfg,
		// Cookie: "state",
		verifier: verifier,
	}, nil

}

// Generates state cookie and returns redirect url
func (s *GoogleOauth) RedirectUrl(w http.ResponseWriter) string {
	state := s.generateStateCookie(w)
	s.ver = oauth2.GenerateVerifier()

	url := s.cfg.AuthCodeURL(state, oauth2.S256ChallengeOption(s.ver))
	return url
}

func (s *GoogleOauth) Login(ctx context.Context, r *http.Request) (*oidc.IDToken, error) {
	code := r.FormValue("code")
	if err := s.validateStateCookie(r); err != nil {
		return nil, err
	}

	token, err := s.cfg.Exchange(ctx, code, oauth2.VerifierOption(s.ver))
	if err != nil {
		return nil, fmt.Errorf("Login exchange: %v", err)
	}

	rawIdToken, ok := token.Extra("id_token").(string)
	if !ok {
		return nil, &model.Error{
			Code:    http.StatusInternalServerError,
			Message: "No id_token field in oauth2 token",
		}
	}
	idToken, err := s.verifier.Verify(ctx, rawIdToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to verify id_token: %w", err)
	}

	return idToken, nil
}

// func (s *GoogleOauth) getUserInfo(ctx context.Context, token *oauth2.Token) (*model.OauthUserInfo, error) {
// 	ctx, cancel := context.WithTimeout(ctx, time.Second*5)
// 	defer cancel()

// 	client := s.cfg.Client(ctx, token)

// 	resp, err := client.Get(s.userInfoUrl)
// 	if err != nil {
// 		return nil, fmt.Errorf("getUserInfo: client.Get %w", err)
// 	}
// 	defer resp.Body.Close()

// 	userInfo, err := jsonutil.ReadJson[*model.OauthUserInfo](resp.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("getUserInfo: read json %w", err)
// 	}

// 	return userInfo, nil
// }

// userInfo, err := s.getUserInfo(ctx, token)
// if err != nil {
// 	return nil, nil, err
// }

//	user := &model.User{
//		Email: userInfo.Email,
//	}
func (s *GoogleOauth) generateStateCookie(w http.ResponseWriter) string {
	state := crand.NewB64Url(128)
	http.SetCookie(w, &http.Cookie{
		Name:  s.stateCookie,
		Value: state,

		HttpOnly: true,
		Secure:   true,
	})
	return state
}

// Validate to protect against csrf
func (s *GoogleOauth) validateStateCookie(r *http.Request) error {
	state := r.FormValue("state")
	stateCookie, err := r.Cookie(s.stateCookie)
	if state != stateCookie.Value || err != nil {
		return errors.New("Invalid google oauth state")
	}
	return nil
}
