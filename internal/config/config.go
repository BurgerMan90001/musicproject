package config

import (
	"fmt"

	"musicproject.com/internal/fileutil"
)

// //go:embed config.dev.yml
// var configFS embed.FS

const (
	DevConfig  = "config.dev.yml"
	ProdConfig = "config.prod.yml"
)

type Env int

const (
	Dev Env = iota
	Prod
)

func LoadEnvironment(env Env) {
	switch env {
	case Dev:

	case Prod:
	default:

	}
}

func LoadConfig(filename string) (*Config, error) {
	cfg, err := fileutil.ReadYaml[Config](filename)
	if err != nil {
		return nil, fmt.Errorf("Load config: %w", err)
	}
	
	return cfg, nil
}
