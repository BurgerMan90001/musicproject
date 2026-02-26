package auth

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	jwt.RegisteredClaims
}
type Controller struct {
	jwtKey []byte
}

func New(secretKey []byte) *Controller {
	return &Controller{
		jwtKey: secretKey,
	}
}
func (c *Controller) GenerateToken(claims *Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(c.jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (c *Controller) VerifyToken(tokenString string) error {
	claims := &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "okapi",
		},
	}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (any, error) {
		return c.jwtKey, nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))
	if err != nil {
		return fmt.Errorf("failed to parse token: %v", err)
	} else if _, ok := token.Claims.(*Claims); ok {
		// token is valid
		return nil
	} else {
		return fmt.Errorf("unknown claims type: %v", err)
	}
}
