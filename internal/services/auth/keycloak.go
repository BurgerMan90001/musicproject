package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/coreos/go-oidc"
	"golang.org/x/oauth2"
	"songsled.com/internal/config/secrets"
	"songsled.com/pkg/crand"
	"songsled.com/pkg/model"
)

type Oidc interface {
	Login(ctx context.Context, w http.ResponseWriter, r *http.Request) (*oidc.IDToken, error)
	RedirectUrl(w http.ResponseWriter) string
}
type Client struct {
	provider  *oidc.Provider
	oidc      *oidc.IDTokenVerifier
	cfg       *oauth2.Config
	ver       string
	stateName string
}

func NewClient(ctx context.Context,
	issuerVar, redirectVar,
	clientIdVar, clientSecretVar string,
) (*Client, error) {

	s, err := secrets.GetenvMap(issuerVar, redirectVar,
		clientIdVar, clientSecretVar)
	if err != nil {
		return nil, err
	}
	provider, err := oidc.NewProvider(ctx, s[issuerVar])
	if err != nil {
		return nil, err
	}

	cfg := &oauth2.Config{
		ClientID:     s[clientIdVar],
		ClientSecret: s[clientSecretVar],
		RedirectURL:  s[redirectVar],
		Endpoint:     provider.Endpoint(),
		Scopes:       []string{oidc.ScopeOpenID, "roles"},
	}
	oidcConfig := &oidc.Config{
		ClientID: s[clientIdVar],
	}
	oidc := provider.Verifier(oidcConfig)

	return &Client{
		provider:  provider,
		cfg:       cfg,
		oidc:      oidc,
		stateName: "state",
	}, nil
}

// Generates state cookie and returns redirect url
func (s *Client) RedirectUrl(w http.ResponseWriter) string {
	state := s.generateStateCookie(w)
	s.ver = oauth2.GenerateVerifier()
	url := s.cfg.AuthCodeURL(state, oauth2.S256ChallengeOption(s.ver))
	return url
}

func (s *Client) Login(ctx context.Context, w http.ResponseWriter, r *http.Request) (*oidc.IDToken, error) {
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
	idToken, err := s.oidc.Verify(ctx, rawIdToken)
	if err != nil {
		return nil, fmt.Errorf("Failed to verify id_token: %w", err)
	}
	http.SetCookie(w, &http.Cookie{
		Name:     s.stateName,
		Value:    "",
		MaxAge:   -1,
		Secure:   true,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	http.Redirect(w, r, "", http.StatusFound)

	return idToken, nil
}

func (s *Client) Logout(ctx context.Context) error {
	// url.JoinPath("realms")
	// url := fmt.Sprintf("%w/realms/%s")
	// http.Post()
	return nil
}
func (s *Client) generateStateCookie(w http.ResponseWriter) string {
	state := crand.NewB64Url(128)
	http.SetCookie(w, &http.Cookie{
		Name:  s.stateName,
		Value: state,

		HttpOnly: true,
		Secure:   true,
	})
	return state
}

// Validate to protect against csrf
func (s *Client) validateStateCookie(r *http.Request) error {
	state := r.FormValue("state")
	stateCookie, err := r.Cookie(s.stateName)
	if state != stateCookie.Value || err != nil {
		return errors.New("Invalid google oauth state")
	}
	return nil
}
