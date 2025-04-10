package main

import (
	"log"

	"github.com/Alwin18/go-unit-test/internal/handlers"
	"github.com/Alwin18/go-unit-test/internal/routes"
)

func main() {
	helloWorldHandler := handlers.NewHelloWorld()
	routes := routes.NewRouteConfig(helloWorldHandler)
	r := routes.SetRoutes()
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
