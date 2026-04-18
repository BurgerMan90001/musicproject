package secrets

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var _ Manager = (*AWSSecretManager)(nil)

type AWSSecretManager struct {
	client *secretsmanager.Client
}

func NewAWS(ctx context.Context, region string) (*AWSSecretManager, error) {
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithRegion(region),
	)
	if err != nil {
		return nil, err
	}
	client := secretsmanager.NewFromConfig(cfg)
	
	return &AWSSecretManager{client: client}, nil
}

func (m *AWSSecretManager) Get(ctx context.Context, name string) (string, error) {
	getInput := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(name),
	}

	out, err := m.client.GetSecretValue(ctx, getInput)
	if err != nil {
		return "", fmt.Errorf("AWSSecretManager get: %w", err)
	}
	if out.SecretString != nil {
		return *out.SecretString, nil
	}
	// For secrets in binary
	return string(out.SecretBinary), nil
}
