package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func TagGetAll(ctx *fiber.Ctx) error {
	var tag []models.TagResponseWithPost

	database.DB.Preload("Post").Find(&tag)

	return ctx.JSON(fiber.Map{
		"tag": tag,
	})
}

func CreateTag(ctx *fiber.Ctx) error {
	tag := new(models.Tag)

	// Parse Body Request to Object Struct
	if err := ctx.BodyParser(tag); err != nil {
		return ctx.Status(http.StatusServiceUnavailable).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	// Manual Validation
	if tag.Name == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	database.DB.Create(&tag)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data succesfully",
		"tag":     tag,
	})

}
