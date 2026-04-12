package config

type Config struct {
	API        API        `yaml:"api"`
	Services   Services   `yaml:"services"`
	Middleware Middleware `yaml:"middleware"`
	Repository Repository `yaml:"repository"`
	Aws        Aws        `yaml:"aws"`
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
}
type Services struct {
	Auth    Auth    `yaml:"auth"`
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
type Buckets struct {
	Audio string `yaml:"audio"`
	Logs  string `yaml:"logs"`
}
type S3 struct {
	Buckets Buckets `yaml:"buckets"`
}
type Aws struct {
	S3 S3 `yaml:"s3"`
}
