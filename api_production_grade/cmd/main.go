package main

import (
	"github.com/go-chi/chi"
	"log"

	"github.com/arkiant/golab/api_production_grade/server"
	"github.com/arkiant/golab/api_production_grade/routes"
)

func main() {
	
	const (
		address = "localhost"
		port    = "5000"
		version = "1"
	)

	// Routes
	r := chi.NewRouter()
	r.Route("/api/v"+version, func(r chi.Router){
		r.Mount("/status", routes.Health() )
	})

	// Function log all routes in handler
	if err := chi.Walk(r, server.WalkFunc); err != nil {
		log.Fatalf("Walk error: %v\n", err)
	}

	// Create a new server
	srv, err := server.NewServer(address, port, r)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	// Listen server
	log.Printf("Server listen on %s:%s", address, port)
	log.Fatal(srv.ListenAndServe())
}
