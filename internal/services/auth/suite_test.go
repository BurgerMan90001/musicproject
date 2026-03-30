package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"musicproject.com/config"
	"musicproject.com/internal/repository"
	"musicproject.com/internal/repository/postgres"
	"musicproject.com/internal/services"
	"musicproject.com/pkg/testutil"
)

type testSuite struct {
	suite.Suite
	ctx context.Context
	pg          *testutil.PostgresContainer
	repo        repository.Repository
	authService services.Auth
}

func (s *testSuite) SetupTest() {

	s.ctx = context.Background()

	cfg := config.LoadConfig()

	pg, err := testutil.NewPostgresContainer(s.ctx, cfg.Repository.Postgres)
	if err != nil {
		panic(err)
	}
	s.pg = pg
	s.repo = postgres.New(s.pg.URL)

	s.authService = New(cfg.Services.Auth, s.repo)
}
func (s *testSuite) TearDownSuite() {
	err := s.pg.Terminate(s.ctx)
	s.Require().NoError(err)
}
func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
