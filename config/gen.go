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
type Jwt struct {
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
	Jwt   Jwt   `yaml:"jwt"`
	Oauth Oauth `yaml:"oauth"`
}
type SMTP struct {
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
	Email    string `yaml:"email"`
	Password string `yaml:"password"`
}
type Services struct {
	Auth Auth `yaml:"auth"`
	SMTP SMTP `yaml:"smtp"`
}
type Middleware struct {
	Ratelimit bool `yaml:"ratelimit"`
}
type Postgres struct {
	Image    string `yaml:"image"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Database string `yaml:"database"`
}
type Repository struct {
	Postgres Postgres `yaml:"postgres"`
}
