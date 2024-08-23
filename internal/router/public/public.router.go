package public

import (
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type PublicRouter struct{}

func (pr *PublicRouter) InitPublicRouter(r *gin.RouterGroup) {
	authController := wire.InitializeAuthController()

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
}