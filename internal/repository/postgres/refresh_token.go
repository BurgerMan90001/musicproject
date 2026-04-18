package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type RefreshToken struct {
	db *sql.DB
}

func NewRefreshToken(db *sql.DB) *RefreshToken {
	return &RefreshToken{db}
}

func (r *RefreshToken) RevokeToken(ctx context.Context, tokenId uuid.UUID) error {
	query := "INSERT INTO revoked_tokens (token_id) VALUES($1)"
	_, err := r.db.ExecContext(ctx, query, tokenId)
	if err != nil {
		return err
	}
	return nil
}

func (r *RefreshToken) Revoked(ctx context.Context, tokenId uuid.UUID) error {
	query := "SELECT EXISTS(SELECT 1 FROM revoked_tokens WHERE token_id=$1)"
	var (
		exists    bool
		existsErr error
	)
	row := r.db.QueryRowContext(ctx, query, tokenId)
	err := row.Scan(&exists)
	if exists {
		existsErr = errors.New("Token revoked")
	}

	return errors.Join(existsErr, err)
}
