package main

import (
	"fmt"
	"log"

	"github.com/rcmendes/crud-example-go/internal/services/core/usecases"
	"github.com/rcmendes/crud-example-go/internal/services/storage/database"
)

//TODO Create log level and init function

func main() {
	database.InitDB()
	database.CreateTables()

	servicesStorage := database.SQLite3ServicesStorage{}
	serviceManager := usecases.NewServiceManager(&servicesStorage)

	for i := 1; i <= 10; i++ {
		name := fmt.Sprintf("Name %d", i)
		description := fmt.Sprintf("Description %d", i)
		command := usecases.CreateServiceCommand{Name: name, Description: &description}

		if err := serviceManager.Create(command); err != nil {
			log.Fatal(err)
		}
	}

	list, err := serviceManager.ListAllServices()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("List of Services:")
	for i, s := range list {
		fmt.Printf("\n%d: %s", i, s.String())
	}
}
