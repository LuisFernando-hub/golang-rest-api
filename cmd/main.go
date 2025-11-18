package main

import (
	"ecom/cmd/api"
	"log"
)

func main() {
	// Application entry point
	server := api.NewAPIServer(":8080", nil)

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}