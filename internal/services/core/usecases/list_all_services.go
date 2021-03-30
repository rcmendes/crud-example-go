package usecases

import (
	"context"
	"time"

	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
)

func (manager *serviceManagerImpl) ListAllServices() (entities.ServiceList, error) {
	// return manager.data, nil

	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, 500*time.Millisecond)

	defer cancelFn()

	return manager.storage.FindAll(ctx)
}
