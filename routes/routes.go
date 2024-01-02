package routes

import (
	"github.com/eliofery/golang-fiber-restapi/controller"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(server *fiber.App) {
	server.Post("/cashiers/login", controller.Login)
	server.Get("/cashiers/:id/logout", controller.Logout)
	server.Get("/cashiers/:id/password", controller.Password)

	server.Post("/cashiers", controller.CreateCashier)
	server.Get("/cashiers", controller.CashiersList)
	server.Get("/cashiers/:id", controller.CashierDetails)
	server.Delete("/cashiers/:id", controller.DeleteCashier)
	server.Put("/cashiers/:id", controller.UpdateCashier)
}
