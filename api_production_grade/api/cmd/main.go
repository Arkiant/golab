package main

import (
	"log"
	"net/http"

	"github.com/arkiant/golab/api_production_grade/api/handler"

	"github.com/arkiant/golab/api_production_grade/api/service"

	"github.com/go-chi/chi"
)

func main() {

	const (
		address = "0.0.0.0"
		port    = "5000"
		version = "1"
	)

	s := service.NewService(nil, version)
	routes := handler.NewHandler(s)

	srv := http.Server{
		Addr:    address + ":" + port,
		Handler: routes,
	}

	// Function log all routes in handler
	if err := chi.Walk(routes, WalkFunc); err != nil {
		log.Fatalf("Walk error: %v\n", err)
	}

	// Listen server
	log.Printf("Server listen on %s:%s", address, port)
	log.Fatal(srv.ListenAndServe())
}

// WalkFunc is a function to log all routes in the handler
func WalkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	log.Printf("%s %s\n", method, route)
	return nil
}
