package services

import (
	"context"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"musicproject.com/pkg/model"
)

type Auth interface {
	Signup(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error)
	Login(ctx context.Context, email string, password string) (*model.User, *model.TokenPair, error)
	Logout(ctx context.Context)
}
type JWT interface {
	GenerateToken(userId uuid.UUID, tokenType string, expireAt time.Time) (string, error)
	GenerateTokenPair(userId uuid.UUID) (*model.TokenPair, error)

	ParseAccessToken(accessToken string) (*model.Claims, error)
	ParseToken(tokenString string, tokenType string) (*jwt.Token, error)

	RevokeToken(ctx context.Context, tokenString string) error
	RefreshTokens(ctx context.Context, refreshToken string) (*model.TokenPair, error)
}
type Oauth interface {
	Login(ctx context.Context, code string) (*model.User, *model.TokenPair, error)
	//GetUserInfo(ctx context.Context, token *oauth2.Token) (*model.OauthUserInfo, error)
	RedirectURL(w http.ResponseWriter) string
	//Exchange(ctx context.Context, code string) (*oauth2.Token, error)
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
