package public

import (
	"github.com/api/global"
	"github.com/api/internal/controller"
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct{}

func (pr *PublicRouter) InitPublicRouter(r *gin.RouterGroup) {
	authController := wire.InitializeAuthController()
	majorController := wire.InitializeMajorController()
	subMajorController := wire.InitializeSubMajorController()
	publicController := controller.NewPublicController()

	if global.Config.Server.Mode == "dev" {
		r.POST("/register", authController.Register)
	}
	r.POST("/login", authController.Login)
	r.POST("/forgot-password", authController.ForgotPassword)
	r.POST("/reset-password", authController.ResetPassword)
	r.GET(
		"/hello-world",
		middleware.AuthMiddleware(),
		publicController.HelloWorld,
	)

	majorGroup := r.Group("/majors")
	{
		majorGroup.GET("/", majorController.GetListMajor)
		majorGroup.GET("/:id", majorController.GetMajor)
	}

	subMajorGroup := r.Group("/sub-majors")
	{
		subMajorGroup.GET("/", subMajorController.GetListSubMajor)
		subMajorGroup.GET("/:id", subMajorController.GetSubMajor)
	}
}
