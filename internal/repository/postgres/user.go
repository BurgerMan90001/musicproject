package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

type User struct {
	db *sql.DB
}

var _ repository.User = (*User)(nil)

func NewUser(db *sql.DB) *User {
	return &User{db}
}

// Gets a user's email and password hash by their uuid
func (r *User) GetByID(ctx context.Context, userId uuid.UUID) (*model.User, error) {
	var (
		email        string
		passwordHash string
	)
	query := "SELECT email, password_hash FROM users WHERE user_id=$1"
	row := r.db.QueryRowContext(ctx, query, userId)
	if err := row.Scan(&email, &passwordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &model.User{
		ID:           userId,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}

func (r *User) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	var (
		userId       uuid.UUID
		passwordHash string
	)
	query := "SELECT user_id, password_hash FROM users WHERE email=$1"
	row := r.db.QueryRowContext(ctx, query, email)
	if err := row.Scan(&userId, &passwordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &model.User{
		ID:           userId,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}
func (r *User) Put(ctx context.Context, user *model.User) (uuid.UUID, error) {
	query := "INSERT INTO users (email, password_hash) VALUES($1, $2) RETURNING user_id"

	row := r.db.QueryRowContext(ctx, query, user.Email, user.PasswordHash)
	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *User) DeleteByID(ctx context.Context, userId uuid.UUID) error {
	query := "DELETE FROM users WHERE user_id=$1"
	_, err := r.db.ExecContext(ctx, query, userId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	return nil
}
