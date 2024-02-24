package main

import (
	"context"
	"log"
	"time"

	pb "root/lib/proto/out"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)




func handlePostRequest(client pb.GetHelloClient) func(c *fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        type Message struct {
            Text string `json:"message"`
        }

        var jsonMessage Message

        if err := c.BodyParser(&jsonMessage); err != nil {
            return err
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()
        r, err := client.HelloWorld(ctx, &pb.HelloWorldResponse{
            Message: jsonMessage.Text,
        })
        if err != nil {
            return err
        }
        return c.SendString(r.GetMessage())
    }
}

func Routes(app *fiber.App, microserviceClient pb.GetHelloClient) {
	app.Post("/", handlePostRequest(microserviceClient))
}

func main() {
    app := fiber.New()
    app.Use(logger.New())

    conn, err := grpc.Dial("localhost:3000", grpc.WithTransportCredentials(insecure.NewCredentials()))
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    microserviceClient := pb.NewGetHelloClient(conn)
	Routes(app, microserviceClient)
    

    log.Fatal(app.Listen(":3001"))
}
