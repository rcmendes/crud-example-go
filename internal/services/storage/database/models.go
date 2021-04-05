package database

import (
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
)

type BaseModel struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type ServiceModel struct {
	BaseModel
	Name        string
	Description *string
}

func NewFromService(service entities.Service) *ServiceModel {
	return &ServiceModel{
		BaseModel: BaseModel{
			ID:        service.ID().String(),
			CreatedAt: service.CreatedAt(),
			UpdatedAt: service.UpdatedAt(),
		},
		Name:        service.Name(),
		Description: service.Description(),
	}
}

func (m *ServiceModel) ToEntity() (entities.Service, error) {
	service, err := entities.NewService(
		uuid.MustParse(m.ID),
		m.CreatedAt,
		m.UpdatedAt,
		m.Name)

	if err != nil {
		return nil, err
	}

	if m.Description != nil {
		service.SetDescription(*m.Description)
	}

	return service, nil
}
