package secrets

import (
	"context"
	"fmt"
	"os"

	secretmanager "cloud.google.com/go/secretmanager/apiv1"
	"cloud.google.com/go/secretmanager/apiv1/secretmanagerpb"
)

var _ Manager = (*GoogleSecretManager)(nil)

type GoogleSecretManager struct {
	client *secretmanager.Client
}

func NewGoogle(ctx context.Context) (*GoogleSecretManager, error) {

	client, err := secretmanager.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("New google secret manager: %w", err)
	}
	return &GoogleSecretManager{client: client}, nil
}
func (sm *GoogleSecretManager) Get(ctx context.Context, key string) (string, error) {
	s, err := sm.client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/666665152595/secrets/%s/versions/latest", key),
	})
	if err != nil {
		return "", fmt.Errorf("Google secret manager get: %w", err)
	}
	return string(s.Payload.Data), nil
}
func (sm *GoogleSecretManager) SetEnv(ctx context.Context, key string) error {
	s, err := sm.client.AccessSecretVersion(ctx, &secretmanagerpb.AccessSecretVersionRequest{
		Name: fmt.Sprintf("projects/666665152595/secrets/%s/versions/latest", key),
	})
	if err != nil {
		return fmt.Errorf("Google secret manager SetEnv: %w", err)
	}
	if err := os.Setenv(key, string(s.Payload.Data)); err != nil {
		return fmt.Errorf("Google secret manager SetEnv: %w", err)
	}
	return nil

}
