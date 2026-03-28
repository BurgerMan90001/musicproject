package services

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"golang.org/x/oauth2"
	"musicproject.com/pkg/model"
)

type Auth interface {
	GenerateToken(userId uuid.UUID, tokenType string, expireAt time.Time) (string, error)
	GenerateTokenPair(userId uuid.UUID) (*model.TokenPair, error)

	ParseAccessToken(accessToken string) (*model.Claims, error)
	ParseToken(tokenString string, tokenType string) (*jwt.Token, error)

	RevokeToken(ctx context.Context, tokenString string) error
	RefreshTokens(ctx context.Context, refreshToken string) (*model.TokenPair, error)

	Signup(ctx context.Context, email string, password string) (*model.TokenPair, error)
	Login(ctx context.Context, email string, password string) (*model.TokenPair, error)
	Logout(ctx context.Context)
}

type Oauth interface {
	GetUserInfo(ctx context.Context, token *oauth2.Token) (*model.OauthUserInfo, error)
	RedirectURL(w http.ResponseWriter) string
}

type File interface {
	Backup(ctx context.Context) error
	UploadSong(ctx context.Context, song *model.Song) error
}

type Rating interface {
	GetAggregatedRating(ctx context.Context, songId uuid.UUID) (float64, error)
	PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error
}

type Search interface {
}
