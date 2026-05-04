package integration

// func (s *testSuite) TestRefresh() {
// 	url := "/v1/auth/refresh"

// 	access, err := s.jwtAccess.GenerateToken(uuid.Nil)
// 	s.Require().NoError(err)

// 	tests := []handlerTest{
// 		{
// 			Name:     "invalid token type",
// 			WantCode: http.StatusUnauthorized,
// 			Req: &request{
// 				refreshToken: access,
// 			},
// 		},
// 		{
// 			Name:        "empty refresh token string",
// 			WantCode:    http.StatusUnauthorized,
// 			Req:         &request{},
// 			WantMessage: "No token present",
// 		},
// 	}
// 	for _, tt := range tests {
// 		s.Run(tt.Name, func() {
// 			w := s.newRequest(url, tt.Req)

// 			_, err := jsonutil.ReadJson[model.Error](w.Result().Body)
// 			s.Require().NoError(err)

// 			s.Equal(tt.WantCode, w.Code, tt.Name)
// 			//s.Equal(tt.WantMessage, resBody.Message)
// 		})
// 	}

// 	validRefresh, err := s.jwtRefresh.GenerateToken(uuid.Nil)
// 	s.Require().NoError(err)
// 	t1 := handlerTest{
// 		Name: "successful refresh with cookie token",
// 		Req: &request{
// 			refreshToken: validRefresh,
// 		},
// 		WantCode: http.StatusOK,
// 	}
// 	s.Run(t1.Name, func() {
// 		w := s.newRequest(url, t1.Req)

// 		s.Equal(t1.WantCode, w.Code, t1.Name)

// 		// Second request, make request with old refresh token
// 		w2 := s.newRequest(url, t1.Req)
// 		// Old refresh token is revoked
// 		s.Equal(http.StatusUnauthorized, w2.Code, t1.Name)

// 		b, err := jsonutil.ReadJson[model.Error](w2.Result().Body)
// 		s.Require().NoError(err)

// 		s.Equal(http.StatusUnauthorized, b.Code, t1.Name)
// 		s.Equal(auth.ErrInvalidToken().Error(), b.Message)
// 	})
// }
