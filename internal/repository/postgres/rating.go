package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

// Song rating methods
func (r *Repository) GetRatings(ctx context.Context, songId uuid.UUID) ([]model.Rating, error) {
	query := "SELECT user_id, value FROM ratings WHERE song_id=$1"
	rows, err := r.db.QueryContext(ctx, query, songId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()

	var (
		ratings []model.Rating
	)
	for rows.Next() {
		var (
			userId uuid.UUID
			value  float64
		)
		if err := rows.Scan(&userId, &value); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, repository.ErrNotFound
			}
			return nil, err
		}
		ratings = append(ratings, model.Rating{
			SongID: songId,
			UserID: userId,
			Value:  value,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ratings, nil
}
func (r *Repository) PutRating(ctx context.Context, songId uuid.UUID, userId uuid.UUID, value float64) error {
	query := "INSERT INTO ratings (song_id, user_id, value) VALUES($1, $2, $3)"
	_, err := r.db.ExecContext(ctx, query, songId, userId, value)
	if err != nil {
		return err
	}
	return nil
}
