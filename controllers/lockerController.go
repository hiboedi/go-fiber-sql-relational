package controllers

import (
	"go-fiber-relationaldb/database"
	"go-fiber-relationaldb/models"

	"github.com/gofiber/fiber/v2"
)

func LockerGetAll(c *fiber.Ctx) error {
	var lockers []models.Locker

	database.DB.Preload("User").Find(&lockers)

	return c.JSON(fiber.Map{"lockers": lockers})
}

func CreateLocker(c *fiber.Ctx) error {
	locker := new(models.Locker)

	//  parse body request to Object struct

	if err := c.BodyParser(locker); err != nil {
		return c.Status(503).JSON(fiber.Map{
			"err": "cant handle request",
		})
	}

	//    manua validation

	if locker.Code == "" {
		return c.Status(400).JSON(fiber.Map{
			"err": "code is required",
		})
	}

	if locker.UserID == 0 {
		return c.Status(400).JSON(fiber.Map{
			"err": "userID is required",
		})
	}

	database.DB.Create(&locker)

	return c.JSON(fiber.Map{
		"message": "success create data",
		"locker":  locker,
	})
}
