package main

import (
	"github.com/krokakrola/emails_sender/internal/api/router"
)

func main() {
	router := router.NewApiRouter()

	router.InitRoutes()
}
