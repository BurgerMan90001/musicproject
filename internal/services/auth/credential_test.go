package auth

func (s *testSuite) TestValidatePassword() {
	tests := []struct {
		name     string
		password string
		wantErr  error
	}{
		{name: "valid password", password: "Gooooop123!", wantErr: nil},
		{name: "no uppercase letters", password: "gooooop123!", wantErr: ErrInvalidPassword},
		{name: "no special characters", password: "Eooooop123", wantErr: ErrInvalidPassword},
		{name: "less than 8 characters", password: "Eee4@", wantErr: ErrInvalidPassword},
		{name: "no numbers", password: "Eeeeeeeeeeee@", wantErr: ErrInvalidPassword},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validatePassword(tt.password)
			s.Equal(tt.wantErr, err, tt.name)
		})
	}
}

func (s *testSuite) TestValidateEmail() {
	tests := []struct {
		name    string
		email   string
		wantErr error
	}{
		{name: "valid email", email: "paulcasiga@gmail.com", wantErr: nil},
		{name: "no @ symbol", email: "yoopgmail.com", wantErr: ErrInvalidEmail},
		{name: "no . seperator", email: "yoop@gmailcom", wantErr: ErrInvalidEmail},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			err := validateEmail(tt.email)
			s.Equal(tt.wantErr, err, tt.name)
		})
	}
}
