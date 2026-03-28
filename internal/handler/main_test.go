package handler

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"go.uber.org/mock/gomock"
	"musicproject.com/config"
	mock_repository "musicproject.com/gen/mocks"
	"musicproject.com/internal/services"
	"musicproject.com/internal/services/auth"
)

var (
	id  uuid.UUID
	cfg *config.Config
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
	cfg = config.LoadConfig()

	var err error
	id, err = uuid.NewV7()
	if err != nil {
		panic(err)
	}

	ctx = context.Background()

	// Run tests
	//code := m.Run()

	// Teardown

	//os.Exit(code)
}

type testContext struct {
	repo *mock_repository.MockRepository

	authService services.Auth

	id uuid.UUID
}

func (c *testContext) beforeEach(t *testing.T) {
	//t.Parallel()

	ctrl := gomock.NewController(t)
	repo := mock_repository.NewMockRepository(ctrl)

	c.authService = auth.New(cfg.Services.Auth, repo)
	c.repo = repo

	server = NewServer(cfg, repo)

}
func (c *testContext) afterEach() {

}

func testCase(test func(t *testing.T, c *testContext)) func(*testing.T) {
	return func(t *testing.T) {
		//t.Parallel()
		context := &testContext{}
		context.beforeEach(t)
		defer context.afterEach()
		//t.Parallel()
		test(t, context)
		//t.Parallel()
	}
}

func newRequest(ctx context.Context, method string, url string, body any, tokenString string) (*httptest.ResponseRecorder, error) {
	var (
		buf []byte
		err error
	)
	if body != nil {
		buf, err = json.Marshal(body)
		if err != nil {
			return nil, err
		}
	}

	w := httptest.NewRecorder()
	r := httptest.NewRequestWithContext(ctx, method, url, bytes.NewBuffer(buf))

	if tokenString != "" {
		r.AddCookie(&http.Cookie{Name: AccessCookie, Value: tokenString})
	}

	if server == nil {
		return nil, errors.New("server is nil")
	}
	server.ServeHTTP(w, r)
	return w, nil
}
