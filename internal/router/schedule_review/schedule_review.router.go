package schedule_review

import (
	"github.com/api/internal/constant"
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type ScheduleReviewRouter struct {}

func (s *ScheduleReviewRouter) InitScheduleReviewRouter(group *gin.RouterGroup) {
	scheduleReviewController := wire.InitializeScheduleReviewController()

	scheduleReviewRouter := group.Group("/schedule-reviews")
	scheduleReviewRouter.Use(middleware.AuthMiddleware())
	{
		scheduleReviewRouter.POST(
			"/",
			middleware.UserTypeMiddleware(constant.UserType.Admin),
			scheduleReviewController.CreateScheduleReview,
		)
		scheduleReviewRouter.PUT(
			"/",
			middleware.UserTypeMiddleware(constant.UserType.Admin),
			scheduleReviewController.UpdateScheduleReview,
		)
		scheduleReviewRouter.DELETE(
			"/:schedule_review_id",
			middleware.UserTypeMiddleware(constant.UserType.Admin),
			scheduleReviewController.DeleteScheduleReview,
		)
		scheduleReviewRouter.GET(
			"/:schedule_review_id",
			scheduleReviewController.GetScheduleReviewDetail,
		)
	} 
}