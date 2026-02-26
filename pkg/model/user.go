package model

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
}

type UserMetadata struct {
	UserID         string `json:"user_id"`
	ProfilePicture string `json:"pfp"` // a link to an image
}
