package integration

func (s *testSuite) TestRateLimitMiddleware() {

	s.Run("", func() {

	})
}
func (s *testSuite) TestAuthMiddleware() {
	// url := "/v1/admin"

	// userId := uuid.Nil

	// expireJwt, err := auth.NewJWTService(
	// 	s.cfg.Auth.Jwt.Issuer,
	// 	s.cfg.Auth.Jwt.Audience,
	// 	model.TokenAccess,
	// 	-1*time.Hour,
	// 	"JWT_ACCESS_KEY",
	// )

	// expired, err := expireJwt.GenerateToken(userId)
	// s.Require().NoError(err)

	// invalidType, err := s.jwtRefresh.GenerateToken(userId)
	// s.Require().NoError(err)

	// userRole, err := s.jwtAccess.GenerateToken(userId, "user")

	// tests := []struct {
	// 	name        string
	// 	wantMessage string
	// 	wantStatus  int
	// 	accessToken string
	// }{
	// 	{
	// 		name:        "expired token",
	// 		wantStatus:  http.StatusUnauthorized,
	// 		wantMessage: auth.ErrInvalidToken().Error(),
	// 		accessToken: expired,
	// 	},
	// 	{
	// 		name:        "invalid token type",
	// 		wantMessage: auth.ErrInvalidToken().Error(),
	// 		wantStatus:  http.StatusUnauthorized,
	// 		accessToken: invalidType,
	// 	},
	// 	{
	// 		name:        "empty token or no header",
	// 		wantMessage: jwt.ErrTokenMalformed.Error(),
	// 		wantStatus:  http.StatusUnauthorized,
	// 		accessToken: "",
	// 	},
	// 	{
	// 		name:        "invalid token",
	// 		wantStatus:  http.StatusUnauthorized,
	// 		wantMessage: jwt.ErrTokenMalformed.Error(),
	// 		accessToken: "aksidoajfjsofuijrngukaernngaerjgknne",
	// 	},
	// 	{
	// 		name:       "invalid role",
	// 		wantStatus: http.StatusUnauthorized,

	// 		accessToken: userRole,
	// 	},
	// }
	// for _, tt := range tests {
	// 	s.Run(tt.name, func() {
	// 		req := &request{
	// 			method:      http.MethodGet,
	// 			accessToken: tt.accessToken,
	// 		}
	// 		w := s.newRequest(url, req)

	// 		resBody, err := jsonutil.ReadJson[map[string]any](w.Result().Body)
	// 		s.Require().NoError(err)

	// 		s.Equal(tt.wantStatus, w.Code, tt.name)
	// 		s.NotEmpty(resBody["message"], tt.name)
	// 	})
	// }

	// s.Run("success", func() {
	// 	valid, err := s.jwtAccess.GenerateToken(uuid.Nil, "admin")
	// 	s.Require().NoError(err)

	// 	req := &request{
	// 		method:      http.MethodGet,
	// 		accessToken: valid,
	// 	}
	// 	w := s.newRequest(url, req)

	// 	s.Equal(http.StatusOK, w.Code)
	// })

}
