package auth

import "golang.org/x/crypto/bcrypt"

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}

func CheckPasswordHash(password string, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(passwordHash),
		[]byte(password),
	)
	return err == nil
}
