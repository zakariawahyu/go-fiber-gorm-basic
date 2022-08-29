package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func LockerGetAll(ctx *fiber.Ctx) error {
	var locker []models.Locker

	if err := database.DB.Preload("User").Find(&locker).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"lockers": locker,
	})
}

func CreateLocker(ctx *fiber.Ctx) error {
	locker := new(models.LockerResponse)

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

	if err := database.DB.Create(&locker).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"err": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data successfully",
		"locker":  locker,
	})
}

func GetLockerByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var locker models.Locker

	if err := database.DB.Where("id = ? ", id).Preload("User").First(&locker).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"locker": locker,
	})
}

func UpdateLocker(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var locker models.LockerResponse

	if err := ctx.BodyParser(&locker); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cant handle request",
		})
	}

	if err := database.DB.Where("id = ?", id).Updates(&locker).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "locker updated successfully",
	})
}

func DeleteLocker(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var locker models.Locker

	if err := database.DB.Where("id = ? ", id).First(&locker).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Delete(&locker, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data deleted successfully",
	})
}
