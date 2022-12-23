package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server encapsulates the logic for initializing the REST API server.
type Server struct {
	handler *Handler
}

// NewServer creates a new instance of Server.
func NewServer(handler *Handler) *Server {
	return &Server{
		handler: handler,
	}
}

// Run sets up the endpoints and starts the server.
// This is a blocking call.
func (s *Server) Run(listenPort int) error {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/accounts", s.handler.CreateAccount)
	r.GET("/accounts/:id", s.handler.GetAccount)
	r.POST("/transactions", s.handler.CreateTransaction)
	return r.Run(fmt.Sprintf(":%d", listenPort))
}
