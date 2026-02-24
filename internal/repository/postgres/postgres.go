package postgres

import (
	"context"
	"database/sql"

	// postgres driver
	_ "github.com/lib/pq"
	"movieexample.com/internal/repository"
	"movieexample.com/pkg/model"
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

func (r *Repository) GetUser(ctx context.Context, id string) (*model.User, error) {
	var (
		username     string
		email        string
		passwordHash string
	)
	query := "SELECT username, email, password_hash FROM users WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&username, &email, &passwordHash); err != nil {
		if err == sql.ErrNoRows {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	return &model.User{
		ID:           id,
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}, nil
}

func (r *Repository) PutUser(ctx context.Context, id string, m *model.User) error {

	query := "INSERT INTO users (id, username, email, password_hash) VALUES($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query,
		id,
		m.Username,
		m.Email,
		m.PasswordHash)
	return err
}
