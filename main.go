package main

import (
	"github.com/eliofery/golang-fiber-restapi/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить переменные среды: ", err)
	}

	if err = database.Connect(); err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	server := fiber.New()

	server.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(fiber.Map{
			"message": "Hello, World!",
		})
	})

	err = server.Listen(":3000")
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
