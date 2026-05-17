package postgres

import (
	"context"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type User struct {
	queries *gensqlc.Queries
}

func NewUser(queries *gensqlc.Queries) *User {
	return &User{queries}
}

// Gets a user's email and password hash by their uuid
func (r *User) GetUserByID(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	// u, err := r.queries.GetUserByID(ctx, userId)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, repository.ErrNotFound
	// 	}
	// }
	// return &model.User{
	// 	ID:       userId,
	// 	Username: u.Username,
	// 	// Email:     u.Email,
	// 	AvatarUrl: u.AvatarUrl.String,

	// 	// PasswordHash: u.PasswordHash.String,
	// }, nil
	return nil, nil
}

func (r *User) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	// u, err := r.queries.GetUserByEmail(ctx, email)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		return nil, repository.ErrUserNotFound
	// 	}
	// 	return nil, err
	// }

	// Null password
	// if !u.PasswordHash.Valid {

	// }
	// return &model.User{
	// 	ID: u.UserID,
	// 	// Email:        email,
	// 	// PasswordHash: u.PasswordHash.String,

	// 	// AvatarURL:    u.AvatarURL,
	// }, nil
	return nil, nil
}
func (r *User) PutUser(ctx context.Context, user *model.User) (uuid.UUID, error) {
	// Password will be null if empty
	// valid := user.PasswordHash != ""
	// userId, err := r.queries.PutUser(ctx, gensqlc.PutUserParams{
	// 	Email: user.Email, PasswordHash: sql.NullString{String: user.PasswordHash, Valid: valid},
	// })

	// if pgerr, ok := err.(*pq.Error); ok {
	// 	// The email is already taken
	// 	if pgerr.Code == pqerror.UniqueViolation {
	// 		return uuid.Nil, repository.ErrEmailTaken
	// 	}
	// }

	return uuid.Nil, nil
}

func (r *User) DeleteUserByID(ctx context.Context, userId uuid.UUID) error {
	// return r.queries.DeleteUserByID(ctx, userId)
	return nil
}
