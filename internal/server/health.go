package server

import "github.com/gin-gonic/gin"

// HealthCheck return running status of server
func healthCheck(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "UP",
	})
}
