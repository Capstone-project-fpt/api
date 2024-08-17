package router

import (
	"github.com/gin-gonic/gin"
	"github.com/go-ecommerce-backend-api/internal/middleware"
	"github.com/go-ecommerce-backend-api/internal/wire"
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
