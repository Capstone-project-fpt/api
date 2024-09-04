package public

import (
	"github.com/api/global"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct{}

func (pr *PublicRouter) InitPublicRouter(r *gin.RouterGroup) {
	authController := wire.InitializeAuthController()
	majorController := wire.InitializeMajorController()

	if global.Config.Server.Mode == "dev" {
		r.POST("/register", authController.Register)
	}
	r.POST("/login", authController.Login)

	majorGroup := r.Group("/majors")
	{
		majorGroup.GET("/", majorController.GetListMajor)
	}
}
