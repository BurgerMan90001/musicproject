package file

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/url"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"songsled.com/internal/config/secrets"
)

var _ Blobstore = (*AWSS3)(nil)

// AWS S3 / R2
type AWSS3 struct {
	client        *s3.Client
	presignClient *s3.PresignClient
	endpoint      string
	public        string
}

func NewS3(ctx context.Context,
	endpoint, region, public, accessKeyIdVar, accessKeySecretVar string) (*AWSS3, error) {

	s, err := secrets.GetenvMap(accessKeyIdVar, accessKeySecretVar)
	if err != nil {
		return nil, err
	}
	cfg, err := config.LoadDefaultConfig(ctx,
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(s[accessKeyIdVar], s[accessKeySecretVar], "")),
		config.WithRegion(region),
		config.WithRequestChecksumCalculation(aws.RequestChecksumCalculationWhenRequired),
		// config.WithChec
	)
	if err != nil {
		return nil, err
	}
	if public == "" {

	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(endpoint)

	})
	presignClient := s3.NewPresignClient(client)

	if _, err := client.ListObjectsV2(ctx, &s3.ListObjectsV2Input{
		Bucket: aws.String("songsled"),
	}); err != nil {
		return nil, fmt.Errorf("S3 list objects: %w", err)
	}
	return &AWSS3{client: client, presignClient: presignClient}, nil
}

func (s *AWSS3) CreateObject(ctx context.Context, bucket string, key string,
	contents []byte, cacheble bool, contentType string) error {

	putInput := s.newPutInput(bucket, key,
		contents, cacheble, contentType)
	if _, err := s.client.PutObject(ctx, putInput); err != nil {
		return err
	}
	return nil
}

// Returns a presigned url to upload files to
func (s *AWSS3) CreateObjectUrl(ctx context.Context,
	bucket, key string, cacheble bool, ttl time.Duration) (string, string, error) {
	putInput := s.newPutInput(bucket, key,
		nil, cacheble, "")

	presignUrl, err := s.presignClient.PresignPutObject(ctx, putInput, func(po *s3.PresignOptions) {
		s3.WithPresignExpires(ttl)

	})
	if err != nil {
		return "", "", err
	}

	objectUrl, err := url.JoinPath(s.endpoint, bucket, key)
	if err != nil {
		return "", "", err
	}
	return presignUrl.URL, objectUrl, nil
}
func (s *AWSS3) newPutInput(bucket string, key string,
	contents []byte, cacheble bool, contentType string) *s3.PutObjectInput {
	// cacheControl := "public, max-age=86400"
	// if !cacheble {
	// 	cacheControl = "no-cache, max-age=0"
	// }
	putInput := &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		
		// CacheControl:      aws.String(cacheControl),
		// ChecksumAlgorithm: types.ChecksumAlgorithmCrc32,
	}
	if contents != nil {
		putInput.Body = bytes.NewReader(contents)
	}
	// if contentType != "" {
	// 	putInput.ContentType = aws.String(contentType)
	// }
	return putInput
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
func (s *AWSS3) GetObjectUrl(ctx context.Context, bucket, key string, ttl time.Duration) (string, error) {
	req, err := s.presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, func(po *s3.PresignOptions) {
		s3.WithPresignExpires(30 * time.Minute)
	})
	if err != nil {
		return "", fmt.Errorf("Get object url: %w", err)
	}
	return req.URL, nil
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
