package types

import "github.com/gin-gonic/gin"

type ContextWithUserContextType struct {
	*gin.Context
	UserContext *UserContext
}