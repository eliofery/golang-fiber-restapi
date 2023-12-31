package controller

import "github.com/gofiber/fiber/v2"

func Login(c *fiber.Ctx) error {
	return c.SendString("Login")
}

func Logout(c *fiber.Ctx) error {
	return c.SendString("Logout")
}

func Password(c *fiber.Ctx) error {
	return c.SendString("Password")
}
