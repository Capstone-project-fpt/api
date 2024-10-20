package context_util

import (
	"github.com/api/internal/types"
	"github.com/gin-gonic/gin"
)

func GetUserContext(ctx *gin.Context) *types.UserContext {
	userContext, isExist := ctx.Get("UserContext")

	if !isExist {
		return nil
	}

	return userContext.(*types.UserContext)
}
