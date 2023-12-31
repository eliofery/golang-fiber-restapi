package routes

import (
	"github.com/eliofery/golang-fiber-restapi/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(server *fiber.App) {
	server.Get("/cashiers/:id/login", controller.Login)
	server.Get("/cashiers/:id/logout", controller.Logout)
	server.Get("/cashiers/:id/password", controller.Password)
}
