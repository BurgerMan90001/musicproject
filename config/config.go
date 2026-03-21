package config

import (
	"fmt"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"musicproject.com/pkg/util/fileutil"
)

type Config struct {
	API        API        `yaml:"api"`
	Repository Repository `yaml:"repository"`
}
type JWT struct {
	AccessKey  string `yaml:"accessKey"`
	RefreshKey string `yaml:"refreshKey"`
	Issuer     string `yaml:"issuer"`
}
type Google struct {
	ClientID     string   `yaml:"clientID"`
	ClientSecret string   `yaml:"clientSecret"`
	RedirectURL  string   `yaml:"redirectURL"`
	Scopes       []string `yaml:"scopes"`
}

type Oauth struct {
	Google Google `yaml:"google"`
}
type Auth struct {
	JWT   JWT   `yaml:"jwt"`
	Oauth Oauth `yaml:"oauth"`
}
type API struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
	Auth    Auth   `yaml:"auth"`
}
type Repository struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

const (
	TypeDev  = "dev"
	TypeProd = "prod"
)

// Reads file from local directory
func ReadConfigFile(env string) Config {
	var fileName string

	switch env {
	case TypeDev:
		fileName = "./config/base.dev.yml"
	case TypeProd:
		fileName = "./config/base.prod.yml"
	default:
		fileName = "./config/base.dev.yml"
	}
	cfg, err := fileutil.ReadYAML[Config](fileName)
	if err != nil {
		panic(err)
	}
	return cfg
}

func DevConfig() Config {
	cfg, err := fileutil.ReadYAML[Config]("./base.dev.yml")
	if err != nil {
		panic(err)
	}
	return cfg
}
func (cfg Config) ApiUrl() string {
	return fmt.Sprintf("%v:%d", cfg.API.Host, cfg.API.Port)
}

func (cfg Config) GoogleOathConfig() *oauth2.Config {
	googleCfg := cfg.API.Auth.Oauth.Google
	return &oauth2.Config{
		ClientID:     googleCfg.ClientID,
		ClientSecret: googleCfg.ClientSecret,
		RedirectURL:  googleCfg.RedirectURL,
		Scopes:       googleCfg.Scopes,
		Endpoint:     google.Endpoint,
	}
}

func (cfg Config) JWTAccessKey() []byte {
	return []byte(cfg.API.Auth.JWT.AccessKey)
}
