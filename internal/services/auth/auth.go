package auth

import (
	"context"
	"errors"
	"time"

	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

const (
	TokenAccess  = "accessKey"
	TokenRefresh = "refreshKey"
)

var ExpiresInOneDay = time.Now().Add(time.Hour * 24)
var ExpiresInOneHour = time.Now().Add(time.Hour * 24)

type Service struct {
	userRepo repository.User

	JWT    *JWTService
	Google Oauth
}

func New(ctx context.Context, cfg config.Auth, repo repository.User, sm secrets.Manager) (*Service, error) {
	google, err := NewGoogle(ctx, cfg.Oauth.Google, sm)
	if err != nil {
		return nil, err
	}

	JWT, err := NewJWTService(ctx, sm)
	if err != nil {
		return nil, err
	}
	return &Service{repo, JWT, google}, nil
}

func (s *Service) Signup(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {
	u, err := s.userRepo.GetByEmail(ctx, email)
	// if a user with that email is found
	if u != nil {
		return nil, nil, err
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
		CreatedAt:    time.Now(),
	}
	// Add the new user
	userId, err := s.userRepo.Put(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	// Set password to empty and update user id
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
	// user is empty
	if user == nil {
		return nil, nil, repository.ErrNotFound
	}
	// compare passwords
	if err := ComparePassword(password, user.PasswordHash); err != nil {
		return nil, nil, ErrIncorrectLogin
	}
	// Generate token pair
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
