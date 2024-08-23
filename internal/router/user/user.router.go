package user

import (
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController := wire.InitializeUserController()

	userRouterPublic := r.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.POST("/otp")	 
	}

	userPrivateRouter := r.Group("/user")
	
	{
		userPrivateRouter.GET("/info")
	}
}