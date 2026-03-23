package postgres

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

type Song struct {
	db *sql.DB
}

func (r *Song) GetSongByID(ctx context.Context, songId uuid.UUID) (*model.Song, error) {
	var (
		name         string
		genre        string
		streams      int
		duration     int
		image        string
		creationDate string
		src          string
	)
	query := "SELECT name, genre, streams, duration, image, creation_date FROM songs WHERE song_id=$1"
	row := r.db.QueryRowContext(ctx, query, songId)
	if err := row.Scan(&name, &genre, &streams, &duration, &image, &creationDate, &src); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, err
	}
	song := &model.Song{
		ID:       songId,
		Name:     name,
		Genre:    genre,
		Streams:  streams,
		Duration: duration,
		Image:    image,
		Source:   src,
	}
	return song, nil
}
func (r *Song) GetSongsByGenre(ctx context.Context, genre string) ([]model.Song, error) {
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
		var ()
		//rows.Scan(&)
		songs = append(songs, model.Song{})
	}
	return nil, nil
}
func (r *Song) PutSong(ctx context.Context, id uuid.UUID, u *model.Song) (uuid.UUID, error) {
	return uuid.Nil, nil
}
