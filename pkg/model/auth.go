package model

type SignupRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// type SignupResponse struct {
// 	AccessToken string `json:"accessToken"`
// }

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type TokenPair struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}
