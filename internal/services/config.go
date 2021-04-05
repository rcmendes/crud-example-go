package services

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rcmendes/crud-example-go/internal/services/core/usecases"
	"github.com/rcmendes/crud-example-go/internal/services/storage/database"
	"github.com/rcmendes/crud-example-go/internal/services/web"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func InitLogger() {
	// zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)

		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}

		// output.FormatLevel = func(i interface{}) string {
		// 	return strings.ToUpper(fmt.Sprintf("| %-6s|", i))
		// }
		// output.FormatMessage = func(i interface{}) string {
		// 	return fmt.Sprintf("***%s****", i)
		// }
		output.FormatFieldName = func(i interface{}) string {
			return fmt.Sprintf("%s:", i)
		}
		output.FormatFieldValue = func(i interface{}) string {
			return strings.ToUpper(fmt.Sprintf("%s", i))
		}
		log.Logger = zerolog.New(output).With().Timestamp().Logger()
	}

}

func InitData() {
	database.CreateTables()
	database.InitDB()
}

func InitRoutes(app *fiber.App) {
	servicesStorage := database.SQLite3ServicesStorage{}
	manager := usecases.NewServiceManager(&servicesStorage)

	controller := web.NewServiceController(manager)

	group := app.Group("services")

	group.Get("/", controller.ListAll)
	group.Post("/", controller.Create)
}
