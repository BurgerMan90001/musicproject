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
		wantErr  bool
	}{
		{name: "valid password", password: "Gooooop123!", wantErr: false},
		{name: "no uppercase letters", password: "gooooop123!", wantErr: true},
		{name: "no special characters", password: "Eooooop123", wantErr: true},
		{name: "less than 8 characters", password: "Eee4@", wantErr: true},
		{name: "no numbers", password: "Eeeeeeeeeeee@", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePassword(tt.password)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
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
		{name: "valid email", email: "paulcasiga@gmail.com"},
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
