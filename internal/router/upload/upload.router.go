package upload

import (
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type UploadRouter struct{}

func (u *UploadRouter) InitUploadRouter(group *gin.RouterGroup) {
	uploadController := wire.InitializeUploadController()

	uploadRouter := group.Group("/uploads")
	uploadRouter.Use(middleware.AuthMiddleware())
	{
		uploadRouter.POST("/presign-url", uploadController.GenerateUploadPresignUrl)
	}
}