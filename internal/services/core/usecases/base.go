package usecases

import (
	"github.com/rcmendes/crud-example-go/internal/services/core/entities"
)

type ServiceManager interface {
	ListAllServices() (entities.ServiceList, error)
	Create(service CreateServiceCommand) error
}

type serviceManagerImpl struct {
	storage ServiceStorage
}

func NewServiceManager(serviceStorage ServiceStorage) ServiceManager {
	return &serviceManagerImpl{serviceStorage}
}
