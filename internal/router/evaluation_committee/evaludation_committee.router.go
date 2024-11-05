package evaluation_committee

import (
	"github.com/api/internal/middleware"
	"github.com/api/internal/wire"
	"github.com/gin-gonic/gin"
)

type EvaluationCommitteeRouter struct {}

func (e *EvaluationCommitteeRouter) InitEvaluationCommitteeRouter(group *gin.RouterGroup) {
	evaluationCommitteeController := wire.InitializeEvaluationCommitteeController()

	evaluationCommitteeRouter := group.Group("/evaluation-committees")
	evaluationCommitteeRouter.Use(middleware.AuthMiddleware())
	{
		evaluationCommitteeRouter.GET("/:id", evaluationCommitteeController.GetEvaluationCommittee)
	}
}