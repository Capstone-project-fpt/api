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
	adminRouter.Use(middleware.UserTypeMiddleware(constant.UserType.Admin))

	adminUserRouter := adminRouter.Group("/users")
	{
		adminUserRouter.GET("/", adminController.GetListUsers)
		adminUserRouter.GET(
			"/:id",
			middleware.PermissionMiddleware(
				constant.PermissionType.ViewAccount,
			),
			adminController.GetUser,
		)
	}

	adminStudentRouter := adminRouter.Group("/students")
	{
		adminStudentRouter.POST(
			"/create-account",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageAccount,
			),
			adminController.CreateStudentAccount,
		)

		adminStudentRouter.POST(
			"/import-data",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageAccount,
			),
			adminController.UploadFileStudentData,
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

		adminTeacherRouter.POST(
			"/import-data",
			middleware.PermissionMiddleware(
				constant.PermissionType.ManageAccount,
			),
			adminController.UploadFileTeacherData,
		)
	}
}
