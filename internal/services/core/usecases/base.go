package usecases

import (
	"github.com/rcmendes/crud-example-go/internal/services/ports"
)

type serviceManagerImpl struct {
	storage ports.ServiceStorage
}

func NewServiceManager(serviceStorage ports.ServiceStorage) ports.ServiceManager {
	return &serviceManagerImpl{serviceStorage}
}
