package secrets

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var sm Manager

func init() {
	ctx := context.Background()
	var err error
	sm, err = NewGoogle(ctx)
	if err != nil {
		panic(err)
	}

	// if err := sm.SetEnv(ctx, "AWS_SECRET_ACCESS_KEY"); err != nil {
	// 	panic(err)
	// }
}

// var sm, err = NewGoogle(ctx)

type Manager interface {
	Get(ctx context.Context, key string) (string, error)
	SetEnv(ctx context.Context, key string) error
}

// type Postgres struct {
// 	database string
// }

// type SMTP struct {
// }
func ReadEnvFile(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		return fmt.Errorf("Read .env file: %w", err)
	}
	return nil
}

// Returns a map of secrets from the names
func GetenvMap(keys ...string) (map[string]string, error) {
	var secrets = make(map[string]string, 4)
	for _, key := range keys {
		secret, err := Getenv(key)
		if err != nil {
			return nil, err
		}
		secrets[key] = secret
	}
	return secrets, nil
}

// return sm.Get(context.Background(), key)
func Getenv(key string) (string, error) {
	s := os.Getenv(key)
	if s == "" {
		var err error
		s, err = sm.Get(context.Background(), key)
		if err != nil {
			return "", fmt.Errorf("Secrets get env var %s: %w", key, err)
		}
	}
	return s, nil
}

func LoadTemplate(envFile string) error {
// 666665152595
	// for _, e := range
	return nil
}
