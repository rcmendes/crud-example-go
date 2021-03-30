package usecases

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
)

type CreateServiceCommand struct {
	Name        string
	Description *string
}

func (manager *serviceManagerImpl) Create(command CreateServiceCommand) error {
	//TODO check if exists
	ctx := context.Background()
	ctx, cancelFn := context.WithTimeout(ctx, 500*time.Millisecond)

	defer cancelFn()

	now := time.Now().UTC()
	service, err := entities.NewService(uuid.New(), now, now, command.Name)
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
