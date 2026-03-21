package auth

import (
	"regexp"
	"unicode"

	"golang.org/x/crypto/bcrypt"
)

const (
	minLength = 8
	maxLength = 255

	cost = 14
)

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
func comparePassword(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	return err
}

func validateCredentials(email string, password string) error {
	if err := validateEmail(email); err != nil {
		return err
	}
	if err := validatePassword(password); err != nil {
		return err
	}
	return nil
}
func validateEmail(email string) error {
	valid, err := regexp.MatchString(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`, email)
	if !valid {
		return ErrInvalidEmail
	}
	return err
}

func validatePassword(password string) error {
	if len(password) < minLength {
		return ErrInvalidPassword
	}
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}
	if !(hasUpper && hasLower && hasNumber && hasSpecial) {
		return ErrInvalidPassword
	}

	return nil
}
