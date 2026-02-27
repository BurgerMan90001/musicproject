package config

import (
	"golang.org/x/oauth2"
	"okapi.com/internal/util/fileutil"
)

type ServiceConfig struct {
	APIConfig        apiConfig        `yaml:"api"`
	RepositoryConfig repositoryConfig `yaml:"repository"`
	Oath2Config      oauth2.Config    `yaml:"oath2"`
}

type apiConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	JWTKey string `yaml:"jwtKey"`
}

type repositoryConfig struct {
	Type    string `yaml:"type"`
	URL     string `yaml:"url"`
	TestURL string `yaml:"testUrl"`
}

type oath2Config struct {
	//oauth2.Config
	ClientID     string   `yaml:"clientId"`
	ClientSecret string   `yaml:"clientSecret"`
	RedirectURL  string   `yaml:"redirectUrl"`
	Scopes       []string `yaml:"scopes"`
}

func ReadConfigFile() ServiceConfig {
	cfg, err := fileutil.ReadYAML[ServiceConfig]("config/base.yml")
	if err != nil {
		panic(err)
	}
	return cfg
}
func getOathConfig(cfg oath2Config) *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		RedirectURL:  cfg.RedirectURL,
		Scopes:       cfg.Scopes,
	}
}
