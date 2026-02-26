package postgres

import (
	"context"
	"database/sql"

	// postgres driver
	_ "github.com/lib/pq"
	"okapi.com/internal/repository"
	"okapi.com/pkg/model"
)

type Repository struct {
	db *sql.DB
}

func New(url string) *Repository {
	db, err := sql.Open("postgres", url)
	if err != nil {
		panic(err)
	}
	return &Repository{db}
}

func (r *Repository) GetUserByID(ctx context.Context, id string) (*model.User, error) {
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
		id           string
		passwordHash string
	)
	query := "SELECT id, password_hash FROM users WHERE email=$1"
	row := r.db.QueryRowContext(ctx, query, email)
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
func (r *Repository) PutUser(ctx context.Context, id string, m *model.User) error {
	query := "INSERT INTO users (id, email, password_hash) VALUES($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query,
		id,
		m.Email,
		m.PasswordHash)
	return err
}

func (r *Repository) DeleteUserByID(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE id=$1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}
