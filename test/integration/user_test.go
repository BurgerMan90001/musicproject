package integration

import (
	"net/http"

	"github.com/google/uuid"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository"
	"musicproject.com/pkg/model"
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

			wantMessage: handler.ErrInvalidMethod.Error(),
			wantCode:    http.StatusMethodNotAllowed,
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			w := s.newRequest(s.ctx, tt.method, URL+tt.userId.String(), nil, "")

			resBody, err := model.ReadJSON[map[string]any](w.Result().Body)
			s.Require().NoError(err)

			// var data map[string]string
			// json.Unmarshal(resBody.Data, &data)

			//s.HTTPBodyContains(handler.HandleUserID(s.repo), tt.method, URL +, v, tt.wantMessage)
			//s.Contains(resBody["message"], tt.wantMessage)
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
