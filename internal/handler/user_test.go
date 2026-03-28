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

func TestHandleUser(t *testing.T) {
	id, err := uuid.NewV7()
	if err != nil {
		t.Error(err)
	}
	url := "/v1/user/"

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

			WantMessage: repository.ErrNotFound.Error(),
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

			c.repo.EXPECT().GetUserByID(ctx, id).Return(tt.RepoItem, tt.RepoErr).AnyTimes()

			w, err := newRequest(ctx, tt.Method, url+id.String(), tt.Body, "")
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
