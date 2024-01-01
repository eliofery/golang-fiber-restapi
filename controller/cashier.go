package controller

import (
	"github.com/eliofery/golang-fiber-restapi/database"
	"github.com/eliofery/golang-fiber-restapi/model"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
	"strconv"
	"time"
)

var (
	validate = validator.New()
)

func CreateCashier(ctx *fiber.Ctx) error {
	var cashier model.Cashier

	if err := ctx.BodyParser(&cashier); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Не удалось создать кассу",
			"error":   err.Error(),
		})
	}

	if err := validate.Struct(cashier); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Не удалось создать кассу",
			"error":   err.Error(),
		})
	}

	cashier.CreatedAt = time.Now()
	cashier.UpdatedAt = cashier.CreatedAt

	database.DB.Create(&cashier)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Касса создана",
		"cashier": cashier,
	})
}

func UpdateCashier(ctx *fiber.Ctx) error {
	return ctx.SendString("UpdateCashier")
}

func DeleteCashier(ctx *fiber.Ctx) error {
	return ctx.SendString("DeleteCashier")
}

func CashiersList(ctx *fiber.Ctx) error {
	var (
		cashiers []model.Cashier
		count    int64
	)

	limit, err := strconv.Atoi(ctx.Query("limit"))
	if err != nil {
		limit = -1
	}

	offset, _ := strconv.Atoi(ctx.Query("offset"))

	database.DB.Select("*").Limit(limit).Offset(offset).Find(&cashiers).Count(&count)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"count":    count,
		"cashiers": cashiers,
	})
}

func CashierDetails(ctx *fiber.Ctx) error {
	return ctx.SendString("CashierDetails")
}
