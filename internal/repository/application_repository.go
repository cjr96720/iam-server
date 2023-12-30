package repository

import (
	"github.com/jinzhu/copier"

	"iam-service/internal/domain"
	"iam-service/internal/model"
)

func (ir *IAMRepository) CreateApplication(app domain.Application) error {
	appModel := model.Application{}
	copier.Copy(&appModel, &app)

	result := ir.db.Create(&appModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (ir *IAMRepository) GetApplications() ([]domain.Application, error) {
	appsModel := []model.Application{}
	result := ir.db.Find(&appsModel)
	if result.Error != nil {
		return nil, result.Error
	}

	apps := []domain.Application{}
	copier.Copy(&apps, &appsModel)

	return apps, nil
}
