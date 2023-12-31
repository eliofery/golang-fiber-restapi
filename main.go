package main

import (
	"github.com/eliofery/golang-fiber-restapi/database"
	"github.com/eliofery/golang-fiber-restapi/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Не удалось загрузить переменные среды: ", err)
	}

	if err = database.Connect(); err != nil {
		log.Fatal("Не удалось подключиться к базе данных: ", err)
	}

	if err = database.Migrations(); err != nil {
		log.Fatal("Не удалось мигрировать базу данных: ", err)
	}

	server := fiber.New()
	routes.RegisterRoutes(server)

	err = server.Listen(":3000")
	if err != nil {
		log.Fatal("Ошибка запуска сервера: ", err)
	}
}
