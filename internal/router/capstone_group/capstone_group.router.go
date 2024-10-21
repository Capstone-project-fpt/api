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
		capstoneGroupRouter.POST("/:id/mentors", capstoneGroupController.InviteMentorToCapstoneGroup)
		capstoneGroupRouter.POST("/:id/mentors/invitation", middleware.UserTypeMiddleware(constant.UserType.Teacher), capstoneGroupController.AcceptInviteMentorToCapstoneGroup)
		capstoneGroupRouter.PUT("/", capstoneGroupController.UpdateCapstoneGroup)
		capstoneGroupRouter.GET("/", capstoneGroupController.GetListCapstoneGroups)
		capstoneGroupRouter.GET("/:id", capstoneGroupController.GetCapstoneGroup)
	}
}
