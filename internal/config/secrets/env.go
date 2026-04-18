package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Env map[string]string

func ReadEnvFile(file string) error {
	err := godotenv.Load(file)

	if err != nil {
		return fmt.Errorf("Read .env file: %w", err)
	}
	return nil
}

// Returns a map of secrets from the names
func GetEnv(keys ...string) (Env, error) {
	var secrets Env = make(Env, 4)
	for _, key := range keys {
		secret := os.Getenv(key)
		if secret == "" {
			return nil, fmt.Errorf("env var %v is empty", key)
		}
		secrets[key] = secret
	}

	return secrets, nil
}
