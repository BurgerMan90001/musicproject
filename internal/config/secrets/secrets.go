package secrets

import (
	"context"
	"errors"
)

// TODO use secret manager. Using config file for now

type Manager interface {
	Get(ctx context.Context, name string) (string, error)
}

// type Postgres struct {
// 	database string
// }

// type SMTP struct {
// }

// Returns a slice of secrets in the same order as the input names
func GetSecrets(ctx context.Context, sm Manager, names ...string) (map[string]string, error) {
	if len(names) == 0 {
		return nil, nil
	}
	var (
		secrets   map[string]string = make(map[string]string)
		errorList []error
	)
	for _, name := range names {
		secret, err := sm.Get(ctx, name)
		secrets[name] = secret
		errorList = append(errorList, err)
	}
	if err := errors.Join(errorList...); err != nil {
		return nil, err
	}
	return secrets, nil
}
