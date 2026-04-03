package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidatePassword(t *testing.T) {
	t.Parallel()
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
		t.Run(tt.name, func(t *testing.T) {
			err := validatePassword(tt.password)
			assert.ErrorIs(t, tt.wantErr, err, tt.name)
		})
	}
}

func TestValidateEmail(t *testing.T) {
	t.Parallel()
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
		t.Run(tt.name, func(t *testing.T) {
			err := validateEmail(tt.email)
			assert.ErrorIs(t, tt.wantErr, err, tt.name)
		})
	}
}
