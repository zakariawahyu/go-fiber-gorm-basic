package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func UserGetAll(ctx *fiber.Ctx) error {
	var users []models.User

	database.DB.Find(&users)

	return ctx.JSON(fiber.Map{
		"users": users,
	})
}

func CreateUser(ctx *fiber.Ctx) error {
	user := new(models.User)

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
		"message": "create data succesfully",
		"user":    user,
	})

}
