package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	jwt.Claims
}

type UserMetadata struct {
	ProfilePicture string `json:"pfp"` // a link to an image
}
