package auth

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"songsled.com/pkg/model"
)

type blocklist struct {
	rdb *redis.Client
}

func newBlocklist(rdb *redis.Client) (*blocklist, error) {
	if rdb == nil {
		return nil, fmt.Errorf("newBlocklist: Redis.Client is nil")
	}
	return &blocklist{rdb}, nil
}

func (r *blocklist) revokeToken(ctx context.Context, claims *model.Claims) error {
	ttl := time.Until(claims.ExpiresAt.Time)
	// Already expired
	if ttl <= 0 {
		return nil
	}
	// jti is the token's id
	key := fmt.Sprintf("blocklist:jti:%s", claims.ID)

	return r.rdb.Set(ctx, key, "1", ttl).Err()
}

func (r *blocklist) revoked(ctx context.Context, jti string) error {
	exists, err := r.rdb.Exists(ctx, fmt.Sprintf("blocklist:jti:%s", jti)).Result()
	if err != nil {
		return fmt.Errorf("blocklist.revoked: %w", err)

	}
	// Token revoked
	if exists > 0 {
		return ErrInvalidToken(err)
	}
	return nil
}

type JWTService struct {
	key       []byte
	tokenType model.TokenType
	issuer    string
	audience  []string
	ttl       time.Duration
}

func NewJWTService(
	issuer string,
	audience []string,
	tokenType model.TokenType,
	ttl time.Duration,
	envVar string,
) (*JWTService, error) {
	key := os.Getenv(envVar)
	if len(key) < 32 {
		return nil, fmt.Errorf("NewJWTService: env var %q must be at least 32 bytes long", envVar)
	}
	if len(audience) == 0 || audience[0] == "" {
		return nil, fmt.Errorf("NewJWTService: jwt audience is empty")
	}
	switch tokenType {
	// Valid types
	case model.TokenAccess, model.TokenRefresh:
	default:
		return nil, fmt.Errorf("NewJWTService: invalid tokenType %v", tokenType)
	}

	return &JWTService{
		key:       []byte(key),
		issuer:    issuer,
		tokenType: tokenType,
		audience:  audience,
		ttl:       ttl,
	}, nil
}
func (s *JWTService) GenerateToken(userId uuid.UUID, roles ...string) (string, error) {
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

func (s *JWTService) ValidateToken(ctx context.Context, tokenString string) (*model.Claims, error) {

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
		return nil, &model.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid jwt token",
			Details: err.Error(),
		}
	}
	claims, ok := token.Claims.(*model.Claims)
	switch {
	case !ok || !token.Valid,
		claims.TokenType != string(s.tokenType):
		return nil, ErrInvalidToken(err)

	}

	return claims, nil
}
