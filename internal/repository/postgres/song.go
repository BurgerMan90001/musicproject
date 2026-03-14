package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

func (r *Repository) GetSongByID(ctx context.Context, id uuid.UUID) (*model.Song, error) {
	var (
		name         string
		genre        string
		streams      int
		duration     int
		image        string
		creationDate string
		src          string
	)
	query := "SELECT name, genre, streams, duration, image, creation_date FROM songs WHERE id=$1"
	row := r.db.QueryRowContext(ctx, query, id)
	if err := row.Scan(&name, &genre, &streams, &duration, &image, &creationDate, &src); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	song := &model.Song{
		ID:       id,
		Name:     name,
		Genre:    genre,
		Streams:  streams,
		Duration: duration,
		Image:    image,
		Source:   src,
	}
	return song, nil
}
func (r *Repository) GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error) {
	query := "SELECT id, WHERE genre=$1"
	rows, err := r.db.QueryContext(ctx, query, genre)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, repository.ErrNotFound
		}
		return nil, err
	}
	defer rows.Close()
	var songs []model.Song
	for rows.Next() {
		//rows.Scan(&)
		songs = append(songs)
	}
	return nil, nil
}
func (r *Repository) PutSong(ctx context.Context, id uuid.UUID, u *model.Song) error {
	return nil
}
