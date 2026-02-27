package model

type User struct {
	ID           string `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

type UserMetadata struct {
	UserID         string `json:"userId"`
	CreatedAt      string `json:"createdAt"`
	ProfilePicture string `json:"pfp"` // a link to an image
}
