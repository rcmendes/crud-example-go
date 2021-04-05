package ports

import (
	"context"

	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
)

type ServiceStorage interface {
	Insert(ctx context.Context, service entities.Service) error
	FindAll(ctx context.Context) (entities.ServiceList, error)
	ExistsByName(ctx context.Context, name string) (bool, error)
}

type CreateServiceRequest interface {
	Name() string
	Description() *string
}

type CreateServiceResponse interface {
	ID() string
}

type ServiceManager interface {
	ListAllServices() (entities.ServiceList, error)
	Create(service CreateServiceRequest) (CreateServiceResponse, error)
}
