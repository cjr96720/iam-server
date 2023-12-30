package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"iam-service/internal/domain"
	"iam-service/internal/request"
	"iam-service/internal/response"
)

type ApplicationController struct {
	iamService domain.IAMServiceInterface
}

func NewApplicationController(is domain.IAMServiceInterface) *ApplicationController {
	return &ApplicationController{iamService: is}
}

// @Summary	Create Application
// @Accept 	json
// @Produce	json
// @Param	application	body	request.CreateApplicationRequest	true	"Application information"
// @Success	200	{object}	response.CreateApplicationResponse
// @Failure	404	{object}	response.ErrorResponse
// @Failure	500	{object}	response.ErrorResponse
// @Router	/api/v1/applications [post]
func (ac *ApplicationController) CreateApplication(c *gin.Context) {
	request := request.CreateApplicationRequest{}
	err := c.ShouldBindJSON(&request)
	if err != nil {
		response := response.ErrorResponse{
			Title: "Invalid Request",
			Error: err,
		}

		c.JSON(http.StatusBadRequest, response)
	}

	app, err := ac.iamService.CreateApplication(request.Name)
	if err != nil {
		response := response.ErrorResponse{
			Title: "Create Application Error",
			Error: err,
		}

		c.JSON(http.StatusInternalServerError, response)
	}

	response := response.CreateApplicationResponse{
		DefaultResponse: response.DefaultResponse{
			Message: fmt.Sprintf("Application `%s` Created", request.Name),
			Status:  response.SUCCESS,
		},
		AppId: app.AppID.String(),
		Token: app.Secret,
	}

	c.JSON(http.StatusOK, response)
}

func (ac *ApplicationController) GetApplications(c *gin.Context) {
	apps, err := ac.iamService.GetApplications()
	if err != nil {
		response := response.ErrorResponse{
			Title: "Get Applications Error",
			Error: err,
		}

		c.JSON(http.StatusInternalServerError, response)
	}

	response := response.GetApplicationsResponse{
		DefaultResponse: response.DefaultResponse{
			Message: "Get Applications",
			Status:  response.SUCCESS,
		},
		Applications: apps,
	}

	c.JSON(http.StatusOK, response)
}
