package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/services/auth"
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
	ctx        context.Context
	cfg        *config.Config
	handler    http.Handler
	sm         secrets.Manager
	jwtService *auth.JWTService
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	t := s.T()

	s.ctx = t.Context()

	cfg, err := config.LoadConfig()
	s.Require().NoError(err)

	s.cfg = cfg

	s.sm, err = secrets.NewEnv()
	s.Require().NoError(err)

	s.jwtService, err = auth.NewJWTService(s.ctx, s.sm)
	s.Require().NoError(err)

	db := postgres.NewTestDB(t, s.ctx, cfg.Repository.Postgres, s.sm)

	s.handler, err = handler.NewMux(s.ctx, cfg, db, s.sm)
	s.Require().NoError(err)
}
func (s *testSuite) TeardownSuite() {

}

type request struct {
	method string
	body   map[string]any
	// Access and refresh tokens are set in cookies
	// accessKey and refreshKey
	accessToken  string
	refreshToken string
}

func (s *testSuite) newRequest(ctx context.Context, url string, req *request) *httptest.ResponseRecorder {
	s.T().Helper()

	s.Require().NotNil(req)

	var buf io.Reader
	if len(req.body) > 0 {
		mar, err := json.Marshal(req.body)
		s.Require().NoError(err)
		buf = bytes.NewBuffer(mar)
	}
	r := httptest.NewRequestWithContext(ctx, req.method, url, buf)

	w := httptest.NewRecorder()

	if req.accessToken != "" {
		r.AddCookie(&http.Cookie{Name: auth.TokenAccess, Value: req.accessToken})
	}
	if req.refreshToken != "" {
		r.AddCookie(&http.Cookie{Name: auth.TokenRefresh, Value: req.refreshToken})
	}

	s.handler.ServeHTTP(w, r)

	return w
}
