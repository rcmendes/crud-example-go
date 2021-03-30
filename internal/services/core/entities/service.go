package entities

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/crud-example-go/internal/services/core/exceptions"
)

type Service interface {
	Entity
	Name() string
	WithDescription(description string) error
	Description() *string
	String() string
}

type service struct {
	entity
	name        string
	description *string
}

type ServiceList = []Service

func (svc *service) Name() string {
	return svc.name
}

func (svc *service) Description() *string {
	return svc.description
}

func (svc *service) String() string {
	// return fmt.Sprintf("Service<name='%s', description='%s', id=%v, created_at=%v, updated_at=%v>",
	return fmt.Sprintf("Service<id='%v', name='%s', description='%s', created_at=%s, updated_at=%s>",
		svc.id, svc.name, *svc.description, svc.createdAt.UTC().Format("2006-01-02T15:04:05-0700"),
		svc.updatedAt.UTC().Format("2006-01-02T15:04:05-0700"))
}

func (svc *service) WithDescription(description string) error {
	if len(description) > 1024 {
		return exceptions.NewInvalidFieldMaxLengthError("description", 1024)
	}

	svc.description = &description

	return nil
}

func NewService(id uuid.UUID, createdAt time.Time, updatedAt time.Time, name string) (Service, error) {
	if len(name) < 1 || len(name) > 20 {
		return nil, exceptions.NewInvalidFieldRangeLengthError("name", 1, 20)
	}

	service := &service{
		entity: entity{
			id:        id,
			createdAt: createdAt,
			updatedAt: updatedAt,
		},
		name: name,
	}

	return service, nil
}
