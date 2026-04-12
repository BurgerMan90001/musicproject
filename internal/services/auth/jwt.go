package auth

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/internal/config"
	"musicproject.com/pkg/model"
)

type JWTService struct {
	key       []byte
	tokenType model.TokenType
	issuer    string
	audience  []string
	ttl       time.Duration
}

func NewJWTService(cfg config.Jwt, envVar string, tokenType model.TokenType, ttl time.Duration) (*JWTService, error) {
	key := os.Getenv(envVar)
	if len(key) < 32 {
		return nil, fmt.Errorf("env var %q must be at least 32 bytes long", envVar)
	}
	if len(cfg.Audience) == 0 || cfg.Audience[0] == "" {
		return nil, fmt.Errorf("jwt audience is empty")
	}
	switch tokenType {
	case model.TokenAccess, model.TokenRefresh:

	default:
		return nil, ErrInvalidTokenType
	}
	return &JWTService{
		key:       []byte(key),
		issuer:    cfg.Issuer,
		tokenType: tokenType,
		audience:  cfg.Audience,
		ttl:       ttl,
	}, nil
}
func (s *JWTService) GenerateToken(userId uuid.UUID, roles []string) (string, error) {
	// if len(opts) > 0 {
	// 	for _, apply := range opts {
	// 		apply(s)
	// 	}
	// }
	if roles == nil {
		roles = defaultRoles
	}
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
		UserID:    userId,
		TokenType: string(s.tokenType),
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

func (s *JWTService) ValidateToken(tokenString string) (*model.Claims, error) {
	if tokenString == "" {
		return nil, ErrNoRefeshToken
	}
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
	case claims.TokenType != string(s.tokenType):
		return nil, ErrInvalidTokenType
	}
	return claims, nil
}

func (s *JWTService) revokeToken(ctx context.Context, tokenString string) error {
	return nil
}
