package auth

import (
	"errors"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/golang-jwt/jwt/v5/request"
	"golang.org/x/crypto/bcrypt"
)

var ErrInvalidToken = errors.New("invalid token")

func JWTParseToken(jwtKey []byte, r *http.Request) (*Claims, error) {
	token, err := request.ParseFromRequest(r, request.AuthorizationHeaderExtractor, func(t *jwt.Token) (any, error) {
		return jwtKey, nil
	}, request.WithClaims(&Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "okapi",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}))
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrInvalidToken
}

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func CheckPasswordHash(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash), 
		[]byte(password),
	)
	return err == nil
}
