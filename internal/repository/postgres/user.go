package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

// type User struct {
// 	db *sql.DB
// }

// Gets a user's email and password hash by their uuid
func (r *Repository) GetUserByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	var (
		email        string
		passwordHash string
	)
	query := "SELECT email, password_hash FROM users WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&email, &passwordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &model.User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}

func (r *Repository) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var (
		id           uuid.UUID
		passwordHash string
	)
	query := "SELECT id, password_hash FROM users WHERE email=$1"
	row := r.db.QueryRowContext(ctx, query, email)
	if err := row.Scan(&id, &passwordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &model.User{
		ID:           id,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}
func (r *Repository) PutUser(ctx context.Context, email string, passwordHash string) (uuid.UUID, error) {
	query := "INSERT INTO users (email, password_hash) VALUES($1, $2) RETURNING id"

	row := r.db.QueryRowContext(ctx, query, email, passwordHash)
	var id uuid.UUID
	if err := row.Scan(&id); err != nil {
		return uuid.Nil, err
	}
	return id, nil
}

func (r *Repository) DeleteUserByID(ctx context.Context, id uuid.UUID) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return repository.ErrNotFound
		}
		return err
	}
	return nil
}
