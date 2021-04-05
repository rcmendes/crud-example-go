package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
	"github.com/rcmendes/crud-example-go/internal/services/core/errors"
	"github.com/rcmendes/crud-example-go/internal/services/ports"
)

type createServiceResponse struct {
	id string
}

func (resp *createServiceResponse) ID() string {
	return resp.id
}

func (manager *serviceManagerImpl) Create(command ports.CreateServiceRequest) (ports.CreateServiceResponse, error) {
	ctx := context.Background()
	//TODO context timeout must be defined in the manager

	// ctx, cancelFn := context.WithTimeout(ctx, 500*time.Millisecond)
	// defer cancelFn()

	name := command.Name()

	if err := entities.ValidateServiceName(name); err != nil {
		return nil, errors.ConstraintError(err)
	}

	exists, err := manager.storage.ExistsByName(ctx, name)
	if err != nil {
		return nil, errors.DatabaseError(err)
	}

	if exists {
		return nil, errors.ServiceAlreadyExistsError(name)
	}

	now := time.Now().UTC()
	id := uuid.New()
	service, err := entities.NewService(id, now, now, name)
	if err != nil {
		return nil, err
	}

	if command.Description() != nil {
		err = service.SetDescription(*command.Description())
		if err != nil {
			return nil, err
		}
	}

	manager.storage.Insert(ctx, service)

	result := createServiceResponse{
		id: id.String(),
	}

	return &result, nil
}
