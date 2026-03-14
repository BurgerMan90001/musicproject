package auth

import (
	"regexp"
	"unicode"
)

func ValidateEmail(email string) error {
	valid, err := regexp.MatchString(`^[\w\-\.]+@([\w-]+\.)+[\w-]{2,}$`, email)
	if !valid {
		return ErrInvalidEmail
	}
	return err
}

func ValidatePassword(password string) error {
	if len(password) < minLength {
		return ErrInvalidPassword
	}
	var (
		hasUpper   bool
		hasLower   bool
		hasNumber  bool
		hasSpecial bool
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
