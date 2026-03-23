package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateToken(t *testing.T) {
	tests := []struct {
		name string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
func TestSignup(t *testing.T) {
	tests := []struct {
		name     string
		email    string
		password string
	}{
		{},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, "", "")
		})
	}
}
func TestParseToken(t *testing.T) {

}
