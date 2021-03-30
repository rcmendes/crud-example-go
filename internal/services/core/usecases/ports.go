package usecases

import (
	"context"

	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
)

type ServiceStorage interface {
	Insert(ctx context.Context, service entities.Service) error
	FindAll(ctx context.Context) (entities.ServiceList, error)
}
