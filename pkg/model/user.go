package model

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password,omitempty"`
	CreatedAt    string    `json:"createdAt"`
	AvatarURL    string    `json:"avatarUrl,omitempty"`
}

type OauthUserInfo struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
}
