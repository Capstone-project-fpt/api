package topic_reference

import (
	"github.com/api/internal/constant"
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type TopicReferenceRouter struct{}

func (tr *TopicReferenceRouter) InitTopicReferenceRouter(r *gin.RouterGroup) {
	topicReferenceController := wire.InitializeTopicReferenceController()

	topicReferenceRouter := r.Group("/topic-references")

	teacherRouter := topicReferenceRouter.Group("/teachers")
	teacherRouter.Use(middleware.AuthMiddleware(constant.UserType.Teacher))
	{
		teacherRouter.POST(
			"/",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageTopicReference,
			),
			topicReferenceController.TeacherCreateTopicReference,
		)
		teacherRouter.PUT(
			"/",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageTopicReference,
			),
			topicReferenceController.TeacherUpdateTopicReference,
		)
		teacherRouter.DELETE(
			"/:id",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageTopicReference,
			),
			topicReferenceController.TeacherDeleteTopicReference,
		)
	}

	adminRouter := topicReferenceRouter.Group("/admins")
	adminRouter.Use(middleware.AuthMiddleware(constant.UserType.Admin))
	{
		adminRouter.POST(
			"/",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageTopicReference,
			),
			topicReferenceController.AdminCreateTopicReference,
		)
	}
}
