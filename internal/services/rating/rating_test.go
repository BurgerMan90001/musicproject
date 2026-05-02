package rating

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"songsled.com/internal/repository/memory"
	"songsled.com/pkg/model"
)

func TestPutRating(t *testing.T) {
	t.Parallel()

	t.Skip("Skipping rating service")

	ctx := t.Context()

	// repo, err := postgres.New(ctx)
	// require.NoError(t, err)

	ratingRepo := memory.NewRating()
	ratingService := New(0, 5, ratingRepo)

	songId := uuid.New()
	tests := []struct {
		name string
		//songId uuid.UUID
		userId uuid.UUID
		value  float64
	}{
		{
			name:  "negative number",
			value: -4,
		},
		{
			name:  "big rating",
			value: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := ratingService.Put(ctx, &model.Rating{
				SongID: songId,
				UserID: tt.userId,
				Value:  tt.value,
			})

			assert.Error(t, err)
		})
	}
	t.Run("put rating success", func(t *testing.T) {
		err := ratingService.Put(ctx, &model.Rating{
			SongID: songId,
			UserID: uuid.New(),
			Value:  2,
		})
		require.NoError(t, err)

		rating, err := ratingService.GetAggregatedRating(ctx, songId)
		require.NoError(t, err)

		assert.Equal(t, float64(2), rating)
	})
	t.Run("update rating", func(t *testing.T) {
		// Same userId
		rating := &model.Rating{
			SongID: songId,
			UserID: uuid.New(),
			Value:  2,
		}
		userId := uuid.New()
		err := ratingService.Put(ctx, &model.Rating{
			SongID: songId,
			UserID: userId,
			Value:  2,
		})
		require.NoError(t, err)

		// Update rating
		rating.Value = 4
		err = ratingService.Put(ctx, &model.Rating{
			SongID: songId,
			UserID: userId,
			Value:  4,
		})
		require.NoError(t, err)

		value, err := ratingService.GetAggregatedRating(ctx, songId)
		require.NoError(t, err)

		assert.Equal(t, float64(4), value)

	})

}

func TestAggregatedRating(t *testing.T) {
	t.Parallel()

	t.Skip("Skipping rating service")

	ctx := t.Context()

	ratingRepo := memory.NewRating()
	ratingService := New(0, 5, ratingRepo)

	songId := uuid.New()

	t.Run("aggregate success", func(t *testing.T) {
		ratings := []float64{4, 3, 4, 1, 3}
		var sum float64 = 0
		for _, rating := range ratings {
			err := ratingService.Put(ctx, &model.Rating{
				SongID: songId,
				Value:  rating,
				UserID: uuid.New(),
			})
			sum += rating
			require.NoError(t, err)
		}
		rating, err := ratingService.GetAggregatedRating(ctx, songId)
		require.NoError(t, err)

		assert.Equal(t, float64(3), rating)
	})
}
