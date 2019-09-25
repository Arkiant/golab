package main

import (
	"log"

	"github.com/arkiant/golab/api_production_grade/api/handler"

	"github.com/arkiant/golab/api_production_grade/api/service"

	"github.com/go-chi/chi"

	"github.com/arkiant/golab/api_production_grade/server"
)

func main() {

	const (
		address = "0.0.0.0"
		port    = "5000"
		version = "1"
	)

	s := service.NewService(nil)
	routes := handler.NewHandler(s)

	// Function log all routes in handler
	if err := chi.Walk(routes, server.WalkFunc); err != nil {
		log.Fatalf("Walk error: %v\n", err)
	}

	// Create a new server
	srv, err := server.NewServer(address, port, routes)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	// Listen server
	log.Printf("Server listen on %s:%s", address, port)
	log.Fatal(srv.ListenAndServe())
}
