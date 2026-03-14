package auth

import (
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

var issuer = "okapi"

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)

const (
	TokenAccess  = "access"
	TokenRefresh = "refresh"
)

type Claims struct {
	UserID    uuid.UUID `json:"userId"`
	Email     string    `json:"email"`
	TokenType string    `json:"tokenType"`
	jwt.RegisteredClaims
}

func GenerateToken(jwtKey []byte, user *model.User, tokenType string, expireAt time.Time) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID:    user.ID,
		Email:     user.Email,
		TokenType: tokenType,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			ExpiresAt: jwt.NewNumericDate(expireAt),
		},
	})
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(jwtKey []byte, r *http.Request) (*Claims, error) {
	extractor := request.MultiExtractor{
		request.AuthorizationHeaderExtractor,
	}

	token, err := request.ParseFromRequest(r, extractor, func(t *jwt.Token) (any, error) {
		return jwtKey, nil
	}, request.WithClaims(&Claims{}))

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}
