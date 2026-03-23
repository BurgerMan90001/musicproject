package handler

import (
	"io"
	"net/http"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
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
			Name:     "get success",
			WantCode: http.StatusOK,
			Method:   http.MethodGet,

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

			WantMessage: repository.ErrUserNotFound.Error(),
		},

		{
			Name:   "method not allowed",
			Method: http.MethodConnect,

			WantMessage: ErrInvalidMethod.Error(),

			WantCode: http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, testCase(func(t *testing.T, c *testContext) {

			c.userRepo.EXPECT().GetUserByID(ctx, id).Return(tt.RepoItem, tt.RepoErr).AnyTimes()

			w, err := newRequest(ctx, tt.Method, tt.URL, tt.Body, false)
			if err != nil {
				t.Error(err)
			}

			t1, err := model.MarshalJSON(tt.WantData, tt.WantCode, tt.WantMessage)
			if err != nil {
				t.Error(err)
			}

			t2, err := io.ReadAll(w.Body)
			if err != nil {
				t.Error(err)
			}

			assert.JSONEq(t, string(t1), string(t2), tt.Name)
			assert.Equal(t, tt.WantCode, w.Code, tt.Name)
		}))
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
