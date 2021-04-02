package usecases

import (
	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
	"github.com/rcmendes/crud-example-go/internal/services/ports"
)

type ServiceManager interface {
	ListAllServices() (entities.ServiceList, error)
	Create(service CreateServiceCommand) error
}

type serviceManagerImpl struct {
	storage ports.ServiceStorage
}

func NewServiceManager(serviceStorage ports.ServiceStorage) ServiceManager {
	return &serviceManagerImpl{serviceStorage}
}
