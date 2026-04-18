package auth

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/oauth2/google"
	"musicproject.com/internal/config"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

var defaultRoles = []string{"user"}

type userRepo interface {
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Put(ctx context.Context, user *model.User) (uuid.UUID, error)
}

//	type refreshTokenRepo interface {
//		RevokeToken(ctx context.Context, refreshToken uui) error
//	}
type Oauth interface {
	Login(ctx context.Context, code string) (*model.User, *model.TokenPair, error)
	RedirectURL(w http.ResponseWriter) string
}

type Service struct {
	userRepo         userRepo
	refreshTokenRepo repository.Token
	jwtAccess        *JWTService
	jwtRefresh       *JWTService
	Google           Oauth
}

func New(ctx context.Context, cfg config.Auth, refreshTokenRepo repository.Token, repo userRepo) (*Service, error) {
	google, err := NewOauth(cfg.Oauth.Google.RedirectURL, cfg.Oauth.Google.Scopes, google.Endpoint)
	if err != nil {
		return nil, err
	}

	if repo == nil {
		return nil, errors.New("Auth service: nil repo")
	}

	jwtAccess, err := NewJWTService(cfg.Jwt, "JWT_ACCESS_KEY", model.TokenAccess, time.Minute*30)
	if err != nil {
		return nil, err
	}
	jwtRefresh, err := NewJWTService(cfg.Jwt, "JWT_REFRESH_KEY", model.TokenRefresh, time.Hour*24*7)
	if err != nil {
		return nil, err
	}
	return &Service{
		userRepo:         repo,
		refreshTokenRepo: refreshTokenRepo,
		jwtAccess:        jwtAccess,
		jwtRefresh:       jwtRefresh,
		Google:           google}, nil
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
func (s *Service) Logout(ctx context.Context, refeshToken string) error {
	claims, err := s.jwtRefresh.ValidateToken(refeshToken)
	if err != nil {
		return err
	}
	userId, err := uuid.Parse(claims.ID)
	if err != nil {
		return err
	}
	if err := s.refreshTokenRepo.RevokeToken(ctx, userId); err != nil {
		return err
	}
	return nil
}

func (s *Service) Refresh(ctx context.Context, refeshToken string) (*model.TokenPair, error) {
	claims, err := s.jwtRefresh.ValidateToken(refeshToken)
	if err != nil {
		return nil, err
	}
	userId, err := uuid.Parse(claims.ID)
	if err != nil {
		return nil, err
	}

	if err := s.refreshTokenRepo.Revoked(ctx, userId); err != nil {
		return nil, err
	}
	if err := s.refreshTokenRepo.RevokeToken(ctx, userId); err != nil {
		return nil, err
	}
	tokenPair, err := s.generateTokenPair(claims.UserID, claims.Roles)
	if err != nil {
		return nil, err
	}

	return tokenPair, nil
}

// Validates access token
func (s *Service) Validate(tokenString string) (*model.Claims, error) {
	return s.jwtAccess.ValidateToken(tokenString)
}

func (s *Service) generateTokenPair(userId uuid.UUID, roles []string) (*model.TokenPair, error) {
	accessToken, err := s.jwtAccess.GenerateToken(userId, roles)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.jwtRefresh.GenerateToken(userId, roles)
	if err != nil {
		return nil, err
	}
	// TODO Revoke refresh

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
