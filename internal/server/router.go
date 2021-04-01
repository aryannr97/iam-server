package server

import (
	"log"

	"github.com/aryannr97/iam-server/internal/controllers"
	"github.com/gin-gonic/gin"
)

// Start initiates the IAM server
func Start(ctr controllers.IAMControllers) {
	routerEngine := gin.Default()

	routerEngine.GET("/health", healthCheck)

	userGroup := routerEngine.Group("/iam/api/1.0/users")

	userGroup.POST("/add", ctr.UserController.AddUser)
	userGroup.GET("/:id", ctr.UserController.GetUser)
	userGroup.DELETE("/:id", ctr.UserController.DeleteUser)
	userGroup.PUT("/:id", ctr.UserController.UpdateUser)

	err := routerEngine.Run(":8080")
	if err != nil {
		log.Fatalf("IAM Server start up failed due to %v", err)
	}
}
