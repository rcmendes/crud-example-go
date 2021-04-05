package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/rcmendes/crud-example-go/internal/services"
	"github.com/rs/zerolog/log"
)

func main() {
	app := fiber.New(fiber.Config{
		Prefork: false,
	})

	app.Use(cors.New())

	services.InitData()
	services.InitLogger()
	services.InitRoutes(app)

	log.Info().Msg("Running CLI app.")
	app.Listen(":8080")

}
