package controller

import "github.com/gofiber/fiber/v2"

func CreateCashier(c *fiber.Ctx) error {
	return c.SendString("CreateCashier")
}

func UpdateCashier(c *fiber.Ctx) error {
	return c.SendString("UpdateCashier")
}

func DeleteCashier(c *fiber.Ctx) error {
	return c.SendString("DeleteCashier")
}

func CashiersList(c *fiber.Ctx) error {
	return c.SendString("DeleteCashier")
}

func CashierDetails(c *fiber.Ctx) error {
	return c.SendString("DeleteCashier")
}
