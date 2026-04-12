package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env map[string]string

func LoadEnv(name string) error {
	err := godotenv.Load(name)
	if err != nil {
		return fmt.Errorf("Read .env file: %w", err)
	}
	return nil
}

// Returns a map of secrets from the names
func GetEnv(names ...string) (Env, error) {

	var secrets Env = make(Env, 4)
	for _, name := range names {
		secret := os.Getenv(name)
		if secret == "" {
			return nil, fmt.Errorf("env var %v is empty", name)
		}
		secrets[name] = secret
	}

	return secrets, nil
}
