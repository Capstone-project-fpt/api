package controller

import (
	"github.com/api/pkg/response"
	"github.com/gin-gonic/gin"
)

type PublicController struct{}

func NewPublicController() *PublicController {
	return &PublicController{}
}

// @Summary HelloWorldController
// @Description Hello World
// @Tags Public
// @Accept json
// @Produce json
// @Router /hello-world [get]
// @Failure 400 {object} response.ResponseErr
// @Success 200 {object} response.ResponseDataSuccess
// @Security ApiKeyAuth
func (p *PublicController) HelloWorld(ctx *gin.Context) {
	response.SuccessResponse(ctx, 200, "Hello World")
}
