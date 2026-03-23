package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"musicproject.com/config"
	mock_repository "musicproject.com/gen/mocks"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/service/auth"
	"musicproject.com/pkg/util/fileutil"
)

var (
	//server *httptest.Server

	id  uuid.UUID
	cfg config.Config
	ctx context.Context

	server *Server
)

type HandlerTest struct {
	Name   string
	Method string
	URL    string
	Body   map[string]any

	WantSuccess bool
	WantCode    int

	WantData    any
	WantMessage string

	RepoItem any
	RepoErr  error
}

func TestMain(m *testing.M) {

	// Run setup
	var err error
	id, err = uuid.NewV7()
	if err != nil {
		panic(err)
	}

	ctx = context.Background()

	// Run tests
	code := m.Run()

	// Teardown

	os.Exit(code)
}

type testContext struct {

	//repo repository.Repository
	repo postgres.Repository
	userRepo *mock_repository.MockUserRepository

	authService *auth.Service

	id uuid.UUID
}

func (c *testContext) beforeEach(t *testing.T) {
	cfg, err := fileutil.ReadYAML[config.Config]("../../config/base.dev.yml")
	if err != nil {
		panic(err)
	}
	ctrl := gomock.NewController(t)

	//c.repo = postgres.Repository{}
	c.userRepo = mock_repository.NewMockUserRepository(ctrl)

	c.authService = auth.New(cfg.Services.Auth, c.userRepo)

}
func (c *testContext) afterEach() {

}

func testCase(test func(t *testing.T, c *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		context := &testContext{}
		context.beforeEach(t)
		defer context.afterEach()
		test(t, context)
	}

}

func newRequest(ctx context.Context, method string, url string, body any, requireAuth bool) (*httptest.ResponseRecorder, error) {
	buf, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	r := httptest.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(buf))
	//r.SetPathValue()
	if requireAuth {
		r.Header.Add("Authorization", "Bearer "+"")
	}
	w := httptest.NewRecorder()
	server.ServeHTTP(w, r)
	return w, nil
}
