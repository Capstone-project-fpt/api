package admin

import (
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type AdminRouter struct {}

func (ar *AdminRouter) InitAdminRouter(r *gin.RouterGroup) {
	adminController := wire.InitializeAdminController()

	adminRouter := r.Group("/admin")

	adminStudentRouter := adminRouter.Group("/students")
	{
		adminStudentRouter.POST("/create-account", adminController.CreateStudentAccount)
	}
}