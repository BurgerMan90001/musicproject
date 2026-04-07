package config

import (
	"bytes"
	"embed"

	"go.yaml.in/yaml/v4"
)

//go:embed config.dev.yml
var configFS embed.FS

const (
	DevConfig  = "config.dev.yml"
	ProdConfig = "config.prod.yml"
)

// Reads file from local directory
func LoadConfig() (*Config, error) {
	f, err := configFS.ReadFile(DevConfig)
	if err != nil {
		return nil, err
	}

	var cfg Config
	if err := yaml.NewDecoder(bytes.NewReader(f)).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, err
}
func (cfg *Config) Write() {
	
}
