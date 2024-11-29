package capstone_group_router

import (
	"github.com/api/internal/constant"
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type CapstoneGroupRouter struct{}

func (cgr *CapstoneGroupRouter) InitCapstoneGroupRouter(r *gin.RouterGroup) {
	capstoneGroupController := wire.InitializeCapstoneGroupController()

	capstoneGroupRouter := r.Group("/capstone-groups")
	capstoneGroupRouter.Use(middleware.AuthMiddleware())
	{
		capstoneGroupRouter.POST("/", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.CreateCapstoneGroup)
		capstoneGroupRouter.POST("/:capstone_group_id/mentors", capstoneGroupController.InviteMentorToCapstoneGroup)
		capstoneGroupRouter.POST("/:capstone_group_id/mentors/invitations", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.ResponseInviteMentorToCapstoneGroup)
		capstoneGroupRouter.PUT("/", capstoneGroupController.UpdateCapstoneGroup)
		capstoneGroupRouter.GET("/", capstoneGroupController.GetListCapstoneGroups)
		capstoneGroupRouter.GET("/current-semester", capstoneGroupController.GetCurrentListCapstoneGroup)
		capstoneGroupRouter.GET("/semesters/:semester_id/students", capstoneGroupController.GetListStudentHaveCapstoneGroup)
		capstoneGroupRouter.GET("/:capstone_group_id", capstoneGroupController.GetCapstoneGroup)
		capstoneGroupRouter.GET("/:capstone_group_id/members", capstoneGroupController.GetMentorAndListMemberCapstoneGroup)
		capstoneGroupRouter.GET("/:capstone_group_id/mentors/invitations", capstoneGroupController.GetListInvitationMentorCapstoneGroups)
		capstoneGroupRouter.PUT("/:capstone_group_id/members", middleware.UserTypeMiddleware(constant.UserType.Admin), capstoneGroupController.UpdateCapstoneGroupStudent)
	}

	capstoneGroupTopicRouter := capstoneGroupRouter.Group("/:capstone_group_id/capstone-group-topics")
	{
		capstoneGroupTopicRouter.POST("/", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.CreateCapstoneGroupTopic)
		capstoneGroupTopicRouter.PUT("/:id", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.UpdateCapstoneGroupTopic)
		capstoneGroupTopicRouter.POST("/:id", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.SelectCapstoneGroupTopic)
		capstoneGroupTopicRouter.DELETE("/:id", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.DeleteCapstoneGroupTopic)
		capstoneGroupTopicRouter.GET("/", capstoneGroupController.GetListCapstoneGroupTopic)
		capstoneGroupTopicRouter.GET("/:id", capstoneGroupController.GetCapstoneGroupTopic)
	}

	capstoneGroupTopicTeacherReviewRouter := capstoneGroupTopicRouter.Group("/:id/teacher-reviews")
	{
		capstoneGroupTopicTeacherReviewRouter.PUT("/", middleware.UserTypeMiddleware(constant.UserType.Teacher), middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.ReviewCapstoneGroupTopic)
	}

	capstoneGroupTopicTeacherFeedbackRouter := capstoneGroupTopicRouter.Group("/:id/feedbacks")
	{
		capstoneGroupTopicTeacherFeedbackRouter.POST("/", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.FeedbackCapstoneGroupTopic)
		capstoneGroupTopicTeacherFeedbackRouter.PUT("/:feedback_id", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.UpdateFeedbackCapstoneGroupTopic)
		capstoneGroupTopicTeacherFeedbackRouter.DELETE("/:feedback_id", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.DeleteFeedbackCapstoneGroupTopic)
		capstoneGroupTopicTeacherFeedbackRouter.GET("/", capstoneGroupController.GetListCapstoneGroupTopicFeedback)
		capstoneGroupTopicTeacherFeedbackRouter.GET("/:feedback_id", capstoneGroupController.GetCapstoneGroupTopicFeedback)
	}

	capstoneGroupReviewRouter := capstoneGroupRouter.Group("/:capstone_group_id/capstone-group-reviews")
	{
		capstoneGroupReviewRouter.PATCH("/", capstoneGroupController.UpdateReportFilesCapstoneGroupReview)
		capstoneGroupReviewRouter.PATCH("/feedback", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.FeedbackCapstoneGroupReview)
		capstoneGroupReviewRouter.GET("/:id", capstoneGroupController.GetCapstoneGroupReview)
	}

	capstoneGroupReportDocumentRouter := capstoneGroupRouter.Group("/:capstone_group_id/report-documents")
	{
		capstoneGroupReportDocumentRouter.POST("/", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.CreateCapstoneGroupReportDocument)
		capstoneGroupReportDocumentRouter.PUT("/", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.UpdateCapstoneGroupReportDocument)
		capstoneGroupReportDocumentRouter.DELETE("/:report_document_id", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.DeleteCapstoneGroupReportDocument)
		capstoneGroupReportDocumentRouter.GET("/", capstoneGroupController.GetCapstoneGroupReportDocuments)
		capstoneGroupReportDocumentRouter.GET("/:report_document_id", capstoneGroupController.GetCapstoneGroupReportDocument)
		capstoneGroupReportDocumentRouter.PUT("/scores", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.MentorUpdateStudentScoreForReportDocument)
		capstoneGroupReportDocumentRouter.PATCH("/scores", middleware.UserTypeMiddleware(constant.UserType.Admin), capstoneGroupController.AdminUpdateStudentScoreForReportDocument)
		capstoneGroupReportDocumentRouter.GET("/scores/:report_document_id", capstoneGroupController.GetReportDocumentStudentsScore)
	}

	capstoneGroupReportDocumentCommentRouter := capstoneGroupRouter.Group("/:capstone_group_id/report-documents/:report_document_id/comments")
	{
		capstoneGroupReportDocumentCommentRouter.POST("/", capstoneGroupController.CommentReport)
		capstoneGroupReportDocumentCommentRouter.PATCH("/", capstoneGroupController.UpdateCommentReport)
		capstoneGroupReportDocumentCommentRouter.DELETE("/", capstoneGroupController.DeleteCommentReport)
		capstoneGroupReportDocumentCommentRouter.GET("/", capstoneGroupController.GetListReportComments)
	}
}
