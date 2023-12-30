package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"iam-service/internal/response"
)

type HealthCheckController struct{}

func NewHealthCheckContoller() *HealthCheckController {
	return &HealthCheckController{}
}

// @Summary		Health Check
// @Success		200	{object}	response.HealthCheckResponse{}
// @Router		/healthz [get]
func (_ *HealthCheckController) HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, response.HealthCheckResponse{Message: response.OK})
}
