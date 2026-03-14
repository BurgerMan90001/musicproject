package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
)

/*
Gets song's amount of streams.
*/
func (r *Repository) GetStreamsByID(ctx context.Context, songId uuid.UUID) (int, error) {
	var value int
	query := "SELECT value FROM streams WHERE song_id=$1"
	row := r.db.QueryRowContext(ctx, query, songId)
	if err := row.Scan(&value); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, repository.ErrNotFound
		}
		return 0, err
	}
	return value, nil
}
