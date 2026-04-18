package secrets

import (
	"context"
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
