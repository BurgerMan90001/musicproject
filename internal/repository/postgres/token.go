package postgres

import (
	"context"
	"database/sql"
)

type Token struct {
	db *sql.DB
}

func (r *Token) GetToken(ctx context.Context, tokenString string) error {
	return nil
}

func (r *Token) PutToken(ctx context.Context, tokenString string) error {
	return nil
}
func (r *Token) DeleteToken(ctx context.Context, tokenString string) error {
	return nil
}
