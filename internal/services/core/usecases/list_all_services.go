package usecases

import (
	"context"
	"time"

	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
	"github.com/rs/zerolog/log"
)

func (manager *serviceManagerImpl) ListAllServices() (entities.ServiceList, error) {
	defer log.Debug().Msg("Fetched all the services from the storage.")
	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, 500*time.Millisecond)

	defer cancelFn()

	return manager.storage.FindAll(ctx)
}
