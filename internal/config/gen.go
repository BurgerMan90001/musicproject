package config

type Config struct {
	API        API        `yaml:"api"`
	Auth       Auth       `yaml:"auth"`
	Upload     Upload     `yaml:"upload"`
	Middleware Middleware `yaml:"middleware"`
	Repository Repository `yaml:"repository"`
}
type API struct {
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Version string `yaml:"version"`
}
type Jwt struct {
	Issuer   string   `yaml:"issuer"`
	Audience []string `yaml:"audience"`
}
type Google struct {
	RedirectURL string   `yaml:"redirectUrl"`
	Scopes      []string `yaml:"scopes"`
}
type Oauth struct {
	Google Google `yaml:"google"`
}
type Auth struct {
	Jwt   Jwt   `yaml:"jwt"`
	Oauth Oauth `yaml:"oauth"`
}
type Encoder struct {
	Logging bool `yaml:"logging"`
	Enabled bool `yaml:"enabled"`
}
type Upload struct {
	Store   string  `yaml:"store"`
	Region  string  `yaml:"region"`
	Bucket  string  `yaml:"bucket"`
	Encoder Encoder `yaml:"encoder"`
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
