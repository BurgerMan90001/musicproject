package model

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refreshToken"`
}

type EmailResetRequest struct {
	Email string `json:"email"`
}

type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type Claims struct {
	UserID    uuid.UUID `json:"userId"`
	TokenType string    `json:"tokenType"`
	jwt.RegisteredClaims
}
