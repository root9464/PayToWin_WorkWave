package core

import (
	"log"
	"sync"

	pb "root/proto"
	"root/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)


var (
	once           		sync.Once
	microserviceClient 	pb.GetHelloClient
)

func initGRPCClient() {
	conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("не получилось соединиться: %v", err)
	}
	microserviceClient = pb.NewGetHelloClient(conn)
}

func AllRoutes(app *fiber.App, microserviceClient pb.GetHelloClient) {
    hello := app.Group("/hello")
	hello.Post("/", routes.HelloWorld(microserviceClient))

}


func Core() {
	app := fiber.New()
	app.Use(logger.New())

	once.Do(initGRPCClient)

	AllRoutes(app, microserviceClient)

	log.Fatal(app.Listen(":6069"))
}