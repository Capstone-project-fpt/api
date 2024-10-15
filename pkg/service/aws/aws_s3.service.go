package aws

import (
	"time"

	"github.com/api/global"
	"github.com/aws/aws-sdk-go/service/s3"
)

type AwsS3Service interface {
	GenerateUploadPresignUrl(key string) (string, error)
}

type awsS3Service struct{}

func NewAwsS3Service() AwsS3Service {
	return &awsS3Service{}
}

func (s *awsS3Service) GenerateUploadPresignUrl(key string) (string, error) {
	req, _ := global.S3Client.PutObjectRequest(&s3.PutObjectInput{
		Bucket: &global.Config.S3.BucketName,
		Key:    &key,
	})

	presignUr, err := req.Presign(15 * time.Minute)

	if err != nil {
		return "", err
	}

	return presignUr, nil
}
