package file

import (
	"bytes"
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type AWSS3 struct {
	client *s3.Client
}

func NewS3(ctx context.Context) (*AWSS3, error) {
	cfg := aws.Config{}
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	client := s3.NewFromConfig(cfg)

	return &AWSS3{client: client}, nil
}

func (s *AWSS3) CreateObject(ctx context.Context, bucket string, key string,
	contents []byte, cacheble bool, contentType string) error {
	cacheControl := "public, max-age=86400"
	if !cacheble {
		cacheControl = "no-cache, max-age=0"
	}

	putInput := &s3.PutObjectInput{
		Bucket:       aws.String(bucket),
		Key:          aws.String(key),
		CacheControl: aws.String(cacheControl),
		Body:         bytes.NewReader(contents),
	}
	if contentType != "" {
		putInput.ContentType = aws.String(contentType)
	}
	if _, err := s.client.PutObject(ctx, putInput); err != nil {
		return err
	}
	return nil
}
func (s *AWSS3) GetObject(ctx context.Context, bucket string, key string) ([]byte, error) {
	getInput := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}

	out, err := s.client.GetObject(ctx, getInput)
	if err != nil {
		return nil, err
	}
	defer out.Body.Close()

	data, err := io.ReadAll(out.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *AWSS3) DeleteObject(ctx context.Context, bucket string, key string) error {
	deleteInput := &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}
	if _, err := s.client.DeleteObject(ctx, deleteInput); err != nil {
		return err
	}

	return nil
}
