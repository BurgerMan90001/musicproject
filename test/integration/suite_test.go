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
	"musicproject.com/internal/services/file"
	"musicproject.com/pkg/model"
)

type HandlerTest struct {
	Name string
	// Required field
	Req *request

	WantCode    int
	WantMessage string
}
type request struct {
	// Will default to GET
	method string
	body   map[string]any
	// Access and refresh tokens are set in cookies as accessKey and refreshKey
	accessToken  string
	refreshToken string
}

type testSuite struct {
	suite.Suite
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

	configFolder := filepath.Join("..", "..", "config")
	var err error
	// Config
	s.cfg, err = config.LoadConfig(filepath.Join(configFolder, "config.dev.yml"))
	s.Require().NoError(err)

	err = secrets.ReadEnvFile(filepath.Join(configFolder, ".env.dev"))
	s.Require().NoError(err)

	// JWT services
	s.jwtAccess, err = auth.NewJWTService(s.cfg.Auth.Jwt, "JWT_ACCESS_KEY", model.TokenAccess, time.Minute*30)
	s.Require().NoError(err)

	s.jwtRefresh, err = auth.NewJWTService(s.cfg.Auth.Jwt, "JWT_REFRESH_KEY", model.TokenRefresh, time.Hour)
	s.Require().NoError(err)

	// Filesystem store
	store, err := file.New(ctx, &file.GoogleCloud{}, s.cfg.Upload.Region)
	s.Require().NoError(err)

	// Test postgres database container
	db := postgres.NewTestDB(t, ctx, s.cfg.Repository.Postgres.Image)

	s.handler, err = handler.NewMux(ctx, s.cfg, store, db)
	s.Require().NoError(err)
}
func (s *testSuite) TeardownSuite() {

}

func (s *testSuite) newRequest(url string, req *request) *httptest.ResponseRecorder {
	//s.T().Helper()

	s.Require().NotNil(req)

	var buf io.Reader
	if len(req.body) > 0 {
		mar, err := json.Marshal(req.body)
		s.Require().NoError(err)
		buf = bytes.NewBuffer(mar)
	}
	r := httptest.NewRequestWithContext(s.T().Context(), req.method, url, buf)

	r.AddCookie(&http.Cookie{Name: string(model.TokenAccess), Value: req.accessToken})
	r.AddCookie(&http.Cookie{Name: string(model.TokenRefresh), Value: req.refreshToken})

	w := httptest.NewRecorder()

	s.handler.ServeHTTP(w, r)

	return w
}
