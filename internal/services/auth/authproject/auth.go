package auth

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"songsled.com/internal/config"
	"songsled.com/pkg/model"
)

type userRepo interface {
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	PutUser(ctx context.Context, user *model.User) (uuid.UUID, error)
}

type Service struct {
	userRepo  userRepo
	access    *JWTService
	refresh   *JWTService
	blocklist *blocklist
	Google    Oauth
}

func New(ctx context.Context,
	cfg config.Auth,
	rdb *redis.Client,
	userRepo userRepo,
) (*Service, error) {

	if userRepo == nil {
		return nil, errors.New("Auth service: nil repo")
	}

	blocklist, err := newBlocklist(rdb)
	if err != nil {
		return nil, err
	}
	// access, err := NewJWTService(
	// 	cfg.Jwt.Issuer,
	// 	cfg.Jwt.Audience,
	// 	model.TokenAccess,
	// 	time.Minute*30,
	// 	"JWT_ACCESS_KEY",
	// )
	// if err != nil {
	// 	return nil, err
	// }

	// refresh, err := NewJWTService(
	// 	cfg.Jwt.Issuer,
	// 	cfg.Jwt.Audience,
	// 	model.TokenRefresh,
	// 	time.Hour*24*7,
	// 	"JWT_REFRESH_KEY",
	// )
	if err != nil {
		return nil, err
	}
	// oauth, err := NewOauth(ctx) // []string{oidc.ScopeOpenID, "profile", "email"},
	// "http://localhost:8080/realms/songsled",
	// "http://localhost:8081/songsled/callback",
	// "OAUTH_CLIENT_ID",
	// "OAUTH_CLIENT_SECRET",
	// google.Endpoint,

	// if err != nil {
	// 	return nil, err
	// }
	return &Service{
		userRepo: userRepo,
		// access:    access,
		// refresh:   refresh,
		blocklist: blocklist,
		// Google:    oauth,
	}, nil
}

func (s *Service) Signup(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {

	// Validate email and password
	if err := validateEmail(email); err != nil {
		return nil, nil, err
	}
	// passwordHash, err := HashPassword(password)
	// if err != nil {
	// 	return nil, nil, ErrInvalidPassword
	// }
	user := &model.User{
		// Roles: defaultRoles,
		// CreatedAt:    time.Now(),
	}
	// Add the new user
	userId, err := s.userRepo.PutUser(ctx, user)
	if err != nil {
		return nil, nil, err
	}
	// Set password to empty and update user id
	user.ID = userId
	// user.PasswordHash = ""

	// Generate token pair
	// tp, err := s.generateTokenPair(userId, user.Roles...)
	if err != nil {
		return nil, nil, err
	}

	return user, nil, nil
}

func (s *Service) Login(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error) {

	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, ErrIncorrectLogin
	}
	// user is empty
	if user == nil {
		return nil, nil, ErrIncorrectLogin
	}
	// compare passwords
	// if err := ComparePassword(password, user.PasswordHash); err != nil {
	// 	return nil, nil, ErrIncorrectLogin
	// }

	// Generate token pair
	// tp, err := s.generateTokenPair(user.ID, user.Roles...)
	// if err != nil {
	// 	return nil, nil, err
	// }
	// Hide password in response
	// user.PasswordHash = ""

	return user, nil, nil
}

func (s *Service) Logout(ctx context.Context, refeshToken string) error {
	claims, err := s.refresh.ValidateToken(ctx, refeshToken)
	if err != nil {
		return err
	}

	if err := s.blocklist.revokeToken(ctx, claims); err != nil {
		return err
	}
	return nil
}

func (s *Service) Refresh(ctx context.Context, refeshToken string) (*model.TokenPair, error) {
	claims, err := s.refresh.ValidateToken(ctx, refeshToken)
	if err != nil {
		return nil, err
	}

	if err := s.blocklist.revoked(ctx, claims.ID); err != nil {
		return nil, err
	}
	if err := s.blocklist.revokeToken(ctx, claims); err != nil {
		return nil, err
	}

	return s.generateTokenPair(claims.UserID, claims.Roles...)
}

// Validates access token
func (s *Service) Validate(ctx context.Context, tokenString string, needRoles ...string) (*model.Claims, error) {
	claims, err := s.access.ValidateToken(ctx, tokenString)
	if err != nil {
		return nil, &model.Error{
			Code:    http.StatusUnauthorized,
			Message: "Invalid access token",
			Details: err.Error(),
		}
	}

	if err := s.blocklist.revoked(ctx, claims.ID); err != nil {
		return nil, err
	}

	for _, role := range needRoles {
		if !slices.Contains(claims.Roles, role) {
			return nil, &model.Error{
				Code:    http.StatusUnauthorized,
				Message: "Invalid access token",
				Details: fmt.Sprintf("Need role %v", role),
			}
		}
	}
	return claims, nil
}

func (s *Service) generateTokenPair(userId uuid.UUID, roles ...string) (*model.TokenPair, error) {
	accessToken, err := s.access.GenerateToken(userId, roles...)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.refresh.GenerateToken(userId, roles...)
	if err != nil {
		return nil, err
	}

	return &model.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
