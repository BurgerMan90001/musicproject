package integration

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/services/auth"
	"musicproject.com/pkg/model"
)

type HandlerTest struct {
	Name string
	// Required field
	Req *request

	WantCode    int
	WantMessage string
}
type testSuite struct {
	suite.Suite
	//ctx        context.Context
	cfg        *config.Config
	handler    http.Handler
	sm         secrets.Manager
	jwtAccess  *auth.JWTService
	jwtRefresh *auth.JWTService
}

func TestIntegrationSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("SHORT: Integration testing skipped")
	}
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	t := s.T()
	ctx := t.Context()

	var err error
	// Config
	s.cfg, err = config.LoadConfig(filepath.Join("..", "..", "config", "config.dev.yml"))
	s.Require().NoError(err)

	err = secrets.LoadEnv(filepath.Join("..", "..", "config", ".env.dev"))
	s.Require().NoError(err)

	// JWT services
	s.jwtAccess, err = auth.NewJWTService(s.cfg.Services.Auth.Jwt, "JWT_ACCESS_KEY", model.TokenAccess, time.Minute*30)
	s.Require().NoError(err)

	s.jwtRefresh, err = auth.NewJWTService(s.cfg.Services.Auth.Jwt, "JWT_REFRESH_KEY", model.TokenRefresh, time.Hour)
	s.Require().NoError(err)

	// Test postgres database container
	db := postgres.NewTestDB(t, ctx, s.cfg.Repository.Postgres)

	s.handler, err = handler.NewMux(ctx, s.cfg, db)
	s.Require().NoError(err)
}
func (s *testSuite) TeardownSuite() {

}

type request struct {
	method string
	body   map[string]any
	// Access and refresh tokens are set in cookies as accessKey and refreshKey
	accessToken  string
	refreshToken string
}

func (s *testSuite) newRequest(url string, req *request) *httptest.ResponseRecorder {
	s.T().Helper()

	s.Require().NotNil(req)

	var buf io.Reader
	if len(req.body) > 0 {
		mar, err := json.Marshal(req.body)
		s.Require().NoError(err)
		buf = bytes.NewBuffer(mar)
	}
	r := httptest.NewRequestWithContext(s.T().Context(), req.method, url, buf)

	w := httptest.NewRecorder()

	if req.accessToken != "" {

		//http.SetCookie(w, &http.Cookie{Name: auth.TokenAccess, Value: req.accessToken})
		r.AddCookie(&http.Cookie{Name: string(model.TokenAccess), Value: req.accessToken})
	}
	if req.refreshToken != "" {
		r.AddCookie(&http.Cookie{Name: string(model.TokenRefresh), Value: req.refreshToken})
	}

	s.handler.ServeHTTP(w, r)

	return w
}
