package rest

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// Server encapsulates the logic for initializing the REST server.
type Server struct {
	handler *Handler
}

// NewServer creates a new instance of Server.
func NewServer(handler *Handler) *Server {
	return &Server{
		handler: handler,
	}
}

func (s *Server) Run(listenPort int) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/accounts", s.handler.CreateAccount)
	r.GET("/accounts/:id", s.handler.GetAccount)
	r.Run(fmt.Sprintf(":%d", listenPort))
}
