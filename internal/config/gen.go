package config

type Config struct {
	API        API        `yaml:"api"`
	Auth       Auth       `yaml:"auth"`
	File       File       `yaml:"file"`
	Repository Repository `yaml:"repository"`
	Test       Test       `yaml:"test"`
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
type Auth struct {
	Jwt Jwt `yaml:"jwt"`
}
type Encoder struct {
	Logging bool `yaml:"logging"`
	Enabled bool `yaml:"enabled"`
}
type Upload struct {
	Encoder Encoder `yaml:"encoder"`
}
type File struct {
	Region   string `yaml:"region"`
	Bucket   string `yaml:"bucket"`
	Endpoint string `yaml:"endpoint"`
	Public   string `yaml:"public"`
	Upload   Upload `yaml:"upload"`
}
type Postgres struct {
	Schema string `yaml:"schema"`
	Image  string `yaml:"image"`
}
type Repository struct {
	Postgres Postgres `yaml:"postgres"`
}
type Test struct {
	ShowErrorDetails bool `yaml:"showErrorDetails"`
	LoadTestdata     bool `yaml:"loadTestdata"`
}
