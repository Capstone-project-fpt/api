package router

import (
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	userController, _ := wire.InitializeUserController()

	r.Use(middleware.AuthMiddleware())

	v1 := r.Group("/api/v1")
	{
		v1.GET("/user", userController.Register)
	}

	return r
}
