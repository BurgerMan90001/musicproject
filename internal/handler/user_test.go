package handler

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	mock_repository "musicproject.com/gen/mocks"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
)

func TestHandleGet(t *testing.T) {
	id, err := uuid.NewV7()
	if err != nil {
		t.Error(err)
	}

	tests := []HandlerTest{
		{
			Name:       "get success",
			WantCode:   http.StatusOK,
			Method:     http.MethodGet,
			WantStatus: StatusSucess,

			WantData: model.User{
				ID:           id,
				Email:        "paulcasigay@gmail.com",
				PasswordHash: "asijdoiojghjflgnhkgfh",
			},

			RepoItem: &model.User{
				ID:           id,
				Email:        "paulcasigay@gmail.com",
				PasswordHash: "asijdoiojghjflgnhkgfh",
			},
		},
		{
			Name:     "user not found",
			Method:   http.MethodGet,
			WantCode: http.StatusNotFound,
			RepoErr:  repository.ErrNotFound,

			WantStatus:  StatusError,
			WantMessage: ErrUserNotFound.Error(),
		},

		{
			Name:   "method not allowed",
			Method: http.MethodConnect,

			WantStatus:  StatusError,
			WantMessage: ErrInvalidMethod.Error(),

			WantCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			mockController := gomock.NewController(t)
			defer mockController.Finish()

			repoMock := mock_repository.NewMockRepository(mockController)

			ctx := context.Background()

			repoMock.EXPECT().GetUserByID(ctx, id).Return(tt.RepoItem, tt.RepoErr).AnyTimes()

			w := httptest.NewRecorder()

			body, err := NewRequestBody(tt.Body)
			if err != nil {
				t.Error(err)
			}

			r := httptest.NewRequestWithContext(ctx, tt.Method, "/user", body)
			r.Header.Set("Content-Type", "application/json")

			r.SetPathValue("id", id.String())

			HandleUserID(repoMock).ServeHTTP(w, r)

			t1, err := MarshalJSON(tt.WantStatus, tt.WantData, tt.WantCode, tt.WantMessage)
			if err != nil {
				t.Error(err)
			}

			t2, err := io.ReadAll(w.Body)
			if err != nil {
				t.Error(err)
			}

			assert.JSONEq(t, string(t1), string(t2), tt.Name)
			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
		})
	}
}
func TestPut(t *testing.T) {
	tests := []HandlerTest{
		{
			Name:     "success",
			WantCode: http.StatusOK,
			//Body:     model.User{},
			RepoItem: model.User{},
		},
	}
	t.Skip()
	for _, tt := range tests {

		t.Run(tt.Name, func(t *testing.T) {

		})
	}
}
