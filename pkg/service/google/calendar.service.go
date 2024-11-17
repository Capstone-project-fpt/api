package google

import (
	"fmt"

	string_util "github.com/api/pkg/utils/string"
)

type IGoogleService interface {
	GenerateGoogleMeeting() string
}

type googleService struct{}

func NewGoogleService() IGoogleService {
	return &googleService{}
}

func (s *googleService) GenerateGoogleMeeting() string {
	meeting_code := string_util.GenerateRandomString(10) // TODO: Get from Google service
	return fmt.Sprintf("https://meet.google.com/" + meeting_code)
}
