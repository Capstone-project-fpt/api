package semester

import (
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

	semesterAdminRouter := semesterRouter.Group("/admin")
	{
		semesterAdminRouter.POST("/", semesterController.AdminCreateSemester)
		semesterAdminRouter.PUT("/", semesterController.AdminUpdateSemester)
		semesterRouter.DELETE("/:id/admin", semesterController.AdminDeleteSemester)
	}
}
