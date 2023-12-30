package response

import "iam-service/internal/domain"

type GeneralResponseStatus string

const (
	SUCCESS GeneralResponseStatus = "SUCCESS"
	FAILED                        = "FAILED"
	OK                            = "OK"
)

type HealthCheckResponse struct {
	Message GeneralResponseStatus `json:"message"`
}

type ErrorResponse struct {
	Title string `json:"title"`
	Error error  `json:"error"`
}

type DefaultResponse struct {
	Message string                `json:"message"`
	Status  GeneralResponseStatus `json:"status"`
}

type CreateApplicationResponse struct {
	DefaultResponse
	AppId string `json:"application_id"`
	Token string `json:"token"`
}

type GetApplicationsResponse struct {
	DefaultResponse
	Applications []domain.Application `json:"applications"`
}
