package auth

import (
	"musicproject.com/internal/repository"
)

// func TestGenerateToken(t *testing.T) {
// 	tests := []struct {
// 		name string
// 	}{}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {

//			})
//		}
//	}
func (s *testSuite) TestLogin() {
	tests := []struct {
		name     string
		email    string
		password string

		wantErr error
	}{
		{
			name:     "no email",
			password: "Dirtycash@123!",
			wantErr:  repository.ErrNotFound,
		},
		{
			name:    "no password",
			email:   "paulcasigay@gmail.com",
			wantErr: ErrIncorrectLogin,
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			user, _, err := s.authService.Login(s.ctx, tt.email, tt.password)

			s.Equal(tt.wantErr, err)
			s.Nil(user)
		})
	}
	s.Run("login success", func() {
		user, tokenPair, err := s.authService.Login(s.ctx, "paulcasigay@gmail.com", "Dirtycash@123!")
		s.NoError(err)

		s.Equal("paulcasigay@gmail.com", user.Email)
		s.Empty(user.PasswordHash)
		s.NotEmpty(tokenPair)
		s.NotEmpty(tokenPair.AccessToken)
		s.NotEmpty(tokenPair.RefreshToken)
	})
}

// func TestParseToken(t *testing.T) {

// }
