package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
	"github.com/rcmendes/crud-example-go/internal/services/core/errors"
)

type CreateServiceCommand struct {
	Name        string
	Description *string
}

func (manager *serviceManagerImpl) Create(command CreateServiceCommand) error {
	ctx := context.Background()
	//TODO context timeout must be defined in the manager

	// ctx, cancelFn := context.WithTimeout(ctx, 500*time.Millisecond)
	// defer cancelFn()

	name := command.Name

	if err := entities.ValidateServiceName(name); err != nil {
		return errors.ConstraintError(err)
	}

	exists, err := manager.storage.ExistsByName(ctx, name)
	if err != nil {
		return errors.DatabaseError(err)
	}

	if exists {
		return errors.ServiceAlreadyExistsError(name)
	}

	now := time.Now().UTC()
	service, err := entities.NewService(uuid.New(), now, now, name)
	if err != nil {
		return err
	}

	if command.Description != nil {
		err = service.WithDescription(*command.Description)
		if err != nil {
			return err
		}
	}

	manager.storage.Insert(ctx, service)

	return nil
}
