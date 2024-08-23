package router

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(middleware.AuthMiddleware())

	// v1 := r.Group("/api/v1")
	// {
	// 	v1.GET("/user", userController.Register)
	// }

	return r
}
