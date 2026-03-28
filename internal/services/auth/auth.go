package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services"
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

	Google services.Oauth
}

func New(cfg config.Auth, repo repository.Repository) *Service {
	google := NewGoogle(cfg.Oauth.Google)

	return &Service{cfg, repo, google}
}

func (s *Service) GenerateToken(userId uuid.UUID, tokenType string, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
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

func (s *Service) GenerateTokenPair(userId uuid.UUID) (*model.TokenPair, error) {
	accessToken, err := s.GenerateToken(userId, TokenAccess, ExpiresInOneHour)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.GenerateToken(userId, TokenAccess, ExpiresInOneHour)
	if err != nil {
		return nil, err
	}
	// TODO Revoke refresh

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *Service) ParseToken(tokenString string, tokenType string) (*jwt.Token, error) {
	key, err := s.keyType(tokenType)
	if err != nil {
		return nil, err
	}
	token, err := jwt.ParseWithClaims(tokenString,
		&model.Claims{},
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
	if err != nil {
		if errors.Is(err, jwt.ErrTokenMalformed) {
			return nil, jwt.ErrTokenMalformed
		}
		return nil, err
	}
	return token, nil
}

func (s *Service) ParseAccessToken(accessToken string) (*model.Claims, error) {
	token, err := s.ParseToken(accessToken, TokenAccess)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, jwt.ErrTokenExpired
		}
		return nil, err
	}

	claims, ok := token.Claims.(*model.Claims)

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
	token, err := s.ParseToken(tokenString, TokenRefresh)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*model.Claims)
	switch {
	case !ok || !token.Valid:
		return nil, jwt.ErrTokenInvalidClaims

	case claims.TokenType != TokenRefresh:
		return nil, ErrInvalidTokenType
		// TODO Check if revoked
	}

	return s.GenerateTokenPair(claims.UserID)
}
func (s *Service) RevokeToken(ctx context.Context, tokenString string) error {
	return nil
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
	passwordHash, err := HashPassword(password)
	if err != nil {
		return nil, err
	}
	// Create new user
	userId, err := s.repo.PutUser(ctx, email, passwordHash)
	if err != nil {
		return nil, err
	}

	// Generate token pair
	pair, err := s.GenerateTokenPair(userId)
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
	if user == nil {
		return nil, repository.ErrNotFound
	}
	if err := comparePassword(password, user.PasswordHash); err != nil {
		return nil, err
	}
	pair, err := s.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, err
	}

	return pair, nil
}

func (s *Service) Logout(ctx context.Context) {

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
