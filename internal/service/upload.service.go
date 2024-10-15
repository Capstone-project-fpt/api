package service

import "github.com/api/pkg/service/aws"

type IUploadService interface {
	GenerateUploadPresignUrl(key string) (string, error)
}

type uploadService struct {
	awsS3Service aws.AwsS3Service
}

func NewUploadService(awsS3Service aws.AwsS3Service) IUploadService {
	return &uploadService{
		awsS3Service: awsS3Service,
	}
}

func (s *uploadService) GenerateUploadPresignUrl(key string) (string, error) {
	preSignUrl, err := s.awsS3Service.GenerateUploadPresignUrl(key)

	if err != nil {
		return "", err
	}

	return preSignUrl, nil
}
