package main

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
    app.Post("/", routes.HelloWorld(microserviceClient))

}


func main() {
	app := fiber.New()
	app.Use(logger.New())

	once.Do(initGRPCClient)

	AllRoutes(app, microserviceClient)

	log.Fatal(app.Listen(":3001"))
}

// server1 := grpc.NewServer()
// server2 := grpc.NewServer()

// // Регистрация сервисов на сервере
// pb.RegisterService1Server(server1, &service1Impl{})
// pb.RegisterService2Server(server2, &service2Impl{})

// // Создание слушателей на разных портах
// lis1, err := net.Listen("tcp", ":3000")
// if err != nil {
//     log.Fatalf("failed to listen: %v", err)
// }
// lis2, err := net.Listen("tcp", ":3001")
// if err != nil {
//     log.Fatalf("failed to listen: %v", err)
// }

// // Запуск серверов
// go func() {
//     if err := server1.Serve(lis1); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }()
// go func() {
//     if err := server2.Serve(lis2); err != nil {
//         log.Fatalf("failed to serve: %v", err)
//     }
// }()