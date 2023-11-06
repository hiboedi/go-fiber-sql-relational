package controllers

import (
	"go-fiber-relationaldb/database"
	"go-fiber-relationaldb/models"

	"github.com/gofiber/fiber/v2"
)

func PostGetAll(c *fiber.Ctx) error {
	var posts []models.PostResponseWithTag

	database.DB.Preload("User").Preload("Tags").Find(&posts)

	return c.JSON(fiber.Map{"posts": posts})
}

func CreatePost(c *fiber.Ctx) error {
	post := new(models.Post)

	//  parse body request to Object struct

	if err := c.BodyParser(post); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	//    manua validation

	if post.Title == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "title is required",
		})
	}

	if post.Body == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "Body is required",
		})
	}

	if post.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "userID is required",
		})
	}

	if len(post.TagsID) == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "tagsID is required",
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

	return c.JSON(fiber.Map{
		"message": "success create data",
		"post":    post,
	})
}
