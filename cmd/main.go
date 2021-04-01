package main

import (
	"log"
	"os"

	"github.com/aryannr97/data-server/pkg/grpc/user"
	"github.com/aryannr97/iam-server/internal/controllers"
	"github.com/aryannr97/iam-server/internal/server"
	"github.com/aryannr97/iam-server/internal/services"
	"google.golang.org/grpc"
)

func main() {
	url, exist := os.LookupEnv("DATA_SERVICE_URL")
	if !exist {
		log.Fatalf("DATA_SERVICE_URL value not set")
	}

	conn, err := grpc.Dial(url, grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Error establishing connection with grpc server: %v", err)
	}

	defer conn.Close()

	svc := services.IAMServices{
		UserService: &services.UserService{
			UserServiceClient: user.NewUserServiceClient(conn),
		},
	}

	ctr := controllers.IAMControllers{
		UserController: &controllers.UserController{
			UserService: svc.UserService,
		},
	}
	server.Start(ctr)
}
