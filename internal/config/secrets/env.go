package secrets

import (
	"bytes"
	"context"
	"embed"
	"fmt"

	"github.com/joho/godotenv"
)

var _ Manager = (*Env)(nil)

//go:embed .env.dev
var secretFS embed.FS

type Env struct {
	env map[string]string
}

func NewEnv() (*Env, error) {
	f, err := secretFS.ReadFile(".env.dev")
	env, err := godotenv.Parse(bytes.NewReader(f))
	if err != nil {
		return nil, fmt.Errorf("Read .env file: %w", err)
	}
	return &Env{env}, nil
}

func (m *Env) Get(ctx context.Context, name string) (string, error) {

	secret, exists := m.env[name]
	if !exists {
		return "", fmt.Errorf("Get secret not found: %v", name)
	}
	if secret == "" {
		return "", fmt.Errorf("Env secret is empty: %v", name)
	}
	return secret, nil
}

func (m *Env) Clear() {

}
