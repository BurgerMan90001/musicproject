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
	Name   string
	Method string
	Body   map[string]any

	WantCode int

	WantData    any
	WantMessage string
}
type testSuite struct {
	suite.Suite
	ctx     context.Context
	cfg     *config.Config
	handler http.Handler
	sm      secrets.Manager
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

	db := postgres.NewTestDB(t, s.ctx, cfg.Repository.Postgres, s.sm)

	s.handler, err = handler.NewMux(s.ctx, cfg, db, s.sm)
	s.Require().NoError(err)
}
func (s *testSuite) TeardownSuite() {

}

func (s *testSuite) newRequest(ctx context.Context, method string, url string, body map[string]any, tokenString string) *httptest.ResponseRecorder {
	s.T().Helper()

	var buf io.Reader
	if len(body) > 0 {
		mar, err := json.Marshal(body)
		s.Require().NoError(err)
		buf = bytes.NewBuffer(mar)
	}
	r := httptest.NewRequestWithContext(ctx, method, url, buf)

	w := httptest.NewRecorder()

	if tokenString != "" {
		r.AddCookie(&http.Cookie{Name: auth.TokenRefresh, Value: tokenString})
	}
	s.handler.ServeHTTP(w, r)

	return w
}
