package main

// import (
// 	"fmt"

// 	"github.com/rcmendes/crud-example-go/internal/services"
// 	"github.com/rcmendes/crud-example-go/internal/services/core/usecases"
// 	"github.com/rcmendes/crud-example-go/internal/services/storage/database"
// 	"github.com/rs/zerolog/log"
// )

// func main() {
// 	services.InitData()
// 	services.InitLogger()
// 	log.Info().Msg("Running CLI app.")

// 	servicesStorage := database.SQLite3ServicesStorage{}
// 	serviceManager := usecases.NewServiceManager(&servicesStorage)

// 	for i := 1; i <= 10; i++ {
// 		name := fmt.Sprintf("Name %d", i)
// 		description := fmt.Sprintf("Description %d", i)
// 		command := usecases.CreateServiceCommand{Name: name, Description: &description}

// 		if err := serviceManager.Create(command); err != nil {
// 			log.Err(err).Send()
// 		}
// 	}

// 	list, err := serviceManager.ListAllServices()

// 	if err != nil {
// 		log.Fatal().Err(err)
// 	}

// 	fmt.Println("List of Services:")
// 	for i, s := range list {
// 		fmt.Printf("%d: %s\n", i, s.String())
// 	}

// 	for i := 1; i <= 2; i++ {
// 		name := fmt.Sprintf("Name %d", i)
// 		description := fmt.Sprintf("Description %d", i)
// 		command := usecases.CreateServiceCommand{Name: name, Description: &description}

// 		if err := serviceManager.Create(command); err != nil {
// 			log.Err(err).Send()
// 		}
// 	}
// }
