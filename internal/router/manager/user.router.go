package manager

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (userRouter *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userPrivateRouter := r.Group("/admin/user")	
	{
		userPrivateRouter.POST("/active-user")
	}
}