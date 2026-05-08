package integration

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/google/uuid"
	"songsled.com/internal/jsonutil"
	"songsled.com/pkg/model"
)

func (s *testSuite) TestHandleUserID() {
	endpoint := "/v1/users"

	validId := "019d34f7-0124-7cd5-8e49-cbfce4c76de4"
	_, err := uuid.Parse(validId)
	s.Require().NoError(err)
	type test struct {
		name     string
		wantCode int
		method   string
		userId   string
	}

	tests := []test{
		{
			name:     "user not found",
			method:   http.MethodGet,
			wantCode: http.StatusNotFound,
			userId:   "aaaaaaaaaaaaaaaaa",
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			req := &request{
				method: tt.method,
			}
			path, err := url.JoinPath(endpoint, tt.userId)
			s.Require().NoError(err)

			w := s.newRequest(path, req)

			r, err := jsonutil.ReadJson[model.Error](w.Result().Body)
			s.Require().NoError(err)

			// s.Nil(resBody["password"])
			data, err := io.ReadAll(w.Result().Body)
			s.Require().NoError(err)

			s.NotContains(data, "password")

			s.Equal(r.Code, tt.wantCode)
			s.Equal(tt.wantCode, w.Code, tt.name, path)
		})
	}

	tests2 := []test{
		{
			name:     "get success",
			wantCode: http.StatusOK,
			method:   http.MethodGet,
			userId:   validId,
		},
		{
			name:     "get success no dashes",
			wantCode: http.StatusOK,
			method:   http.MethodGet,
			userId:   strings.Trim(validId, "-"),
		},
	}

	for _, tt := range tests2 {
		s.Run(tt.name, func() {

			// req := &request{
			// 	method: tt.method,
			// }
			// path, err := url.JoinPath(endpoint, tt.userId)
			// s.Require().NoError(err)

			// w := s.newRequest(path, req)

			// r, err := jsonutil.ReadJson[model.User](w.Result().Body)
			// s.Require().NoError(err)

			// data, err := io.ReadAll(w.Result().Body)
			// s.Require().NoError(err)

			// s.Empty(r.PasswordHash, data)

			// s.Equal(tt.userId, r.ID.String(), data)
			// s.NotEmpty(r.Email)
		})
	}

}
