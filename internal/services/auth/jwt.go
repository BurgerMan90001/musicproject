package auth

import (
	"context"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/pkg/model"
)

type JWTService struct {
	issuer     string
	accessKey  []byte
	refreshKey []byte
}

func NewJWTService(ctx context.Context, sm secrets.Manager) (*JWTService, error) {
	var (
		issuer, issuerErr      = sm.Get(ctx, "JWT_ISSUER")
		accessKey, accessErr   = sm.Get(ctx, "JWT_ACCESS_KEY")
		refreshKey, refreshErr = sm.Get(ctx, "JWT_REFRESH_KEY")
	)
	if err := errors.Join(issuerErr, accessErr, refreshErr); err != nil {
		return nil, err
	}
	return &JWTService{
		issuer:     issuer,
		accessKey:  []byte(accessKey),
		refreshKey: []byte(refreshKey),
	}, nil
}
func (s *JWTService) generateToken(userId uuid.UUID, tokenType string, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &model.Claims{
		UserID:    userId,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    s.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	})

	key, err := s.keyFunc(tokenType)
	if err != nil {
		return "", err
	}
	day := time.Now().Weekday()
	switch day {
	case time.Thursday:

	}

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *JWTService) generateTokenPair(userId uuid.UUID) (*model.TokenPair, error) {
	accessToken, err := s.generateToken(userId, TokenAccess, ExpiresInOneHour)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateToken(userId, TokenAccess, ExpiresInOneHour)
	if err != nil {
		return nil, err
	}
	// TODO Revoke refresh

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
func (s *JWTService) revokeToken(ctx context.Context, tokenString string) error {
	return nil
}

func (s *JWTService) validateToken(tokenString string, tokenType string) (*jwt.Token, error) {
	key, err := s.keyFunc(tokenType)
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
		jwt.WithIssuer(s.issuer),
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

func (s *JWTService) validateAccessToken(accessToken string) (*model.Claims, error) {
	token, err := s.validateToken(accessToken, TokenAccess)
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

func (s *JWTService) refreshTokens(ctx context.Context, refeshToken string) (*model.TokenPair, error) {
	token, err := s.validateToken(refeshToken, TokenRefresh)
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

	return s.generateTokenPair(claims.UserID)
}

func (s *JWTService) keyFunc(tokenType string) ([]byte, error) {
	switch tokenType {
	case TokenAccess:
		return s.accessKey, nil
	case TokenRefresh:
		return []byte(s.refreshKey), nil
	default:
		return nil, ErrInvalidTokenType
	}
}
