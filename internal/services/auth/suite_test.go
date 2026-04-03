package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"musicproject.com/config"
	"musicproject.com/internal/repository/memory"
)

type testSuite struct {
	suite.Suite
	ctx context.Context
	//pg          *testutil.PostgresContainer
	authService *Service
}

func (s *testSuite) SetupTest() {
	//t := s.T()

	s.ctx = context.Background()

	cfg, err := config.LoadConfig()
	s.Require().NoError(err)

	//db, _ := postgres.NewTestDB(t, s.ctx, cfg.Repository.Postgres)

	userRepo := memory.NewUser()
	s.authService = New(cfg.Services.Auth, userRepo)
}
func (s *testSuite) TearDownSuite() {

}
func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}
