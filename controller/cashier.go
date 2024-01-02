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
	var cashier model.Cashier

	id := ctx.Params("id")

	database.DB.Find(&cashier, "id = ?", id)
	if cashier.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Кассир не найден",
		})
	}

	var updateCashier model.Cashier
	err := ctx.BodyParser(&updateCashier)
	if err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Не удалось обновить кассу",
			"error":   err.Error(),
		})
	}

	if err := validate.Struct(updateCashier); err != nil {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"message": "Не удалось обновить кассу",
			"error":   err.Error(),
		})
	}

	cashier.Name = updateCashier.Name
	cashier.Password = updateCashier.Password
	cashier.UpdatedAt = time.Now()

	database.DB.Save(&cashier)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Касса обновлена",
		"cashier": cashier,
	})
}

func DeleteCashier(ctx *fiber.Ctx) error {
	var cashier model.Cashier

	id := ctx.Params("id")

	database.DB.Where("id = ?", id).First(&cashier)
	if cashier.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Касса не найдена",
		})
	}

	database.DB.Delete(&cashier)

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Касса удалена",
	})
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
	var cashier model.Cashier

	id := ctx.Params("id")

	database.DB.Select("*").Where("id = ?", id).First(&cashier)
	if cashier.ID == 0 {
		return ctx.Status(http.StatusNotFound).JSON(fiber.Map{
			"message": "Касса не найдена",
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"cashier": cashier,
	})
}
