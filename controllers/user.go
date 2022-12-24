package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func UserGetAll(ctx *fiber.Ctx) error {
	var users []models.UserResponse

	if err := database.DB.Find(&users).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(users)
}

func GetUserDetailAll(ctx *fiber.Ctx) error {
	var users []models.User

	if err := database.DB.Preload("Locker").Preload("Post.Tags").Find(&users).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.JSON(users)
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.UserResponse)

	// Parse Body Request to Object Struct
	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(http.StatusServiceUnavailable).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	// Manual Validation
	if user.Name == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	database.DB.Create(&user)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data successfully",
		"user":    user,
	})
}

func GetUserByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.UserResponse

	if err := database.DB.Where("id = ? ", id).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func GetUserDetailByID(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	if err := database.DB.Where("id = ? ", id).Preload("Post.Tags").Preload("Locker").First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(user)
}

func UpdateUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.UserResponse

	if err := ctx.BodyParser(&user); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "cant handle request",
		})
	}

	if err := database.DB.Where("id = ?", id).Updates(&user).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user updated successfully",
	})
}

func DeleteUser(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var user models.User

	if err := database.DB.Where("id = ? ", id).First(&user).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Delete(&user, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "data deleted successfully",
	})
}
