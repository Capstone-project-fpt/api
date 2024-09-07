package user

import (
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (ur *UserRouter) InitUserRouter(r *gin.RouterGroup) {
	userController := wire.InitializeUserController()

	userRouter := r.Group("/users")
	userRouter.Use(middleware.AuthMiddleware())
	{
		userRouter.GET("/", userController.GetUser)
		userRouter.GET("/me", userController.GetMe)
	}
}
