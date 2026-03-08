package config

import (
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"musicproject.com/pkg/util/fileutil"
)

type Config struct {
	API        apiCfg        `yaml:"api"`
	Repository repositoryCfg `yaml:"repository"`
	Oauth      oauthCfg      `yaml:"oauth"`
}

type apiCfg struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Jwt  jwtCfg `yaml:"jwt"`
}

type jwtCfg struct {
	AccessKey  string `yaml:"accessKey"`
	RefreshKey string `yaml:"refreshKey"`
}

type repositoryCfg struct {
	Type    string `yaml:"type"`
	URL     string `yaml:"url"`
	TestURL string `yaml:"testUrl"`
}

type oauthCfg struct {
	Google oauthGoogle `yaml:"google"`
}

type oauthGoogle struct {
	ClientID     string   `yaml:"clientId"`
	ClientSecret string   `yaml:"clientSecret"`
	RedirectURL  string   `yaml:"redirectUrl"`
	Scopes       []string `yaml:"scopes"`
}

// Reads file from local directory
func ReadConfigFile(fileName string) Config {
	cfg, err := fileutil.ReadYAML[Config](fileName)
	if err != nil {
		panic(err)
	}
	return cfg
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
