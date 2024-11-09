package controller

import "github.com/api/internal/service"

type ScheduleReviewController struct {
	scheduleReviewService service.IScheduleReviewService
}

func NewScheduleReviewController(scheduleReviewService service.IScheduleReviewService) *ScheduleReviewController {
	return &ScheduleReviewController{
		scheduleReviewService: scheduleReviewService,
	}
}
