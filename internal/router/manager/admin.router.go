package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct {}

func (adminRouter *AdminRouter) InitAdminRouter(r *gin.RouterGroup) {
	adminPublicRouter := r.Group("/admin")
	{
		adminPublicRouter.POST("/login")
	}

	adminPrivateRouter := r.Group("/admin/user")
	{
		adminPrivateRouter.POST("/active-user")
	}
}