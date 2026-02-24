package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           string `json:"id"`
	Username     string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	jwt.Claims
}
