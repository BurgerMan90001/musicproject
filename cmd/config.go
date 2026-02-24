package main

import (
	"movieexample.com/internal/repository"
	"movieexample.com/internal/repository/memory"
	"movieexample.com/internal/repository/postgres"
)

type serviceConfig struct {
	APIConfig        apiConfig        `yaml:"api"`
	RepositoryConfig repositoryConfig `yaml:"repository"`
}

type apiConfig struct {
	Port   int    `yaml:"port"`
	JWTKey string `yaml:"jwt_key"`
}

type repositoryConfig struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

func newRepository(cfg repositoryConfig) repository.Repository {
	var repo repository.Repository

	switch cfg.Type {
	case "memory":
		repo = memory.New()
	case "postgres":
		repo = postgres.New(cfg.URL)
	default:
		repo = memory.New()
	}
	return repo
}
