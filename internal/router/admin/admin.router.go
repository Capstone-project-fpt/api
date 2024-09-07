package admin

import (
	"github.com/api/internal/constant"
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct{}

func (ar *AdminRouter) InitAdminRouter(r *gin.RouterGroup) {
	adminController := wire.InitializeAdminController()

	adminRouter := r.Group("/admin")
	adminRouter.Use(middleware.AuthMiddleware())

	adminStudentRouter := adminRouter.Group("/students")
	{
		adminStudentRouter.POST(
			"/create-account",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageAccount,
			),
			adminController.CreateStudentAccount,
		)
	}

	adminTeacherRouter := adminRouter.Group("/teachers")
	{
		adminTeacherRouter.POST(
			"/create-account", 
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageAccount,
			),
			adminController.CreateTeacherAccount,
		)
	}
}
