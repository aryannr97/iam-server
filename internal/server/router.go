package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Start initiates the IAM server
func Start() {
	routerEngine := gin.Default()

	routerEngine.GET("/health", healthCheck)

	err := routerEngine.Run(":8080")
	if err != nil {
		log.Fatalf("IAM Server start up failed due to %v", err)
	}
}
