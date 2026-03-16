package model

import "github.com/google/uuid"

type User struct {
	ID           uuid.UUID `json:"id"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"passwordHash,omitempty"`
}

type GoogleUserInfo struct {
	Email   string `json:"email"`
	Picture string `json:"picture"`
}

type UserMetadata struct {
	UserID         uuid.UUID `json:"userId"`
	CreatedAt      string    `json:"createdAt"`
	ProfilePicture string    `json:"pfp,omitempty"` // a link to an image
}
