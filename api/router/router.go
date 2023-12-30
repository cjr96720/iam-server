package router

import (
	"github.com/casbin/casbin/v2"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"iam-service/api/controller"
	"iam-service/internal/repository"
	"iam-service/internal/service"
)

func Setup(db *gorm.DB, enforcer *casbin.Enforcer, gin *gin.Engine) {
	baseRouter := gin.Group("")
	// setup health check router
	healthCheckController := controller.NewHealthCheckContoller()
	baseRouter.GET("healthz", healthCheckController.HealthCheck)

	v1ApiRouter := gin.Group("/api/v1")
	iamRepository := repository.NewIAMRepository(db, enforcer)
	iamService := service.NewIAMService(iamRepository)

	// setup application router
	applicationController := controller.NewApplicationController(iamService)
	v1ApiRouter.POST("/applications", applicationController.CreateApplication)
	v1ApiRouter.GET("/applications", applicationController.GetApplications)
}
