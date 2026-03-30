package auth

import (
	"testing"
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

		//wantUser model.User
		wantErr error

		repoErr error
	}{
		{
			name:     "success",
			email:    "paulcasigay@gmail.com",
			password: "Dirtycash@123!",
		},
	}
	for _, tt := range tests {
		s.Run(tt.name, func() {

			user, _, err := s.authService.Login(s.ctx, tt.email, tt.password)

			s.Equal(tt.email, user.Email)
			s.Equal(tt.wantErr, err)
			s.Empty(user.PasswordHash)
		})
	}
}

func TestParseToken(t *testing.T) {

}
