package auth

import "golang.org/x/crypto/bcrypt"

var cost = 14

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
	return nil
}
func ComparePassword(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	return err == nil
}
