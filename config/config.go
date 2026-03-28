package config

import (
	"bytes"
	"embed"
	"fmt"

	"go.yaml.in/yaml/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type JWT struct {
	AccessKey  string `yaml:"accessKey"`
	RefreshKey string `yaml:"refreshKey"`
	Issuer     string `yaml:"issuer"`
}
type Config struct {
	API        API        `yaml:"api"`
	Services   Services   `yaml:"services"`
	Repository Repository `yaml:"repository"`
}
type API struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
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
type Services struct {
	Auth Auth `yaml:"auth"`
}
type Repository struct {
	URL string `yaml:"url"`
}

//go:embed config.dev.yml
var configFS embed.FS

const (
	DevConfig  = "config.dev.yml"
	ProdConfig = "config.prod.yml"
)

// Reads file from local directory
func LoadConfig() *Config {

	f, err := configFS.ReadFile(DevConfig)
	if err != nil {
		panic(err)
	}
	var cfg Config
	if err := yaml.NewDecoder(bytes.NewReader(f)).Decode(&cfg); err != nil {
		panic(err)
	}
	return &cfg
}

func (cfg Config) URL() string {
	return fmt.Sprintf("%v:%d", cfg.API.Host, cfg.API.Port)
}

func (cfg Config) GoogleOathConfig() *oauth2.Config {
	googleCfg := cfg.Services.Auth.Oauth.Google
	return &oauth2.Config{
		ClientID:     googleCfg.ClientID,
		ClientSecret: googleCfg.ClientSecret,
		RedirectURL:  googleCfg.RedirectURL,
		Scopes:       googleCfg.Scopes,
		Endpoint:     google.Endpoint,
	}
}
