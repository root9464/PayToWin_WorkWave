package routes

import (
	"context"
	"log"
	pb "root/proto"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
)


var mu sync.Mutex

func HelloWorld(microserviceClient pb.GetHelloClient) func(c *fiber.Ctx) error {
    return func(c *fiber.Ctx) error {
        type Message struct {
            Text string `json:"message"`
        }

        var jsonMessage Message

        if err := c.BodyParser(&jsonMessage); err != nil {
            log.Fatal("ошибка: не удалось прочитать данные из body", err)
        }

        ctx, cancel := context.WithTimeout(context.Background(), time.Second)
        defer cancel()

        mu.Lock()
        defer mu.Unlock()

        r, err := microserviceClient.HelloWorld(ctx, &pb.HelloWorldResponse{
            Message: jsonMessage.Text,
        })
        if err != nil {
            log.Fatal("ошибка: не удалось получить ответ от микросервиса", err)
        }
        return c.SendString(r.GetMessage())
    }
}