package config

import (
	"bytes"
	"embed"

	"go.yaml.in/yaml/v4"
)

type Config struct {
	API        API        `yaml:"api"`
	Services   Services   `yaml:"services"`
	Middleware Middleware `yaml:"middleware"`
	Repository Repository `yaml:"repository"`
}
type API struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
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
type Services struct {
	Auth Auth `yaml:"auth"`
}
type Middleware struct {
	Ratelimit bool `yaml:"ratelimit"`
}
type Postgres struct {
	Image    string `yaml:"image"`
	Database string `yaml:"database"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
}
type Repository struct {
	Postgres Postgres `yaml:"postgres"`
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
