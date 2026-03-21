package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

const (
	TokenAccess  = "access"
	TokenRefresh = "refresh"
)

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)
var ExpiresInOneHour = time.Now().Add(time.Hour * 24)

type Service struct {
	cfg  config.Auth
	repo repository.Repository

	Google Oauth
}

type Claims struct {
	UserID    uuid.UUID `json:"userId"`
	TokenType string    `json:"tokenType"`
	jwt.RegisteredClaims
}

type Oauth interface {
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*model.OauthUserInfo, error)
	RedirectURL(w http.ResponseWriter) string
}

func New(cfg config.Auth, repo repository.Repository) *Service {

	google := NewGoogle(cfg.Oauth.Google)

	return &Service{cfg, repo, google}
}

func (s *Service) generateToken(userId uuid.UUID, tokenType string, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:    userId,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.cfg.JWT.Issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	})

	key, err := s.keyType(tokenType)
	if err != nil {
		return "", err
	}
	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (s *Service) generateTokenPair(userId uuid.UUID) (*model.TokenPair, error) {
	accessToken, err := s.generateToken(userId, TokenAccess, ExpiresInOneHour)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateToken(userId, TokenAccess, ExpiresInOneHour)
	if err != nil {
		return nil, err
	}
	// Revoke refresh

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *Service) parseToken(tokenString string, tokenType string) (*jwt.Token, error) {
	key, err := s.keyType(tokenType)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString,
		&Claims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrSignatureInvalid
			}
			return key, nil
		},
		jwt.WithIssuer(s.cfg.JWT.Issuer),
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
	)
	return token, err
}
func (s *Service) keyType(tokenType string) ([]byte, error) {
	switch tokenType {
	case TokenAccess:
		return []byte(s.cfg.JWT.AccessKey), nil
	case TokenRefresh:
		return []byte(s.cfg.JWT.RefreshKey), nil
	default:
		return nil, ErrInvalidTokenType
	}
}
func (s *Service) ValidateAccessToken(tokenString string) (*Claims, error) {
	token, err := s.parseToken(tokenString, TokenAccess)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, err
		}
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	switch {
	case !ok || !token.Valid:
		return nil, jwt.ErrTokenInvalidClaims
	case claims.TokenType != TokenAccess:
		return nil, ErrInvalidTokenType
	default:
		return claims, nil
	}
}

func (s *Service) RefreshTokens(ctx context.Context, tokenString string) (*model.TokenPair, error) {
	token, err := s.parseToken(tokenString, TokenRefresh)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	switch {
	case !ok || !token.Valid:
		return nil, jwt.ErrTokenInvalidClaims

	case claims.TokenType != TokenRefresh:
		return nil, ErrInvalidTokenType
		// TODO Check if revoked
	}

	return s.generateTokenPair(claims.UserID)
}

func setToken(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:  "",
		Value: "",
	})
}

func (s *Service) Signup(ctx context.Context, email string, password string) (*model.TokenPair, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	// if a user with that email is found
	if user != nil {
		return nil, ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, err
	}

	// Validate email and password
	if err := validateEmail(email); err != nil {
		return nil, err
	}
	if err := validatePassword(password); err != nil {
		return nil, err
	}

	// Create new user
	userId, err := s.repo.PutUser(ctx, email, password)
	if err != nil {
		return nil, err
	}

	// Generate token pair
	pair, err := s.generateTokenPair(userId)
	if err != nil {
		return nil, err
	}

	return pair, nil
}

func (s *Service) Login(ctx context.Context, email string, password string) (*model.TokenPair, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}
	if err := comparePassword(password, user.PasswordHash); err != nil {
		return nil, err
	}

	pair := &model.TokenPair{
		AccessToken:  "",
		RefreshToken: "",
	}
	return pair, nil
}

func (s *Service) Logout(ctx context.Context) {

}
