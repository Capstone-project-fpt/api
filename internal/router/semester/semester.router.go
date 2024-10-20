package semester

import (
	"github.com/api/internal/constant"
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type SemesterRouter struct{}

func (ur *SemesterRouter) InitSemesterRouter(r *gin.RouterGroup) {
	semesterController := wire.InitializeSemesterController()

	semesterRouter := r.Group("/semesters")
	semesterRouter.Use(middleware.AuthMiddleware())

	semesterCommonRouter := semesterRouter
	{
		semesterCommonRouter.GET("/", semesterController.GetListSemesters)
		semesterCommonRouter.GET("/:id", semesterController.GetSemester)
	}

	semesterAdminRouter := semesterRouter.Group("")
	{
		semesterAdminRouter.POST("/", middleware.UserTypeMiddleware(constant.UserType.Admin), semesterController.AdminCreateSemester)
		semesterAdminRouter.PUT("/", middleware.UserTypeMiddleware(constant.UserType.Admin), semesterController.AdminUpdateSemester)
		semesterRouter.DELETE("/:id", middleware.UserTypeMiddleware(constant.UserType.Admin), semesterController.AdminDeleteSemester)
	}
}
