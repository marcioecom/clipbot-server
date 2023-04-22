package storage

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/marcioecom/clipbot-server/helper"
)

func New() (*Storage, error) {
	var (
		key      = helper.GetEnv("key").String()
		secret   = helper.GetEnv("secret").String()
		endpoint = helper.GetEnv("aws_endpoint").String()
		region   = helper.GetEnv("aws_region").FallBack("us-east-1")
	)

	s3Config := aws.NewConfig().
		WithCredentials(credentials.NewStaticCredentials(key, secret, "")).
		WithEndpoint(endpoint).
		WithRegion(region).
		WithS3ForcePathStyle(true)

	sess, err := session.NewSession(s3Config)
	if err != nil {
		return nil, err
	}

	return &Storage{
		client: s3.New(sess),
	}, nil
}

func (s *Storage) Upload(key string, body io.ReadSeeker) error {
	object := &s3.PutObjectInput{
		Key:    aws.String(key),
		Body:   body,
		ACL:    aws.String("private"),
		Bucket: aws.String(helper.GetEnv("aws_bucket").String()),
	}

	_, err := s.client.PutObject(object)
	if err != nil {
		return err
	}

	return nil
}

func (s *Storage) Download(key, path string) error {
	input := &s3.GetObjectInput{
		Key:    aws.String(key),
		Bucket: aws.String(helper.GetEnv("aws_bucket").String()),
	}

	result, err := s.client.GetObject(input)
	if err != nil {
		return err
	}

	out, err := os.Create(path)
	defer out.Close()

	_, err = io.Copy(out, result.Body)
	if err != nil {
		return err
	}

	return nil
}
