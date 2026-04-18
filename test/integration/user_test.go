package integration

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/jsonutil"
	"musicproject.com/internal/repository"
)

func (s *testSuite) TestHandleUserID() {
	URL := "/v1/users/"
	validId, err := uuid.Parse("019d34f7-0124-7cd5-8e49-cbfce4c76de4")
	s.Require().NoError(err)

	tests := []struct {
		name        string
		wantCode    int
		method      string
		wantMessage string
		userId      uuid.UUID
	}{
		{
			name:     "get success",
			wantCode: http.StatusOK,
			method:   http.MethodGet,
			userId:   validId,
		},
		{
			name:     "user not found",
			method:   http.MethodGet,
			wantCode: http.StatusNotFound,

			wantMessage: repository.ErrNotFound.Error(),
		},
		{
			name:   "method not allowed",
			method: http.MethodConnect,

			wantMessage: "Method not allowed",
			wantCode:    http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			req := &request{
				method: tt.method,
			}
			w := s.newRequest(URL+tt.userId.String(), req)

			resBody, err := jsonutil.ReadJson[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			if resBody["message"] != nil {
				s.Equal(tt.wantMessage, resBody["message"])
			}
			s.Nil(resBody["password"])

			if resBody["userId"] != nil {
				s.Equal(tt.userId.String(), resBody["userId"])
			}

			s.Equal(tt.wantCode, w.Code, tt.name)
		})
	}
}
