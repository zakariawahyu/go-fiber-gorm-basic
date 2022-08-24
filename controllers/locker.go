package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func LockerGetAll(ctx *fiber.Ctx) error {
	var locker []models.Locker

	database.DB.Preload("User").Find(&locker)

	return ctx.JSON(fiber.Map{
		"locker": locker,
	})
}

func CreateLocker(ctx *fiber.Ctx) error {
	locker := new(models.Locker)

	// Parse Body Request to Object Struct
	if err := ctx.BodyParser(locker); err != nil {
		return ctx.Status(http.StatusServiceUnavailable).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	// Manual Validation
	if locker.Code == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "code is required",
		})
	}

	if locker.UserID == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "user_id is required",
		})
	}

	database.DB.Create(&locker)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data succesfully",
		"locker":  locker,
	})

}
