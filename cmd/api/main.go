package main

import "github.com/krokakrola/emails_sender/internal/api"

func main() {
	router := api.NewApiRouter()

	router.InitRoutes()
}
