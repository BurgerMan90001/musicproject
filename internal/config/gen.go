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
	Issuer string `yaml:"issuer"`
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
type Buckets struct {
	Audio string `yaml:"audio"`
	Log   string `yaml:"log"`
}
type S3 struct {
	Buckets Buckets `yaml:"buckets"`
}
type Aws struct {
	S3 S3 `yaml:"s3"`
}
