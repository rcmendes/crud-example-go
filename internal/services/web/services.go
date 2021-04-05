package web

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/crud-example-go/internal/services/ports"
	"github.com/rs/zerolog/log"
)

type ServiceController struct {
	manager ports.ServiceManager
}

type CreateServiceRequest struct {
	NameKey        string  `json:"name"`
	DescriptionKey *string `json:"description"`
}

type ServiceDTO struct {
	ID          string    `json:"id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	Name        string    `json:"name"`
	Description *string   `json:"description"`
}

func (req *CreateServiceRequest) Name() string {
	return req.NameKey
}

func (req *CreateServiceRequest) Description() *string {
	return req.DescriptionKey
}

func NewServiceController(manager ports.ServiceManager) *ServiceController {
	return &ServiceController{
		manager,
	}
}

func (controller ServiceController) Create(ctx *fiber.Ctx) error {
	var request CreateServiceRequest
	if err := ctx.BodyParser(&request); err != nil {
		log.Err(err).Send()
		return fiber.ErrBadRequest
	}

	data, err := controller.manager.Create(&request)
	if err != nil {
		log.Err(err).Send()

		//TODO handler business error
		return fiber.ErrInternalServerError
	}

	payload := struct {
		ID string `json:"id"`
	}{
		ID: data.ID(),
	}

	return ctx.Status(fiber.StatusCreated).JSON(payload)
}

func (controller ServiceController) ListAll(ctx *fiber.Ctx) error {
	data, err := controller.manager.ListAllServices()

	if err != nil {
		log.Err(err).Send()

		//TODO handler business error
		return fiber.ErrInternalServerError
	}

	list := make([]*ServiceDTO, 0, len(data))

	for _, svc := range data {
		parsed := &ServiceDTO{
			ID:          svc.ID().String(),
			CreatedAt:   svc.CreatedAt(),
			UpdatedAt:   svc.UpdatedAt(),
			Name:        svc.Name(),
			Description: svc.Description(),
		}

		list = append(list, parsed)
	}

	return ctx.JSON(list)
}
