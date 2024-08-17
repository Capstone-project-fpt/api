package user

import "github.com/gin-gonic/gin"

type UserRouter struct {}

func (userRouter *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.POST("/otp")	 
	}

	userPrivateRouter := r.Group("/user")
	
	{
		userPrivateRouter.GET("/info")
	}
}