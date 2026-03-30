package auth

import (
	"context"
	"errors"
	"time"

	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/services"
	"musicproject.com/pkg/model"
)

const (
	TokenAccess  = "access"
	TokenRefresh = "refresh"
)

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)
var ExpiresInOneHour = time.Now().Add(time.Hour * 24)

type Service struct {
	cfg  config.Auth
	repo repository.Repository

	JWT    services.JWT
	Google services.Oauth
}

func New(cfg config.Auth, repo repository.Repository) *Service {
	google := NewGoogle(cfg.Oauth.Google)
	JWT := NewJWTService(cfg.JWT)
	return &Service{cfg, repo, JWT, google}
}

func (s *Service) Signup(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {
	u, err := s.repo.GetUserByEmail(ctx, email)
	// if a user with that email is found
	if u != nil {
		return nil, nil, ErrUserAlreadyExists
	}
	if err != nil && !errors.Is(err, repository.ErrNotFound) {
		return nil, nil, err
	}

	// Validate email and password
	if err := validateEmail(email); err != nil {
		return nil, nil, err
	}
	passwordHash, err := HashPassword(password)
	if err != nil {
		return nil, nil, ErrInvalidPassword
	}
	// Add the new user
	userId, err := s.repo.PutUser(ctx, email, passwordHash)
	if err != nil {
		return nil, nil, err
	}

	// Generate token pair
	tokenPair, err := s.JWT.GenerateTokenPair(userId)
	if err != nil {
		return nil, nil, err
	}
	user := &model.User{
		ID:    userId,
		Email: email,
	}

	return user, tokenPair, nil
}

func (s *Service) Login(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}
	if user == nil {
		return nil, nil, repository.ErrNotFound
	}
	if err := ComparePassword(password, user.PasswordHash); err != nil {
		return nil, nil, ErrIncorrectLogin
	}
	pair, err := s.JWT.GenerateTokenPair(user.ID)
	if err != nil {
		return nil, nil, err
	}
	// Hide password in response
	user.PasswordHash = ""

	return user, pair, nil
}

func (s *Service) Logout(ctx context.Context) {

}
