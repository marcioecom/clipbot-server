package storage

import "github.com/aws/aws-sdk-go/service/s3"

type Storage struct {
	client *s3.S3
}
