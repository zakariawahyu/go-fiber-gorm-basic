package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func PostGetAll(ctx *fiber.Ctx) error {
	var post []models.Post

	database.DB.Preload("User").Find(&post)

	return ctx.JSON(fiber.Map{
		"post": post,
	})
}

func CreatePost(ctx *fiber.Ctx) error {
	post := new(models.Post)

	// Parse Body Request to Object Struct
	if err := ctx.BodyParser(post); err != nil {
		return ctx.Status(http.StatusServiceUnavailable).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	// Manual Validation
	if post.Title == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "title is required",
		})
	}

	if post.Body == "" {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "body is required",
		})
	}

	database.DB.Create(&post)
	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data succesfully",
		"post":    post,
	})

}
