package server

import (
	"net/http"
)

type server struct {
	version string
	http.Server
}

// VERSION actual server
const VERSION = "1"

// NewServer create a new api server
func NewServer(addr string, port string) *server {
	s := new(server)
	s.version = VERSION
	s.Addr = addr + ":" + port
	return s
}
