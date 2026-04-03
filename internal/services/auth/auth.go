package auth

import (
	"context"
	"errors"
	"time"

	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

const (
	TokenAccess  = "access"
	TokenRefresh = "refresh"
)

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)
var ExpiresInOneHour = time.Now().Add(time.Hour * 24)

type Service struct {
	cfg      config.Auth
	userRepo repository.User

	JWT    *JWTService
	Google Oauth
}

func New(cfg config.Auth, repo repository.User) *Service {
	google := NewGoogle(cfg.Oauth.Google)

	JWT := NewJWTService(cfg.Jwt)
	return &Service{cfg, repo, JWT, google}
}

func (s *Service) Signup(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
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
	user := &model.User{
		Email:        email,
		PasswordHash: passwordHash,
	}
	// Add the new user
	userId, err := s.userRepo.Put(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	// Set password to empty and set user id
	user.ID = userId
	user.PasswordHash = ""

	// Generate token pair
	tokenPair, err := s.JWT.GenerateTokenPair(userId)
	if err != nil {
		return nil, nil, err
	}

	return user, tokenPair, nil
}

func (s *Service) Login(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {
	user, err := s.userRepo.GetByEmail(ctx, email)
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

func (s *Service) Logout(ctx context.Context, refeshToken string) error {

	return nil
}

func (s *Service) Refresh(ctx context.Context, refeshToken string) (*model.TokenPair, error) {
	tokenPair, err := s.JWT.refreshTokens(ctx, refeshToken)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}
