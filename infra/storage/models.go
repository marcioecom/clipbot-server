package storage

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Storage struct {
	client   *s3.S3
	uploader *s3manager.Uploader
}
