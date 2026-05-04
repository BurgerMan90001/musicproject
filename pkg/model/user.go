package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID uuid.UUID `json:"id"`
	// Email        string    `json:"email"`
	Username string `json:"username"`
	// PasswordHash string    `json:" - "`
	// Roles     []string  `json:"roles"`
	CreatedAt time.Time `json:"createdAt"`
	AvatarUrl string    `json:"avatarUrl,omitempty"`
}

// type OauthUserInfo struct {
// 	Email   string `json:"email"`
// 	Picture string `json:"picture"`
// }
