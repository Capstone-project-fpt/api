package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/global"
	"github.com/go-ecommerce-backend-api/internal/router"
)

func  InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	r.Use() // logging
	r.Use() // cross
	r.Use() // limit rate

	managerRouter := router.RouterGroupApp.Manager
	userRouter := router.RouterGroupApp.User

	MainGroup := r.Group("/api/v1")
	{
		MainGroup.GET("/health-check")
	}
	{
		userRouter.InitUserRouter(MainGroup)
		userRouter.InitProductRouter(MainGroup)
	}
	{
		managerRouter.InitAdminRouter(MainGroup)
		managerRouter.InitUserRouter(MainGroup)
	}

	return r
}