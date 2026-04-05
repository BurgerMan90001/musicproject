package config

type Config struct {
	API        API         `yaml:"api"`
	Services   Services    `yaml:"services"`
	Middleware Middleware  `yaml:"middleware"`
	Repository Repository  `yaml:"repository"`
	Aws        interface{} `yaml:"aws"`
}
type API struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
}
type Google struct {
	RedirectURL string   `yaml:"redirectURL"`
	Scopes      []string `yaml:"scopes"`
}
type Oauth struct {
	Google Google `yaml:"google"`
}
type Auth struct {
	Oauth Oauth `yaml:"oauth"`
}
type Services struct {
	Auth Auth `yaml:"auth"`
}
type Middleware struct {
	Logger    bool `yaml:"logger"`
	Ratelimit bool `yaml:"ratelimit"`
}
type Postgres struct {
	Image string `yaml:"image"`
}
type Repository struct {
	Postgres Postgres `yaml:"postgres"`
}
