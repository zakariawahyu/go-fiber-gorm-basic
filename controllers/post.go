package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/zakariawahyu/go-fiber-gorm-basic/database"
	"github.com/zakariawahyu/go-fiber-gorm-basic/models"
	"net/http"
)

func PostGetAll(ctx *fiber.Ctx) error {
	var post []models.PostResponseWithTag

	if err := database.DB.Preload("User.Locker").Preload("Tags").Find(&post).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"post": post,
	})
}

func CreatePost(ctx *fiber.Ctx) error {
	post := new(models.Post)

	// Parse Body Request to Object Struct
	if err := ctx.BodyParser(post); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
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

	if post.UserID == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "user_id is required",
		})
	}

	if len(post.TagsID) == 0 {
		return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{
			"err": "tags_id is required",
		})
	}

	database.DB.Create(&post)
	if len(post.TagsID) > 0 {
		for _, tagID := range post.TagsID {
			postTag := new(models.PostTag)
			postTag.PostID = post.ID
			postTag.TagID = tagID
			database.DB.Create(&postTag)
		}
	}

	return ctx.Status(http.StatusOK).JSON(fiber.Map{
		"message": "create data succesfully",
		"post":    post,
	})
}

func GetPostById(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var post models.PostResponseWithTag

	if err := database.DB.Where("id = ?", id).Preload("User.Locker").Preload("Tags").First(&post).Error; err != nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"post": post,
	})
}

func UpdatePost(ctx *fiber.Ctx) error {
	id := ctx.Params("id")
	var post models.Post

	if err := ctx.BodyParser(&post); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := database.DB.Where("id = ?", id).Updates(&post).Error; err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "post updated successfully",
	})
}
