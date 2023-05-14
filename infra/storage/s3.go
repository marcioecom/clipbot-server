package storage

import (
	"io"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/marcioecom/clipbot-server/helper"
)

func New() (*Storage, error) {
	var (
		key      = helper.GetEnv("aws_key").String()
		secret   = helper.GetEnv("aws_secret").String()
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

	uploader := s3manager.NewUploader(sess, func(u *s3manager.Uploader) {
		u.PartSize = 5 * 1024 * 1024                                      // 5MB per part
		u.BufferProvider = s3manager.NewBufferedReadSeekerWriteToPool(20) // 20MB buffer
	})

	return &Storage{
		client:   s3.New(sess),
		uploader: uploader,
	}, nil
}

func (s *Storage) Upload(key string, body io.ReadSeeker) error {
	object := &s3manager.UploadInput{
		Key:    aws.String(key),
		Body:   body,
		ACL:    aws.String("private"),
		Bucket: aws.String(helper.GetEnv("aws_bucket").String()),
	}

	_, err := s.uploader.Upload(object)
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
