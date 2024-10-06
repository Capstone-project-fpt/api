package initialize

import (
	"github.com/api/global"
	"github.com/api/internal/middleware"
	"github.com/api/internal/router"
	"github.com/api/internal/service"

	"time"

	swaggerDocs "github.com/api/docs"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/hellofresh/health-go/v5"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine

	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	healthCheck := service.NewHealthCheckService()

	swaggerDocs.SwaggerInfo.BasePath = "/api/v1"
	r.GET("/health-check", healthCheckHandle(healthCheck.HealthCheck()))

	// Configure CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{global.Config.Server.WebURL}, // TODO: Use config variable
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	r.Use(middleware.I18nMiddleware())
	r.Use(middleware.ErrorHandleMiddleware())

	publicRouter := router.RouterGroupApp.Public
	userRouter := router.RouterGroupApp.User
	adminRouter := router.RouterGroupApp.Admin
	topicReferenceRouter := router.RouterGroupApp.TopicReference

	MainGroup := r.Group("/api/v1")
	{
		publicRouter.InitPublicRouter(MainGroup)
	}
	{
		userRouter.InitUserRouter(MainGroup)
	}
	{
		adminRouter.InitAdminRouter(MainGroup)
	}
	{
		topicReferenceRouter.InitTopicReferenceRouter(MainGroup)
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

func healthCheckHandle(h *health.Health) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		h.Handler().ServeHTTP(ctx.Writer, ctx.Request)
	}
}
