package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/config"
)

const (
	TokenAccess  = "access"
	TokenRefresh = "refresh"
)

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)

type Service struct {
	AccessKey  []byte
	RefreshKey []byte
	Issuer     string
}

type Claims struct {
	UserID    uuid.UUID `json:"userId"`
	TokenType string    `json:"tokenType"`
	jwt.RegisteredClaims
}

func New(cfg config.Config) *Service {
	return &Service{
		[]byte(cfg.API.Jwt.AccessKey),
		[]byte(cfg.API.Jwt.RefreshKey),
		cfg.API.Jwt.Issuer,
	}
}

func (s *Service) GenerateToken(userId uuid.UUID, tokenType string, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:    userId,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.Issuer,
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

func (s *Service) RefreshTokens(ctx context.Context, tokenString string) {
	token, err := s.parseToken(tokenString, TokenRefresh)
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

	case claims.TokenType != TokenRefresh:
		return nil, ErrInvalidTokenType

	default:
		return claims, nil
	}
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
		jwt.WithIssuer(s.Issuer),
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
	)
	return token, err
}
func (s *Service) keyType(tokenType string) ([]byte, error) {
	switch tokenType {
	case TokenAccess:
		return s.AccessKey, nil
	case TokenRefresh:
		return s.RefreshKey, nil
	default:
		return nil, ErrInvalidTokenType
	}
}

func SetToken() {

}

func (s *Service) Signup() {

}

func (s *Service) Login() {

}
