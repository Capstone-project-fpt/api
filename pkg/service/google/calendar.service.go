package google

import "fmt"

type IGoogleService interface {
	GenerateGoogleMeeting() string
}

type googleService struct {}

func NewGoogleService() IGoogleService {
	return &googleService{}
}

func(s *googleService) GenerateGoogleMeeting() string {
	return fmt.Sprintf("https://meet.google.com" + "/123456789")
}