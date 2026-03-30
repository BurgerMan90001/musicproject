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
	"musicproject.com/config"
	"musicproject.com/internal/handler"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/pkg/testutil"
)

type HandlerTest struct {
	Name   string
	Method string
	//URL    string
	Body map[string]any

	WantCode int

	WantData    any
	WantMessage string
}
type testSuite struct {
	suite.Suite
	ctx    context.Context
	pg     *testutil.PostgresContainer
	cfg    *config.Config
	repo   repository.Repository
	server *handler.Server
}

func TestIntegrationSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupSuite() {
	s.ctx = context.Background()
	s.cfg = config.LoadConfig()

	pg, err := testutil.NewPostgresContainer(s.ctx, s.cfg.Repository.Postgres)
	if err != nil {
		panic(err)
	}
	s.pg = pg
	s.repo = postgres.New(s.pg.URL)

	s.server = handler.NewServer(s.cfg, s.repo)

}
func (s *testSuite) TeardownSuite() {
	err := s.pg.Terminate(s.ctx)
	s.Require().NoError(err)
}

func (s *testSuite) newRequest(ctx context.Context, method string, url string, body map[string]any, tokenString string) *httptest.ResponseRecorder {

	var buf io.Reader

	if len(body) > 0 {
		mar, err := json.Marshal(body)
		s.Require().NoError(err)
		buf = bytes.NewBuffer(mar)
	}
	r := httptest.NewRequestWithContext(ctx, method, url, buf)

	w := httptest.NewRecorder()

	if tokenString != "" {
		r.AddCookie(&http.Cookie{Name: handler.AccessCookie, Value: tokenString})
	}
	s.server.ServeHTTP(w, r)

	return w
}
