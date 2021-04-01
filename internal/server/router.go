package server

import (
	"log"

	// Swagger docs
	_ "github.com/aryannr97/iam-server/docs"
	"github.com/aryannr97/iam-server/internal/controllers"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Start initiates the IAM server
func Start(ctr controllers.IAMControllers) {
	routerEngine := gin.Default()

	routerEngine.GET("/health", healthCheck)
	routerEngine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
