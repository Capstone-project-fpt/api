package service

type IScheduleReviewService interface {}

type scheduleReviewService struct {}

func NewScheduleReviewService() IScheduleReviewService {
	return &scheduleReviewService{}	
}