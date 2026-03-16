package config

import (
	"fmt"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"musicproject.com/pkg/util/fileutil"
)

type Config struct {
	API struct {
		Port    int    `yaml:"port"`
		Host    string `yaml:"host"`
		Version string `yaml:"version"`
		Jwt     struct {
			AccessKey  string `yaml:"accessKey"`
			RefreshKey string `yaml:"refreshKey"`
			Issuer     string `yaml:"issuer"`
		} `yaml:"jwt"`
	} `yaml:"api"`
	Repository struct {
		Type string `yaml:"type"`
		URL  string `yaml:"url"`
	} `yaml:"repository"`
	Oauth struct {
		Google struct {
			ClientID     string   `yaml:"clientId"`
			ClientSecret string   `yaml:"clientSecret"`
			RedirectURL  string   `yaml:"redirectUrl"`
			Scopes       []string `yaml:"scopes"`
		} `yaml:"google"`
	} `yaml:"oauth"`
}

// Reads file from local directory
func ReadConfigFile() Config {
	var fileName string

	env := os.Getenv("env")
	switch env {
	case "dev", "test":
		fileName = "./base.dev.yml"
	case "prod":
		fileName = "./base.prod.yml"
	default:
		fileName = "./base.dev.yml"
	}
	cfg, err := fileutil.ReadYAML[Config](fileName)
	if err != nil {
		panic(err)
	}
	return cfg
}
func (cfg Config) ApiUrl() string {
	return fmt.Sprintf("%v:%d", cfg.API.Host, cfg.API.Port)
}

func (cfg Config) GoogleOathConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     cfg.Oauth.Google.ClientID,
		ClientSecret: cfg.Oauth.Google.ClientSecret,
		RedirectURL:  cfg.Oauth.Google.RedirectURL,
		Scopes:       cfg.Oauth.Google.Scopes,
		Endpoint:     google.Endpoint,
	}
}

func (cfg Config) JWTAccessKey() []byte {
	return []byte(cfg.API.Jwt.AccessKey)
}
