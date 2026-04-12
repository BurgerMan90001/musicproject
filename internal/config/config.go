package config

import (
	"bytes"
	"fmt"
	"os"

	"go.yaml.in/yaml/v4"
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

// Reads file from local directory
func LoadConfig(name string) (*Config, error) {
	f, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.NewDecoder(bytes.NewReader(f)).Decode(&cfg); err != nil {
		return nil, fmt.Errorf("config load: %w", err)
	}
	return &cfg, nil
}

// func (cfg *Config) Write() {

// }
