package jsonutil

import (
	"encoding/json"
	"errors"
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

	tests := []struct {
		name    string
		data    any
		code    int
		details []string

		wantCode int
	}{
		{
			name: "invalid status code",
			data: model.User{
				ID:    uuid.New(),
				Email: "paulcasigay@gmail.com",
			},
			code:     1000,
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "nil data",
			code:     http.StatusOK,
			wantCode: http.StatusInternalServerError,
		},
		{
			name:     "empty error string",
			data:     errors.New(""),
			code:     http.StatusBadGateway,
			wantCode: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()

			WriteJSON(w, tt.data, tt.code, tt.details...)

			var errorRes model.ErrorResponse
			err := json.NewDecoder(w.Body).Decode(&errorRes)
			require.NoError(t, err)

			// Assert all codes match
			assert.Equal(t, tt.wantCode, w.Code, tt.name)
			assert.Equal(t, tt.wantCode, errorRes.Code)
			assert.NotEmpty(t, errorRes.Message)
		})
	}

	t.Run("much details", func(t *testing.T) {
		w := httptest.NewRecorder()

		WriteJSON(w, model.User{
			ID:    uuid.New(),
			Email: "paulcasigay@gmail.com",
		}, http.StatusOK, "asd", "asd", "asd", "asd",
			"asd", "asd", "asd", "asd", "asd", "asd")

		assert.Equal(t, http.StatusOK, w.Code)
	})
}
func TestWriteJSONStatusOK(t *testing.T) {
	t.Parallel()

	w := httptest.NewRecorder()

	user := model.User{
		ID:    uuid.New(),
		Email: "paulcasigay@gmail.com",
	}
	WriteJSON(w, user, http.StatusOK)

	assert.Equal(t, http.StatusOK, w.Code)

	var resUser model.User
	err := json.NewDecoder(w.Body).Decode(&resUser)
	require.NoError(t, err)

	assert.Equal(t, user, resUser)
}
