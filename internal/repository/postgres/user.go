package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lib/pq/pqerror"
	"songsled.com/internal/repository"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type UserRepo struct {
	queries *gensqlc.Queries
}

func NewUser(queries *gensqlc.Queries) *UserRepo {
	return &UserRepo{queries}
}

// Gets a user's email and password hash by their uuid
func (r *UserRepo) GetUserByID(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	u, err := r.queries.GetUserByID(ctx, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
	}

	return &model.User{
		ID:           userId,
		Email:        u.Email,
		PasswordHash: u.PasswordHash.String,
	}, nil
}

func (r *UserRepo) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	u, err := r.queries.GetUserByEmail(ctx, email)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrUserNotFound
		}
		return nil, err
	}

	// Null password
	// if !u.PasswordHash.Valid {

	// }
	return &model.User{
		ID:           u.UserID,
		Email:        email,
		PasswordHash: u.PasswordHash.String,

		// AvatarURL:    u.AvatarURL,
	}, nil
}
func (r *UserRepo) PutUser(ctx context.Context, user *model.User) (uuid.UUID, error) {
	// Password will be null if empty
	valid := user.PasswordHash != ""
	userId, err := r.queries.PutUser(ctx, gensqlc.PutUserParams{
		Email: user.Email, PasswordHash: sql.NullString{String: user.PasswordHash, Valid: valid},
	})

	if pgerr, ok := err.(*pq.Error); ok {
		// The email is already taken
		if pgerr.Code == pqerror.UniqueViolation {
			return uuid.Nil, repository.ErrEmailTaken
		}
	}

	return userId, nil
}

func (r *UserRepo) DeleteUserByID(ctx context.Context, userId uuid.UUID) error {
	return r.queries.DeleteUserByID(ctx, userId)
}
