package api

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	router *gin.Engine
}

func NewServer() *Server {
	server := &Server{}
	router := gin.Default()

	router.GET("/dollar", server.getDollar)
	router.GET("/health", server.checkHealth)
	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	go server.periodicHealthCheck()
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
