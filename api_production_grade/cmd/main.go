package main

import (
	"log"

	"github.com/arkiant/golab/api_production_grade/server"
)

func main() {
	const (
		address = "localhost"
		port    = "5000"
	)
	log.Printf("Server listen on %s:%s", address, port)
	srv := server.NewServer(address, port)
	log.Fatal(srv.ListenAndServe())
}
