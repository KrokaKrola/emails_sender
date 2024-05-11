package main

import (
	"log"

	"github.com/krokakrola/emails_sender/internal/api/infra/config"
	"github.com/krokakrola/emails_sender/internal/api/ui/router"
)

func main() {
	env := config.NewEnvironment()

	err := env.Load()

	if err != nil {
		log.Fatalf("Error loading env file %v", err)
	}

	router := router.NewApiRouter()

	router.InitRoutes()
}
