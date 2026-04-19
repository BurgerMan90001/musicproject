package secrets

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func ReadEnvFile(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		return fmt.Errorf("Read .env file: %w", err)
	}
	return nil
}

// Returns a map of secrets from the names
func GetEnvMap(keys ...string) (map[string]string, error) {
	var secrets = make(map[string]string, 4)
	for _, key := range keys {
		secret := os.Getenv(key)
		if secret == "" {
			return nil, fmt.Errorf("Env var %v is empty", key)
		}
		secrets[key] = secret
	}
	return secrets, nil
}

// func GetEnv(key string) (string, error) {
// 	secret := os.Getenv(key)
// 	if secret == "" {
// 		return "", fmt.Errorf("Env var %v is empty", key)
// 	}
// 	return secret, nil
// }
