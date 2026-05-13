package postgres

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"songsled.com/internal/repository/postgres/gensqlc"
	"songsled.com/pkg/model"
)

type Genre struct {
	q *gensqlc.Queries
}

func NewGenreRepo(q *gensqlc.Queries) *Genre {
	return &Genre{q}
}
func (r *Genre) NewGenre(ctx context.Context, name string) (uuid.UUID, error) {
	return r.q.NewGenre(ctx, name)
}
func (r *Genre) GetGenres(ctx context.Context, n int32) ([]*model.Genre, error) {
	l, err := r.q.GetGenres(ctx, n)
	if err != nil {
		return nil, fmt.Errorf("Get genres: %w", err)
	}
	var genres []*model.Genre
	for _, g := range l {
		genres = append(genres, &model.Genre{
			GenreId: g.GenreID,
			Name:    g.GenreName,
		})
	}
	return genres, nil
}
func (r *Genre) GetGenreById(ctx context.Context, genreId uuid.UUID) (string, error) {
	g, err := r.q.GetGenreById(ctx, genreId)
	if err != nil {
		return "", fmt.Errorf("Get genre by id: %w", err)
	}
	return g.GenreName, nil
}
func (r *Genre) GetGenreByName(ctx context.Context, genreName string) (uuid.UUID, error) {
	g, err := r.q.GetGenreByName(ctx, genreName)
	if err != nil {
		return uuid.Nil, fmt.Errorf("Get genre by name: %w", err)
	}
	return g.GenreID, nil
}
