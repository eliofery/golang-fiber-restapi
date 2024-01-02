package controller

import (
	"github.com/eliofery/golang-fiber-restapi/database"
	"github.com/eliofery/golang-fiber-restapi/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"strconv"
	"time"
)

func Login(ctx *fiber.Ctx) error {
	var cashier model.Cashier

	err := ctx.BodyParser(&cashier)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Не удалось войти",
			"error":   err.Error(),
		})
	}

	err = validate.Struct(cashier)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Не удалось войти",
			"error":   err.Error(),
		})
	}

	database.DB.Where("name = ? AND password = ?", cashier.Name, cashier.Password).First(&cashier)
	if cashier.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Не удалось войти",
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"issuer": strconv.Itoa(int(cashier.ID)),
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return ctx.Status(http.StatusUnauthorized).JSON(fiber.Map{
			"message": "Токен не действителен",
			"error":   err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Вход выполнен",
		"cashier": cashier,
		"token":   jwtToken,
	})
}

func Logout(ctx *fiber.Ctx) error {
	return ctx.SendString("Logout")
}

func Password(ctx *fiber.Ctx) error {
	return ctx.SendString("Password")
}
