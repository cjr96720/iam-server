package service

import (
	"github.com/google/uuid"

	"iam-service/internal/domain"
)

type IAMService struct {
	iamRepository domain.IAMRepositoryInterface
}

func NewIAMService(ir domain.IAMRepositoryInterface) domain.IAMServiceInterface {
	return &IAMService{iamRepository: ir}
}

func (is *IAMService) CreateApplication(appName string) (domain.Application, error) {
	appID := uuid.New()

	// TODO: implement JWT Token
	secret := "temporarySecret"

	app := domain.Application{
		Name:   appName,
		AppID:  appID,
		Secret: secret,
	}

	err := is.iamRepository.CreateApplication(app)
	if err != nil {
		return domain.Application{}, err
	}
	return app, nil
}

func (is *IAMService) GetApplications() ([]domain.Application, error) {
	return is.iamRepository.GetApplications()
}
