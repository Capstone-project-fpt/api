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
		capstoneGroupRouter.POST("/:capstone_group_id/mentors/invitation", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.AcceptInviteMentorToCapstoneGroup)
		capstoneGroupRouter.PUT("/", capstoneGroupController.UpdateCapstoneGroup)
		capstoneGroupRouter.GET("/", capstoneGroupController.GetListCapstoneGroups)
		capstoneGroupRouter.GET("/:capstone_group_id", capstoneGroupController.GetCapstoneGroup)
	}

	capstoneGroupTopicRouter := capstoneGroupRouter.Group("/:capstone_group_id/capstone-group-topics")
	{
		capstoneGroupTopicRouter.POST("/", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.CreateCapstoneGroupTopic)
		capstoneGroupTopicRouter.PUT("/:id", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.UpdateCapstoneGroupTopic)
		capstoneGroupTopicRouter.DELETE("/:id", middleware.UserTypeMiddleware(constant.UserType.Student), capstoneGroupController.DeleteCapstoneGroupTopic)
		capstoneGroupTopicRouter.GET("/", capstoneGroupController.GetListCapstoneGroupTopic)
		capstoneGroupTopicRouter.GET("/:id", capstoneGroupController.GetCapstoneGroupTopic)
	}
}
