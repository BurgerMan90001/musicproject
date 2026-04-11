package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

type JWTService struct {
	key       []byte
	tokenType string
	issuer    string
	audience  []string
	ttl       time.Duration
}

func NewJWTService(envVar, issuer, tokenType string, audience []string, ttl time.Duration) (*JWTService, error) {
	key := os.Getenv(envVar)
	if len(key) < 32 {
		return nil, fmt.Errorf("env var %q must be at least 32 bytes long", envVar)
	}
	switch tokenType {
	case TokenAccess, TokenRefresh:
	default:
		return nil, ErrInvalidTokenType
	}
	return &JWTService{
		key:       []byte(key),
		issuer:    issuer,
		tokenType: tokenType,
		audience:  audience,
		ttl:       ttl,
	}, nil
}
func (s *JWTService) generateToken(userId uuid.UUID, roles []string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
		UserID:    userId,
		TokenType: s.tokenType,
		Roles:     roles,
		RegisteredClaims: jwt.RegisteredClaims{
			// ID is for revocation
			ID:        uuid.NewString(),
			Issuer:    s.issuer,
			Audience:  s.audience,
			Subject:   userId.String(),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
			ExpiresAt: jwt.NewNumericDate(now.Add(s.ttl)),
		},
	})
	return token.SignedString(s.key)
}

func (s *JWTService) validateToken(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&model.Claims{},
		func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
			}
			return s.key, nil
		},
		jwt.WithIssuer(s.issuer),
		jwt.WithValidMethods([]string{"HS256"}),
		jwt.WithExpirationRequired(),
		jwt.WithAudience(s.audience[0]),
	)
	if err != nil {
		return nil, fmt.Errorf("validate token: %v", err)
	}
	claims, ok := token.Claims.(*model.Claims)
	switch {
	case !ok || !token.Valid:
		return nil, jwt.ErrTokenInvalidClaims
	case claims.TokenType != s.tokenType:
		return nil, ErrInvalidTokenType
	}
	return claims, nil
}

func (s *JWTService) revokeToken(ctx context.Context, tokenString string) error {
	return nil
}

// 	return s.generateTokenPair(claims.UserID)
// }

// func (s *JWTService) keyFunc(tokenType string) ([]byte, error) {
// 	switch tokenType {
// 	case TokenAccess:
// 		return s.accessKey, nil
// 	case TokenRefresh:
// 		return []byte(s.refreshKey), nil
// 	default:
// 		return nil, ErrInvalidTokenType
// 	}
// }
