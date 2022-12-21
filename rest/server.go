package rest

import (
	"fmt"
	"log"
	"net/http"
)

// Server encapsulates the logic for initializing the REST server.
type Server struct {
	handlers []handler
}

type handler struct {
	path string
	fn   http.HandlerFunc
}

// NewServer creates a new instance of Server.
func NewServer() *Server {
	return &Server{}
}

// AddHandlerFunc adds a new endpoint to the Server.
// Call this function one or more times to add endpoints to the server
// before starting the server.
func (s *Server) AddHandlerFunc(path string, fn http.HandlerFunc) {
	s.handlers = append(s.handlers, handler{
		path: path,
		fn:   fn,
	})
}

func (s *Server) ListenAndServe(port int) {
	mux := http.NewServeMux()
	for _, h := range s.handlers {
		mux.Handle(h.path, h.fn)
	}
	listenAddr := fmt.Sprintf(":%d", port)
	log.Printf("Server listening on %s", listenAddr)
	http.ListenAndServe(listenAddr, mux)
}
