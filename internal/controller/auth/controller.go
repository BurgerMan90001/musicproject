package auth

import (
	"golang.org/x/crypto/bcrypt"
	"movieexample.com/internal/repository"
)

type Controller struct {
	repo repository.Repository
}

func New(repo repository.Repository) *Controller {
	return &Controller{repo: repo}
}


func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		panic(err)
	}

	return string(bytes)
}

func CheckPasswordHash(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
