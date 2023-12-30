package main

import (
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
)

func main() {
	server := fiber.New()

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	err := server.Listen(":3000")
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
