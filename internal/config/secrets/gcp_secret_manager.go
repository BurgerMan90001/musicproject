package secrets

import "context"

var _ Manager = (*GoogleSecretManager)(nil)

type GoogleSecretManager struct {
}

func NewGoogle() *GoogleSecretManager {
	return &GoogleSecretManager{}
}
func (m *GoogleSecretManager) Get(ctx context.Context, key string) (string, error) {
	return "", nil
}
