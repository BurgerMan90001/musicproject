package auth

import (
	"fmt"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	cost = 14

	minLength = 8
	maxLength = 255
)

func HashPassword(password string) (string, error) {
	// Validate password first
	if err := ValidatePassword(password); err != nil {
		return "", err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
func ValidatePassword(password string) error {
	if len(password) < minLength {
		return fmt.Errorf("password less than %d characters", minLength)
	}
	for _, char := range password {
		switch {
		case unicode.IsDigit(char):

		}
	}
	return nil
}
func ComparePassword(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	return err == nil
}
