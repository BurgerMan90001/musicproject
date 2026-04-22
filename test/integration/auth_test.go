package integration

import (
	"net/http"

	"musicproject.com/internal/jsonutil"
	"musicproject.com/pkg/model"
)

/* Oauth tests */
func (s *testSuite) TestHandleOathGoogleLogin() {
	//t := s.T()
	url := "/v1/auth/google/login"
	tests := []handlerTest{
		{
			Name: "login google oauth",
			Req: &request{
				method: http.MethodGet,
			},

			WantCode: http.StatusOK,
		},
	}
	s.T().Skip()
	for _, tt := range tests {
		s.Run(tt.Name, func() {
			w := s.newRequest(url, tt.Req)

			userInfo, err := jsonutil.ReadJson[model.OauthUserInfo](w.Result().Body)
			s.Require().NoError(err)

			s.Equal(tt.WantCode, w.Code, tt.Name)
			s.Empty(userInfo.Email)
		})
	}

}
