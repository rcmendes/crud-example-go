package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/rcmendes/crud-example-go/internal/services/core/usecases"
	"github.com/rcmendes/crud-example-go/internal/services/storage/database"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func initLogger() {
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

//TODO Create log level and init function
func init() {
	initLogger()

	database.InitDB()
	database.CreateTables()
}

func main() {
	log.Info().Msg("Running CLI app.")

	servicesStorage := database.SQLite3ServicesStorage{}
	serviceManager := usecases.NewServiceManager(&servicesStorage)

	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Name %d", i)
		description := fmt.Sprintf("Description %d", i)
		command := usecases.CreateServiceCommand{Name: name, Description: &description}

		if err := serviceManager.Create(command); err != nil {
			log.Err(err).Send()
		}
	}

	list, err := serviceManager.ListAllServices()

	if err != nil {
		log.Fatal().Err(err)
	}

	fmt.Println("List of Services:")
	for i, s := range list {
		fmt.Printf("%d: %s\n", i, s.String())
	}

	for i := 1; i <= 2; i++ {
		name := fmt.Sprintf("Name %d", i)
		description := fmt.Sprintf("Description %d", i)
		command := usecases.CreateServiceCommand{Name: name, Description: &description}

		if err := serviceManager.Create(command); err != nil {
			log.Err(err).Send()
		}
	}
}
