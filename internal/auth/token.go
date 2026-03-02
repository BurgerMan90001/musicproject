package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"okapi.com/pkg/model"
)

var issuer = "okapi"

var ErrInvalidToken = errors.New("invalid token")

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)

type Claims struct {
	UserID string `json:"userId"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(jwtKey []byte, user *model.User, expireAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		UserID: user.ID,
		Email:  user.Email,
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
		request.OAuth2Extractor,
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
