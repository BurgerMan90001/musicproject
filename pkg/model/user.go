package model

type User struct {
	ID           UUID   `json:"id"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}
type GoogleUserInfo struct {
	ID      UUID   `json:"id"`
	Email   string `json:"email"`
	Picture string `json:"picture"`
}
type UserMetadata struct {
	UserID         UUID   `json:"userId"`
	CreatedAt      string `json:"createdAt"`
	ProfilePicture string `json:"pfp,omitempty"` // a link to an image
}

type Login struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
	User         User   `json:"user"`
}
