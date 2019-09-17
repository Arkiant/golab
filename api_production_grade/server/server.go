package server

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type server struct {
	http.Server
}

// NewServer create a new api server
func NewServer(addr string, port string, routes *chi.Mux) (*server, error) {
	s := new(server)
	s.Addr = addr + ":" + port
	s.Handler = routes

	return s, nil
}

// WalkFunc is a function to log all routes in the handler
func WalkFunc(method string, route string, handler http.Handler, middlewares ...func(http.Handler) http.Handler) error {
	log.Printf("%s %s\n", method, route)
	return nil
}
