package config

type ServiceConfig struct {
	APIConfig        apiConfig        `yaml:"api"`
	RepositoryConfig repositoryConfig `yaml:"repository"`
}

type apiConfig struct {
	Host        string      `yaml:"host"`
	Port        int         `yaml:"port"`
	JWTKey      string      `yaml:"jwtKey"`
	Oath2Config oath2Config `yaml:"oath2"`
}

type repositoryConfig struct {
	Type string `yaml:"type"`
	URL  string `yaml:"url"`
}

type oath2Config struct {
	ClientID     string   `yaml:"clientID"`
	ClientSecret string   `yaml:"clientSecret"`
	RedirectURL  string   `yaml:"redirectUrl"`
	Scopes       []string `yaml:"scopes"`
}
