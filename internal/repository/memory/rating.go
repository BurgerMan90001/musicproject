package memory

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

var _ repository.Rating = (*Rating)(nil)

type Rating struct {
	mu   sync.RWMutex
	data map[uuid.UUID][]*model.Rating
}

func NewRating() *Rating {
	return &Rating{data: make(map[uuid.UUID][]*model.Rating, 10)}
}

func (r *Rating) GetRatings(_ context.Context, songId uuid.UUID) ([]*model.Rating, error) {
	ratings := r.data[songId]
	if len(ratings) == 0 {
		return nil, repository.ErrNotFound
	}
	return ratings, nil
}
func (r *Rating) Put(_ context.Context, rating *model.Rating) (uuid.UUID, error) {
	r.data[rating.SongID] = append(r.data[rating.SongID], rating)
	return rating.SongID, nil
}
func (r *Rating) DeleteByID(_ context.Context, songId uuid.UUID) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	delete(r.data, songId)
	return nil
}
func (r *Rating) Update(_ context.Context, rating *model.Rating) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	ratings := r.data[rating.SongID]
	if len(ratings) == 0 {
		return repository.ErrNotFound
	}
	var err error = repository.ErrNotFound
	for i, rt := range ratings {
		//log.Println(rt.UserID, rating.UserID)
		//log.Println(rt.UserID == rating.UserID)
		if rt.UserID == rating.UserID {
			err = nil

			// Update rating
			r.data[rating.SongID][i].Value = rating.Value
			break
		}
	}
	//log.Println(len(r.data[rating.SongID]))
	return err
}
