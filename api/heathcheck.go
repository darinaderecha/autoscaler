package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Server) checkHealth(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Server is ok"})
}
func (server *Server) periodicHealthCheck() {
	for {
		resp, err := http.Get("http://localhost:8080/health")
		if err != nil {
			fmt.Println("Server is down:", err)
		} else if resp.StatusCode != http.StatusOK {
			fmt.Println("Server is not ok, status code:", resp.StatusCode)
		} else {
			fmt.Println("Server is up and running")
		}
		time.Sleep(2 * time.Minute)
	}
}
