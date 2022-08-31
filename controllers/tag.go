package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func TagGetAll(ctx *fiber.Ctx) error {
	var tag []models.Tag

	if err := database.DB.Find(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"tags": tag,
	})
}

func TagGetAllWithPost(ctx *fiber.Ctx) error {
	var tag []models.TagResponseWithPost

	if err := database.DB.Preload("Post").Find(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"tags": tag,
	})
}

func CreateTag(ctx *fiber.Ctx) error {
	tag := new(models.Tag)

	// Parse Body Request to Object Struct
	if err := ctx.BodyParser(tag); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	// Manual Validation
	if tag.Name == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	if err := database.DB.Create(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data successfully",
		"tag":     tag,
	})
}

func GetTagById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var tag models.Tag

	if err := database.DB.Where("id = ? ", id).First(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"tag": tag,
	})
}

func GetTagByIdWithPost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var tag models.TagResponseWithPost

	if err := database.DB.Where("id = ? ", id).Preload("Post").First(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"tag": tag,
	})
}

func UpdateTag(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var tag models.Tag

	if err := ctx.BodyParser(&tag); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Where("id = ?", id).Updates(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "tag updated successfully",
	})
}

func DeleteTag(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var tag models.Tag

	if err := database.DB.Where("id = ? ", id).First(&tag).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Delete(&tag, id).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "tag deleted successfully",
	})
}
