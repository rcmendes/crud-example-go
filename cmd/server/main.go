package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/rcmendes/crud-example-go/internal/services/core/usecases"
	"github.com/rcmendes/crud-example-go/internal/services/storage/database"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

//TODO Create log level and init function
func init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	debug := flag.Bool("debug", false, "sets log level to debug")

	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

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
