package domain

import "github.com/google/uuid"

type Application struct {
	Name   string    `json:"name" binding:"required,min=3"`
	AppID  uuid.UUID `json:"application_id"`
	Secret string    `json:"secret"`
}
