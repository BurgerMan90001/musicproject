package integration

import (
	"bytes"
	"embed"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/suite"
	"songsled.com/internal/config"
	"songsled.com/internal/config/secrets"
	"songsled.com/internal/handler"
	"songsled.com/internal/repository/postgres"
	"songsled.com/internal/services/file"
	"songsled.com/pkg/model"
)

//go:embed testdata
var testdataFS embed.FS

type handlerTest struct {
	Name string
	// Required field
	Req *request

	WantCode    int
	WantMessage string
}
type request struct {
	// Will default to GET
	method string
	// Optional
	body map[string]any
	// Access and refresh tokens are set in cookies as accessKey and refreshKey
	accessToken  string
	refreshToken string
}

type testSuite struct {
	suite.Suite
	cfg     *config.Config
	handler http.Handler
	repo    *postgres.Repo
	// jwtAccess  *auth.JWTService
	// jwtRefresh *auth.JWTService
}

func TestIntegrationSuite(t *testing.T) {
	// if testing.Short() {
	// 	t.Skip("SHORT: Integration testing skipped")
	// }
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	t := s.T()
	ctx := t.Context()

	// Env variables for testing
	err := secrets.ReadEnvFile(".env")
	s.Require().NoError(err)

	// Config
	s.cfg, err = config.LoadConfig(filepath.Join("..", "..", "config.dev.yml"))
	s.Require().NoError(err)

	err = secrets.ReadEnvFile(filepath.Join("..", "..", ".env.dev"))
	s.Require().NoError(err)

	store, err := file.New(ctx, &file.AWSS3{}, s.cfg.File)
	s.Require().NoError(err)

	// Test postgres database container
	s.repo = postgres.NewTest(t, ctx, s.cfg.Repository.Postgres.Image)

	
	if os.Getenv("LOAD_TESTDATA") == "true" {
		err = s.repo.ExecFile(ctx, filepath.Join("..", "..", "database", "test.sql"))
		s.Require().NoError(err)
	}

	s.handler, err = handler.New(ctx, s.cfg, store, s.repo, nil)
	s.Require().NoError(err)
}

func (s *testSuite) newRequest(url string, req *request) *httptest.ResponseRecorder {
	s.T().Helper()
	if req == nil {
		req = &request{
			method: "GET",
		}
	}
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
