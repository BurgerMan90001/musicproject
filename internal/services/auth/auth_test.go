package auth

import (
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
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
	
	s.T().Skip()
	tests := []struct {
		name     string
		email    string
		password string

		wantUser *model.User
		wantErr  error

		repoErr error
	}{
		{
			name:     "success",
			email:    "paulcasigay@gmail.com",
			password: "Dirtycash@123!",
			wantUser: &model.User{
				Email: "paulcasigay@gmail.com",
			},
		},
		{
			name:     "no email",
			password: "Dirtycash@123!",

			wantErr: repository.ErrNotFound,
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

			if tt.wantUser != nil {
				s.Equal(tt.wantUser.Email, user.Email)
				s.Empty(user.PasswordHash)
			}

		})
	}
}

// func TestParseToken(t *testing.T) {

// }
