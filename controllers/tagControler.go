package controllers

import (
	"go-fiber-relationaldb/database"
	"go-fiber-relationaldb/models"

	"github.com/gofiber/fiber/v2"
)

func TagGetAll(c *fiber.Ctx) error {
	var tags []models.TagResponseWithPost

	database.DB.Preload("Posts").Find(&tags)

	return c.JSON(fiber.Map{"tags": tags})
}

func CreateTag(c *fiber.Ctx) error {
	tag := new(models.Tag)

	//  parse body request to Object struct

	if err := c.BodyParser(tag); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	//    manua validation

	if tag.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "name is required",
		})
	}

	database.DB.Create(&tag)

	return c.JSON(fiber.Map{
		"message": "success create data",
		"tag":     tag,
	})
}
