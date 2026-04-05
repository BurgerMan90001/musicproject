package jsonutil

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"musicproject.com/pkg/model"
)

func TestWriteJSON(t *testing.T) {
	t.Parallel()

	errorTests := []struct {
		name string
		data any
		code int
		args []string

		wantCode int
	}{
		{
			name: "no status code / invalid code",
			data: model.User{
				ID:    uuid.New(),
				Email: "paulcasigay@gmail.com",
			},
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range errorTests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			WriteJSON(w, tt.data, tt.code, tt.args...)

			var errorRes model.Error
			err := json.NewDecoder(w.Body).Decode(&errorRes)
			require.NoError(t, err)

			// Assert all codes match
			assert.Equal(t, tt.wantCode, w.Code, tt.name)
			assert.Equal(t, tt.wantCode, errorRes.Code)

			assert.NotEmpty(t, errorRes.Message)
		})
	}
	t.Run("status ok", func(t *testing.T) {
		w := httptest.NewRecorder()
		data := model.User{
			ID:    uuid.New(),
			Email: "paulcasigay@gmail.com",
		}
		WriteJSON(w, data, http.StatusOK)
		assert.Equal(t, http.StatusOK, w.Code)

		var user model.User
		err := json.NewDecoder(w.Body).Decode(&user)
		require.NoError(t, err)

		assert.Equal(t, data, user)
	})
}
