package auth

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

const (
	TokenAccess  = "accessKey"
	TokenRefresh = "refreshKey"
)

var defaultRoles = []string{"user"}

type userRepo interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Put(ctx context.Context, user *model.User) (uuid.UUID, error)
}

type Service struct {
	userRepo   userRepo
	jwtAccess  *JWTService
	jwtRefresh *JWTService
	Google     Oauth
}

func New(ctx context.Context, cfg config.Auth, repo userRepo, sm secrets.Manager) (*Service, error) {
	google, err := NewGoogle(ctx, cfg.Oauth.Google, sm)
	if err != nil {
		return nil, err
	}

	if repo == nil {
		return nil, errors.New("Auth service: nil repo")
	}
	if sm == nil {
		return nil, errors.New("Auth service: nil secret manager")
	}

	jwtAccess, err := NewJWTService("JWT_ACCESS_KEY", "", TokenAccess, []string{}, time.Minute*30)
	if err != nil {
		return nil, err
	}
	jwtRefresh, err := NewJWTService("JWT_REFRESH_KEY", "", TokenRefresh, []string{}, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	return &Service{
		userRepo:   repo,
		jwtAccess:  jwtAccess,
		jwtRefresh: jwtRefresh,
		Google:     google}, nil
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
	tokenPair, err := s.generateTokenPair(userId, defaultRoles)
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
	pair, err := s.generateTokenPair(user.ID, nil)
	if err != nil {
		return nil, nil, err
	}
	// Hide password in response
	user.PasswordHash = ""

	return user, pair, nil
}

// TODO
func (s *Service) Logout(ctx context.Context, accessToken, refeshToken string) error {
	//err := s.jwtRefresh.revokeToken(ctx, refeshToken)

	return nil
}

func (s *Service) Refresh(ctx context.Context, refeshToken string) (*model.TokenPair, error) {
	claims, err := s.jwtRefresh.validateToken(refeshToken)
	if err != nil {
		return nil, err
	}
	if err := s.jwtRefresh.revokeToken(ctx, refeshToken); err != nil {
		return nil, err
	}
	tokenPair, err := s.generateTokenPair(claims.UserID, claims.Roles)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

func (s *Service) ValidateAccess(tokenString string) (*model.Claims, error) {
	return s.jwtAccess.validateToken(tokenString)
}

func (s *Service) generateTokenPair(userId uuid.UUID, roles []string) (*model.TokenPair, error) {
	accessToken, err := s.jwtAccess.generateToken(userId, roles)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwtRefresh.generateToken(userId, roles)
	if err != nil {
		return nil, err
	}
	// TODO Revoke refresh

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
