package auth

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"musicproject.com/internal/config"
	"musicproject.com/internal/config/secrets"
	"musicproject.com/internal/repository/postgres"
)

type testSuite struct {
	suite.Suite
	ctx         context.Context
	authService *Service
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(testSuite))
}

func (s *testSuite) SetupTest() {
	t := s.T()

	s.ctx = t.Context()

	cfg, err := config.LoadConfig()
	s.Require().NoError(err)

	//userRepo := memory.NewUser()
	sm, err := secrets.NewEnv()
	s.Require().NoError(err)

	db := postgres.NewTestDB(t, s.ctx, cfg.Repository.Postgres, sm)
	userRepo := postgres.NewUser(db)

	s.authService, err = New(s.ctx, cfg.Services.Auth, userRepo, sm)
	s.Require().NoError(err)
}

func (s *testSuite) TearDownSuite() {

}
